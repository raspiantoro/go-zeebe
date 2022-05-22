package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDB(dbUrl string) (db *gorm.DB, err error) {
	return gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
}
