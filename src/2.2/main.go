package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>welcome to native cloud</h1>"))

	//fmt.Println("request header:", r.Header)
	for k, v := range r.Header {
		//fmt.Println(w, k+": ", v)
		for _, vv := range v {
			w.Header().Set(k, vv)
		}
	}

	//读取当前系统的环境变量中的 VERSION 配置
	var VERSION string
	VERSION = os.Getenv("VERSION")
	//fmt.Println(VERSION)
	w.Header().Set("VERSION:", VERSION)

	clientIP := getCurrentIP(r)
	httpCode := http.StatusOK
	log.Printf("clientip: %s \n", clientIP)
	log.Printf("status code: %d \n", httpCode)
}

func getCurrentIP(r *http.Request) string {
	ip := r.Header.Get("X-Real-IP")
	if ip == "" {
		ip = strings.Split(r.RemoteAddr, ":")[0]
	}
	return ip
}

func healthz(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "I am ok")
}

func main() {
	// 注册路由和路由函数，将url规则与处理器函数绑定做一个map映射存起来，并且会实现ServeHTTP方法，使处理器函数变成Handler函数
	http.HandleFunc("/", index)
	http.HandleFunc("/healthz", healthz)
	fmt.Println("服务器已经启动，请在浏览器地址栏中输入 http://localhost:8900/")
	// 启动 HTTP 服务，并监听端口号，开始监听，处理请求，返回响应
	err := http.ListenAndServe(":8900", nil)
	fmt.Println("监听之后")
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
