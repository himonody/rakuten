package handler

import (
	"rakuten_backend/internal/api/request"
	"rakuten_backend/internal/api/response/code"
	"rakuten_backend/internal/api/response/web"
	"rakuten_backend/internal/context"
	"rakuten_backend/internal/service"
	"strconv"
)

func GetDailyStatistics(c *context.Context) {
	req := new(request.DailyStats)
	req.StatDate = c.Query("statDate")
	req.Page, _ = strconv.Atoi(c.DefaultQuery("page", "1"))
	req.PageSize, _ = strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	res, err := service.GetDailyStatsList(req)
	if err != nil {
		c.Errorf("日统计报表查询错误: %v", err)
		web.Fail(c, code.InternalError)
		return
	}
	web.Success(c, res)
	return
}

func GetWeeklyStatistics(c *context.Context) {
	req := new(request.WeeklyStats)
	req.StatWeek = c.Query("statWeek")
	req.Page, _ = strconv.Atoi(c.DefaultQuery("page", "1"))
	req.PageSize, _ = strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	res, err := service.GetWeeklyStatsList(req)
	if err != nil {
		c.Errorf("周统计报表查询错误: %v", err)
		web.Fail(c, code.InternalError)
		return
	}
	web.Success(c, res)
	return
}
func GetMonthlyStatistics(c *context.Context) {
	req := new(request.MonthlyStats)
	req.StatMonth = c.Query("statMonth")
	req.Page, _ = strconv.Atoi(c.DefaultQuery("page", "1"))
	req.PageSize, _ = strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	res, err := service.GetMonthlyStatsList(req)
	if err != nil {
		c.Errorf("月统计报表查询错误: %v", err)
		web.Fail(c, code.InternalError)
		return
	}
	web.Success(c, res)
	return
}
