package models

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

func CreateGormDatabase(dbfile string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dbfile), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
