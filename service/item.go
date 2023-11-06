package service

import (
	"context"
	"fmt"
	"github.com/bnb-chain/mind-marketplace-backend/dao"
	"github.com/bnb-chain/mind-marketplace-backend/models"
	"gorm.io/gorm"
)

type Item interface {
	Get(context context.Context, id int64) (*models.Item, error)
	GetByGroup(context context.Context, groupId int64) (*models.Item, error)
	Search(context context.Context, request *models.SearchItemRequest) (int64, []*models.Item, error)
}

type ItemService struct {
	itemDao dao.ItemDao
}

func NewItemService(itemDao dao.ItemDao) Item {
	return &ItemService{
		itemDao: itemDao,
	}
}

func (s *ItemService) Get(context context.Context, id int64) (*models.Item, error) {
	item, err := s.itemDao.Get(context, id, false)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, NotFoundErr
		} else {
			return nil, fmt.Errorf("fail to get item")
		}
	}

	return convertItem(item), nil
}

func (s *ItemService) GetByGroup(context context.Context, groupId int64) (*models.Item, error) {
	item, err := s.itemDao.GetByGroupId(context, groupId, false)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, NotFoundErr
		} else {
			return nil, fmt.Errorf("fail to get item")
		}
	}

	return convertItem(item), nil
}

func (s *ItemService) Search(context context.Context, request *models.SearchItemRequest) (int64, []*models.Item, error) {
	address, keyword := "", ""
	if request.Filter != nil {
		address = request.Filter.Address
		keyword = request.Filter.Keyword
	}

	if len(address) != 0 {
		if err := validateAddress(address); err != nil {
			return 0, nil, err
		}
	}
	if len(keyword) != 0 {
		if err := validateKeyword(keyword); err != nil {
			return 0, nil, err
		}
	}

	offset := 0
	if request.Offset != nil {
		offset = int(*request.Offset)
	}
	sort := models.SearchItemRequestSortCREATIONDESC
	if request.Sort != nil {
		sort = *request.Sort
	}
	limit := defaultSearchLimit
	if request.Limit > 0 {
		limit = int(request.Limit)
	}
	if limit > maxSearchLimit {
		return 0, nil, TooBigLimitErr
	}

	total, items, err := s.itemDao.Search(context, request.Filter.CategoryID, address, keyword, false, sort, offset, limit)
	if err != nil {
		return 0, nil, fmt.Errorf("fail to search item")
	}

	if len(items) == 0 {
		return total, []*models.Item{}, nil
	}

	page := make([]*models.Item, 0)
	for _, c := range items {
		r := convertItem(*c)
		page = append(page, r)
	}

	return total, page, nil
}
