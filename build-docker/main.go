package main

import (
	"net/http"
)

func main() {
	//注册一个函数，响应某一个路由
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello CICD"))
	})

	//启动端口服务
	http.ListenAndServe(":8000", nil)


}