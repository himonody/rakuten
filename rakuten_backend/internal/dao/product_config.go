package dao

import (
	"rakuten_backend/internal/db"
	"rakuten_backend/internal/model"
)

func GetAdminProductConfigList(name string, page, pageSize int) (productList []*model.ProductConfig, count int64, err error) {
	productList = make([]*model.ProductConfig, 0)
	res := db.MySQL.Model("product_config").
		Select("id,product_code,name,price,image_url,status")
	if name != "" {
		res = res.Where("like %?%", name)
	}
	if res.Count(&count) == nil {
		return productList, 0, err
	}
	if err = res.Offset((page - 1) * pageSize).Limit(pageSize).Find(&productList).Error; err != nil {
		return productList, 0, err
	}
	return productList, count, nil
}

func GetUserProductConfigList(name string, page, pageSize int) (productList []*model.ProductConfig, count int64, err error) {
	productList = make([]*model.ProductConfig, 0)
	res := db.MySQL.Model("product_config").
		Select("id,product_code,name,price,image_url").
		Where("status=0")
	if name != "" {
		res = res.Where("like %?%", name)
	}
	if res.Count(&count) == nil {
		return productList, 0, err
	}
	if err = res.Offset((page - 1) * pageSize).Limit(pageSize).Find(&productList).Error; err != nil {
		return productList, 0, err
	}
	return productList, count, nil
}
