package user_service

import (
	"data_monitor/common"
	"data_monitor/dao"
	"data_monitor/pkg/user/user_dao"
	"data_monitor/pkg/user/user_model"
)

func CreateUser(m *user_model.User) (innerErrCode common.InnerErrCode, errMsg string) {
	tx := dao.GetDB().Begin()
	defer tx.Rollback()

	user, err := user_dao.GetUserByName(tx, m.Name)
	if err != nil{
		return common.INNER_ERR_CODE_UNKNOWN_ERR, err.Error()
	}
	if user.Id != 0 {
		return common.INNER_ERR_CODE_UNKNOWN_ERR, "user already existed"
	}

	if err = user_dao.CreateUser(tx, m); err !=nil{
		return common.INNER_ERR_CODE_UNKNOWN_ERR, err.Error()
	}

	if err = tx.Commit().Error; err != nil{
		return common.INNER_ERR_CODE_UNKNOWN_ERR, err.Error()
	}
	return common.INNER_ERR_CODE_SUCCESS, ""
}