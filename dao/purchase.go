package dao

import (
	"context"
	"fmt"
	"gorm.io/gorm"

	"github.com/bnb-chain/mind-marketplace-backend/database"
)

const (
	PurchaseSortCreationAsc  = "CREATION_ASC"
	PurchaseSortCreationDesc = "CREATION_DESC"
	PurchaseSortPriceAsc     = "PRICE_ASC"
	PurchaseSortPriceDesc    = "PRICE_DESC"
)

type PurchaseDao interface {
	Create(context context.Context, purchase *database.Purchase) error
	Update(context context.Context, purchase *database.Purchase) error
	Get(context context.Context, id int64) (database.Purchase, error)
	Search(context context.Context, itemId int64, address string, sort string, offset, limit int) (int64, []*database.Purchase, error)
	Query(context context.Context, itemIds []int64, bucketIds []int64, objectIds []int64, address string, sort string, offset, limit int) (int64, []*database.Purchase, error)
}

type dbPurchaseDao struct {
	db *gorm.DB
}

func NewDbPurchaseDao(db *gorm.DB) PurchaseDao {
	return &dbPurchaseDao{
		db: db,
	}
}

func (dao *dbPurchaseDao) Create(context context.Context, purchase *database.Purchase) error {

	if err := dao.db.Create(purchase).Error; err != nil {
		return err
	}
	return nil
}

func (dao *dbPurchaseDao) Update(context context.Context, purchase *database.Purchase) error {
	if err := dao.db.Save(purchase).Error; err != nil {
		return err
	}

	return nil
}

func (dao *dbPurchaseDao) Get(context context.Context, id int64) (database.Purchase, error) {
	var purchase = database.Purchase{}
	if err := dao.db.Preload("Item").Where("id = ?", id).Take(&purchase).Error; err != nil {
		return purchase, err
	}
	return purchase, nil
}

func (dao *dbPurchaseDao) Search(context context.Context, itemId int64, address string, sort string, offset, limit int) (total int64, purchases []*database.Purchase, err error) {
	rawSql := fmt.Sprintf(" inner join items i on p.item_id = i.id where i.status = %d", database.ItemListed)
	parameters := make([]interface{}, 0)

	if itemId > 0 {
		rawSql = rawSql + ` and item_id = ?`
		parameters = append(parameters, itemId)
	}

	if len(address) > 0 {
		rawSql = rawSql + ` and buyer_address = ?`
		parameters = append(parameters, address)
	}

	countSql := "select count(1) from purchases p " + rawSql

	err = dao.db.Raw(countSql, parameters...).Scan(&total).Error
	if err != nil {
		return 0, nil, err
	}
	if total == 0 {
		return
	}

	dataSql := "select * from purchases p " + rawSql
	dataSql = dataSql + " order by "
	switch sort {
	case PurchaseSortCreationAsc:
		dataSql = dataSql + "p.id asc "
	case PurchaseSortCreationDesc:
		dataSql = dataSql + "p.id desc "
	case PurchaseSortPriceAsc:
		dataSql = dataSql + "p.price asc "
	case PurchaseSortPriceDesc:
		dataSql = dataSql + "p.price desc "
	default:
		return 0, nil, fmt.Errorf("unsupported sort string: %s", sort)
	}
	dataSql = dataSql + fmt.Sprintf("limit %d, %d", offset, limit)

	err = dao.db.Preload("Item").Raw(dataSql, parameters...).Scan(&purchases).Error
	if err != nil {
		return 0, nil, err
	}

	itemIds := []int64{}
	for _, purchase := range purchases {
		itemIds = append(itemIds, purchase.ItemId)
	}

	var items []*database.Item
	err = dao.db.Preload("Stats").Where("id in ?", itemIds).Find(&items).Error
	if err != nil {
		return 0, nil, err
	}

	for _, purchase := range purchases {
		for _, item := range items {
			if purchase.ItemId == item.Id {
				purchase.Item = item
			}
		}
	}

	return
}

func (dao *dbPurchaseDao) Query(context context.Context, itemIds []int64, bucketIds []int64, objectIds []int64, address string, sort string, offset, limit int) (total int64, purchases []*database.Purchase, err error) {
	rawSql := fmt.Sprintf(" inner join items i on p.item_id = i.id where i.status = %d", database.ItemListed)
	parameters := make([]interface{}, 0)

	if len(itemIds) > 0 {
		rawSql = rawSql + ` and item_id in ?`
		parameters = append(parameters, itemIds)
	} else if len(bucketIds) > 0 {
		rawSql = rawSql + ` and resource_id in ? and type = ?`
		parameters = append(parameters, bucketIds)
		parameters = append(parameters, database.COLLECTION)
	} else if len(objectIds) > 0 {
		rawSql = rawSql + ` and resource_id in ? and type = ?`
		parameters = append(parameters, objectIds)
		parameters = append(parameters, database.OBJECT)
	}

	if len(address) > 0 {
		rawSql = rawSql + ` and buyer_address = ?`
		parameters = append(parameters, address)
	}

	countSql := "select count(1) from purchases p " + rawSql

	err = dao.db.Raw(countSql, parameters...).Scan(&total).Error
	if err != nil {
		return 0, nil, err
	}
	if total == 0 {
		return
	}

	dataSql := "select * from purchases p " + rawSql
	dataSql = dataSql + " order by "
	switch sort {
	case PurchaseSortCreationAsc:
		dataSql = dataSql + "p.id asc "
	case PurchaseSortCreationDesc:
		dataSql = dataSql + "p.id desc "
	case PurchaseSortPriceAsc:
		dataSql = dataSql + "p.price asc "
	case PurchaseSortPriceDesc:
		dataSql = dataSql + "p.price desc "
	default:
		return 0, nil, fmt.Errorf("unsupported sort string: %s", sort)
	}
	dataSql = dataSql + fmt.Sprintf("limit %d, %d", offset, limit)

	err = dao.db.Preload("Item").Raw(dataSql, parameters...).Scan(&purchases).Error
	if err != nil {
		return 0, nil, err
	}

	itemIds = []int64{}
	for _, purchase := range purchases {
		itemIds = append(itemIds, purchase.ItemId)
	}

	var items []*database.Item
	err = dao.db.Preload("Stats").Where("id in ?", itemIds).Find(&items).Error
	if err != nil {
		return 0, nil, err
	}

	for _, purchase := range purchases {
		for _, item := range items {
			if purchase.ItemId == item.Id {
				purchase.Item = item
			}
		}
	}

	return
}
