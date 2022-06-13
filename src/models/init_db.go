package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//定义数据库表映射结构体

type User struct {
	ID            uint   `json:"user_id,omitempty" gorm:"primaryKey"` //omitempty表示将结构体对象转成json时，如果没有值则不需要赋默认值
	Name          string `json:"name"`
	Password      string `json:"password,omitempty"`
	Description   string
	Image         string
	FollowCount   int64 `json:"follow_count"`
	FollowerCount int64 `json:"follower_count"`
	IsFollow      bool  `json:"is_follow"`
}

func init() {
	//连接数据库
	dsn := "root:123456@(localhost:3306)/dousheng?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("数据库连接失败....")
		panic(err)
	}
	//自动前移(已经存在同样表时，判断表字段是否一致，如果不存在表，则自动创建表)
	//表名默认就是结构体的复数
	//启用自动迁移模式可以保持mysql表更新到最新。
	db.AutoMigrate(&User{})
}
func GetDB() *gorm.DB {
	dsn := "root:123456@(localhost:3306)/dousheng?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("数据库连接失败....")
		panic(err)
	}
	return db
}
