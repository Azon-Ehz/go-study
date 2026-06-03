package main

import (
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
	dsn := "root:123456@tcp(localhost:3307)/mysql?charset=utf8mb4&parseTime=True&loc=Local"

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  true,        // Disable color
		},
	)

	// Globally mode
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	// 自动建表
	_ = db.AutoMigrate(&Product{}) //此处应该返回true

	// 创建
	err = gorm.G[Product](db).Create(ctx, &Product{Code: "D42", Price: 100})
	// 查询
	product, err := gorm.G[Product](db).Where("id = ?", 1).First(ctx) // 查找对应主键的产品
	_, err := gorm.G[Product](db).Where("code = ?", "D42").Find(ctx)  // 查找 code 为 D42 的所有产品
	// 更新 - 将产品价格更新为 200
	_, _ = gorm.G[Product](db).Where("id = ?", product.ID).Update(ctx, "Price", 200)
	// 更新 - 更新多个字段
	_, _ = gorm.G[Product](db).Where("id = ?", product.ID).Updates(ctx, Product{Code: sql.NullString{String: "", Valid: true}, Price: 100})
	// 删除 - 删除产品
	_, _ = gorm.G[Product](db).Where("id = ?", product.ID).Delete(ctx)
}
