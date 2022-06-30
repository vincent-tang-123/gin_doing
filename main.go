package main

import "fmt"

//import (
//	"github.com/jinzhu/gorm"
//	_ "github.com/jinzhu/gorm/dialects/mysql"
//)
//
//// 1 定义模型
//
//type User struct {
//	gorm.Model
//	Name string `gorm:"default:'小王子'"`  // 用 tag 给模型设置 默认值
//	Age int64
//}
//
//func (User)TableName() string{
//	return "user"
//}
//
//func main(){
//	// 连接 Mysql 数据库
//	db, err := gorm.Open("mysql",
//		"root:521314td@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
//	if err !=nil {
//		panic(err)
//	}
//	defer db.Close()
//	db.SingularTable(true) // 禁用表名复数
//
//	// 2 创建表 自动迁移 (把结构体和数据表进行对应)  把模型与数据库中的表对应起来
//	db.AutoMigrate(&User{})
//
//	// 3 新增数据 创建结构体实例
//	//u1 := User{Name: "vincent", Age: 38} // 在代码层面创建一个 user 对象
//	//db.Create(&u1)
//	//u2 := User{Name: "qimi", Age: 67}
//	//db.Create(&u2)
//
//	// 4 查询
//	//var user User // 声明模型结构体类型变量 user
//	//user := new(User) // new 和 make 的区别, new 返回一个指针， make 只能对 map channel slice 进行初始化
//	//db.First(&user) // 根据 ID 主键主键 取第一条
//	//fmt.Printf("user:%#v\n", user) // user:main.User{Model:gorm.Model{ID:0x1, CreatedAt:time.Date(2022, time.June, 15, 17, 48, 21, 0, time.Local), UpdatedAt:time.Date(2022, time.June, 15, 17, 48, 21, 0, time.Local), DeletedAt:<nil>}, Name:"vincent", Age:38}
//	//
//	//var users []User
//	//db.Debug().Find(&users) // SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL
//	//fmt.Printf("users:%#v\n", users)
//
//	// FirstOrInit
//	//var user User
//	//db.Debug().FirstOrInit(&user, User{Name: "echo"}) // SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL AND ((`user`.`name` = 'echo')) ORDER BY `user`.`id` ASC LIMIT 1
//	//
//	//// Attrs 如果找 echo1 找不到 就会常见一天 echo1 的数据 且将 Age 字段属性赋值为 88
//	////db.Debug().Attrs(User{Age: 88}).FirstOrInit(&user, User{Name: "echo1"}) // SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL AND ((`user`.`name` = 'echo1')) ORDER BY `user`.`id` ASC LIMIT 1
//	//
//	//// Assign 无论 echo1 是否找的到 都会为该数据的 Age 字段赋一个值
//	//
//	//db.Debug().Assign(User{Age: 11}).FirstOrInit(&user, User{Name: "echo1"}) // SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL AND ((`user`.`name` = 'echo1')) ORDER BY `user`.`id` ASC LIMIT 1
//	//fmt.Printf("user:%#v\n", user)
//
//	var user User
//	db.First(&user)
//	//user.Name = "suiguohui"
//	//user.Age = 17
//	//db.Debug().Save(&user) // 默认会修改所有字段  UPDATE `user` SET `created_at` = '2022-06-15 17:48:21',
//	// `updated_at` = '2022-06-16 21:00:11', `deleted_at` = NULL, `name` = 'suiguohui', `age` = 17  WHERE `user`.`deleted_at` IS NULL AND `user`.`id` = 1
//
//	// 更新单独字段
//	db.Debug().Model(&user).Update("Name", "vincent111") //  UPDATE `user` SET `name` = 'vincent111', `updated_at` = '2022-06-16 21:08:21'  WHERE `user`.`deleted_at` IS NULL AND `user`.`id` = 1
//
//	// 更新多个字段
//	db.Debug().Model(&user).Updates(map[string]interface{}{"name": "hello", "age": 18}) // UPDATE `user` SET `age` = 18, `name` = 'hello', `updated_at` = '2022-06-16 21:12:40'  WHERE `user`.`deleted_at` IS NULL AND `user`.`id` = 1
//
//	db.Debug().Model(&user).Updates(map[string]interface{}{"name": "", "age": 0}) // UPDATE `user` SET `age` = 0, `name` = '', `updated_at` = '2022-06-16 21:29:48'  WHERE `user`.`deleted_at` IS NULL AND `user`.`id` = 1
//
//	// 去掉钩子函数影响 只更新 age 字段值
//	db.Debug().Model(&user).UpdateColumn("age", 100) // UPDATE `user` SET `age` = 100  WHERE `user`.`deleted_at` IS NULL AND `user`.`id` = 1
//
//	// 让 users 表中所有用户的年龄在原来的基础上 + 2
//	db.Debug().Model(&User{}).Update("age", gorm.Expr("age+?", 2)) // UPDATE `user` SET `age` = age+2, `updated_at` = '2022-06-16 21:50:02'  WHERE `user`.`deleted_at` IS NULL
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//}


func main(){
	var a int = 3
	fmt.Println(a) // 3
	fmt.Println(&a) // 0xc0000180a8
	fmt.Println(*&a) // 3

	b :=new(int) // 接受一个类型 返回一个类型的指针
	fmt.Println(b) // 0xc0000b2018
	fmt.Println(*b) // 0

}
