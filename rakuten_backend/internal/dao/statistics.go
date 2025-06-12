package dao

import (
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"rakuten_backend/internal/api/request"
	"rakuten_backend/internal/db"
	"rakuten_backend/internal/model"
)

// GetTotalStats 总统计表
func GetTotalStats() (stats *model.TotalStats, err error) {
	if err = db.MySQL.Model("total_stats").Where("id=1").First(&stats).Error; err != nil {
		return nil, err
	}
	return stats, nil
}

// GetDailyStats 日统计表
func GetDailyStats(daily string) (dailyStats *model.DailyStats, err error) {
	res := db.MySQL.Model("daily_stats")
	if daily != "" {
		res.Where("stat_date = ?", daily)
	}

	if err = res.First(&dailyStats).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return dailyStats, nil
}

// GetWeeklyStats 周统计表
func GetWeeklyStats(weekly string) (weeklyStats *model.WeeklyStats, err error) {

	res := db.MySQL.Model("weekly_stats")
	if weekly != "" {
		res.Where("stat_week = ?", weekly)
	}
	if err = res.First(&weeklyStats).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return weeklyStats, nil
}

// GetMonthlyStats 月统计表
func GetMonthlyStats(monthly string) (monthlyStats *model.MonthlyStats, err error) {

	res := db.MySQL.Model("monthly_stats")
	if monthly != "" {
		res.Where("stat_week = ?", monthly)
	}
	if err = res.First(&monthlyStats).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return monthlyStats, nil
}

// GetDailyStatsList 日统计表
func GetDailyStatsList(req *request.DailyStats) (dailyStatsList []*model.DailyStats, count int64, err error) {
	dailyStatsList = make([]*model.DailyStats, 0)
	res := db.MySQL.Model("daily_stats")
	if req.StatDate != "" {
		res.Where("stat_date = ?", req.StatDate)
	}
	if err = res.Count(&count).Error; err != nil || count == 0 {
		return nil, 0, nil
	}
	res.Order("id desc")
	if err = res.Offset((req.Page - 1) * req.PageSize).Limit(req.PageSize).Find(&dailyStatsList).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, 0, nil
		}
		return nil, 0, err
	}
	return dailyStatsList, count, nil
}

// GetWeeklyStatsList 周统计表
func GetWeeklyStatsList(req *request.WeeklyStats) (weeklyStatsList []*model.WeeklyStats, count int64, err error) {
	weeklyStatsList = make([]*model.WeeklyStats, 0)
	res := db.MySQL.Model("weekly_stats")
	if req.StatWeek != "" {
		res.Where("stat_week = ?", req.StatWeek)
	}
	if err = res.Count(&count).Error; err != nil || count == 0 {
		return nil, 0, nil
	}
	res.Order("id desc")
	if err = res.Offset((req.Page - 1) * req.PageSize).Limit(req.PageSize).Find(&weeklyStatsList).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, 0, nil
		}
		return nil, 0, err
	}
	return weeklyStatsList, count, nil
}

