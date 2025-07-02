package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnenction() (*gorm.DB, error) {

	// user := "root"
	// pass := "mysql"
	// ip := "127.0.0.1"
	// port := "3306"

	// punya neva
	user := "neva"
	pass := "252525"
	ip := "127.0.0.1"
	port := "3306"
	dbname := "db_ytl"

	// user := "haydar"
	// pass := "252525"
	// ip := "127.0.0.1"
	// port := "3306"
	// dbname := "belajarapi"

	dsn := user + ":" + pass + "@tcp(" + ip + ":" + port + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	DB = db
	return db, nil
}
