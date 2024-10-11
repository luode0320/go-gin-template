package log

import (
	"github.com/natefinch/lumberjack"
	"go-gin-template/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var Logger *zap.SugaredLogger

// LogConfig 用于配置日志参数
type LogConfig struct {
	Level          zapcore.Level `yaml:"level"`           // 日志等级: -1:debug < 0:info < 1:warm < 2:error < 3:panic
	LogFilePath    string        `yaml:"log_file_path"`   // 日志文件路径
	MaxSize        int           `yaml:"max_size"`        // 单个文件的最大大小（单位：MB）
	MaxAge         int           `yaml:"max_age"`         // 文件最大保留时间（单位：天）
	MaxBackups     int           `yaml:"max_backups"`     // 最多保留的备份文件数量
	Compress       bool          `yaml:"compress"`        // 是否压缩旧的日志文件
	DisableConsole bool          `yaml:"disable_console"` // 是否禁用控制台输出
}

func InitLogger() {
	// 初始化 zap 日志记录器
	logger, err := setupZapLogger(&LogConfig{
		Level:          zapcore.Level(config.ServiceConfig.Logger.LogLevel),
		LogFilePath:    filepath.Join(strings.Split(strings.ReplaceAll(strings.ReplaceAll(config.ServiceConfig.Logger.FileName, "/", ","), "\\", ","), ",")...),
		MaxSize:        config.ServiceConfig.Logger.MaxSize,
		MaxAge:         config.ServiceConfig.Logger.MaxAge,
		MaxBackups:     config.ServiceConfig.Logger.MaxBackups,
		Compress:       config.ServiceConfig.Logger.Compress,
		DisableConsole: config.ServiceConfig.Logger.DisableConsole,
	})
	if err != nil {
		log.Fatalf("Failed to initialize zap logger: %v", err)
	}

	Logger = logger.Sugar()
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
	return zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1)), nil
}

// Debug 输出debug 级别的日志
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

func Infof(fmt string, arg ...interface{}) {
	Logger.Infof(fmt, arg...)
}

// Warn 输出warm 级别的日志
func Warn(message string) {
	Logger.Warn(message)
}

func Warmf(fmt string, arg ...interface{}) {
	Logger.Warnf(fmt, arg...)
}

// Error 输出error 级别的日志, 不会退出程序
func Error(message string) {
	Logger.Error(message)
}

func Errorf(fmt string, arg ...interface{}) {
	Logger.Errorf(fmt, arg...)
}

// Panic 输出panic 级别的日志，打印完信息后会抛出panic中断程序，谨慎使用
func Panic(message string) {
	Logger.Panic(message)
}

func Panicf(fmt string, arg ...interface{}) {
	Logger.Panicf(fmt, arg...)
}

// Fatal 输出fatal 级别的日志，打印完信息后会调用os.Exit退出服务，谨慎使用
func Fatal(message string) {
	Logger.Fatal(message)
}

func Fatalf(fmt string, arg ...interface{}) {
	Logger.Fatalf(fmt, arg...)
}
