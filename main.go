package main

import (
	"go_web/route"
	"log"
	"net/http"
)

const port = ":8083"

func main() {
	// 静态文件服务器
	http.Handle("/", http.FileServer(http.Dir("html")))
	// 注册路由
	route.RegisterRoute()
	//http://127.0.0.1:8083/login.html
	println("web服务启动成功,在浏览器中输入ip或域名+端口号形式访问，如:https://127.0.0.1:8083/login.html")
	//util.TestGenerateSnowFlakeId()
	// 监听端口
	/*err := http.ListenAndServe(port, nil)*/
	err := http.ListenAndServeTLS(port, "server.crt",
		"server.key", nil)
	if err != nil {
		log.Printf("web服务 error: %v", err.Error())
	}
}
