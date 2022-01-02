package device_service

import (
	"data_monitor/common"
	"data_monitor/dao"
	"data_monitor/pkg/device/device_dao"
	"data_monitor/pkg/device/device_model"
)

func AssignDevice2User(m *device_model.DeviceUser) (innerErrCode common.InnerErrCode, errMsg string) {
	tx := dao.GetDB().Begin()
	defer tx.Rollback()

	if err := device_dao.AssignDevice2User(tx, m); err != nil{
		return common.INNER_ERR_CODE_UNKNOWN_ERR, err.Error()
	}

	if err := tx.Commit().Error; err != nil{
		return common.INNER_ERR_CODE_UNKNOWN_ERR, err.Error()
	}
	return common.INNER_ERR_CODE_SUCCESS, ""
}

func DeleteDevice2User(m *device_model.DeviceUser) (innerErrCode common.InnerErrCode, errMsg string) {
	tx := dao.GetDB().Begin()
	defer tx.Rollback()

	if err := device_dao.DeleteDevice2User(tx, m); err != nil{
		return common.INNER_ERR_CODE_UNKNOWN_ERR, err.Error()
	}

	if err := tx.Commit().Error; err != nil{
		return common.INNER_ERR_CODE_UNKNOWN_ERR, err.Error()
	}
	return common.INNER_ERR_CODE_SUCCESS, ""
}

func IsRelationExisted(userId int64, deviceId int64) (bool, error) {
	r, err := device_dao.Get2ById(userId, deviceId)
	if err != nil{
		return false, err
	}
	if r == nil {
		return false, nil
	}
	return true, nil
}