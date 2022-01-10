package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/pprof"
	"os"
	"strings"
)

func index(w http.ResponseWriter, r *http.Request) {

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

	w.Write([]byte("welcome to native cloud"))
}

func getCurrentIP(r *http.Request) string {
	ip := r.Header.Get("X-Real-IP")
	if ip == "" {
		ip = strings.Split(r.RemoteAddr, ":")[0]
	}
	return ip
}

// ClientIP 尽最大努力实现获取客户端 IP 的算法。
//解析 X-Real-IP 和 X-Forwarded-For 以便于反向代理（nginx 或 haproxy）可以正常工作。
func ClientIP(r *http.Request) string {
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}
	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}
	return ""
}

func healthz(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "I am ok")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/", index)
	mux.HandleFunc("/healthz", healthz)
	// 启动 HTTP 服务，并监听端口号，开始监听，处理请求，返回响应
	err := http.ListenAndServe("localhost:8080", mux)
	if err != nil {
		log.Fatalf("start http server failed, error: %s\n", err.Error())
	}
}
