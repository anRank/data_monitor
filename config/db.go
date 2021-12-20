package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	DB_HOSTNAME = "127.0.0.1"
	DB_DATABASE_NAME = "data_monitor"
	DB_PORT = "3306"
	DB_USERNAME = "root"
	DB_PWD = "000220"

)

var db *gorm.DB

func InitDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DB_USERNAME, DB_PWD, DB_HOSTNAME, DB_PORT, DB_DATABASE_NAME)

	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

}

func GetDB() *gorm.DB{
	return db
}
