package user_service

import (
	"data_monitor/common"
	"data_monitor/dao"
	"data_monitor/pkg/user/user_dao"
)

func DeleteUser(id int) (innerErrCode common.InnerErrCode,errMsg string)  {
	tx := dao.GetDB().Begin()
	defer tx.Rollback()

	if err := user_dao.DeleteUser(tx, id); err !=nil{
		return common.INNER_ERR_CODE_UNKNOWN_ERR, err.Error()
	}

	if err := tx.Commit().Error; err != nil{
		return common.INNER_ERR_CODE_UNKNOWN_ERR, err.Error()
	}
	return common.INNER_ERR_CODE_SUCCESS, ""
}