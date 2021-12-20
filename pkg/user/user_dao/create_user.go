package user_dao

import (
	"data_monitor/pkg/user/user_model"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, m *user_model.User) error {
	if err := db.Create(m).Error; err != nil{
		return err
	}
	return db.Save(m).Error
}