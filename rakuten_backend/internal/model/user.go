package model

import (
	"github.com/shopspring/decimal"
	"time"
)

// User 会员表
type User struct {
	ID         int    `gorm:"column:id;primaryKey" json:"id"`                                 //会员ID
	Username   string `gorm:"column:username;comment:会员用户名" json:"username"`                  //会员用户名
	Password   string `gorm:"column:password;comment:会员密码 md5+盐 然后逆序 再md5+盐" json:"password"` //会员密码
	Telegram   string `gorm:"column:telegram;comment:telegram json" json:"telegram"`          //会员飞机
	Google     string `gorm:"column:google;comment:谷歌信息 json" json:"google"`                  //会员谷歌
	GoogleAuth string `gorm:"column:google_auth;comment:谷歌验证信息" json:"google_auth"`           //会员谷歌验证码
	Apple      string `gorm:"column:apple;comment:apple信息 json" json:"apple"`                 //会员苹果
	Email      string `gorm:"column:email;comment:邮箱信息 json" json:"email"`                    //会员邮箱
	InviteCode string `gorm:"column:invite_code;comment:邀请code" json:"invite_code"`           //邀请码
	InviteId   int    `gorm:"column:invite_id;comment:邀请id" json:"invite_id"`                 //邀请人id
	IsOnline   int    `gorm:"column:is_online;comment:是否在线 0 离线 1 在线" json:"is_online"`       //是否在线 0 离线 1 在线

	Status   int    `gorm:"column:status;comment:会员状态 0 正常 1 禁用" json:"status"`                           //会员状态 0 正常 1 禁用
	VIPLevel int    `gorm:"column:vip_level;comment:VIP等级 0普通会员 1 一级会员 2 二级会员 3 三级会员 4" json:"vip_level"` //VIP等级 0普通会员 1 一级会员 2 二级会员 3 三级会员 4
	Avatar   string `gorm:"column:avatar;comment:头像" json:"avatar"`                                       //头像

	PayoutPassword string `gorm:"column:payout_password;comment:提款密码 md5+盐 然后逆序 再md5+盐" json:"payout_password"` //提款密码
	PayoutStatus   int    `gorm:"column:payout_status;comment:提款冻结状态 0正常 1冻结" json:"payout_status"`             //提款冻结状态 0正常 1冻结

	FrozenBalance    decimal.Decimal `gorm:"column:frozen_balance;comment:冻结金额" json:"frozen_balance"`       //冻结金额
	FrozenStatus     int             `gorm:"column:frozen_status;comment:冻结状态 0正常 1冻结" json:"frozen_status"` //冻结状态 0正常 1冻结
	TotalBalance     decimal.Decimal `gorm:"column:total_balance;comment:总余额（可用 + 冻结）" json:"total_balance"` //总余额（可用 + 冻结）
	AvailableBalance decimal.Decimal `gorm:"column:available_balance;comment:可用余额" json:"available_balance"` //可用余额

	RegisterIP     string    `gorm:"column:register_ip;comment:注册 IP（支持 IPv4/IPv6）" json:"register_ip"` //注册 IP（支持 IPv4/IPv6）
	RegisterTime   time.Time `gorm:"column:register_time;comment:注册时间" json:"register_time"`            //注册时间
	RegisterDevice string    `gorm:"column:register_device;comment:注册设备" json:"register_device"`        //注册设备

	CustomerService1 string    `gorm:"column:customer_service1;comment:客服飞机链接1" json:"customer_service1"` //客服飞机链接1
	CustomerService2 string    `gorm:"column:customer_service2;comment:客服飞机链接2" json:"customer_service2"` //客服飞机链接2
	Version          int       `gorm:"column:version;comment:版本号" json:"version"`                         //版本号
	CreateAt         time.Time `gorm:"column:create_at;comment:创建时间" json:"create_at"`
	UpdateAt         time.Time `gorm:"column:update_at;comment:更新时间" json:"update_at"`
}

