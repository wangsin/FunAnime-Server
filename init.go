package main

import (
	"sinblog.cn/FunAnime-Server/router"
	"sinblog.cn/FunAnime-Server/util/conf"
)

func initHandler() {
	conf.Init("dev")
	router.Init()
}
