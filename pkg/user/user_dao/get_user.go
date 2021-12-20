package user_dao

import (
	"data_monitor/pkg/user/user_model"
	"gorm.io/gorm"
)

func GetUserById(db *gorm.DB, id int) (user *user_model.User, err error) {
	query := db.Model(&user_model.User{})

	err = query.Where("id = ?", id).Find(&user).Error
	if err != nil{
		return nil, err
	}

	return user, nil
}

func GetUserByName(db *gorm.DB, name string)  (user *user_model.User, err error) {
	query := db.Model(&user_model.User{})

	err = query.Where("name = ?", name).Find(&user).Error
	if err != nil{
		return nil, err
	}

	return user, nil
}