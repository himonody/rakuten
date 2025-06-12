package handler

import (
	"rakuten_backend/internal/api/response/code"
	"rakuten_backend/internal/api/response/web"
	"rakuten_backend/internal/context"
	"rakuten_backend/internal/service"
)

func GetDashboard(c *context.Context) {
	dashboard, err := service.GetDashboard()
	if err != nil {
		c.Errorf("仪表盘查询错误: %v", err)
		web.Fail(c, code.InternalError)
		return
	}
	web.Success(c, dashboard)
	return
}
