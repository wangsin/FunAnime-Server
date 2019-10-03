package main

import "sinblog.cn/FunAnime-Server/router"

func main() {
	initHandler()
	_ = router.NewRouter().Run(":8080")
}
