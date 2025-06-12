package dao

import (
	"errors"
	"gorm.io/gorm"
	"rakuten_backend/internal/db"
	"rakuten_backend/internal/model"
)

func GetAdminUserByUserName(username string) (user *model.AdminUser, err error) {
	user = new(model.AdminUser)
	err = db.MySQL.Model("admin_user").
		Where("username = ?", username).
		First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return user, nil
}