/*
CREATE TABLE `user` (
                        `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '会员ID，自增主键',

    -- 基本信息
                        `username` VARCHAR(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '会员用户名',
                        `password` VARCHAR(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '会员密码，md5+盐 -> reverse -> md5+盐',
                        `telegram` TEXT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT 'Telegram信息（JSON）',
                        `google` TEXT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT 'Google信息（JSON）',
                        `google_auth` TEXT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT 'Google验证信息',
                        `apple` TEXT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT 'Apple信息（JSON）',
                        `email` TEXT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '邮箱信息（JSON）',

    -- 邀请、代理、状态
                        `invite_code` VARCHAR(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '邀请码',
                        `invite_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '邀请人ID',
                        `is_online`  TINYINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否在线：0=离线，1=在线',
                        `status` TINYINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '会员状态：0=正常，1=禁用',
                        `vip_level_id` TINYINT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'VIP等级：0普通，1~4等级',

    -- 头像、客服、注册信息
                        `avatar` VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '头像',

                        `register_ip` VARCHAR(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '注册IP（支持IPv4/IPv6）',
                        `register_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '注册时间',
                        `register_device` VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '注册设备信息',

                        `customer_service1` VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '客服Telegram链接（多个用逗号分隔）',
                        `customer_service2` VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '客服Telegram链接（多个用逗号分隔）',
                        `payout_password` VARCHAR(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '提款密码，md5+盐 -> reverse -> md5+盐',
                        `payout_status` TINYINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '提现状态：0正常，1冻结',
                        `frozen_balance` DECIMAL(18, 6) NOT NULL DEFAULT 0 COMMENT '冻结余额',
                        `frozen_status` TINYINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '冻结状态：0正常，1冻结',
                        `total_balance` DECIMAL(18, 6) NOT NULL DEFAULT 0 COMMENT '总余额（可用+冻结）',
                        `available_balance` DECIMAL(18, 6) NOT NULL DEFAULT 0 COMMENT '可用余额',

    -- 乐观锁
                        `version` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '乐观锁版本号',
                        `operator_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '管理员ID',

    -- 通用字段
                        `create_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                        `update_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',

                        PRIMARY KEY (`id`),
                        UNIQUE KEY `uniq_username` (`username`),
                        UNIQUE KEY `uniq_invite_code` (`invite_code`),
                        KEY `idx_invite_id` (`invite_id`),
                        KEY `idx_register_time` (`register_time`),
                        KEY `idx_register_ip` (`register_ip`),
                        KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='会员表';
*/

// UserLoginLog 登录日志
type UserLoginLog struct {
	Id           int        `gorm:"column:id;primaryKey" json:"id"`                  //主键id
	UserId       int        `gorm:"column:user_id" json:"user_id"`                   //会员id
	LoginIp      string     `gorm:"column:login_ip" json:"login_ip"`                 //会员登录ip
	LoginAddress string     `gorm:"column:login_address" json:"login_address"`       //会员登录地址
	LoginAt      time.Time  `gorm:"column:login_at" json:"login_at"`                 //会员登录时间
	LogoutAt     *time.Time `gorm:"column:logout_at" json:"logout_at"`               //会员登出时间
	CreateAt     time.Time  `gorm:"column:create_at;comment:<UNK>" json:"create_at"` //创建时间
	UpdateAt     time.Time  `gorm:"column:update_at;comment:<UNK>" json:"update_at"` //更新时间
}

// VipLevelConfig VIP等级配置表
type VipLevelConfig struct {
	ID             int             `gorm:"column:id;primaryKey" json:"id"`                // 主键ID，自增
	Level          int             `gorm:"column:level;" json:"level"`                    // 会员等级 0 普通，1 一级...
	Name           string          `gorm:"column:name" json:"name"`                       // 等级名称
	Icon           string          `gorm:"column:icon" json:"icon"`                       // 等级图标 URL
	TopUpAmount    decimal.Decimal `gorm:"column:top_up_amount" json:"top_up_amount"`     // 升级所需最低累计充值金额
	CommissionRate decimal.Decimal `gorm:"column:commission_rate" json:"commission_rate"` // 佣金分成比例（百分比，如 12.5 表示12.5%）
	Benefits       string          `gorm:"column:benefits" json:"benefits"`               // 福利说明
	Status         int             `gorm:"column:status" json:"status"`                   // 启用状态 1 启用 0 禁用
	OperatorId     string          `gorm:"column:operator_id;comment:操作者id" json:"operator_id"`
	CreateAt       time.Time       `gorm:"column:create_at" json:"create_at"` // 创建时间
	UpdateAt       time.Time       `gorm:"column:update_at" json:"update_at"` // 更新时间
}

