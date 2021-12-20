package user_dao

import (
	"data_monitor/pkg/user/user_model"
	"gorm.io/gorm"
)

type UserFinder struct {

}

func NewUserFinder() *UserFinder {
	return &UserFinder{}
}

func (f *UserFinder)DoFind(db *gorm.DB, userId int64) (*user_model.User, error) {
	user := &user_model.User{}
	if err := db.Find(user, userId).Error; err!=nil{
		return nil, err
	}
	return user, nil
}