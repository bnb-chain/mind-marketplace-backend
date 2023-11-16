package service

import (
	"context"
	"fmt"
	"gorm.io/gorm"

	"github.com/bnb-chain/mind-marketplace-backend/dao"
	"github.com/bnb-chain/mind-marketplace-backend/models"
)

type Purchase interface {
	Get(context context.Context, id int64) (*models.Purchase, error)
	Search(context context.Context, request *models.SearchPurchaseRequest) (int64, []*models.Purchase, error)
	Query(context context.Context, request *models.QueryPurchaseRequest) (int64, []*models.Purchase, error)
}

type PurchaseService struct {
	purchaseDao dao.PurchaseDao
	itemDao     dao.ItemDao
}

func NewPurchaseService(purchaseDao dao.PurchaseDao, itemDao dao.ItemDao) Purchase {
	return &PurchaseService{
		purchaseDao: purchaseDao,
		itemDao:     itemDao,
	}
}

func (s *PurchaseService) Get(context context.Context, id int64) (*models.Purchase, error) {
	purchase, err := s.purchaseDao.Get(context, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, NotFoundErr
		} else {
			return nil, fmt.Errorf("fail to get purchase")
		}
	}

	return convertPurchase(purchase), nil
}

func (s *PurchaseService) Search(context context.Context, request *models.SearchPurchaseRequest) (int64, []*models.Purchase, error) {
	itemIds, bucketIds, objectIds := make([]int64, 0), make([]int64, 0), make([]int64, 0)
	address := ""
	if request.Filter != nil {
		if request.Filter.ItemID > 0 {
			itemIds = append(itemIds, request.Filter.ItemID)
		} else if request.Filter.BucketID > 0 {
			bucketIds = append(bucketIds, request.Filter.BucketID)
		} else if request.Filter.ObjectID > 0 {
			objectIds = append(objectIds, request.Filter.ObjectID)
		}
		address = request.Filter.Address
	}

	if len(address) != 0 {
		if err := validateAddress(address); err != nil {
			return 0, nil, err
		}
	}

	offset := 0
	if request.Offset != nil {
		offset = int(*request.Offset)
	}
	sort := models.SearchPurchaseRequestSortCREATIONDESC
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

	total, purchases, err := s.purchaseDao.Query(context, itemIds, bucketIds, objectIds, address, sort, offset, limit)
	if err != nil {
		return 0, nil, fmt.Errorf("fail to search purchase")
	}

	if len(purchases) == 0 {
		return total, []*models.Purchase{}, nil
	}

	page := make([]*models.Purchase, 0)
	for _, p := range purchases {
		r := convertPurchase(*p)
		page = append(page, r)
	}

	return total, page, nil
}

func (s *PurchaseService) Query(context context.Context, request *models.QueryPurchaseRequest) (int64, []*models.Purchase, error) {
	itemIds, bucketIds, objectIds := make([]int64, 0), make([]int64, 0), make([]int64, 0)
	address := ""
	if request.Filter != nil {
		itemIds = request.Filter.ItemIds
		bucketIds = request.Filter.BucketIds
		objectIds = request.Filter.ObjectIds
		address = request.Filter.Address
	}

	if len(address) != 0 {
		if err := validateAddress(address); err != nil {
			return 0, nil, err
		}
	}

	offset := 0
	if request.Offset != nil {
		offset = int(*request.Offset)
	}
	sort := models.SearchPurchaseRequestSortCREATIONDESC
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

	total, purchases, err := s.purchaseDao.Query(context, itemIds, bucketIds, objectIds, address, sort, offset, limit)
	if err != nil {
		return 0, nil, fmt.Errorf("fail to search purchase")
	}

	if len(purchases) == 0 {
		return total, []*models.Purchase{}, nil
	}

	page := make([]*models.Purchase, 0)
	for _, p := range purchases {
		r := convertPurchase(*p)
		page = append(page, r)
	}

	return total, page, nil
}
