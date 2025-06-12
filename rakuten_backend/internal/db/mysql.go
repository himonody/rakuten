package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"log"
	"os"
	"rakuten_backend/config"
	"sync"
	"time"
)

// Config 是数据库配置结构体

var (
	MySQL *gorm.DB  //  全局变量
	once  sync.Once // 单例保证
	err   error
)

// InitMySQL 初始化数据库连接（只执行一次）
func InitMySQL() error {
	once.Do(func() {

		newLogger := gormLogger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			gormLogger.Config{
				SlowThreshold:             time.Second,
				LogLevel:                  4,
				IgnoreRecordNotFoundError: true,
				Colorful:                  true,
			},
		)

		MySQL, err = gorm.Open(mysql.Open(config.GetDSN()), &gorm.Config{
			Logger: newLogger,
		})

		if err != nil {
			err = fmt.Errorf("gorm Open 失败: %w", err)
			return
		}

		sqlDB, e := MySQL.DB()
		if e != nil {
			err = fmt.Errorf("获取底层 sqlDB 失败: %w", e)
			return
		}

		sqlDB.SetMaxOpenConns(config.GlobalConfig.MySQL.MaxOpenConns)
		sqlDB.SetMaxIdleConns(config.GlobalConfig.MySQL.MaxIdleConns)
		sqlDB.SetConnMaxLifetime(config.GlobalConfig.MySQL.ConnMaxLifetime)

		if e := sqlDB.Ping(); e != nil {
			err = fmt.Errorf("数据库 ping 失败: %w", e)
		}
	})
	return err
}
