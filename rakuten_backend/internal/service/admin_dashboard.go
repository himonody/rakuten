package service

import (
	"fmt"
	"rakuten_backend/internal/dao"
	"time"
)

func GetDashboard() (map[string]interface{}, error) {
	now := time.Now()
	today := now.Format("2006-01-02")
	year, week := now.ISOWeek()
	weekStr := fmt.Sprintf("%d-%02d", year, week)
	monthStr := now.Format("2006-01")

	dailyStats, err := dao.GetDailyStats(today)
	if err != nil {
		return nil, err
	}
	weeklyStats, err := dao.GetWeeklyStats(weekStr)
	if err != nil {
		return nil, err
	}
	monthlyStats, err := dao.GetMonthlyStats(monthStr)
	if err != nil {
		return nil, err
	}
	totalStats, err := dao.GetTotalStats()
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{
		"day":   dailyStats,
		"week":  weeklyStats,
		"month": monthlyStats,
		"total": totalStats,
	}, nil
}
