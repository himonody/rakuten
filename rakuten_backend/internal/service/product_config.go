package service

import (
	"rakuten_backend/internal/api/response/webres"
	"rakuten_backend/internal/dao"
)

func GetUserProductConfigList(name string, page, pageSize int) (map[string]interface{}, error) {
	productList, count, err := dao.GetUserProductConfigList(name, page, pageSize)
	if err != nil {
		return nil, err
	}
	return webres.DataRsp(productList, count, page, pageSize), err
}