/*
CREATE TABLE `vip_level_config` (
                                  `id` TINYINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID，自增',
                                  `level` TINYINT UNSIGNED NOT NULL COMMENT '会员等级，例如 0 普通，1 一级，2 二级',
                                  `name` VARCHAR(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '等级名称，例如 普通会员、黄金会员',
                                  `icon` VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '等级图标URL',
                                  `top_up_amount` DOUBLE NOT NULL DEFAULT 0 COMMENT '升级所需最低累计充值金额',
                                  `benefits` TEXT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '会员福利说明，例如返佣、提现加速等',
                                  `commission_rate` DOUBLE NOT NULL DEFAULT 0 COMMENT '佣金分成比例（百分比）',
                                  `status` TINYINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '状态 0启用 1禁用',
                                  `operator_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '管理员ID',
                                  `create_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                  `update_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                  PRIMARY KEY (id),
                                  UNIQUE KEY uniq_level (level)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='会员等级配置表';

*/

// TopUpRecord 会员充值记录
type TopUpRecord struct {
	ID          int             `gorm:"primaryKey;column:id" json:"id"`          // 自增主键
	UserID      int             `gorm:"column:user_id" json:"user_id"`           // 用户ID
	Amount      decimal.Decimal `gorm:"column:amount" json:"amount"`             // 充值金额
	Status      int             `gorm:"column:status" json:"status"`             // 状态：0待审核 1成功 2失败
	AddressId   int             `gorm:"column:address_id" json:"address_id"`     //充值配置表id
	TxID        string          `gorm:"column:tx_id" json:"tx_id"`               // 区块链 TXID 或流水号
	ProofImage  string          `gorm:"column:proof_image" json:"proof_image"`   // 转账截图URL（可选）
	SubmitTime  time.Time       `gorm:"column:submit_time" json:"submit_time"`   // 提交时间
	SuccessTime *time.Time      `gorm:"column:success_time" json:"success_time"` // 成功时间
	OperatorId  string          `gorm:"column:operator_id;comment:操作者id" json:"operator_id"`
	Remark      string          `gorm:"column:remark" json:"remark"` // 备注
}

/*
-- 会员充值记录表
CREATE TABLE `top_up_record` (
                                 `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增主键',
                                 `user_id` BIGINT UNSIGNED NOT NULL COMMENT '用户ID',
                                 `amount` DECIMAL(18,6) NOT NULL COMMENT '充值金额',
                                 `status` TINYINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '状态：0待审核 1成功 2失败',
                                 `address_id` BIGINT UNSIGNED NOT NULL COMMENT '充值配置表id',
                                 `tx_id` VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '区块链 TXID 或流水号',
                                 `proof_image` VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '转账截图URL（可选）',
                                 `submit_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '提交时间',
                                 `success_time` TIMESTAMP NULL DEFAULT NULL COMMENT '成功时间',
                                 `operator_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '管理员ID',
                                 `remark` VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '备注',
                                 PRIMARY KEY (`id`),
                                 KEY `idx_user_id` (`user_id`),
                                 KEY `idx_address_id` (`address_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='会员充值记录表';
*/

// TopUpAddress 充值地址配置
type TopUpAddress struct {
	ID         int       `gorm:"primaryKey;column:id" json:"id"`
	Network    string    `gorm:"column:network" json:"network"` // 网络类型：TRC20 / ERC20
	Coin       string    `gorm:"column:coin" json:"coin"`       // 币种：USDT 等
	Address    string    `gorm:"column:address" json:"address"` // 共用地址
	Memo       string    `gorm:"column:memo" json:"memo"`       // 标签（如XRP需要）
	Status     int       `gorm:"column:status" json:"status"`   // 启用状态
	Remark     string    `gorm:"column:remark" json:"remark"`   // 备注
	OperatorId string    `gorm:"column:operator_id;comment:操作者id" json:"operator_id"`
	CreateAt   time.Time `gorm:"column:create_at" json:"create_at"`
	UpdateAt   time.Time `gorm:"column:update_at" json:"update_at"`
}

