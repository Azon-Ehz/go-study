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
	MyName       string         `gorm:"column:name"` // A regular string field
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

	var user User
	var users User
	// Get first matched record
	db.Where("name = ?", "jinzhu").First(&user)
	db.Where(&User{MyName: "jinzhu"}).First(&user)
	// !=
	db.Where("name <> ?", "jinzhu").Find(&users) //不加limit 多条记录
	//IN
	db.Where("name IN ?", []string{"jinzhu", "jinzhu 2"}).Find(&users)
	//Like
	db.Where("name LIKE ?", "jin%").Find(&users)
	//AND
	db.Where("name = ? AND age >= ?", "jinzhu", "22").Find(&users)
	// Time
	lastWeek := time.Now().AddDate(0, 0, -7)
	db.Where("updated_at > ?", lastWeek).Find(&users)
	today := time.Now().Truncate(24 * time.Hour)
	//BETWEEN
	db.Where("created_at BETWEEN ? AND ?", lastWeek, today).Find(&users)
	db.Where(map[string]interface{}{"name": "jinzhu"}).First(&users)
	fmt.Println(users)

	//gorm查询方式条件有三种 1.string 2.struct 3.map
	//第一种最灵活(后面两种满足时不考虑) 第二种会屏蔽细节(比如0值)但好维护 第三种可以解决第二种的坑 可读性也更强
	//中高级开发必须要掌握sql 1.groupBy 2.Having子句 3.子查询
}
