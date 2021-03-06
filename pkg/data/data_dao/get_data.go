package data_dao

import (
	"data_monitor/pkg/data/data_model"
	"gorm.io/gorm"
)

func GetDataList(db *gorm.DB) (data []data_model.Data, err error) {
	query := db.Model(&data_model.Data{})

	err = query.Find(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func GetDataByDeviceId(db *gorm.DB, deviceId int64) (data []data_model.Data, err error) {
	query := db.Model(&data_model.Data{})

	err = query.Where("device_id = ?", deviceId).Find(&data).Error
	if err != nil{
		return nil, err
	}
	return data, nil
}