package database

import (
	"github.com/shopspring/decimal"
)

type Listing struct {
	Id            int64           `json:"id" gorm:"primaryKey"`
	Price         decimal.Decimal `json:"price" gorm:"type:decimal(65,0);"`
	ListBscHeight uint64          `gorm:"NOT NULL;index:idx_list_bsc_height"`
	GroupId       int64           `json:"group_id" gorm:"uniqueIndex:idx_list_group_id;not null;"`
}
