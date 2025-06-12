package model

import (
	"github.com/shopspring/decimal"
	"time"
)

// TotalStats 总统计表
type TotalStats struct {
	ID                   int             `gorm:"primaryKey;column:id" json:"id"`                                                                         // 主键ID
	TotalAdminCount      int             `gorm:"column:total_admin_count;not null;default:0" json:"total_admin_count"`                                   // 总管理员人数
	TotalMemberCount     int             `gorm:"column:total_member_count;not null;default:0" json:"total_member_count"`                                 // 总会员人数
	TotalAgentCount      int             `gorm:"column:total_agent_count;not null;default:0" json:"total_agent_count"`                                   // 总代理人数
	RechargeAmount       decimal.Decimal `gorm:"column:recharge_amount;type:decimal(18,6);not null;default:0.0000" json:"recharge_amount"`               // 累计充值金额
	WithdrawAmount       decimal.Decimal `gorm:"column:withdraw_amount;type:decimal(18,6);not null;default:0.0000" json:"withdraw_amount"`               // 累计提款金额
	BonusAmount          decimal.Decimal `gorm:"column:bonus_amount;type:decimal(18,6);not null;default:0.0000" json:"bonus_amount"`                     // 彩金赠送金额
	RechargeWithdrawDiff decimal.Decimal `gorm:"column:recharge_withdraw_diff;type:decimal(18,6);not null;default:0.0000" json:"recharge_withdraw_diff"` // 充值提款差值（充值 - 提款）
	RebateAmount         decimal.Decimal `gorm:"column:rebate_amount;type:decimal(18,6);not null;default:0.0000" json:"rebate_amount"`                   // 累计总返利
	FlashOrderCount      int             `gorm:"column:flash_order_count;not null;default:0" json:"flash_order_count"`                                   // 总抢购订单数
	FlashOrderDoneCount  int             `gorm:"column:flash_order_done_count;not null;default:0" json:"flash_order_done_count"`                         // 总抢购订单完成数
	GroupOrderCount      int             `gorm:"column:group_order_count;not null;default:0" json:"group_order_count"`                                   // 总拼团订单数
	GroupOrderDoneCount  int             `gorm:"column:group_order_done_count;not null;default:0" json:"group_order_done_count"`                         // 总拼团订单完成数
	CreateAt             time.Time       `gorm:"column:create_at;autoCreateTime" json:"create_at"`                                                       // 创建时间
	UpdateAt             time.Time       `gorm:"column:update_at;autoUpdateTime" json:"update_at"`                                                       // 更新时间
}

// DailyStats 日统计表
type DailyStats struct {
	ID                   int             `gorm:"primaryKey;column:id;autoIncrement" json:"id"`                                                  // 主键ID
	StatDate             string          `gorm:"column:stat_date;type:date;uniqueIndex" json:"stat_date"`                                       // 统计日期
	NewAdminCount        int             `gorm:"column:new_admin_count;default:0" json:"new_admin_count"`                                       // 日新增管理员人数
	NewMemberCount       int             `gorm:"column:new_member_count;default:0" json:"new_member_count"`                                     // 日新增会员人数
	NewAgentCount        int             `gorm:"column:new_agent_count;default:0" json:"new_agent_count"`                                       // 日新增代理人数
	RechargeAmount       decimal.Decimal `gorm:"column:recharge_amount;type:decimal(18,6);default:0.0000" json:"recharge_amount"`               // 日累计充值金额
	WithdrawAmount       decimal.Decimal `gorm:"column:withdraw_amount;type:decimal(18,6);default:0.0000" json:"withdraw_amount"`               // 日累计提款金额
	BonusAmount          decimal.Decimal `gorm:"column:bonus_amount;type:decimal(18,6);default:0.0000" json:"bonus_amount"`                     // 彩金赠送金额
	RechargeWithdrawDiff decimal.Decimal `gorm:"column:recharge_withdraw_diff;type:decimal(18,6);default:0.0000" json:"recharge_withdraw_diff"` // 充值提款差值
	RebateAmount         decimal.Decimal `gorm:"column:rebate_amount;type:decimal(18,6);default:0.0000" json:"rebate_amount"`                   // 日累计总返利
	FlashOrderCount      int             `gorm:"column:flash_order_count;default:0" json:"flash_order_count"`                                   // 日抢购订单数
	FlashOrderDoneCount  int             `gorm:"column:flash_order_done_count;default:0" json:"flash_order_done_count"`                         // 日抢购订单完成数
	GroupOrderCount      int             `gorm:"column:group_order_count;default:0" json:"group_order_count"`                                   // 日拼团订单数
	GroupOrderDoneCount  int             `gorm:"column:group_order_done_count;default:0" json:"group_order_done_count"`                         // 日拼团订单完成数
	CreatedAt            time.Time       `gorm:"column:create_at;autoCreateTime" json:"create_at"`                                              // 创建时间
	UpdatedAt            time.Time       `gorm:"column:update_at;autoUpdateTime" json:"update_at"`                                              // 更新时间
}

