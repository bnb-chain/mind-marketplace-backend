package database

import (
	"github.com/shopspring/decimal"
	"time"
)

type Purchase struct {
	Id               int64           `json:"id" gorm:"primaryKey"`
	BuyerAddress     string          `json:"buyer_address" gorm:"index:idx_purchase_buyer_address"`
	ItemId           int64           `json:"item_id" gorm:"index:idx_purchase_item_id"`
	Price            decimal.Decimal `json:"price" gorm:"type:decimal(65,0);"`
	CreatedAt        time.Time       `json:"created_at" gorm:"NOT NULL;type:TIMESTAMP;default:CURRENT_TIMESTAMP;<-:create"`
	UpdatedBscHeight int64           `json:"updated_bsc_height"`
}
