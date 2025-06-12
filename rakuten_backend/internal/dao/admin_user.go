package dao

import (
	"errors"
	"gorm.io/gorm"
	"rakuten_backend/internal/db"
	"rakuten_backend/internal/model"
)

func InsertAdminUser(admin *model.AdminUser) (err error) {
	res := db.MySQL.Model("admin_user").
		Create(admin)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("创建用户失败")
	}
	return nil
}

func GetAdminUser(username string, id, role, page, pageSize, isAgent int) (admins []*model.AdminUser, count int64, err error) {
	admins = make([]*model.AdminUser, 0)
	res := db.MySQL.Model("admin_user")
	res = res.Where("is_agent = ? ", isAgent)
	if id > 0 {
		res = res.Where("id = ? ", id)
	}
	if username != "" {
		res = res.Where("username = ? ", username)
	}
	if role > -1 {
		res = res.Where("role = ? ", role)
	}

	if err = res.Count(&count).Error; err != nil {
		return nil, 0, err
	}
	res.Order("id desc")
	if err = res.Offset((page - 1) * pageSize).Limit(pageSize).Find(&admins).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return admins, 0, nil
		}
		return nil, 0, err
	}
	return admins, count, nil
}

func EditAdminUser(id int, adminUserMap map[string]interface{}) (err error) {
	res := db.MySQL.Model("admin_user").Where("id = ? ", id).Updates(adminUserMap)
	if res.RowsAffected == 0 {
		return errors.New("更新失败")
	}
	if res.Error != nil {
		return res.Error
	}
	return nil

}

func DeleteAdminUser(id int) (err error) {
	res := db.MySQL.Model("admin_user").Delete(&model.AdminUser{}, id)
	if res.RowsAffected == 0 {
		return errors.New("更新失败")
	}
	if res.Error != nil {
		return res.Error
	}
	return nil

}
