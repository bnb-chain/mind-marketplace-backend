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
	itemId, bucketId, objectId := int64(0), int64(0), int64(0)
	address := ""
	if request.Filter != nil {
		itemId = request.Filter.ItemID
		bucketId = request.Filter.BucketID
		objectId = request.Filter.ObjectID
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

	if itemId <= 0 { //item id is not provided
		if bucketId > 0 { // bucket id is provided
			item, _ := s.itemDao.GetByBucketId(context, bucketId, true)
			itemId = item.Id
		} else if objectId > 0 { // object id is provided
			item, _ := s.itemDao.GetByObjectId(context, objectId, true)
			itemId = item.Id
		}
	}

	total, purchases, err := s.purchaseDao.Search(context, itemId, address, sort, offset, limit)
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
