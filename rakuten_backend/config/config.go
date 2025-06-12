package config

import (
	"fmt"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	Server ServerConfig `mapstructure:"server"`
	MySQL  MySQL        `mapstructure:"mysql"`
	JWT    JWTConfig    `mapstructure:"jwt"`
	HTTP   HTTP         `mapstructure:"jwt"`
	Log    LogConfig    `mapstructure:"log"`
}
type HTTP struct {
	Port    int `mapstructure:"port"`
	Timeout int `mapstructure:"timeout"`
}

type ServerConfig struct {
	Port string `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`      // 日志级别
	Filename   string `mapstructure:"filename"`   // 日志文件路径
	MaxSize    int    `mapstructure:"maxSize"`    // 单个日志文件最大尺寸，单位MB
	MaxBackups int    `mapstructure:"maxBackups"` // 最大保留多少个备份
	MaxAge     int    `mapstructure:"maxAge"`     // 日志文件保留天数
	Compress   bool   `mapstructure:"compress"`   // 是否压缩
	Console    bool   `mapstructure:"console"`    // 是否输出到控制台
}

type MySQL struct {
	Host            string        `mapstructure:"host"`
	Port            string        `mapstructure:"port"`
	Username        string        `mapstructure:"username"`
	Password        string        `mapstructure:"password"`
	DBName          string        `mapstructure:"dbname"`
	Charset         string        `mapstructure:"charset"`
	ParseTime       bool          `mapstructure:"parseTime"`
	Loc             string        `mapstructure:"loc"`
	MaxOpenConns    int           `mapstructure:"maxOpenConns"`    // 最大打开连接数
	MaxIdleConns    int           `mapstructure:"maxIdleConns"`    // 最大空闲连接数
	ConnMaxLifetime time.Duration `mapstructure:"connMaxLifetime"` // 连接最大生命周期
}

type JWTConfig struct {
	Secret string `mapstructure:"secret"`
	Expire int    `mapstructure:"expire"`
	Issuer string `mapstructure:"issuer"`
}

// ConfigManager 配置管理器
type ConfigManager struct {
	config Config
	viper  *viper.Viper
}

var (
	// 全局配置实例
	GlobalConfig Config
	// 用于确保只初始化一次
	initOnce sync.Once
	// 用于确保线程安全
	mu sync.RWMutex
	// viper实例
	v *viper.Viper
)

// OnConfigChange 配置变更时的回调函数
var OnConfigChange func()

// InitConfig 初始化全局配置
func InitConfig() error {
	var initErr error
	initOnce.Do(func() {
		v = viper.New()

		// 设置默认值
		setDefaults()

		// 设置配置文件
		v.SetConfigName("config")
		v.SetConfigType("yaml")
		v.AddConfigPath("./config")

		// 启用配置热更新
		v.WatchConfig()
		v.OnConfigChange(func(e fsnotify.Event) {
			mu.Lock()
			defer mu.Unlock()

			// 重新加载配置
			if err := v.Unmarshal(&GlobalConfig); err != nil {
				panic(fmt.Sprintf("Failed to reload config: %v", err))
			}

			// 验证配置
			if err := validateConfig(); err != nil {
				panic(fmt.Sprintf("Invalid config after reload: %v", err))
			}

			// 如果设置了回调函数，则执行
			if OnConfigChange != nil {
				OnConfigChange()
			}
		})

		if err := v.ReadInConfig(); err != nil {
			initErr = err
			return
		}

		if err := v.Unmarshal(&GlobalConfig); err != nil {
			initErr = err
			return
		}

		// 验证配置
		if err := validateConfig(); err != nil {
			initErr = err
			return
		}
	})

	return initErr
}

// setDefaults 设置默认配置值
func setDefaults() {
	// 服务器默认配置
	v.SetDefault("server.port", "8080")
	v.SetDefault("server.mode", "debug")

	// 数据库默认配置

	v.SetDefault("mysql.host", "localhost")
	v.SetDefault("mysql.port", "3306")
	v.SetDefault("mysql.charset", "utf8mb4")
	v.SetDefault("mysql.parseTime", true)
	v.SetDefault("mysql.loc", "Local")
	v.SetDefault("mysql.maxOpenConns", 100)
	v.SetDefault("mysql.maxIdleConns", 20)
	v.SetDefault("mysql.connMaxLifetime", 3600)

	// JWT默认配置
	v.SetDefault("jwt.expire", 24) // 默认24小时
	v.SetDefault("jwt.issuer", "gin-demo")

	// 日志默认配置
	v.SetDefault("log.level", "debug")
	v.SetDefault("log.filename", "logs/app.log")
	v.SetDefault("log.maxSize", 100)
	v.SetDefault("log.maxBackups", 10)
	v.SetDefault("log.maxAge", 30)
	v.SetDefault("log.compress", true)
	v.SetDefault("log.console", true)
}

// validateConfig 验证配置
func validateConfig() error {
	// 验证服务器配置
	if GlobalConfig.Server.Port == "" {
		return fmt.Errorf("server port is required")
	}

	if GlobalConfig.MySQL.Host == "" {
		return fmt.Errorf("database host is required")
	}
	if GlobalConfig.MySQL.Port == "" {
		return fmt.Errorf("database port is required")
	}
	if GlobalConfig.MySQL.Username == "" {
		return fmt.Errorf("database username is required")
	}
	if GlobalConfig.MySQL.DBName == "" {
		return fmt.Errorf("database name is required")
	}

	// 验证JWT配置
	if GlobalConfig.JWT.Secret == "" {
		return fmt.Errorf("jwt secret is required")
	}
	if GlobalConfig.JWT.Expire <= 0 {
		return fmt.Errorf("jwt expire time must be greater than 0")
	}

	// 验证日志配置
	if GlobalConfig.Log.Level == "" {
		return fmt.Errorf("log level is required")
	}
	if GlobalConfig.Log.Filename == "" {
		return fmt.Errorf("log filename is required")
	}
	if GlobalConfig.Log.MaxSize <= 0 {
		return fmt.Errorf("log maxSize must be greater than 0")
	}
	if GlobalConfig.Log.MaxBackups <= 0 {
		return fmt.Errorf("log maxBackups must be greater than 0")
	}
	if GlobalConfig.Log.MaxAge <= 0 {
		return fmt.Errorf("log maxAge must be greater than 0")
	}
	return nil
}

// GetDSN 获取数据库连接字符串
func GetDSN() string {
	cfg := GlobalConfig.MySQL
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%v&loc=%s",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
		cfg.Charset,
		cfg.ParseTime,
		cfg.Loc,
	)
}

// IsDebug 判断是否为调试模式
func IsDebug() bool {
	return GlobalConfig.Server.Mode == "debug"
}

// GetServerAddr 获取服务器地址
func GetServerAddr() string {
	return ":" + GlobalConfig.Server.Port
}

func GetJWT() JWTConfig {
	return GlobalConfig.JWT
}
