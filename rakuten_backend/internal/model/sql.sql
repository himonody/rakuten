DROP TABLE IF EXISTS `user`;
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


CREATE TABLE `user_login_log` (
                        `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID，自增主键',
                        `user_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '会员Id',
                        `login_ip` VARCHAR(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '登录ip',
                        `login_address` VARCHAR(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '登陆地址',
                        `login_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '登录时间',
                        `logout_at` TIMESTAMP DEFAULT NULL COMMENT '登出时间',
                        `create_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                        `update_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',

                        PRIMARY KEY (`id`),
                        KEY `idx_user_id` (`user_id`),
                        KEY `idx_login_ip` (`login_ip`),
                        KEY `idx_login_at` (`login_at`),
                        KEY `idx_logout_at` (`logout_at`),
                        KEY `idx_login_address` (`login_address`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='会员登录日志';

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

-- 会员提款记录表
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
DROP TABLE `product_config`;
# product_config 已建
CREATE TABLE `product_config` (
                                  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
                                  `product_code` VARCHAR(64) NOT NULL COMMENT '商品编码，唯一',
                                  `name` VARCHAR(255) NOT NULL COMMENT '商品名称',
                                  `price` DECIMAL(18,6) NOT NULL COMMENT '价格',
                                  `image_url` VARCHAR(255) DEFAULT NULL COMMENT '商品主图URL',
                                  `status` TINYINT NOT NULL DEFAULT 1 COMMENT '商品状态(1上架，0下架)',
                                  `description` TEXT DEFAULT NULL COMMENT '商品描述',
                                  `operator_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '管理员ID',
                                  `create_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                  `update_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                  PRIMARY KEY (`id`),
                                  UNIQUE KEY `uniq_product_code` (`product_code`),
                                  KEY `idx_status` (`status`),
                                  KEY `idx_name` (`name`),
                                  KEY `idx_update_at` (`update_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='商品配置表';

# product_config 已建
CREATE TABLE `admin_dashboard` (
                                  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
                                  `product_code` VARCHAR(64) NOT NULL COMMENT '商品编码，唯一',
                                  `name` VARCHAR(255) NOT NULL COMMENT '商品名称',
                                  `price` DECIMAL(18,6) NOT NULL COMMENT '价格',
                                  `image_url` VARCHAR(255) DEFAULT NULL COMMENT '商品主图URL',
                                  `status` TINYINT NOT NULL DEFAULT 1 COMMENT '商品状态(1上架，0下架)',
                                  `description` TEXT DEFAULT NULL COMMENT '商品描述',
                                  `operator_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '管理员ID',
                                  `create_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                  `update_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                  PRIMARY KEY (`id`),
                                  UNIQUE KEY `uniq_product_code` (`product_code`),
                                  KEY `idx_status` (`status`),
                                  KEY `idx_name` (`name`),
                                  KEY `idx_update_at` (`update_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='商品配置表';

# admin_user 已建
CREATE TABLE `admin_user` (
                              `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '管理员ID，自增主键',
                              `username` VARCHAR(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '管理员用户名',
                              `password` VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '管理员密码，md5+盐 -> reverse -> md5+盐',
                              `google_auth` TEXT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL  COMMENT 'Google验证信息',
                              `auth_ip` TEXT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '授权ip',
                              `role`  TINYINT NOT NULL DEFAULT 0 COMMENT '0普通管理1超级管理员',
                              `status`  TINYINT NOT NULL DEFAULT 0 COMMENT '0正常 1禁用',
                              `is_agent`  TINYINT NOT NULL DEFAULT 0 COMMENT '0不是1是',
                              `create_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                              `update_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                              PRIMARY KEY (`id`),
                              UNIQUE KEY `uniq_username` (`username`)

) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='管理员表';

# admin_user_operation_log 已建
CREATE TABLE `admin_user_operation_log` (
                      `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '操作id，自增主键',
                      `operator_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '管理员ID',
                      `operation_type` TINYINT NOT NULL DEFAULT 0 COMMENT '操作类型',
                      `operation_ip` VARCHAR(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT  '授权ip',
                      `create_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                      `update_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                      PRIMARY KEY (`id`),
                      KEY `idx_admin_id` (`operator_id`),
                      KEY `idx_operation_type` (`operation_type`),
                      KEY `idx_create_at` (`create_at`)

) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='管理员操作记录';


-- 消息表：存储系统站内信消息体
CREATE TABLE `messages` (
                            `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '消息唯一 ID',
                            `title` VARCHAR(255) NOT NULL COMMENT '消息标题',
                            `content` TEXT NOT NULL COMMENT '消息内容',
                            `sender_id` BIGINT UNSIGNED NOT NULL COMMENT '发送者用户 ID',
                            `popup` TINYINT(1) DEFAULT 0 COMMENT '是否弹窗（0 否，1 是）',
                            `created_at` DATETIME NOT NULL COMMENT '创建时间',
                            PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='系统消息表';


-- 接收者表：记录每个消息的接收状态
CREATE TABLE `recipients` (
                              `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '唯一 ID',
                              `message_id` BIGINT UNSIGNED NOT NULL COMMENT '关联的消息 ID',
                              `receiver_id` BIGINT UNSIGNED NOT NULL COMMENT '接收者用户 ID',
                              `read` TINYINT(1) DEFAULT 0 COMMENT '是否已读（0 未读，1 已读）',
                              `read_at` DATETIME DEFAULT NULL COMMENT '读取时间',
                              `delivered` TINYINT(1) DEFAULT 0 COMMENT '是否已送达（0 未送达，1 已送达）',
                              `popup_shown` TINYINT(1) DEFAULT 0 COMMENT '弹窗是否已显示（0 否，1 是）',
                              PRIMARY KEY (`id`),
                              KEY `idx_receiver_id` (`receiver_id`),
                              KEY `idx_message_id` (`message_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='消息接收状态表';
DROP TABLE IF EXISTS `total_stats`;
-- total_stats 总统计表
CREATE TABLE `total_stats` (
                               `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
                               `total_admin_count` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '总管理员人数',
                               `total_member_count` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '总会员人数',
                               `total_agent_count` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '总代理人数',
                               `recharge_amount` DECIMAL(18,6) NOT NULL DEFAULT 0.0000 COMMENT '累计充值金额',
                               `withdraw_amount` DECIMAL(18,6) NOT NULL DEFAULT 0.0000 COMMENT '累计提款金额',
                               `bonus_amount` DECIMAL(18,6) NOT NULL DEFAULT 0.0000 COMMENT '彩金赠送金额',
                               `recharge_withdraw_diff` DECIMAL(18,6) NOT NULL DEFAULT 0.0000 COMMENT '充值提款差值（充值 - 提款）',
                               `rebate_amount` DECIMAL(18,6) NOT NULL DEFAULT 0.0000 COMMENT '累计总返利',
                               `flash_order_count` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '总抢购订单数',
                               `flash_order_done_count` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '总抢购订单完成数',
                               `group_order_count` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '总拼团订单数',
                               `group_order_done_count` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '总拼团订单完成数',
                               `create_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                               `update_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                               PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='总统计表';
DROP TABLE IF EXISTS `daily_stats`;
-- daily_stats 日统计表
CREATE TABLE `daily_stats` (
                               `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
                               `stat_date` DATE NOT NULL COMMENT '统计日期',
                               `new_admin_count` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '日新增管理员人数',
                               `new_member_count` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '日新增会员人数',
                               `new_agent_count` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '日新增代理人数',
                               `recharge_amount` DECIMAL(18,6) NOT NULL DEFAULT 0.0000 COMMENT '日累计充值金额',
                               `withdraw_amount` DECIMAL(18,6) NOT NULL DEFAULT 0.0000 COMMENT '日累计提款金额',
                               `bonus_amount` DECIMAL(18,6) NOT NULL DEFAULT 0.0000 COMMENT '彩金赠送金额',
                               `recharge_withdraw_diff` DECIMAL(18,6) NOT NULL DEFAULT 0.0000 COMMENT '充值提款差值（充值 - 提款）',
                               `rebate_amount` DECIMAL(18,6) NOT NULL DEFAULT 0.0000 COMMENT '日累计总返利',
                               `flash_order_count` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '日抢购订单数',
                               `flash_order_done_count` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '日抢购订单完成数',
                               `group_order_count` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '日拼团订单数',
                               `group_order_done_count` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '日拼团订单完成数',
                               `create_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                               `update_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                               PRIMARY KEY (`id`),
                               UNIQUE KEY `uk_stat_date` (`stat_date`),
                               KEY `idx_id_stat_date` (`id`, `stat_date`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='日统计表';
DROP TABLE IF EXISTS `weekly_stats`;
-- weekly_stats 周统计表
CREATE TABLE `weekly_stats` (
                                `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
                                `stat_week` CHAR(8) NOT NULL COMMENT '统计周（格式：YYYY-WW）',
                                `new_admin_count` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '周新增管理员人数',
                                `new_member_count` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '周新增会员人数',
                                `new_agent_count` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '周新增代理人数',
                                `recharge_amount` DECIMAL(18,6) NOT NULL DEFAULT 0.0000 COMMENT '周累计充值金额',
                                `withdraw_amount` DECIMAL(18,6) NOT NULL DEFAULT 0.0000 COMMENT '周累计提款金额',
                                `bonus_amount` DECIMAL(18,6) NOT NULL DEFAULT 0.0000 COMMENT '彩金赠送金额',
                                `recharge_withdraw_diff` DECIMAL(18,6) NOT NULL DEFAULT 0.0000 COMMENT '充值提款差值（充值 - 提款）',
                                `rebate_amount` DECIMAL(18,6) NOT NULL DEFAULT 0.0000 COMMENT '周累计总返利',
                                `flash_order_count` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '周抢购订单数',
                                `flash_order_done_count` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '周抢购订单完成数',
                                `group_order_count` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '周拼团订单数',
                                `group_order_done_count` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '周拼团订单完成数',
                                `create_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                `update_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                PRIMARY KEY (`id`),
                                UNIQUE KEY `uk_stat_week` (`stat_week`),
                                KEY `idx_id_stat_week` (`id`, `stat_week`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='周统计表';
DROP TABLE IF EXISTS `monthly_stats`;
-- monthly_stats 月统计表
CREATE TABLE `monthly_stats` (
                                 `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
                                 `stat_month` CHAR(7) NOT NULL COMMENT '统计月份（格式：YYYY-MM）',
                                 `new_admin_count` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '月新增管理员人数',
                                 `new_member_count` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '月新增会员人数',
                                 `new_agent_count` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '月新增代理人数',
                                 `recharge_amount` DECIMAL(18,6) NOT NULL DEFAULT 0.0000 COMMENT '月累计充值金额',
                                 `withdraw_amount` DECIMAL(18,6) NOT NULL DEFAULT 0.0000 COMMENT '月累计提款金额',
                                 `bonus_amount` DECIMAL(18,6) NOT NULL DEFAULT 0.0000 COMMENT '彩金赠送金额',
                                 `recharge_withdraw_diff` DECIMAL(18,6) NOT NULL DEFAULT 0.0000 COMMENT '充值提款差值（充值 - 提款）',
                                 `rebate_amount` DECIMAL(18,6) NOT NULL DEFAULT 0.0000 COMMENT '月累计总返利',
                                 `flash_order_count` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '月抢购订单数',
                                 `flash_order_done_count` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '月抢购订单完成数',
                                 `group_order_count` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '月拼团订单数',
                                 `group_order_done_count` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '月拼团订单完成数',
                                 `create_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                 `update_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                 PRIMARY KEY (`id`),
                                 UNIQUE KEY `uk_stat_month` (`stat_month`),
                                 KEY `idx_id_stat_month` (`id`, `stat_month`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='月统计表';

