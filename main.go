package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)
				// 响应					请求
func sayHello(w http.ResponseWriter, r *http.Request){
	//_, _ = fmt.Fprintln(w, "<h1>Hello Goland</h1>") // Fprintln 往 响应体中写入数据
	b, _ := ioutil.ReadFile("./hello.txt") // ioutil.ReadFile 读取文件 返回 二进制数据， 和错误
	_, _ = fmt.Fprintln(w, string(b)) // 打印数据
	t , _ := ioutil.ReadFile("./test.txt")
	_, _ = fmt.Fprintln(w, string(t))


}

func main(){
	http.HandleFunc("/hello", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil{
		fmt.Printf("http server failed , err:%v\n", err)
		return
	}

}

