package dao

import (
	"context"
	"gorm.io/gorm"

	"github.com/bnb-chain/mind-marketplace-backend/database"
)

type CategoryDao interface {
	Create(context context.Context, category *database.Category) error
	Get(context context.Context, name string) (database.Category, error)
	GetAll(context context.Context) ([]database.Category, error)
}

type dbCategoryDao struct {
	db *gorm.DB
}

func NewDbCategoryDao(db *gorm.DB) CategoryDao {
	return &dbCategoryDao{
		db: db,
	}
}

func (dao *dbCategoryDao) Create(context context.Context, category *database.Category) error {
	if err := dao.db.Omit("Items").Create(category).Error; err != nil {
		return err
	}
	return nil
}

func (dao *dbCategoryDao) Get(context context.Context, name string) (database.Category, error) {
	var category = database.Category{}
	if err := dao.db.Raw("select from categories where lower(name) = ?", name).Take(&category).Error; err != nil {
		return category, err
	}
	return category, nil
}

func (dao *dbCategoryDao) GetAll(context context.Context) ([]database.Category, error) {
	var categories []database.Category
	if err := dao.db.Omit("Items").Find(&categories).Error; err != nil {
		return categories, err
	}
	return categories, nil
}
