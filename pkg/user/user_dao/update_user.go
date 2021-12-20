package user_dao

import (
	"data_monitor/pkg/user/user_model"
	"gorm.io/gorm"
)

func UpdateUser(db *gorm.DB, m *user_model.User) error {
	var user user_model.User
	db.Where("id = ?", m.Id).First(&user)
	user = *m
	if err := db.Save(user).Error; err != nil{
		return err
	}
	return nil
}