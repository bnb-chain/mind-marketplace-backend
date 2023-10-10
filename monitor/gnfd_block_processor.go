package monitor

import (
	"context"
	"cosmossdk.io/math"
	"errors"
	"fmt"
	"github.com/bnb-chain/greenfield-data-marketplace-backend/dao"
	"github.com/bnb-chain/greenfield-data-marketplace-backend/database"
	"github.com/bnb-chain/greenfield/types/resource"
	"github.com/bnb-chain/greenfield/x/permission/types"
	abciTypes "github.com/cometbft/cometbft/abci/types"
	"gorm.io/gorm"
	"regexp"
	"strings"
)

type GnfdBlockProcessor struct {
	client   *GnfdCompositeClients
	blockDao dao.GnfdBlockDao
	itemDao  dao.ItemDao
	db       *gorm.DB
}

func NewGnfdBlockProcessor(client *GnfdCompositeClients, blockDao dao.GnfdBlockDao, itemDao dao.ItemDao, db *gorm.DB) *GnfdBlockProcessor {
	return &GnfdBlockProcessor{
		client:   client,
		blockDao: blockDao,
		itemDao:  itemDao,
		db:       db,
	}
}

func (g *GnfdBlockProcessor) GetDatabaseBlockHeight() (uint64, error) {
	block, err := g.blockDao.Max(context.Background())
	if err != nil {
		return 0, err
	}
	return block.Height, nil
}

func (g *GnfdBlockProcessor) GetBlockchainBlockHeight() (uint64, error) {
	return g.client.GetLatestBlockHeight()
}

func (g *GnfdBlockProcessor) Process(blockHeight uint64) error {
	results, err := g.client.GetBlockResults(int64(blockHeight))
	if err != nil {
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
				rawSql, err = g.handleEventCreateGroup(blockHeight, event)
				if err != nil {
					return err
				}
				if rawSql != "" {
					rawCreateSqls = append(rawCreateSqls, rawSql)
				}
			case "greenfield.storage.EventDeleteGroup":
				rawSql, err = g.handleEventDeleteGroup(blockHeight, event)
				if err != nil {
					return err
				}
				if rawSql != "" {
					rawDeleteSqls = append(rawDeleteSqls, rawSql)
				}
			case "greenfield.storage.EventUpdateGroupExtra":
				rawSql, err = g.handleEventUpdateGroupExtra(blockHeight, event)
				if err != nil {
					return err
				}
				if rawSql != "" {
					rawUpdateSqls = append(rawUpdateSqls, rawSql)
				}
			case "greenfield.permission.EventPutPolicy":
				rawSql, err = g.handleEventPutPolicy(blockHeight, event)
				if err != nil {
					return err
				}
				if rawSql != "" {
					rawUpdateSqls = append(rawUpdateSqls, rawSql)
				}
			}
		}
	}

	rawDeleteSqls = append(rawDeleteSqls, fmt.Sprintf("insert into gnfd_blocks (height) values (%d)", blockHeight))

	return g.db.Transaction(func(tx *gorm.DB) error {
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
}

func (g *GnfdBlockProcessor) handleEventCreateGroup(blockHeight uint64, event abciTypes.Event) (string, error) {
	rawSql := ""
	createGroup, err := parseEventCreateGroup(event)
	if err != nil {
		return rawSql, err
	}

	resourceName := ""
	matchBucket, err := regexp.MatchString(groupBucketRegex, createGroup.GroupName)
	if err != nil {
		return rawSql, err
	}

	matchObject, err := regexp.MatchString(groupObjectRegex, createGroup.GroupName)
	if err != nil {
		return rawSql, err
	}

	// the group we do not care about
	if !matchBucket && !matchObject {
		return rawSql, nil
	}

	if matchBucket {
		resourceName = strings.Replace(createGroup.GroupName, groupBucketPrefix, "", 1)
	} else {
		bucketNameObjectName := strings.Replace(createGroup.GroupName, groupObjectPrefix, "", 1)
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

	_, err = g.itemDao.GetByGroupId(context.Background(), int64(createGroup.GroupId.Uint64()))
	if err != nil && err != gorm.ErrRecordNotFound {
		return rawSql, err
	}
	// the group already created
	if err == nil {
		return rawSql, nil
	}

	return fmt.Sprintf("insert into items (group_id, group_name, name, owner_address, status, description, url, price, updated_gnfd_height)"+
		" values (%d, '%s', '%s', '%s', %d, '%s', '%s', '%s', %d)",
		createGroup.GroupId.Uint64(), createGroup.GroupName, resourceName, createGroup.Owner,
		database.ItemPending, extra.Desc, extra.Url, extra.Price, blockHeight), nil
}

func (g *GnfdBlockProcessor) handleEventDeleteGroup(blockHeight uint64, event abciTypes.Event) (string, error) {
	rawSql := ""
	deleteGroup, err := parseEventDeleteGroup(event)
	if err != nil {
		return rawSql, err
	}

	return fmt.Sprintf("update items set status = %d, updated_gnfd_height = %d where group_id = %d",
		database.ItemDelisted, blockHeight, deleteGroup.GroupId.Uint64()), nil
}

func (g *GnfdBlockProcessor) handleEventUpdateGroupExtra(blockHeight uint64, event abciTypes.Event) (string, error) {
	rawSql := ""
	updateGroupExtra, err := parseEventUpdateGroupExtra(event)
	if err != nil {
		return rawSql, err
	}

	extra, err := parseExtra(updateGroupExtra.Extra)
	if err != nil {
		return rawSql, err
	}

	return fmt.Sprintf("update items set description = '%s', url = '%s', price = %s, updated_gnfd_height = %d where group_id = %d",
		extra.Desc, extra.Url, extra.Price, blockHeight, updateGroupExtra.GroupId.Uint64()), nil
}

func (g *GnfdBlockProcessor) handleEventPutPolicy(blockHeight uint64, event abciTypes.Event) (string, error) {
	rawSql := ""
	putPolicy, err := parseEventPutPolicy(event)
	if err != nil {
		return rawSql, err
	}

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
	if putPolicy.Principal.Type == types.PRINCIPAL_TYPE_GNFD_GROUP {
		groupId, err = putPolicy.Principal.GetGroupID()
		if err != nil {
			return rawSql, err
		}
	}
	if groupId.IsZero() {
		return rawSql, nil
	}

	return fmt.Sprintf("update items set `type` = %d, resource_id = %d, status = %d, updated_gnfd_height = %d where group_id = %d",
		resourceType, putPolicy.ResourceId.Uint64(), database.ItemListed, blockHeight, groupId.Uint64()), nil
}
