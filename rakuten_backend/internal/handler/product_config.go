package handler

import (
	"rakuten_backend/internal/api/response/code"
	"rakuten_backend/internal/api/response/web"
	"rakuten_backend/internal/context"
	"rakuten_backend/internal/service"
	"strconv"
)

func GetUserProductConfigList(c *context.Context) {
	productName := c.DefaultQuery("name", "")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	res, err := service.GetUserProductConfigList(productName, page, pageSize)
	if err != nil {
		c.Errorf("用户端商品列表查询失败: %v", err)
		web.Fail(c, code.InternalError)
		return
	}
	web.Success(c, res)
	return
}

func GetAdminProductConfigList(c *context.Context) {
	productName := c.DefaultQuery("name", "")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	res, err := service.GetUserProductConfigList(productName, page, pageSize)
	if err != nil {
		c.Errorf("用户端商品列表查询失败: %v", err)
		web.Fail(c, code.InternalError)
		return
	}
	web.Success(c, res)
	return
}
