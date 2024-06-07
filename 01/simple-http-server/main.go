package main

import "net/http"

func main() {
	// 通过 http.HandleFunc 为 web 服务设置一个处理函数
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World"))
	})
	// 通过 http 包提供的 ListenAndServe 函数，建立起一个 HTTP 服务，这个服务监听本地的 8080 端口
	http.ListenAndServe(":8080", nil)
}
