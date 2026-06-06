package main

import (
	"database/sql"
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
	// 自动建表
	_ = db.AutoMigrate(&User{}) //此处应该返回true

	user := User{
		Name:     "John Doe",
		Email:    nil,
		Age:      30,
		Birthday: nil,
	}
	result := db.Create(&user)
	user = User{
		Name:     "Li Doe",
		Email:    nil,
		Age:      24,
		Birthday: nil,
	}
	db.Select("Name", "Age", "CreatedAt").Create(&user)

	user = User{
		Name:     "Jiang Doe",
		Email:    nil,
		Age:      24,
		Birthday: nil,
	}
	db.Omit("Name", "Age", "CreatedAt").Create(&user)
	// INSERT INTO `users` (`name`,`age`,`created_at`) VALUES ("jinzhu", 18, "2020-07-04 11:05:21.775")

	fmt.Println(user.ID)             //// 返回插入数据的主键
	fmt.Println(result.Error)        //返回 error
	fmt.Println(result.RowsAffected) //返回插入记录的条数
	//db.Model(&User{ID: 1}).Update("name", "Zinon2")

	//update可以更新零值 但updates则不行
	//empty := ""
	//db.Model(&User{ID: 1}).Updates(User{Email: &empty})
	//解决仅更新非零值的方法有两种 1.将string类型改为指针 *string 或者 sql.nullXXX(数据类型)
}
