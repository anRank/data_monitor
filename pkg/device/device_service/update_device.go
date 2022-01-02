package device_service

import (
	"data_monitor/common"
	"data_monitor/dao"
	"data_monitor/pkg/device/device_dao"
	"data_monitor/pkg/device/device_model"
)

func UpdateDevice(m *device_model.Device) (innerErrCode common.InnerErrCode,errMsg string) {
	tx := dao.GetDB().Begin()
	defer tx.Rollback()

	if err := device_dao.UpdateDevice(tx, m); err != nil {
		return common.INNER_ERR_CODE_UNKNOWN_ERR, err.Error()
	}

	if err := tx.Commit().Error; err != nil{
		return common.INNER_ERR_CODE_UNKNOWN_ERR, err.Error()
	}
	return common.INNER_ERR_CODE_SUCCESS, ""
}