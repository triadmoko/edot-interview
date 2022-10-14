package configs

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func DB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("sqlite.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	
	return db, nil
}
