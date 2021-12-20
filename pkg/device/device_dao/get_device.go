package device_dao

import (
	"data_monitor/pkg/device/device_model"
	"gorm.io/gorm"
)

func GetDevice(db *gorm.DB, id int) (device *device_model.Device, err error) {
	query := db.Model(&device_model.Device{})

	err = query.Where("id = ?", id).Find(&device).Error
	if err != nil{
		return nil, err
	}

	return device, nil
}