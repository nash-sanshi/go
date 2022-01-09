package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// HelloServer 函数实现了处理器的签名，所以这是一个处理器函数
func HelloServer(w http.ResponseWriter, r *http.Request) {

	fmt.Println("request header:", r.Header)

	for k, v := range r.Header {

		fmt.Fprintln(w, k+": ", v)
	}

	//读取当前系统的环境变量中的 VERSION 配置
	var VERSION string
	VERSION = os.Getenv("VERSION")
	fmt.Println(VERSION)
	w.Header().Add("VERISON", VERSION)
}

func main() {
	// 注册路由和路由函数，将url规则与处理器函数绑定做一个map映射存起来，并且会实现ServeHTTP方法，使处理器函数变成Handler函数
	http.HandleFunc("/healthz", HelloServer)
	fmt.Println("服务器已经启动，请在浏览器地址栏中输入 http://localhost:8900/healthz")
	// 启动 HTTP 服务，并监听端口号，开始监听，处理请求，返回响应
	err := http.ListenAndServe(":8900", nil)
	fmt.Println("监听之后")
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
