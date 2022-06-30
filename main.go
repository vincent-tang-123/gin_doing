package main

import "github.com/gin-gonic/gin"

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

	// 启动服务
	r.Run(":9090")


	//
	//
	//

}



