package user_dao

import "gorm.io/gorm"

func DeleteUser(db *gorm.DB, id int) error {
	m, err := GetUserById(db, id)
	if err != nil{
		return err
	}
	m.IsDeleted = 1
	return db.Save(m).Error
}