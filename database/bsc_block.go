package database

import "time"

type BscBlock struct {
	Id        int64     `json:"id" gorm:"primaryKey"`
	Height    uint64    `gorm:"NOT NULL;index:idx_bsc_block_height"`
	CreatedAt time.Time `json:"created_at" gorm:"NOT NULL;type:TIMESTAMP;default:CURRENT_TIMESTAMP;<-:create"`
}
