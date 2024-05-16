package monitor

import (
	"context"
	"fmt"
	"github.com/bnb-chain/mind-marketplace-backend/dao"
	"github.com/bnb-chain/mind-marketplace-backend/database"
	"github.com/bnb-chain/mind-marketplace-backend/metric"
	"github.com/bnb-chain/mind-marketplace-backend/monitor/contracts"
	"github.com/bnb-chain/mind-marketplace-backend/util"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"sort"
	"strings"
)

type BscBlockProcessor struct {
	client              *BscCompositeClients
	marketplaceContract ethcommon.Address
	marketplaceAbi      *abi.ABI
	blockDao            dao.BscBlockDao
	itemDao             dao.ItemDao
	db                  *gorm.DB
	metricServer        *metric.MetricService
}

func NewBscBlockProcessor(client *BscCompositeClients, marketplaceContract string,
	blockDao dao.BscBlockDao, itemDao dao.ItemDao, db *gorm.DB,
	metricServer *metric.MetricService) *BscBlockProcessor {
	marketplaceAbi, err := abi.JSON(strings.NewReader(contracts.MarketplaceMetaData.ABI))
	if err != nil {
		panic("marshal abi error")
	}

	return &BscBlockProcessor{
		client:              client,
		marketplaceContract: ethcommon.HexToAddress(marketplaceContract),
		marketplaceAbi:      &marketplaceAbi,
		blockDao:            blockDao,
		itemDao:             itemDao,
		db:                  db,
		metricServer:        metricServer,
	}
}

func (p *BscBlockProcessor) Name() string {
	return "bsc"
}

func (p *BscBlockProcessor) GetDatabaseBlockHeight() (uint64, error) {
	block, err := p.blockDao.Max(context.Background())
	if err != nil {
		return 0, err
	}
	return block.Height, nil
}

func (p *BscBlockProcessor) GetBlockchainBlockHeight() (uint64, error) {
	return p.client.GetLatestBlockHeightWithRetry()
}

