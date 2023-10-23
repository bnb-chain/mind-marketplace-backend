package database

import "time"

type Account struct {
	Id                int64     `json:"id" gorm:"primaryKey"`
	Address           string    `json:"address" gorm:"index:idx_account_address"`
	TwitterUserName   string    `json:"twitter"`
	TwitterVerified   bool      `json:"twitter_verified"`
	InstagramUserName string    `json:"instagram"`
	InstagramVerified bool      `json:"instagram_verified"`
	Bio               string    `json:"bio" gorm:"size:1024"`
	CreatedAt         time.Time `json:"created_at" gorm:"NOT NULL;type:TIMESTAMP;default:CURRENT_TIMESTAMP;<-:create"`
	UpdatedAt         time.Time `json:"updated_at" gorm:"NOT NULL;type:TIMESTAMP;default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP"`
}
