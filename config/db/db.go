// db/db.go
package db

import (
	"database/sql"
	"fmt"
	"github.com/robfig/cron"
	"go-gin-template/config"
	"go-gin-template/config/log"
	"go-gin-template/model"
	"gorm.io/gorm/logger"
	syslog "log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Conn 连接实例
var Conn *gorm.DB

// InitDB 初始化数据库连接
func InitDB() {
	// 检查数据库是否存在，如果不存在则创建
	createDB()
	// 连接数据库
	connectDB()
	// 自动创建或更新数据库表结构
	migrate()
	// 数据库连接检查
	connectionCheck()
}

// 检查数据库是否存在，如果不存在则创建
func createDB() {
	// 将yaml配置参数拼接成连接数据库的url
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8mb4&parseTime=True&loc=Local",
		config.ServiceConfig.Db.UserName,
		config.ServiceConfig.Db.Password,
		config.ServiceConfig.Db.Url,
		config.ServiceConfig.Db.Port,
	)

	// 连接数据库
	dbTemp, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Errorf("连接系统数据库失败 : %s", err.Error())
		return
	}
	defer dbTemp.Close()

	// 检查数据库是否存在，如果不存在则创建
	_, err = dbTemp.Exec("CREATE DATABASE IF NOT EXISTS " + config.ServiceConfig.Db.DbName)
	if err != nil {
		log.Errorf("创建系统数据库失败: %s", err.Error())
		return
	}

	log.Infof("===========数据库检查完成==============")
}

// 连接数据库
func connectDB() {
	//将yaml配置参数拼接成连接数据库的url
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.ServiceConfig.Db.UserName,
		config.ServiceConfig.Db.Password,
		config.ServiceConfig.Db.Url,
		config.ServiceConfig.Db.Port,
		config.ServiceConfig.Db.DbName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.New(
			syslog.New(os.Stdout, "\r\n", syslog.LstdFlags), // io.writer
			logger.Config{
				SlowThreshold: time.Second,                                       // 慢 SQL 阈值
				LogLevel:      logger.LogLevel(config.ServiceConfig.Db.LogLevel), // 日志级别
			},
		),
	})
	if err != nil {
		log.Errorf("Failed to connect to the database:", err)
		os.Exit(1)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Errorf("Failed to get underlying sql.DB instance:", err)
		os.Exit(1)
	}

	// 设置连接池大小
	sqlDB.SetMaxIdleConns(config.ServiceConfig.Db.MaxIdleConns) // 空闲连接的最大数量
	sqlDB.SetMaxOpenConns(config.ServiceConfig.Db.MaxOpenConns) // 最大打开的连接数

	Conn = db

	log.Infof("===========数据库连接成功==============")
}

// 自动创建或更新数据库表结构
func migrate() {
	// 可以传多个, 逗号拼接
	Conn.AutoMigrate(&model.User{})
	log.Infof("===========数据库表结构检查完成==============")
}

// 数据库连接检查
func connectionCheck() {
	log.Infof("===========启动数据库连接检查任务==============")
	go func() {
		c := cron.New()
		_ = c.AddFunc("0/30 * * * * ?", func() {
			Db, _ := Conn.DB()
			// 如果ping正常，就返回
			if err := Db.Ping(); err != nil {
				// 检查数据库是否存在，如果不存在则创建
				createDB()
				// 连接数据库
				connectDB()
				// 自动创建或更新数据库表结构
				migrate()
			}
		})
	}()
}
