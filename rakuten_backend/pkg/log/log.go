package log

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"rakuten_backend/config"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	// ZapLog 全局日志对象
	ZapLog      *zap.SugaredLogger
	uuidRequest string
)

func GetUUid(uid string) {
	uuidRequest = uid
}

// InitLog 初始化日志
func InitLog() error {
	// 创建日志目录
	if err := os.MkdirAll(filepath.Dir(config.GlobalConfig.Log.Filename), 0755); err != nil {
		return err
	}

	// 设置日志级别
	level := zap.InfoLevel
	switch config.GlobalConfig.Log.Level {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	}

	// 配置编码器
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// 配置输出
	var cores []zapcore.Core

	// 文件输出
	fileWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   config.GlobalConfig.Log.Filename,
		MaxSize:    config.GlobalConfig.Log.MaxSize,
		MaxBackups: config.GlobalConfig.Log.MaxBackups,
		MaxAge:     config.GlobalConfig.Log.MaxAge,
		Compress:   config.GlobalConfig.Log.Compress,
	})
	fileCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		fileWriter,
		level,
	)
	cores = append(cores, fileCore)

	// 控制台输出
	if config.GlobalConfig.Log.Console {
		consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)
		consoleCore := zapcore.NewCore(
			consoleEncoder,
			zapcore.AddSync(os.Stdout),
			level,
		)
		cores = append(cores, consoleCore)
	}

	// 创建Logger
	core := zapcore.NewTee(cores...)
	ZapLog = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1)).Sugar()
	return nil
}

// GinLogger 接收gin框架默认的日志
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		url := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		cost := time.Since(start)
		ZapLog.Info(url,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", url),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.Duration("cost", cost),
		)
	}
}

// GinRecovery 恢复项目可能出现的panic
func GinRecovery(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				ZapLog.Error("[Recovery from panic]",
					zap.Any("error", err),
					zap.String("request", fmt.Sprintf("%s %s", c.Request.Method, c.Request.URL.Path)),
				)
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}

func Info(args ...interface{}) {
	args = append(args, fmt.Sprintf(" %s ", uuidRequest))
	ZapLog.Named(funcName()).Info(args...)
}

func Infof(template string, args ...interface{}) {
	template += " %s "
	args = append(args, uuidRequest)
	ZapLog.Named(funcName()).Infof(template, args...)
}

func Warn(args ...interface{}) {
	args = append(args, fmt.Sprintf(" commitId:%s ", uuidRequest))
	ZapLog.Named(funcName()).Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	template += " %s "
	args = append(args, uuidRequest)
	ZapLog.Named(funcName()).Warnf(template, args...)
}

func Error(args ...interface{}) {
	args = append(args, fmt.Sprintf(" %s", uuidRequest))
	ZapLog.Named(funcName()).Error(args...)
}

func Errorf(template string, args ...interface{}) {
	template += " %s "
	args = append(args, uuidRequest)
	ZapLog.Named(funcName()).Errorf(template, args...)
}

func funcName() string {
	pc, _, _, _ := runtime.Caller(2)
	funcName := runtime.FuncForPC(pc).Name()
	return path.Base(funcName)
}
