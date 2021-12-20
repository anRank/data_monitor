package dao

import (
	"data_monitor/config"
	"gorm.io/gorm"
)

func GetDB() *gorm.DB {
	return config.DB
}
