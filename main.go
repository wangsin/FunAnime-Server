package main

import (
	"sinblog.cn/FunAnime-Server/router"
	barrage "sinblog.cn/FunAnime-Server/service/websocket"
	"sinblog.cn/FunAnime-Server/util/logger"
)

func main() {
	initHandler("dev")
	// websocket服务 监听8090
	go barrage.Main()
	err := router.NewRouter().Run(":8088")
	if err != nil {
		logger.Fatal("start_serve_failed", logger.Fields{"err": err})
		return
	}
}