// GetMonthlyStatsList 月统计表
func GetMonthlyStatsList(req *request.MonthlyStats) (monthlyStatsList []*model.MonthlyStats, count int64, err error) {
	monthlyStatsList = make([]*model.MonthlyStats, 0)
	res := db.MySQL.Model("monthly_stats")
	if req.StatMonth != "" {
		res.Where("stat_week = ?", req.StatMonth)
	}
	if err = res.Count(&count).Error; err != nil || count == 0 {
		return nil, 0, nil
	}
	res.Order("id desc")
	if err = res.Offset((req.Page - 1) * req.PageSize).Limit(req.PageSize).Find(&monthlyStatsList).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, 0, nil
		}
		return nil, 0, err
	}
	return monthlyStatsList, count, nil
}
func UpdateOrInsertTransactionStats(req *request.Stats) error {
	return db.MySQL.Transaction(func(tx *gorm.DB) error {
		//更新总统计表
		totalStats := &model.TotalStats{
			ID:                   1,
			TotalAdminCount:      req.NewAgentCount,
			TotalMemberCount:     req.NewMemberCount,
			TotalAgentCount:      req.NewAgentCount,
			RechargeAmount:       req.RechargeAmount,
			WithdrawAmount:       req.WithdrawAmount,
			BonusAmount:          req.BonusAmount,
			RechargeWithdrawDiff: req.RechargeWithdrawDiff,
			RebateAmount:         req.RebateAmount,
			FlashOrderCount:      req.FlashOrderCount,
			FlashOrderDoneCount:  req.FlashOrderDoneCount,
			GroupOrderCount:      req.GroupOrderCount,
			GroupOrderDoneCount:  req.GroupOrderDoneCount,
		}

		if err := tx.Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "id"}},
			DoUpdates: clause.Assignments(map[string]interface{}{
				"total_admin_count":      gorm.Expr("total_admin_count + ?", totalStats.TotalAdminCount),
				"total_member_count":     gorm.Expr("total_member_count + ?", totalStats.TotalMemberCount),
				"total_agent_count":      gorm.Expr("total_agent_count + ?", totalStats.TotalAgentCount),
				"recharge_amount":        gorm.Expr("recharge_amount + ?", totalStats.RechargeAmount),
				"withdraw_amount":        gorm.Expr("withdraw_amount + ?", totalStats.WithdrawAmount),
				"bonus_amount":           gorm.Expr("bonus_amount + ?", totalStats.BonusAmount),
				"recharge_withdraw_diff": gorm.Expr("recharge_withdraw_diff + ?", totalStats.RechargeWithdrawDiff),
				"rebate_amount":          gorm.Expr("rebate_amount + ?", totalStats.RebateAmount),
				"flash_order_count":      gorm.Expr("flash_order_count + ?", totalStats.FlashOrderCount),
				"flash_order_done_count": gorm.Expr("flash_order_done_count + ?", totalStats.FlashOrderDoneCount),
				"group_order_count":      gorm.Expr("group_order_count + ?", totalStats.GroupOrderCount),
				"group_order_done_count": gorm.Expr("group_order_done_count + ?", totalStats.GroupOrderDoneCount),
			}),
		}).Create(totalStats).Error; err != nil {
			return err
		}
		//更新日总统计表
		dailyStats := &model.DailyStats{
			StatDate:             req.StatDate,
			NewAdminCount:        req.NewAdminCount,
			NewMemberCount:       req.NewMemberCount,
			NewAgentCount:        req.NewAgentCount,
			RechargeAmount:       req.RechargeAmount,
			WithdrawAmount:       req.WithdrawAmount,
			BonusAmount:          req.BonusAmount,
			RechargeWithdrawDiff: req.RechargeWithdrawDiff,
			RebateAmount:         req.RebateAmount,
			FlashOrderCount:      req.FlashOrderCount,
			FlashOrderDoneCount:  req.FlashOrderDoneCount,
			GroupOrderCount:      req.GroupOrderCount,
			GroupOrderDoneCount:  req.GroupOrderDoneCount,
		}
		if err := tx.Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "stat_date"}},
			DoUpdates: clause.Assignments(map[string]interface{}{
				"new_admin_count":        gorm.Expr("new_admin_count + ?", dailyStats.NewAdminCount),
				"new_member_count":       gorm.Expr("new_member_count + ?", dailyStats.NewMemberCount),
				"new_agent_count":        gorm.Expr("new_agent_count + ?", dailyStats.NewAgentCount),
				"recharge_amount":        gorm.Expr("recharge_amount + ?", dailyStats.RechargeAmount),
				"withdraw_amount":        gorm.Expr("withdraw_amount + ?", dailyStats.WithdrawAmount),
				"bonus_amount":           gorm.Expr("bonus_amount + ?", dailyStats.BonusAmount),
				"recharge_withdraw_diff": gorm.Expr("recharge_withdraw_diff + ?", dailyStats.RechargeWithdrawDiff),
				"rebate_amount":          gorm.Expr("rebate_amount + ?", dailyStats.RebateAmount),
				"flash_order_count":      gorm.Expr("flash_order_count + ?", dailyStats.FlashOrderCount),
				"flash_order_done_count": gorm.Expr("flash_order_done_count + ?", dailyStats.FlashOrderDoneCount),
				"group_order_count":      gorm.Expr("group_order_count + ?", dailyStats.GroupOrderCount),
				"group_order_done_count": gorm.Expr("group_order_done_count + ?", dailyStats.GroupOrderDoneCount),
			}),
		}).Create(dailyStats).Error; err != nil {
			return err
		}
		//更新周总统计表
		weeklyStats := &model.WeeklyStats{
			StatWeek:             req.StatWeek,
			NewAdminCount:        req.NewAdminCount,
			NewMemberCount:       req.NewMemberCount,
			NewAgentCount:        req.NewAgentCount,
			RechargeAmount:       req.RechargeAmount,
			WithdrawAmount:       req.WithdrawAmount,
			BonusAmount:          req.BonusAmount,
			RechargeWithdrawDiff: req.RechargeWithdrawDiff,
			RebateAmount:         req.RebateAmount,
			FlashOrderCount:      req.FlashOrderCount,
			FlashOrderDoneCount:  req.FlashOrderDoneCount,
			GroupOrderCount:      req.GroupOrderCount,
			GroupOrderDoneCount:  req.GroupOrderDoneCount,
		}
		if err := tx.Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "stat_week"}},
			DoUpdates: clause.Assignments(map[string]interface{}{
				"new_admin_count":        gorm.Expr("new_admin_count + ?", weeklyStats.NewAdminCount),
				"new_member_count":       gorm.Expr("new_member_count + ?", weeklyStats.NewMemberCount),
				"new_agent_count":        gorm.Expr("new_agent_count + ?", weeklyStats.NewAgentCount),
				"recharge_amount":        gorm.Expr("recharge_amount + ?", weeklyStats.RechargeAmount),
				"withdraw_amount":        gorm.Expr("withdraw_amount + ?", weeklyStats.WithdrawAmount),
				"bonus_amount":           gorm.Expr("bonus_amount + ?", weeklyStats.BonusAmount),
				"recharge_withdraw_diff": gorm.Expr("recharge_withdraw_diff + ?", weeklyStats.RechargeWithdrawDiff),
				"rebate_amount":          gorm.Expr("rebate_amount + ?", weeklyStats.RebateAmount),
				"flash_order_count":      gorm.Expr("flash_order_count + ?", weeklyStats.FlashOrderCount),
				"flash_order_done_count": gorm.Expr("flash_order_done_count + ?", weeklyStats.FlashOrderDoneCount),
				"group_order_count":      gorm.Expr("group_order_count + ?", weeklyStats.GroupOrderCount),
				"group_order_done_count": gorm.Expr("group_order_done_count + ?", weeklyStats.GroupOrderDoneCount),
			}),
		}).Create(weeklyStats).Error; err != nil {
			return err
		}
		//更新月总统计表
		monthlyStats := &model.MonthlyStats{
			StatMonth:            req.StatMonth,
			NewAdminCount:        req.NewAdminCount,
			NewMemberCount:       req.NewMemberCount,
			NewAgentCount:        req.NewAgentCount,
			RechargeAmount:       req.RechargeAmount,
			WithdrawAmount:       req.WithdrawAmount,
			BonusAmount:          req.BonusAmount,
			RechargeWithdrawDiff: req.RechargeWithdrawDiff,
			RebateAmount:         req.RebateAmount,
			FlashOrderCount:      req.FlashOrderCount,
			FlashOrderDoneCount:  req.FlashOrderDoneCount,
			GroupOrderCount:      req.GroupOrderCount,
			GroupOrderDoneCount:  req.GroupOrderDoneCount,
		}
		if err := tx.Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "stat_month"}},
			DoUpdates: clause.Assignments(map[string]interface{}{
				"new_admin_count":        gorm.Expr("new_admin_count + ?", monthlyStats.NewAdminCount),
				"new_member_count":       gorm.Expr("new_member_count + ?", monthlyStats.NewMemberCount),
				"new_agent_count":        gorm.Expr("new_agent_count + ?", monthlyStats.NewAgentCount),
				"recharge_amount":        gorm.Expr("recharge_amount + ?", monthlyStats.RechargeAmount),
				"withdraw_amount":        gorm.Expr("withdraw_amount + ?", monthlyStats.WithdrawAmount),
				"bonus_amount":           gorm.Expr("bonus_amount + ?", monthlyStats.BonusAmount),
				"recharge_withdraw_diff": gorm.Expr("recharge_withdraw_diff + ?", monthlyStats.RechargeWithdrawDiff),
				"rebate_amount":          gorm.Expr("rebate_amount + ?", monthlyStats.RebateAmount),
				"flash_order_count":      gorm.Expr("flash_order_count + ?", monthlyStats.FlashOrderCount),
				"flash_order_done_count": gorm.Expr("flash_order_done_count + ?", monthlyStats.FlashOrderDoneCount),
				"group_order_count":      gorm.Expr("group_order_count + ?", monthlyStats.GroupOrderCount),
				"group_order_done_count": gorm.Expr("group_order_done_count + ?", monthlyStats.GroupOrderDoneCount),
			}),
		}).Create(monthlyStats).Error; err != nil {
			return err
		}
		return nil
	})
}
