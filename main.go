package main

import (
	"sinblog.cn/FunAnime-Server/router"
	"sinblog.cn/FunAnime-Server/util/logger"
)

func main() {
	// todo 开发环境 生产环境配置待完善
	initHandler("dev")
	err := router.NewRouter().Run(":8080")
	if err != nil {
		logger.Fatal("start_serve_failed", logger.Fields{"err": err})
		return
	}
}
