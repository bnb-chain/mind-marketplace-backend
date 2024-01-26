package database

type Category struct {
	Id   int64  `json:"id" gorm:"primaryKey;not null;autoIncrement:false"`
	Name string `json:"name" gorm:"not null;size:32"`

	//Items []*Item `json:"items" gorm:"foreignKey:CategoryId"`
}
