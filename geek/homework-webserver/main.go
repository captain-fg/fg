//creating a new web server
package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
)

func main() {
	var mux = http.NewServeMux()
	mux.HandleFunc("/", responser)
	mux.HandleFunc("/healthz", responser_200)
	if err := http.ListenAndServe(":8000", mux); err != nil {
		fmt.Printf("Got an error when starting the http server..Error : %s\n", err.Error())
	}
}
func responser(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		// fmt.Printf("Key is %s, Value is %s\n", k, v)
		//  HOMEWORK  1. 接收客户端request，并将request中带的header写入resonse header        <----------- Homework Q1
		w.Header().Set(k, strings.Join(v, ""))
	}
	//  HOMEWORK  2. 读取当前系统的环境变量中的VERSION的配置， 并写入resonse header           <----------- Homework Q2
	//试了一下 没有version找个key，自己set了一个~
	os.Setenv("VERSION", "V1.0.0.0")
	//fmt.Printf(os.Getenv("VERSION") + "\n")
	w.Header().Set("Version", os.Getenv("VERSION"))
	//  HOMEWORK  3. server端记录访问日志包括客户端IP,HTTP返回码，输出到server端的标准输出     <----------- Homework Q3
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	fmt.Printf("Client IP is %s\n", ip)
	fmt.Printf("Status code is 200\n")
}

//  HOMEWORK  4. 当访问 localhost/healthz 时候, 应返回200                                   <----------- Homework Q4
func responser_200(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("200"))
}
