package monitor

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"strings"

	"cosmossdk.io/math"
	"github.com/bnb-chain/greenfield-data-marketplace-backend/dao"
	"github.com/bnb-chain/greenfield-data-marketplace-backend/database"
	"github.com/bnb-chain/greenfield-data-marketplace-backend/metric"
	"github.com/bnb-chain/greenfield-data-marketplace-backend/util"
	"github.com/bnb-chain/greenfield/types/resource"
	permTypes "github.com/bnb-chain/greenfield/x/permission/types"
	storageTypes "github.com/bnb-chain/greenfield/x/storage/types"
	abciTypes "github.com/cometbft/cometbft/abci/types"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"gorm.io/gorm"
)

type GnfdBlockProcessor struct {
	client       *GnfdCompositeClients
	blockDao     dao.GnfdBlockDao
	itemDao      dao.ItemDao
	db           *gorm.DB
	metricServer *metric.MetricService

	groupBucketRegex  string
	groupBucketPrefix string
	groupObjectRegex  string
	groupObjectPrefix string
}

func NewGnfdBlockProcessor(client *GnfdCompositeClients,
	blockDao dao.GnfdBlockDao, itemDao dao.ItemDao, db *gorm.DB,
	metricServer *metric.MetricService,
	groupBucketRegex string,
	groupBucketPrefix string,
	groupObjectRegex string,
	groupObjectPrefix string) *GnfdBlockProcessor {
	return &GnfdBlockProcessor{
		client:            client,
		blockDao:          blockDao,
		itemDao:           itemDao,
		db:                db,
		metricServer:      metricServer,
		groupBucketRegex:  groupBucketRegex,
		groupBucketPrefix: groupBucketPrefix,
		groupObjectRegex:  groupObjectRegex,
		groupObjectPrefix: groupObjectPrefix,
	}
}

func (p *GnfdBlockProcessor) Name() string {
	return "gnfd"
}

func (p *GnfdBlockProcessor) GetDatabaseBlockHeight() (uint64, error) {
	block, err := p.blockDao.Max(context.Background())
	if err != nil {
		return 0, err
	}
	return block.Height, nil
}

func (p *GnfdBlockProcessor) GetBlockchainBlockHeight() (uint64, error) {
	return p.client.GetLatestBlockHeight()
}

