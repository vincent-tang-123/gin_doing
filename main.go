package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// querystring

func main(){
	r := gin.Default()

	// GET 请求 URL ？后面是 querystring 参数
	// key=value 格式 ，多个 key-value 用 & 符号连接
	// eg: http://127.0.0.1:9090/web?query=vincent&age=19
	r.GET("/web", func(context *gin.Context) {
		// 获取浏览器那边发请求携带的的 query string 参数
		name := context.Query("query") // 通过 Query 获取请求中的 querystring
		age := context.Query("age")
		//name := context.DefaultQuery("query", "somebody") // 取不到就用指定的默认值
		//name, ok := context.GetQuery("query") // 取不到 第二个参数就返回 false
		//if !ok{
		//	// 取不到
		//	name = "somebody"
		//}
		context.JSON(http.StatusOK, gin.H{
			"name": name,
			"age": age,
		})

	})

	r.Run(":9090")
}