func (p *BscBlockProcessor) Process(blockHeight uint64) error {
	topics := []ethcommon.Hash{
		ethcommon.HexToHash(eventBuyTopic),
		ethcommon.HexToHash(eventListTopic),
		ethcommon.HexToHash(eventDelistTopic),
		ethcommon.HexToHash(eventPriceUpdatedTopic),
	}
	logs, err := p.client.QueryChainLogs(blockHeight, blockHeight, topics, p.marketplaceContract)
	if err != nil {
		util.Logger.Errorf("processor: %s, fail to query chain logs err: %s", p.Name(), err)
		return err
	}

	// to process logs in order
	sort.SliceStable(logs, func(i, j int) bool {
		return logs[i].Index < logs[j].Index
	})

	rawSqls := []string{}
	for _, l := range logs {
		sqls, err := p.handleEventBuy(blockHeight, l)
		if err != nil {
			util.Logger.Errorf("processor: %s, fail to handle EventBuy err: %s", p.Name(), err)
			return err
		}
		rawSqls = append(rawSqls, sqls...)

		sql, err := p.handleEventList(blockHeight, l)
		if err != nil {
			util.Logger.Errorf("processor: %s, fail to handle EventList err: %s", p.Name(), err)
			return err
		}
		if sql != "" {
			rawSqls = append(rawSqls, sql)
		}

		sql, err = p.handleEventDelist(blockHeight, l)
		if err != nil {
			util.Logger.Errorf("processor: %s, fail to handle EventDelist err: %s", p.Name(), err)
			return err
		}
		if sql != "" {
			rawSqls = append(rawSqls, sql)
		}

		sql, err = p.handleEventPriceUpdated(blockHeight, l)
		if err != nil {
			util.Logger.Errorf("processor: %s, fail to handle EventPriceUpdated err: %s", p.Name(), err)
			return err
		}
		if sql != "" {
			rawSqls = append(rawSqls, sql)
		}
	}

	rawSqls = append(rawSqls, fmt.Sprintf("insert into bsc_blocks (height) values (%d)", blockHeight))

	err = p.db.Transaction(func(tx *gorm.DB) error {
		for _, rawSql := range rawSqls {
			if err = tx.Exec(rawSql).Error; err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		util.Logger.Errorf("processor: %s, fail to update database err: %s", p.Name(), err)
		return err
	}

	p.metricServer.SetBscSavedBlockHeight(blockHeight)
	return nil
}

func (p *BscBlockProcessor) handleEventBuy(blockHeight uint64, l types.Log) ([]string, error) {
	event, err := parseBuyEvent(l)
	if err != nil {
		util.Logger.Errorf("processor: %s, fail to parse BuyEvent err: %s", p.Name(), err)
		return nil, err
	}
	if event == nil {
		return nil, nil
	}

	item, err := p.itemDao.GetByGroupId(context.Background(), event.GroupId.Int64(), true)
	if err != nil {
		return nil, err
	}

	if item.Stats == nil {
		item.Stats = &database.ItemStats{
			ItemId: item.Id,
			Sale:   0,
			Volume: decimal.Zero,
		}
	}

	volume := item.Stats.Volume.Add(item.Price)
	sale := item.Stats.Sale + 1

	var rawSqls []string

	blockHeader, err := p.client.GetBlockHeader(blockHeight)
	if err != nil {
		util.Logger.Errorf("processor: %s, fail to get block header err: %s", p.Name(), err)
		return nil, err
	}

	rawSql1 := fmt.Sprintf("insert into purchases (item_id, buyer_address, price, purchased_at, updated_bsc_height) "+
		" values (%d, '%s', '%s', %d,  %d)", item.Id, event.Buyer.String(), item.Price, blockHeader.Time, blockHeight)
	rawSqls = append(rawSqls, rawSql1)

	if sale == 1 { // create new statistics record
		rawSql2 := fmt.Sprintf("insert into item_stats (item_id, volume, sale) "+
			" values (%d, %s, %d)", item.Id, volume, sale)
		rawSqls = append(rawSqls, rawSql2)
	} else {
		rawSql2 := fmt.Sprintf("update item_stats set volume = %s, sale = %d where item_id = %d ",
			volume, sale, item.Id)
		rawSqls = append(rawSqls, rawSql2)
	}

	return rawSqls, nil
}

func (p *BscBlockProcessor) handleEventList(blockHeight uint64, l types.Log) (string, error) {
	event, err := parseListEvent(l)
	if err != nil {
		util.Logger.Errorf("processor: %s, fail to parse ListEvent err: %s", p.Name(), err)
		return "", err
	}
	if event == nil {
		return "", nil
	}

	// item should be existed
	_, err = p.itemDao.GetByGroupId(context.Background(), event.GroupId.Int64(), true)
	if err != nil {
		util.Logger.Errorf("processor: %s, fail to find item %d err: %s, write to listing table", p.Name(), event.GroupId.Int64(), err)

		// insert into a temp table: listing
		rawSql := fmt.Sprintf("insert into listing (price, list_bsc_height, group_id) values (%s, %d, %d)",
			event.Price, blockHeight, event.GroupId)
		return rawSql, nil
	}

	// only list Objects, i.e., pictures
	rawSql := fmt.Sprintf("update items set status = %d, price = %s, updated_bsc_height = %d where group_id = %d and `type` = %d",
		database.ItemListed, event.Price, blockHeight, event.GroupId, database.OBJECT)

	return rawSql, nil
}

func (p *BscBlockProcessor) handleEventDelist(blockHeight uint64, l types.Log) (string, error) {
	event, err := parseDelistEvent(l)
	if err != nil {
		util.Logger.Errorf("processor: %s, fail to parse DelistEvent err: %s", p.Name(), err)
		return "", err
	}
	if event == nil {
		return "", nil
	}

	// item should be existed
	_, err = p.itemDao.GetByGroupId(context.Background(), event.GroupId.Int64(), true)
	if err != nil {
		util.Logger.Errorf("processor: %s, fail to find item %d err: %s", p.Name(), event.GroupId.Int64(), err)
		return "", err
	}

	rawSql := fmt.Sprintf("update items set status = %d, updated_bsc_height = %d where group_id = %d ",
		database.ItemDelisted, blockHeight, event.GroupId)

	return rawSql, nil
}

func (p *BscBlockProcessor) handleEventPriceUpdated(blockHeight uint64, l types.Log) (string, error) {
	event, err := parsePriceUpdatedEvent(l)
	if err != nil {
		util.Logger.Errorf("processor: %s, fail to parse PriceUpdatedEvent err: %s", p.Name(), err)
		return "", err
	}
	if event == nil {
		return "", nil
	}

	rawSql := fmt.Sprintf("update items set pricegi = '%s', updated_bsc_height = %d where group_id = %d ",
		event.Price, blockHeight, event.GroupId)

	return rawSql, nil
}