// WeeklyStats 周统计表
type WeeklyStats struct {
	ID                   int             `gorm:"primaryKey;column:id;autoIncrement" json:"id"`                                                  // 主键ID
	StatWeek             string          `gorm:"column:stat_week;size:8;uniqueIndex" json:"stat_week"`                                          // 统计周（格式：YYYY-WW）
	NewAdminCount        int             `gorm:"column:new_admin_count;default:0" json:"new_admin_count"`                                       // 周新增管理员人数
	NewMemberCount       int             `gorm:"column:new_member_count;default:0" json:"new_member_count"`                                     // 周新增会员人数
	NewAgentCount        int             `gorm:"column:new_agent_count;default:0" json:"new_agent_count"`                                       // 周新增代理人数
	RechargeAmount       decimal.Decimal `gorm:"column:recharge_amount;type:decimal(18,6);default:0.0000" json:"recharge_amount"`               // 周累计充值金额
	WithdrawAmount       decimal.Decimal `gorm:"column:withdraw_amount;type:decimal(18,6);default:0.0000" json:"withdraw_amount"`               // 周累计提款金额
	BonusAmount          decimal.Decimal `gorm:"column:bonus_amount;type:decimal(18,6);default:0.0000" json:"bonus_amount"`                     // 彩金赠送金额
	RechargeWithdrawDiff decimal.Decimal `gorm:"column:recharge_withdraw_diff;type:decimal(18,6);default:0.0000" json:"recharge_withdraw_diff"` // 充值提款差值
	RebateAmount         decimal.Decimal `gorm:"column:rebate_amount;type:decimal(18,6);default:0.0000" json:"rebate_amount"`                   // 周累计总返利
	FlashOrderCount      int             `gorm:"column:flash_order_count;default:0" json:"flash_order_count"`                                   // 周抢购订单数
	FlashOrderDoneCount  int             `gorm:"column:flash_order_done_count;default:0" json:"flash_order_done_count"`                         // 周抢购订单完成数
	GroupOrderCount      int             `gorm:"column:group_order_count;default:0" json:"group_order_count"`                                   // 周拼团订单数
	GroupOrderDoneCount  int             `gorm:"column:group_order_done_count;default:0" json:"group_order_done_count"`                         // 周拼团订单完成数
	CreatedAt            time.Time       `gorm:"column:create_at;autoCreateTime" json:"create_at"`                                              // 创建时间
	UpdatedAt            time.Time       `gorm:"column:update_at;autoUpdateTime" json:"update_at"`                                              // 更新时间
}

// MonthlyStats 月统计表
type MonthlyStats struct {
	ID                   int             `gorm:"primaryKey;column:id;autoIncrement" json:"id"`                                                  // 主键ID
	StatMonth            string          `gorm:"column:stat_month;size:7;uniqueIndex" json:"stat_month"`                                        // 统计月份（格式：YYYY-MM）
	NewAdminCount        int             `gorm:"column:new_admin_count;default:0" json:"new_admin_count"`                                       // 月新增管理员人数
	NewMemberCount       int             `gorm:"column:new_member_count;default:0" json:"new_member_count"`                                     // 月新增会员人数
	NewAgentCount        int             `gorm:"column:new_agent_count;default:0" json:"new_agent_count"`                                       // 月新增代理人数
	RechargeAmount       decimal.Decimal `gorm:"column:recharge_amount;type:decimal(18,6);default:0.0000" json:"recharge_amount"`               // 月累计充值金额
	WithdrawAmount       decimal.Decimal `gorm:"column:withdraw_amount;type:decimal(18,6);default:0.0000" json:"withdraw_amount"`               // 月累计提款金额
	BonusAmount          decimal.Decimal `gorm:"column:bonus_amount;type:decimal(18,6);default:0.0000" json:"bonus_amount"`                     // 彩金赠送金额
	RechargeWithdrawDiff decimal.Decimal `gorm:"column:recharge_withdraw_diff;type:decimal(18,6);default:0.0000" json:"recharge_withdraw_diff"` // 充值提款差值
	RebateAmount         decimal.Decimal `gorm:"column:rebate_amount;type:decimal(18,6);default:0.0000" json:"rebate_amount"`                   // 月累计总返利
	FlashOrderCount      int             `gorm:"column:flash_order_count;default:0" json:"flash_order_count"`                                   // 月抢购订单数
	FlashOrderDoneCount  int             `gorm:"column:flash_order_done_count;default:0" json:"flash_order_done_count"`                         // 月抢购订单完成数
	GroupOrderCount      int             `gorm:"column:group_order_count;default:0" json:"group_order_count"`                                   // 月拼团订单数
	GroupOrderDoneCount  int             `gorm:"column:group_order_done_count;default:0" json:"group_order_done_count"`                         // 月拼团订单完成数
	CreatedAt            time.Time       `gorm:"column:create_at;autoCreateTime" json:"create_at"`                                              // 创建时间
	UpdatedAt            time.Time       `gorm:"column:update_at;autoUpdateTime" json:"update_at"`                                              // 更新时间
}