/*
-- 充值地址配置表
CREATE TABLE `top_up_address` (
                                  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增主键',
                                  `network` VARCHAR(20) NOT NULL COMMENT '网络类型：TRC20 / ERC20',
                                  `coin` VARCHAR(20) NOT NULL COMMENT '币种：USDT 等',
                                  `address` VARCHAR(100) NOT NULL COMMENT '共用地址',
                                  `memo` VARCHAR(100) DEFAULT NULL COMMENT '标签（如XRP需要）',
                                  `status` TINYINT UNSIGNED NOT NULL DEFAULT 1 COMMENT '启用状态，默认1',
                                  `remark` VARCHAR(255) DEFAULT NULL COMMENT '备注',
                                  `operator_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '管理员ID',
                                  `create_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                  `update_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                  PRIMARY KEY (`id`),
                                  UNIQUE KEY `uniq_address` (`address`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='充值地址配置表';

*/

// PayoutRecord 会员提款记录
type PayoutRecord struct {
	ID           int             `gorm:"primaryKey;column:id" json:"id"`            // 自增主键
	UserID       int             `gorm:"column:user_id" json:"user_id"`             // 用户ID
	Amount       decimal.Decimal `gorm:"column:amount" json:"amount"`               // 提现金额
	RealAmount   decimal.Decimal `gorm:"column:real_amount" json:"real_amount"`     //实际金额
	BeforeAmount decimal.Decimal `gorm:"column:before_amount" json:"before_amount"` //提现后余额
	Fee          decimal.Decimal `gorm:"column:fee" json:"fee"`                     //手续费金额
	Rate         decimal.Decimal `gorm:"column:rate" json:"rate"`                   //手续费比例，如 0.005 表示0.5%
	Status       int             `gorm:"column:status" json:"status"`               // 状态：0待审核 1已审核 2退回 3失败
	Method       string          `gorm:"column:method" json:"method"`               // 提现方式
	Address      string          `gorm:"column:address" json:"address"`             // 提现地址或账户
	TxID         string          `gorm:"column:tx_id" json:"tx_id"`                 // 区块链 TXID 或流水号
	TxURL        string          `gorm:"column:tx_url" json:"tx_url"`               // 交易详情url
	SubmitTime   time.Time       `gorm:"column:submit_time" json:"submit_time"`     // 提交时间
	SuccessTime  *time.Time      `gorm:"column:success_time" json:"success_time"`   // 成功时间，指针类型表示可空
	OperatorId   string          `gorm:"column:operator_id;comment:操作者id" json:"operator_id"`
	Remark       string          `gorm:"column:remark" json:"remark"` // 备注
}

/*
CREATE TABLE `payout_record` (
                                 `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增主键',
                                 `user_id` BIGINT UNSIGNED NOT NULL COMMENT '用户ID',
                                 `amount` DECIMAL(18,6) NOT NULL COMMENT '提现金额',
                                 `real_amount` DECIMAL(18,6) NOT NULL COMMENT '实际金额',
                                 `before_amount` DECIMAL(18,6) NOT NULL COMMENT '提现后余额',
                                 `fee` DECIMAL(18,6) NOT NULL COMMENT '手续费金额',
                                 `rate` DECIMAL(18,6) NOT NULL COMMENT '手续费比例，如 0.005 表示0.5%',
                                 `status` TINYINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '状态：0待审核 1已审核 2退回 3失败',
                                 `method` VARCHAR(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '提现方式',
                                 `address` VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '提现地址或账户',
                                 `tx_id` VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '区块链 TXID 或流水号',
                                 `tx_url` VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '交易详情url',
                                 `submit_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '提交时间',
                                 `success_time` TIMESTAMP DEFAULT NULL COMMENT '成功时间',
                                 `operator_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '管理员ID',
                                 `remark` VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '备注',
                                 PRIMARY KEY (`id`),
                                 KEY `idx_user_id` (`user_id`),
                                 KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='会员提款记录表';

*/

