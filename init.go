package main

import (
	"sinblog.cn/FunAnime-Server/cache"
	"sinblog.cn/FunAnime-Server/model"
	"sinblog.cn/FunAnime-Server/util/conf"
)

func initHandler() {
	conf.Init("dev")
	model.DatabaseInit()
	cache.Redis()
}
