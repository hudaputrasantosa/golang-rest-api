package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func getConnection() (*gorm.DB, error) {
	dsn := "root:password@tcp(localhost:3306)/database?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil

}
