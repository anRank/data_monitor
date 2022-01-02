package device_dao

import (
	"data_monitor/pkg/device/device_model"
	"gorm.io/gorm"
)

func UpdateDevice(db *gorm.DB, m *device_model.Device) error {
	var device device_model.Device
	db.Where("id = ?", m.Id).First(&device)
	device = *m
	if err := db.Save(device).Error; err != nil{
		return err
	}
	return nil
}