package dao

import (
	"context"
	"github.com/bnb-chain/mind-marketplace-backend/database"
	"gorm.io/gorm"
)

type GnfdBlockDao interface {
	Create(context context.Context, block *database.GnfdBlock) error
	Max(context context.Context) (database.GnfdBlock, error)
}

type dbGnfdBlockDao struct {
	db *gorm.DB
}

func NewDbGnfdBlockDao(db *gorm.DB) GnfdBlockDao {
	return &dbGnfdBlockDao{
		db: db,
	}
}

func (dao *dbGnfdBlockDao) Create(context context.Context, block *database.GnfdBlock) error {
	if err := dao.db.Create(block).Error; err != nil {
		return err
	}
	return nil
}

func (dao *dbGnfdBlockDao) Max(context context.Context) (database.GnfdBlock, error) {
	var block database.GnfdBlock
	if err := dao.db.Raw("select * from gnfd_blocks order by height desc limit 1").First(&block).Error; err != nil {
		return block, err
	}
	return block, nil
}
