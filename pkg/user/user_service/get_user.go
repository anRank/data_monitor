package user_service

import (
	"data_monitor/dao"
	"data_monitor/pkg/user/user_dao"
	"data_monitor/pkg/user/user_model"
)

func GetUserById(id int64) (user *user_model.User, err error) {
	return user_dao.GetUserById(dao.GetDB(), id)
}