package dao

import (
	"context"
	"gorm.io/gorm"

	"github.com/bnb-chain/greenfield-data-marketplace-backend/database"
)

type AccountDao interface {
	Create(context context.Context, account *database.Account) error
	Get(context context.Context, id int64) (database.Account, error)
	GetByAddress(context context.Context, address string) (database.Account, error)
	Update(context context.Context, account *database.Account) error
}

type dbAccountDao struct {
	db *gorm.DB
}

func NewDbAccountDao(db *gorm.DB) AccountDao {
	return &dbAccountDao{
		db: db,
	}
}

func (dao *dbAccountDao) Create(context context.Context, account *database.Account) error {
	if err := dao.db.Create(account).Error; err != nil {
		return err
	}
	return nil
}

func (dao *dbAccountDao) Get(context context.Context, id int64) (database.Account, error) {
	var account database.Account
	if err := dao.db.Where("id = ?", id).Take(&account).Error; err != nil {
		return account, err
	}
	return account, nil
}

func (dao *dbAccountDao) GetByAddress(context context.Context, address string) (database.Account, error) {
	var account database.Account

	if err := dao.db.Where(&database.Account{Address: address}).Take(&account).Error; err != nil {
		return account, err
	}
	return account, nil
}

func (dao *dbAccountDao) Update(context context.Context, account *database.Account) error {
	if err := dao.db.Save(account).Error; err != nil {
		return err
	}
	return nil
}
