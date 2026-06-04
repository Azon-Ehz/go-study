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
	result := db.Where("id = ?", 33333).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		fmt.Println("未找到")
	}
	user.MyName = "Zinon"
	user.Age = 25
	// 1. 通过save更新
	// 根据主键判断执行 INSERT 还是 UPDATE
	// 更新时会写入所有字段（包括零值）
	db.Save(&user) //save方法是一个集create和update为一体d函数
	// 2.通过update更新
	// 更新单个字段
	db.Model(&User{}).Update("name", "Ehz")
	//3. updates更新
	db.Model(&User{}).Select("name", "age").Where("name like ?", "Z%").Updates(&User{MyName: "Zinon", Age: 25}) // 指定更新哪些字段
	db.Model(&User{}).Omit("name", "age").Where("name like ?", "").Updates(&User{MyName: "Zinon", Age: 25})     //声明不更新的字段
	db.Model(&User{}).Select("name", "age").Where("name like ?", "").Updates(&User{MyName: "Zinon", Age: 25})   //声明可以写入0值的字段

}
func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	if u.Age == 25 {
		if u.Age == 25 {
			return errors.New("你已经25岁了，不允许更新")
		}
	}
	return nil
}
