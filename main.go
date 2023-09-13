package main

import "gin-demo/routers"

func main() {
	r := routers.Router()
	r.Run(":8888") // 监听并在 0.0.0.0:8080 上启动服务
}
