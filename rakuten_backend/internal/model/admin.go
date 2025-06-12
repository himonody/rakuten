package model

import "time"

type AdminUser struct {
	ID         uint64    `gorm:"column:id" json:"id"`                   // 管理员ID
	Username   string    `gorm:"column:username" json:"username"`       // 管理员用户名
	Password   string    `gorm:"column:password" json:"password"`       // 管理员密码，md5+盐 -> reverse -> md5+盐
	GoogleAuth string    `gorm:"column:google_auth" json:"google_auth"` // Google验证信息
	AuthIP     string    `gorm:"column:auth_ip" json:"auth_ip"`         // 授权IP
	Role       uint8     `gorm:"column:role" json:"role"`               // 0普通管理员，1超级管理员
	IsAgent    uint8     `gorm:"column:is_agent" json:"is_agent"`       // 0不是，1是
	Status     uint8     `gorm:"column:status" json:"status"`           // 0正常 1禁用
	CreatedAt  time.Time `gorm:"column:create_at" json:"create_at"`     // 创建时间
	UpdatedAt  time.Time `gorm:"column:update_at" json:"update_at"`     // 更新时间
}

type AdminUserOperationLog struct {
	ID            uint64    `gorm:"column:id" json:"id"` // 操作id
	OperatorId    string    `gorm:"column:operator_id;comment:操作者id" json:"operator_id"`
	OperationType uint64    `gorm:"column:operation_type" json:"operation_type"` //操作类型
	OperationIp   string    `gorm:"column:operation_ip" json:"operation_ip"`     //操作ip
	Content       string    `gorm:"column:content" json:"content"`               //操作内容
	CreatedAt     time.Time `gorm:"column:create_at" json:"create_at"`           // 创建时间
	UpdatedAt     time.Time `gorm:"column:update_at" json:"update_at"`           // 更新时间
}

/*
CREATE TABLE `admin_user_operation_log` (
                      `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '操作id，自增主键',
                      `operator_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '管理员ID',
                      `operation_type` TINYINT NOT NULL DEFAULT 0 COMMENT '操作类型',
                      `operation_ip` VARCHAR(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT  '授权ip',
                      `content` VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT  '操作内容',
                      `create_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                      `update_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                      PRIMARY KEY (`id`),
                      KEY `idx_admin_id` (`admin_id`)

) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='管理员操作记录';
*/
