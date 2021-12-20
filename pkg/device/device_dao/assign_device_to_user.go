package device_dao

import (
	"data_monitor/pkg/device/device_model"
	"gorm.io/gorm"
)

func AssignDevice2User(db *gorm.DB, m *device_model.DeviceUser) error {
	if err := db.Create(m).Error; err != nil{
		return err
	}
	return db.Save(m).Error
}
