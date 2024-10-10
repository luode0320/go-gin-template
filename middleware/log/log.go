package log

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.SugaredLogger

// LogConfig 用于配置日志参数
type LogConfig struct {
	Level          zapcore.Level `yaml:"level"`           // 日志等级
	LogFilePath    string        `yaml:"log_file_path"`   // 日志文件路径
	MaxSize        int           `yaml:"max_size"`        // 单个文件的最大大小（单位：MB）
	MaxAge         int           `yaml:"max_age"`         // 文件最大保留时间（单位：天）
	MaxBackups     int           `yaml:"max_backups"`     // 最多保留的备份文件数量
	Compress       bool          `yaml:"compress"`        // 是否压缩旧的日志文件
	DisableConsole bool          `yaml:"disable_console"` // 是否禁用控制台输出
}

// 日志配置: 可以移动到配置文件
var logConfig = LogConfig{
	Level:          zap.InfoLevel,
	LogFilePath:    filepath.Join("app.log"),
	MaxSize:        10, // 10 MB
	MaxAge:         7,  // 7 days
	MaxBackups:     5,  // 5 backup files
	Compress:       true,
	DisableConsole: false,
}

// 设置 zap 日志记录器
func setupZapLogger(config *LogConfig) (*zap.Logger, error) {
	// 创建 lumberjack 对象
	hook := lumberjack.Logger{
		Filename:   config.LogFilePath,
		MaxSize:    config.MaxSize, // MB
		MaxAge:     config.MaxAge,  // days
		MaxBackups: config.MaxBackups,
		Compress:   config.Compress,
	}

	// 设置日志级别
	lvl := zap.NewAtomicLevelAt(config.Level)

	// 设置编码器
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	// 创建核心组件
	core := zapcore.NewTee(
		//输出到控制台
		zapcore.NewCore(encoder, zapcore.AddSync(&hook), lvl),
		//输出到指定文件
		zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), lvl),
	)
	// 创建 zap.Logger 实例
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	Logger = logger.Sugar()
	return logger, nil
}

// CustomLogger 自定义日志中间件
func CustomLogger() gin.HandlerFunc {
	// 初始化 zap 日志记录器
	logger, err := setupZapLogger(&logConfig)
	if err != nil {
		log.Fatalf("Failed to initialize zap logger: %v", err)
	}

	return func(c *gin.Context) {
		startTime := time.Now()

		c.Next()

		endTime := time.Now()

		method := c.Request.Method
		path := c.Request.URL.Path
		statusCode := c.Writer.Status()
		duration := endTime.Sub(startTime)

		// 使用结构化日志记录请求信息
		logger.Info("请求信息:",
			zap.String("method", method),
			zap.String("URL", path),
			zap.Int("状态码", statusCode),
			zap.Float64("持续时间(s)", duration.Seconds()),
		)
	}
}

// Debug 输出debug 级别的日志
//
// @Author: zhaoruobo
// @Date: 2023/9/26
func Debug(message string) {
	Logger.Debug(message)
}

func Debugf(tmp string, arg ...interface{}) {
	Logger.Debugf(tmp, arg...)
}

// Info 输出info 级别的日志
func Info(message string) {
	Logger.Info(message)
}

func Infof(tmp string, arg ...interface{}) {
	Logger.Infof(tmp, arg...)
}

// Warn 输出warm 级别的日志
func Warn(message string) {
	Logger.Warn(message)
}

func Warmf(tmp string, arg ...interface{}) {
	Logger.Warnf(tmp, arg...)
}

// Error 输出error 级别的日志
func Error(message string) {
	Logger.Error(message)
}

func Errorf(tmp string, arg ...interface{}) {
	Logger.Errorf(tmp, arg...)
}

// Panic 输出panic 级别的日志，打印完信息后会抛出panic中断程序，谨慎使用
func Panic(message string) {
	Logger.Panic(message)
}

func Panicf(tmp string, arg ...interface{}) {
	Logger.Panicf(tmp, arg...)
}

// Fatal 输出fatal 级别的日志，打印完信息后会调用os.Exit退出服务，谨慎使用
func Fatal(message string) {
	Logger.Fatal(message)
}

func Fatalf(tmp string, arg ...interface{}) {
	Logger.Fatalf(tmp, arg...)
}
