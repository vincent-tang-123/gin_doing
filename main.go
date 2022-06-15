package main

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

// 定义模型

type User struct {
	// 内嵌 gorm.Model
	gorm.Model // 继承 id ， 创建时间， 更新时间 删除时间字段
	Name string
	Age sql.NullInt64
	Birthday *time.Time
	Email string `gorm:"type:varchar(100);unique_index"` // 唯一约束
	Role string `gorm:"size:255"` // 设置字段大小为 255
	MemberNumber *string `gorm:"unique;not null"` // 设置会员号 唯一且不为空
	Num int `gorm:"AUTO_INCREMENT"` // 设置 字段为自增类型
	Address string `gorm:"index:addr"` // 给 address 字段创建名为 addr 的索引
	IgnoreMe int `gorm:"-"` // 忽略本字段
}

// 使用 AnimalID 作为主键 默认生成的表名复数 Animals

type Animal struct {
	AnimalID int64 `gorm:"primary_key;column:beast_id"` // gorm 自定义列名
	Name 	 string `gorm:"column:day_of_beast"` // gorm 自定义列名
	Age      int64 `gorm:"column:age_of_the_beast"`
}

// 重命名表名 将 Animal 表名转换成 vincent

func (Animal)TableName() string{
	return "vincent"
}

func main(){
	// 连接 Mysql 数据库
	db, err := gorm.Open("mysql",
		"root:521314td@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err !=nil {
		panic(err)
	}
	defer db.Close()
	db.SingularTable(true) // 禁用表名复数

	// 创建表 自动迁移 (把结构体和数据表进行对应)
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Animal{})

	// 使用 User 结构体创建表名叫 xiaowangzi 的表
	db.Table("xiaowangzi").CreateTable(&User{})





}

