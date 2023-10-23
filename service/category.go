package service

import (
	"context"
	"fmt"

	"gorm.io/gorm"

	"github.com/bnb-chain/mind-marketplace-backend/dao"
	"github.com/bnb-chain/mind-marketplace-backend/models"
)

type Category interface {
	GetAll(context context.Context) ([]*models.Category, error)
}

type CategoryService struct {
	categoryDao dao.CategoryDao
}

func NewCategoryService(categoryDao dao.CategoryDao) Category {
	return &CategoryService{
		categoryDao: categoryDao,
	}
}

func (s *CategoryService) GetAll(context context.Context) ([]*models.Category, error) {
	categories, err := s.categoryDao.GetAll(context)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, NotFoundErr
		} else {
			return nil, fmt.Errorf("fail to get category")
		}
	}

	all := make([]*models.Category, 0)
	for _, c := range categories {
		all = append(all, convertCategory(c))
	}

	return all, nil
}
