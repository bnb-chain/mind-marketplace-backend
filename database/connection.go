package database

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/bnb-chain/greenfield-data-marketplace-backend/util"
)

func Connect() (*gorm.DB, error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/marketplace?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db.Debug(), err
}

func ConnectDBWithConfig(config *util.DBConfig) (*gorm.DB, error) {
	if config.DBDialect == "sqlite3" {
		db, err := gorm.Open(sqlite.Open(config.DBPath), &gorm.Config{})
		return db.Debug(), err
	} else if config.DBDialect == "mysql" {
		userPass := os.Getenv("MYSQL_USER_PASS") //username and password, e.g., root:123456
		db, err := gorm.Open(mysql.Open(userPass+config.DBPath), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		if err = db.AutoMigrate(&Account{}); err != nil {
			panic(err)
		}
		if err = db.AutoMigrate(&Item{}); err != nil {
			panic(err)
		}
		if err = db.AutoMigrate(&ItemStats{}); err != nil {
			panic(err)
		}
		if err = db.AutoMigrate(&Purchase{}); err != nil {
			panic(err)
		}
		if err = db.AutoMigrate(&BscBlock{}); err != nil {
			panic(err)
		}
		if err = db.AutoMigrate(&GnfdBlock{}); err != nil {
			panic(err)
		}
		return db.Debug(), nil
	} else {
		return nil, fmt.Errorf("dialect %s not supported", config.DBDialect)
	}
}
