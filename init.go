package main

import (
	"sinblog.cn/FunAnime-Server/util/conf"
	"sinblog.cn/FunAnime-Server/util/logger"
)

func initHandler(runType string) {
	logger.Init()
	conf.Init(runType)
	//model.DatabaseInit()
	//cache.Redis()
}
