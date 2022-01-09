package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("test", "welcome to native cloud")

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
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/healthz", healthz)
	// 启动 HTTP 服务，并监听端口号，开始监听，处理请求，返回响应
	if err := http.ListenAndServe("localhost:8080", mux); err != nil {
		log.Fatalf("start http server failed, error: %s\n", err.Error())
	}
}
