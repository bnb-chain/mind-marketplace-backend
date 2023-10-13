package database

import (
	"github.com/shopspring/decimal"
	"time"
)

type ItemType int

const (
	COLLECTION ItemType = 0
	OBJECT     ItemType = 1
	UNKNOWN    ItemType = 2
)

type ItemStatus int

const (
	ItemPending  ItemStatus = 0  // the group is created
	ItemListed   ItemStatus = 1  // the item is listed
	ItemDelisted ItemStatus = 2  // the item is delisted
	ItemBlocked  ItemStatus = 10 // the item is blocked by admin
)

type Item struct {
	Id                int64           `json:"id" gorm:"primaryKey"`
	Type              int8            `json:"type"`                                              // collection (bucket) or data (object)
	Name              string          `json:"name" gorm:"index:idx_item_name;not null;size:256"` // bucket name or object name
	ResourceId        int64           `json:"resource_id" gorm:"index:idx_item_resource_id"`     // bucket id or object id
	OwnerAddress      string          `json:"owner_address" gorm:"index:idx_item_owner_address;not null;size:42"`
	Description       string          `json:"description" gorm:"size:1024"`
	Url               string          `json:"string" gorm:"size:1024"`
	Price             decimal.Decimal `json:"price" gorm:"type:decimal(65,0);"`
	GroupId           int64           `json:"group_id" gorm:"uniqueIndex:idx_item_group_id;not null;"`
	GroupName         string          `json:"group_name" gorm:"not null;size:256"`
	Status            ItemStatus      `json:"status" gorm:"index:idx_item_status;not null;size:2"`
	ListedAt          int64           `json:"listed_at" gorm:"not null"`
	CreatedAt         time.Time       `json:"created_at" gorm:"index:idx_item_created_at;NOT NULL;type:TIMESTAMP;default:CURRENT_TIMESTAMP;<-:create"`
	UpdatedAt         time.Time       `json:"updated_at" gorm:"NOT NULL;type:TIMESTAMP;default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP"`
	CreatedGnfdHeight int64           `json:"created_gnfd_height"`
	UpdatedGnfdHeight int64           `json:"updated_gnfd_height"`
	UpdatedBscHeight  int64           `json:"updated_bsc_height"`

	Stats *ItemStats `json:"stats" gorm:"foreignKey:ItemId"`
}

type ItemStats struct {
	ItemId    int64           `json:"item_id" gorm:"primaryKey;not null"`
	Sale      int64           `json:"sale"`
	Volume    decimal.Decimal `json:"volume" gorm:"type:decimal(65,0);"`
	UpdatedAt time.Time       `json:"updated_at" gorm:"NOT NULL;type:TIMESTAMP;default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP"`
}
