package dao

import (
	"context"
	"github.com/bnb-chain/mind-marketplace-backend/database"
	"gorm.io/gorm"
)

type ListingDao interface {
	Create(context context.Context, block *database.Listing) error
	GetByGroupId(context context.Context, groupId int64) (database.Listing, error)
}

type DbListingDao struct {
	db *gorm.DB
}

func NewDbListingDao(db *gorm.DB) *DbListingDao {
	return &DbListingDao{
		db: db,
	}
}

func (dao *DbListingDao) Create(context context.Context, listing *database.Listing) error {
	if err := dao.db.Create(listing).Error; err != nil {
		return err
	}
	return nil
}

func (dao *DbListingDao) GetByGroupId(context context.Context, groupId int64) (database.Listing, error) {
	var item = database.Listing{}
	if err := dao.db.Preload("Stats").Where("group_id = ?", groupId).Take(&item).Error; err != nil {
		return item, err
	}
	return item, nil
}
