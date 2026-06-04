package main

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// `User` 属于 `Company`，`CompanyID` 是外键
type User struct {
	gorm.Model
	Name      string
	CompanyID int //目的是新建数据表
	Company   Company
}

type Company struct {
	ID   int
	Name string
}

func main() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:123456@tcp(localhost:3307)/gorm?charset=utf8mb4&parseTime=True&loc=Local"

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // 慢查询阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 查询结果为nil算不算错误
			ParameterizedQueries:      false,       // 是否隐藏入参
			Colorful:                  true,        // 是否彩色打印
		},
	)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         255,   // 所有 string 字段默认最大长度
		DisableDatetimePrecision:  true,  // 关闭 datetime 的“毫秒/微秒精度”
		DontSupportRenameIndex:    false, // 如果数据库不支持改索引名字，就用删除重建代替”
		DontSupportRenameColumn:   false, // 如果不能 rename column，就用 modify/change 方式替代”
		SkipInitializeWithVersion: false, // 是否跳过 MySQL 版本检测
	}), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&User{})

	db.Create(&User{
		Name: "Ehz",
		Company: Company{
			ID:   1,
			Name: "isens",
		},
	})
}