func (p *GnfdBlockProcessor) Process(blockHeight uint64) error {
	results, err := p.client.GetBlockResults(int64(blockHeight))
	if err != nil {
		util.Logger.Errorf("processor: %s, fail to block results err: %s", p.Name(), err)
		return err
	}

	rawCreateSqls := []string{}
	rawUpdateSqls := []string{}
	rawDeleteSqls := []string{}

	for _, result := range results.TxsResults {
		for _, event := range result.Events {
			rawSql := ""
			switch event.Type {
			case "greenfield.storage.EventCreateGroup":
				rawSql, err = p.handleEventCreateGroup(blockHeight, event)
				if err != nil {
					util.Logger.Errorf("processor: %s, fail to handle EventCreateGroup err: %s", p.Name(), err)
					return err
				}
				if rawSql != "" {
					rawCreateSqls = append(rawCreateSqls, rawSql)
				}
			case "greenfield.storage.EventDeleteGroup":
				rawSql, err = p.handleEventDeleteGroup(blockHeight, event)
				if err != nil {
					util.Logger.Errorf("processor: %s, fail to handle EventDeleteGroup err: %s", p.Name(), err)
					return err
				}
				if rawSql != "" {
					rawDeleteSqls = append(rawDeleteSqls, rawSql)
				}
			case "greenfield.storage.EventUpdateGroupExtra":
				rawSql, err = p.handleEventUpdateGroupExtra(blockHeight, event)
				if err != nil {
					util.Logger.Errorf("processor: %s, fail to handle EventUpdateGroupExtra err: %s", p.Name(), err)
					return err
				}
				if rawSql != "" {
					rawUpdateSqls = append(rawUpdateSqls, rawSql)
				}
			case "greenfield.permission.EventPutPolicy":
				rawSql, err = p.handleEventPutPolicy(blockHeight, event)
				if err != nil {
					util.Logger.Errorf("processor: %s, fail to handle EventPutPolicy err: %s", p.Name(), err)
					return err
				}
				if rawSql != "" {
					rawUpdateSqls = append(rawUpdateSqls, rawSql)
				}
			}
		}
	}

	rawDeleteSqls = append(rawDeleteSqls, fmt.Sprintf("insert into gnfd_blocks (height) values (%d)", blockHeight))

	err = p.db.Transaction(func(tx *gorm.DB) error {
		for _, rawSql := range rawCreateSqls {
			if err = tx.Exec(rawSql).Error; err != nil {
				return err
			}
		}
		for _, rawSql := range rawUpdateSqls {
			if err = tx.Exec(rawSql).Error; err != nil {
				return err
			}
		}
		for _, rawSql := range rawDeleteSqls {
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

	p.metricServer.SetGnfdSavedBlockHeight(blockHeight)
	return nil
}

func (p *GnfdBlockProcessor) handleEventCreateGroup(blockHeight uint64, event abciTypes.Event) (string, error) {
	rawSql := ""
	e, err := sdkTypes.ParseTypedEvent(event)
	if err != nil {
		util.Logger.Errorf("processor: %s, fail to parse EventCreateGroup err: %s", p.Name(), err)
		return rawSql, err
	}
	createGroup := e.(*storageTypes.EventCreateGroup)

	resourceName := ""
	matchBucket, err := regexp.MatchString(p.groupBucketRegex, createGroup.GroupName)
	if err != nil {
		return rawSql, err
	}

	matchObject, err := regexp.MatchString(p.groupObjectRegex, createGroup.GroupName)
	if err != nil {
		return rawSql, err
	}

	// the group we do not care about
	if !matchBucket && !matchObject {
		return rawSql, nil
	}

	if matchBucket {
		resourceName = strings.Replace(createGroup.GroupName, p.groupBucketPrefix, "", 1)
	} else {
		bucketNameObjectName := strings.Replace(createGroup.GroupName, p.groupObjectPrefix, "", 1)
		_, objectName, found := strings.Cut(bucketNameObjectName, "_")
		if !found {
			return rawSql, errors.New("cannot parse object name")
		}
		resourceName = objectName
	}

	extra, err := parseExtra(createGroup.Extra)
	if err != nil {
		return rawSql, err
	}

	_, err = p.itemDao.GetByGroupId(context.Background(), int64(createGroup.GroupId.Uint64()), true)
	if err != nil && err != gorm.ErrRecordNotFound {
		return rawSql, err
	}
	// the group already created
	if err == nil {
		return rawSql, nil
	}

	block, err := p.client.GetBlock(int64(blockHeight))
	if err != nil {
		util.Logger.Errorf("processor: %s, fail to get block err: %s", p.Name(), err)
		return rawSql, err
	}

	return fmt.Sprintf("insert into items (group_id, group_name, name, owner_address, status, description, url, price, listed_at, created_gnfd_height)"+
		" values (%d, '%s', '%s', '%s', %d, '%s', '%s', '%s', %d, %d)",
		createGroup.GroupId.Uint64(), createGroup.GroupName, resourceName, createGroup.Owner,
		database.ItemPending, extra.Desc, extra.Url, extra.Price, block.Time.Unix(), blockHeight), nil
}

func (p *GnfdBlockProcessor) handleEventDeleteGroup(blockHeight uint64, event abciTypes.Event) (string, error) {
	rawSql := ""
	e, err := sdkTypes.ParseTypedEvent(event)
	if err != nil {
		util.Logger.Errorf("processor: %s, fail to parse EventCreateGroup err: %s", p.Name(), err)
		return rawSql, err
	}
	deleteGroup := e.(*storageTypes.EventDeleteGroup)

	return fmt.Sprintf("update items set status = %d, updated_gnfd_height = %d where group_id = %d",
		database.ItemDelisted, blockHeight, deleteGroup.GroupId.Uint64()), nil
}

func (p *GnfdBlockProcessor) handleEventUpdateGroupExtra(blockHeight uint64, event abciTypes.Event) (string, error) {
	rawSql := ""
	e, err := sdkTypes.ParseTypedEvent(event)
	if err != nil {
		util.Logger.Errorf("processor: %s, fail to parse EventCreateGroup err: %s", p.Name(), err)
		return rawSql, err
	}
	updateGroupExtra := e.(*storageTypes.EventUpdateGroupExtra)

	_, err = p.itemDao.GetByGroupId(context.Background(), int64(updateGroupExtra.GroupId.Uint64()), true)
	if err != nil && err != gorm.ErrRecordNotFound {
		return rawSql, err
	}

	if err != nil && err == gorm.ErrRecordNotFound { // the group we do not care about
		return rawSql, nil
	}

	extra, err := parseExtra(updateGroupExtra.Extra)
	if err != nil {
		return rawSql, err
	}

	return fmt.Sprintf("update items set description = '%s', url = '%s', price = %s, updated_gnfd_height = %d where group_id = %d",
		extra.Desc, extra.Url, extra.Price, blockHeight, updateGroupExtra.GroupId.Uint64()), nil
}

func (p *GnfdBlockProcessor) handleEventPutPolicy(blockHeight uint64, event abciTypes.Event) (string, error) {
	rawSql := ""
	e, err := sdkTypes.ParseTypedEvent(event)
	if err != nil {
		util.Logger.Errorf("processor: %s, fail to parse EventCreateGroup err: %s", p.Name(), err)
		return rawSql, err
	}
	putPolicy := e.(*permTypes.EventPutPolicy)

	resourceType := database.UNKNOWN
	if putPolicy.ResourceType == resource.RESOURCE_TYPE_BUCKET {
		resourceType = database.COLLECTION
	} else if putPolicy.ResourceType == resource.RESOURCE_TYPE_OBJECT {
		resourceType = database.OBJECT
	}
	if resourceType == database.UNKNOWN {
		return rawSql, nil
	}

	groupId := math.NewUint(0)
	if putPolicy.Principal.Type == permTypes.PRINCIPAL_TYPE_GNFD_GROUP {
		groupId, err = putPolicy.Principal.GetGroupID()
		if err != nil {
			return rawSql, err
		}
	}
	if groupId.IsZero() {
		return rawSql, nil
	}

	return fmt.Sprintf("update items set `type` = %d, resource_id = %d, updated_gnfd_height = %d where group_id = %d",
		resourceType, putPolicy.ResourceId.Uint64(), blockHeight, groupId.Uint64()), nil
}
