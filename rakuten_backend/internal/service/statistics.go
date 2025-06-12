package service

import (
	"rakuten_backend/internal/api/request"
	"rakuten_backend/internal/api/response/webres"
	"rakuten_backend/internal/dao"
)

func GetDailyStatsList(req *request.DailyStats) (map[string]interface{}, error) {
	dailyStatsList, count, err := dao.GetDailyStatsList(req)
	if err != nil {
		return nil, err
	}
	return webres.DataRsp(dailyStatsList, count, req.Page, req.PageSize), nil
}

func GetWeeklyStatsList(req *request.WeeklyStats) (map[string]interface{}, error) {
	weeklyStatsList, count, err := dao.GetWeeklyStatsList(req)
	if err != nil {
		return nil, err
	}
	return webres.DataRsp(weeklyStatsList, count, req.Page, req.PageSize), nil
}
func GetMonthlyStatsList(req *request.MonthlyStats) (map[string]interface{}, error) {
	monthlyStatsList, count, err := dao.GetMonthlyStatsList(req)
	if err != nil {
		return nil, err
	}
	return webres.DataRsp(monthlyStatsList, count, req.Page, req.PageSize), nil
}

func UpdateOrInsertTransactionStats(req *request.Stats) error {
	defer func() {
		if err := recover(); err != nil {
			return
		}
	}()
	err := dao.UpdateOrInsertTransactionStats(req)
	if err != nil {
		return err
	}
	return nil
}
