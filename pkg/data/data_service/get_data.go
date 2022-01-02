package data_service

import (
	"data_monitor/dao"
	"data_monitor/pkg/data/data_dao"
	"data_monitor/pkg/data/data_model"
)

func GetDataList() (data []data_model.Data, err error) {
	return data_dao.GetDataList(dao.GetDB())
}

func GetDataByDeviceId(deviceId int64) (data []data_model.Data, err error) {
	return data_dao.GetDataByDeviceId(dao.GetDB(), deviceId)
}
