package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// gin 框架 函数参数指定为 *gin.Context
func sayHello(c *gin.Context){
	c.JSON(200, gin.H{
		"message": "Hello golang",
	})

}

func main(){
	// 创建一个默认的路由引擎
	r := gin.Default()
	// 指定用户使用GET请求访问/hello 时，执行 sayHello 函数
	r.GET("/hello", sayHello)

	// restful 风格路由
	r.GET("/book", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"method": "get",
		})
	})
	r.POST("/book", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"method": "post",
		})
	})
	r.PUT("/book", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"method": "put",
		})
	})
	r.DELETE("/book", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"method": "delete",
		})
	})

	// go 语言返回 json 数据
	r.GET("/json", func(context *gin.Context) {
		// 方法1 ：使用的是 map
		//data := map[string]interface{}{
		//	"name": "vincent",
		//	"message": "hello world",
		//	"age": 18,
		//}

		//data := gin.H{
		//	"name": "小王子",
		//	"message": "hello world",
		//	"age": 18,
		//}

		// 方法二： 使用结构体 灵活使用 tag 来对结构体字段做定制化操作
		type msg struct{ // 定义结构体字段
			 Name string `json:"name"` // 首字母小写 不可导出 可以打 tag
			 Message string `json:"message"`
			 Age int `json:"age"`
		}

		data := msg{ // 示例化结构体
			"小王子",
			"hello world",
			18,
		}

		context.JSON(http.StatusOK, data) // json 的序列化
		/*
		{
		    "name": "小王子",
		    "message": "hello world",
		    "age": 18
		}
		*/

	})
	// 启动服务
	r.Run(":9090")

}

