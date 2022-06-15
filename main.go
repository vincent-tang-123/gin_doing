package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 1 定义模型

type User struct {
	ID int64
	Name string `gorm:"default:'小王子'"`  // 用 tag 给模型设置 默认值
	Age int64
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

	// 2 创建表 自动迁移 (把结构体和数据表进行对应)  把模型与数据库中的表对应起来
	db.AutoMigrate(&User{})

	// 3 新增数据 创建结构体实例
	u := User{Name: "", Age: 38} // 在代码层面创建一个 user 对象
	fmt.Println(db.NewRecord(&u)) // 判断主键是否为空  true
	db.Debug().Create(&u)  // 创建一条数据 Debug 打印sql 日志
	fmt.Println(db.NewRecord(&u)) // 判断主键是否为空 false







}

