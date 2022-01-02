package device_dao

import "gorm.io/gorm"

func DeleteDevice(db *gorm.DB, id int64) error {
	m, err := GetDevice(db, id)
	if err != nil{
		return err
	}
	m.IsDeleted = 1
	return db.Save(m).Error
}