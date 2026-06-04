package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type User struct {
	ID           uint           // Standard field for the primary key
	Name         string         // A regular string field
	Email        *string        // A pointer to a string, allowing for null values
	Age          uint8          // An unsigned 8-bit integer
	Birthday     *time.Time     // A pointer to time.Time, can be null
	MemberNumber sql.NullString // Uses sql.NullString to handle nullable strings
	ActivatedAt  sql.NullTime   // Uses sql.NullTime for nullable time fields
	CreatedAt    time.Time      // Automatically managed by GORM for creation time
	UpdatedAt    time.Time      // Automatically managed by GORM for update time
	ignored      string         // fields that aren't exported are ignored
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

	var users User
	result := db.First(&users)
	fmt.Println(users.ID)
	errors.Is(result.Error, gorm.ErrRecordNotFound)

	result = db.Take(&users)
	fmt.Println(users.ID)

	result = db.Last(&users)
	fmt.Println(users.ID)

	result = db.First(&users, 10)
	fmt.Println(users.ID)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		fmt.Println("未找到")
	}

	result = db.First(&users, "12")
	fmt.Println(users.ID)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		fmt.Println("未找到")
	}

	result = db.First(&users, []int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	fmt.Println(users.ID)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		fmt.Println("未找到")
	}

	var user []User
	result = db.Find(&user, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}) // where in
	result = db.Find(&user)                                   // no where
	fmt.Printf("总记录数: %d\n", result.RowsAffected)

	for _, u := range user {
		fmt.Println(u.ID)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			fmt.Println("未找到")
		}
	}

}
