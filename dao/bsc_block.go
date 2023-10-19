package dao

import (
	"context"
	"github.com/bnb-chain/mind-marketplace-backend/database"
	"gorm.io/gorm"
)

type BscBlockDao interface {
	Create(context context.Context, block *database.BscBlock) error
	Max(context context.Context) (database.BscBlock, error)
}

type dbBscBlockDao struct {
	db *gorm.DB
}

func NewDbBscBlockDao(db *gorm.DB) BscBlockDao {
	return &dbBscBlockDao{
		db: db,
	}
}

func (dao *dbBscBlockDao) Create(context context.Context, block *database.BscBlock) error {
	if err := dao.db.Create(block).Error; err != nil {
		return err
	}
	return nil
}

func (dao *dbBscBlockDao) Max(context context.Context) (database.BscBlock, error) {
	var block database.BscBlock
	if err := dao.db.Raw("select * from bsc_blocks order by height desc limit 1").First(&block).Error; err != nil {
		return block, err
	}
	return block, nil
}
