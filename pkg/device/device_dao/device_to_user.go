package device_dao

import (
	"data_monitor/dao"
	"data_monitor/pkg/device/device_model"
	"gorm.io/gorm"
)

func AssignDevice2User(db *gorm.DB, m *device_model.DeviceUser) error {
	if err := db.Create(m).Error; err != nil{
		return err
	}
	return db.Save(m).Error
}

func DeleteDevice2User(db *gorm.DB, m *device_model.DeviceUser) error {
	m.IsDeleted = 1
	return db.Save(m).Error
}

func Get2ById(userID int64, deviceID int64) (device *device_model.DeviceUser,err error) {
	db := dao.GetDB()
	query := db.Model(&device_model.DeviceUser{})

	err = query.Where("user_id = ? and device_id = ?", userID, deviceID).Find(&device).Error
	if err != nil{
		return nil, err
	}

	return device, nil
}