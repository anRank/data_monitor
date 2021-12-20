package device_service

import (
	"data_monitor/dao"
	"data_monitor/pkg/device/device_dao"
	"data_monitor/pkg/device/device_model"
)

func GetDevice(id int) (device *device_model.Device, err error) {
	return device_dao.GetDevice(dao.GetDB(), id)
}
