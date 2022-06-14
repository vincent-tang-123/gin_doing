package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// UserInfo --> 数据表
type UserInfo struct {
	ID uint
	Name string
	Gender string
	Hobby string
}

func main(){
	// 连接 Mysql 数据库
	db, err := gorm.Open("mysql",
		"root:521314td@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err !=nil {
		panic(err)
	}
	defer db.Close()

	// 创建表 自动迁移 (把结构体和数据表进行对应)
	//db.AutoMigrate(&UserInfo{})

	//// 创建数据行
	//u1 := UserInfo{1, "提米", "男", "play"}
	//db.Create(&u1)

	// 查询
	var u UserInfo
	db.First(&u) // 查询表中的第一条数据 保存到 u 中
	fmt.Printf("%#v\n", u) //main.UserInfo{ID:0x1, Name:"提米", Gender:"男", Hobby:"play"}
	fmt.Println(u.ID, u.Name, u.Gender, u.Hobby)

	// 更新
	db.Model(&u).Update("hobby", "篮球")
	fmt.Println(u.ID, u.Name, u.Gender, u.Hobby) // 1 提米 男 篮球

	// 删除
	db.Delete(&u)




}

