package device_dao

import (
	"data_monitor/pkg/device/device_model"
	"gorm.io/gorm"
)

func CreateDevice(db *gorm.DB, m *device_model.Device) error {
	if err := db.Create(m).Error; err!= nil{
		return err
	}
	return db.Save(m).Error
}