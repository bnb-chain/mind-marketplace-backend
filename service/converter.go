package service

import (
	"github.com/bnb-chain/mind-marketplace-backend/database"
	"github.com/bnb-chain/mind-marketplace-backend/models"
	"github.com/bnb-chain/mind-marketplace-backend/util"
)

func convertAccount(account database.Account) *models.Account {
	id := account.Id
	return &models.Account{
		ID:                &id,
		Address:           &account.Address,
		CreatedAt:         account.CreatedAt.Unix(),
		TwitterUserName:   account.TwitterUserName,
		TwitterVerified:   account.TwitterVerified,
		InstagramUserName: account.InstagramUserName,
		InstagramVerified: account.InstagramVerified,
	}
}

func convertItem(item database.Item) *models.Item {
	id := item.Id
	typ := formatItemType(database.ItemType(item.Type))
	result := models.Item{
		ID:           &id,
		Type:         &typ,
		Name:         &item.Name,
		GroupID:      item.GroupId,
		GroupName:    item.GroupName,
		CreatedAt:    item.ListedAt,
		Description:  item.Description,
		URL:          item.Url,
		Price:        util.Decimal(item.Price),
		OwnerAddress: item.OwnerAddress,
	}

	if item.Stats != nil {
		result.TotalSale = item.Stats.Sale
		result.TotalVolume = util.Decimal(item.Stats.Volume)
	}

	if item.Status == database.ItemListed {
		result.Status = "LISTED"
	} else if item.Status == database.ItemBlocked {
		result.Status = "BLOCKED"
	} else {
		result.Status = "PENDING"
	}

	return &result
}

func convertPurchase(purchase database.Purchase) *models.Purchase {
	id := purchase.Id
	item := convertItem(*purchase.Item)

	return &models.Purchase{
		ID:           &id,
		BuyerAddress: &purchase.BuyerAddress,
		Price:        util.Decimal(purchase.Price),
		CreatedAt:    purchase.PurchasedAt,
		Item:         item,
	}
}

func formatItemType(typ database.ItemType) string {
	switch typ {
	case database.COLLECTION:
		return "COLLECTION"
	case database.OBJECT:
		return "OBJECT"
	default:
		return "UNKNOWN"
	}
}