// UserBalanceLog 用户余额变动日志
type UserBalanceLog struct {
	ID            int             `gorm:"column:id;primaryKey;comment:主键ID" json:"id"`
	UserID        int             `gorm:"column:user_id;comment:用户ID" json:"user_id"`
	TxNo          string          `gorm:"column:tx_no" json:"tx_no"` // 区块链 TXID 或流水号
	Type          int             `gorm:"column:type;comment:帐变类型" json:"type"`
	ChangeAmount  decimal.Decimal `gorm:"column:change_amount;comment:帐变金额，正为增加，负为减少" json:"change_amount"`
	BeforeBalance decimal.Decimal `gorm:"column:before_balance;comment:变动前总余额" json:"before_balance"`
	AfterBalance  decimal.Decimal `gorm:"column:after_balance;comment:变动后可用余额" json:"after_balance"`
	Remark        string          `gorm:"column:remark;comment:备注说明" json:"remark"`
	OperatorId    string          `gorm:"column:operator_id;comment:操作者id" json:"operator_id"`
	CreatedAt     time.Time       `gorm:"column:created_at;comment:创建时间" json:"created_at"`
	UpdateAt      time.Time       `gorm:"column:update_at" json:"update_at"`
}

/*
CREATE TABLE `user_balance_log` (
                                    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
                                    `user_id` BIGINT UNSIGNED NOT NULL COMMENT '用户ID',
                                    `tx_no` VARCHAR(64) NOT NULL COMMENT '业务订单号',
                                    `type` TINYINT NOT NULL COMMENT '帐变类型',
                                    `change_amount` DECIMAL(18,6) NOT NULL COMMENT '帐变金额，正为增加，负为减少',
                                    `before_balance` DECIMAL(18,6) NOT NULL COMMENT '变动前总余额',
                                    `after_balance` DECIMAL(18,6) NOT NULL COMMENT '变动后总余额',
                                    `remark` VARCHAR(255) DEFAULT NULL COMMENT '备注说明',
                                    `operator_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '管理员ID',
                                    `create_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                    `update_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                    PRIMARY KEY (`id`),
                                    KEY `idx_tx_no` (`tx_no`),
                                    KEY `idx_user_id` (`user_id`),
                                    KEY `idx_type` (`type`),
                                    KEY `idx_create_at` (`create_at`),
                                    KEY `idx_update_at` (`update_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户余额变动日志表';
*/

// 余额变动类型枚举
const (
	BalanceLogTypeWithdraw     = 1 // 提款
	BalanceLogTypeWithdrawFail = 2 // 提款失败
	BalanceLogTypeAdminAdd     = 3 // 后台充值
	BalanceLogTypeAdminSub     = 4 // 后台减少
	BalanceLogTypeOnlineAdd    = 5 // 在线充值
	BalanceLogTypeReward       = 6 // 返奖
	BalanceLogTypeBonus        = 7 // 送彩金
)

// ProductConfig 商品配置
type ProductConfig struct {
	ID          int       `gorm:"primaryKey;column:id;comment:主键ID" json:"id"`
	ProductCode string    `gorm:"column:product_code;comment:商品编码，唯一" json:"product_code"`
	Name        string    `gorm:"column:name;comment:商品名称" json:"name"`
	Price       float64   `gorm:"column:price;comment:价格" json:"price"`
	ImageURL    string    `gorm:"column:image_url;comment:商品主图URL" json:"image_url"`
	Status      int       `gorm:"column:status;comment:商品状态(0上架，1下架)" json:"status"`
	Description string    `gorm:"column:description;comment:商品描述" json:"description"`
	OperatorId  string    `gorm:"column:operator_id;comment:操作者id" json:"operator_id"`
	CreatedAt   time.Time `gorm:"column:created_at;comment:创建时间" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at;comment:更新时间" json:"updated_at"`
}

/*
CREATE TABLE `product_config` (
                                  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
                                  `product_code` VARCHAR(64) NOT NULL COMMENT '商品编码，唯一',
                                  `name` VARCHAR(255) NOT NULL COMMENT '商品名称',
                                  `price` DECIMAL(18,6) NOT NULL COMMENT '价格',
                                  `image_url` VARCHAR(255) DEFAULT NULL COMMENT '商品主图URL',
                                  `status` TINYINT NOT NULL DEFAULT 1 COMMENT '商品状态(1上架，0下架)',
                                  `description` TEXT DEFAULT NULL COMMENT '商品描述',
                                  `operator_id` BIGINT  NOT NULL DEFAULT 0 COMMENT '操作id',
                                  `create_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                  `update_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                  PRIMARY KEY (`id`),
                                  UNIQUE KEY `uniq_product_code` (`product_code`),
                                  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='商品配置表';
*/
