package dao

import (
	"context"
	"fmt"
	"github.com/bnb-chain/greenfield-data-marketplace-backend/database"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

const (
	ItemSortCreationAsc     = "CREATION_ASC"
	ItemSortCreationDesc    = "CREATION_DESC"
	ItemSortTotalVolumeAsc  = "TOTAL_VOLUME_ASC"
	ItemSortTotalVolumeDesc = "TOTAL_VOLUME_DESC"
	ItemSortTotalSaleAsc    = "TOTAL_SALE_ASC"
	ItemSortTotalSaleDesc   = "TOTAL_SALE_DESC"
)

type ItemDao interface {
	Create(context context.Context, collection *database.Item) error
	Update(context context.Context, collection *database.Item) error
	Get(context context.Context, id int64) (database.Item, error)
	GetByGroupId(context context.Context, groupId int64) (database.Item, error)
	Search(context context.Context, address, keyword string, hideBlocked bool, sort string, offset, limit int) (int64, []*database.Item, error)
}

type dbItemDao struct {
	db *gorm.DB
}

func NewDbItemDao(db *gorm.DB) ItemDao {
	return &dbItemDao{
		db: db,
	}
}

func (dao *dbItemDao) Create(context context.Context, item *database.Item) error {
	dbTx := dao.db.Begin()
	stats := item.Stats

	if err := dbTx.Model(database.Item{}).Omit("Stats").Create(item).Error; err != nil {
		dbTx.Rollback()
		return err
	}

	if stats == nil {
		stats = &database.ItemStats{
			Sale:   0,
			Volume: decimal.NewFromInt(0),
		}
	}
	stats.ItemId = item.Id

	if err := dbTx.Model(database.ItemStats{}).Create(stats).Error; err != nil {
		dbTx.Rollback()
		return err
	}
	return dbTx.Commit().Error
}

func (dao *dbItemDao) Update(context context.Context, collection *database.Item) error {
	if err := dao.db.Omit("Stats").Save(collection).Error; err != nil {
		return err
	}

	return nil
}

func (dao *dbItemDao) Get(context context.Context, id int64) (database.Item, error) {
	var item = database.Item{}
	if err := dao.db.Preload("Stats").Where("id = ?", id).Take(&item).Error; err != nil {
		return item, err
	}
	return item, nil
}

func (dao *dbItemDao) GetByGroupId(context context.Context, groupId int64) (database.Item, error) {
	var item = database.Item{}
	if err := dao.db.Preload("Stats").Where("group_id = ?", groupId).Take(&item).Error; err != nil {
		return item, err
	}
	return item, nil
}

func (dao *dbItemDao) Search(context context.Context, address, keyword string, hideBlocked bool, sort string, offset, limit int) (total int64, items []*database.Item, err error) {
	rawSql := " where 1 = 1 "
	parameters := make([]interface{}, 0)

	if len(address) > 0 {
		rawSql = rawSql + ` and owner_address = ?`
		parameters = append(parameters, address)
	}

	if len(keyword) > 0 {
		rawSql = rawSql + ` and name like ?`
		parameters = append(parameters, "%"+keyword+"%")
	}

	if hideBlocked {
		rawSql = rawSql + ` and status not in ( ?, ? ) `
		parameters = append(parameters, database.ItemDelisted)
		parameters = append(parameters, database.ItemBlocked)
	} else {
		rawSql = rawSql + ` and status not in ( ? ) `
		parameters = append(parameters, database.ItemDelisted)
	}

	countSql := "select count(1) from items " + rawSql

	err = dao.db.Raw(countSql, parameters...).Scan(&total).Error
	if err != nil {
		return 0, nil, err
	}
	if total == 0 {
		return
	}

	dataSql := "select * from items inner join item_stats on items.id = item_stats.item_id " + rawSql
	dataSql = dataSql + " order by "
	switch sort {
	case ItemSortCreationAsc:
		dataSql = dataSql + "id asc "
	case ItemSortCreationDesc:
		dataSql = dataSql + "id desc "
	case ItemSortTotalVolumeAsc:
		dataSql = dataSql + "volume asc "
	case ItemSortTotalVolumeDesc:
		dataSql = dataSql + "volume desc "
	case ItemSortTotalSaleAsc:
		dataSql = dataSql + "sale asc "
	case ItemSortTotalSaleDesc:
		dataSql = dataSql + "sale desc "
	default:
		return 0, nil, fmt.Errorf("unsupported sort string: %s", sort)
	}
	dataSql = dataSql + fmt.Sprintf("limit %d, %d", offset, limit)

	err = dao.db.Raw(dataSql, parameters...).Scan(&items).Error
	if err != nil {
		return 0, nil, err
	}

	itemIds := []int64{}
	for _, item := range items {
		itemIds = append(itemIds, item.Id)
	}

	var itemStats []*database.ItemStats
	err = dao.db.Where("item_id in ?", itemIds).Find(&itemStats).Error
	if err != nil {
		return 0, nil, err
	}

	for _, item := range items {
		for _, stats := range itemStats {
			if item.Id == stats.ItemId {
				item.Stats = stats
			}
		}
	}

	return
}
