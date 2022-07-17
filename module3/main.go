package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

const PNameVersion = "VERSION"

func main() {
	//注册路由处理器
        log.Printf("main---------start to handler http request")
	http.HandleFunc("/", reqHandler)
        log.Printf("111")
	http.HandleFunc("/healthz", health)
        log.Printf("222")
	//监听80端口,使用GO内置的HTTP Server
	err := http.ListenAndServe(":80", nil)
        log.Printf("333")
	if err != nil {
		log.Printf("startup http server failed, %v", err)
	}
        log.Printf("end main ---------start to handler http request")

}

/**
1.接收客户端 request，并将 request 中带的 header 写入 response header
2.读取当前系统的环境变量中的 VERSION 配置，并写入 response header
3.Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
4.当访问 localhost/healthz 时，应返回200
*/
func reqHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("start to handler http request")
	writeBackReqHeader(w, r)
	writeVersionHeader(w, r)
	logRemoteIPAndStatus(w, r)
	io.WriteString(w, "simple http server for fun!\n")
	log.Printf("handle request done")
}

//1.接收客户端 request，并将 request 中带的 header 写入 response header
func writeBackReqHeader(w http.ResponseWriter, r *http.Request) {
	log.Printf("fun1: write back request header to response")
	for k, v := range r.Header {
		w.Header().Set(k, strings.Join(v, ""))
	}
}

//2.读取当前系统的环境变量中的 VERSION 配置，并写入 response header
func writeVersionHeader(w http.ResponseWriter, r *http.Request) {
	versionValue := os.Getenv(PNameVersion)
	log.Printf("env param: version = %s ", versionValue)
	w.Header().Set(PNameVersion, versionValue)
}

//3.Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
func logRemoteIPAndStatus(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	log.Printf("remote addr:%s, resp status code %d", r.RemoteAddr, http.StatusOK)
}

//4.当访问 localhost/healthz 时，应返回200
func health(w http.ResponseWriter, r *http.Request) {
	log.Printf("handle request: /healthz, and response 200_OK")
	reqHandler(w, r)
	w.WriteHeader(http.StatusOK)
        w.Header().Set("Custom-Header", "Awesome")
	io.WriteString(w, "200")
}
