package main

import (
	"context"
	"database/sql"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Product struct {
	Code  sql.NullString
	Price uint
	gorm.Model
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

	// 自动建表
	_ = db.AutoMigrate(&Product{}) //此处应该返回true

	ctx := context.Background()
	// 创建
	err = gorm.G[Product](db).Create(ctx, &Product{Code: sql.NullString{String: "D42", Valid: false}, Price: 100})
	// 查询
	product, err := gorm.G[Product](db).Where("id = ?", 1).First(ctx) // 查找对应主键的产品
	_, err = gorm.G[Product](db).Where("code = ?", "D42").Find(ctx)   // 查找 code 为 D42 的所有产品
	// 更新 - 将产品价格更新为 200
	_, _ = gorm.G[Product](db).Where("id = ?", product.ID).Update(ctx, "Price", 200)
	// 更新 - 更新多个字段
	_, _ = gorm.G[Product](db).Where("id = ?", product.ID).Updates(ctx, Product{Code: sql.NullString{String: "D43", Valid: false}, Price: 100})
	// 删除 - 删除产品
	_, _ = gorm.G[Product](db).Where("id = ?", product.ID).Delete(ctx)
}
