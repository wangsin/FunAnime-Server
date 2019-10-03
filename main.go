package main

import (
	"fmt"
	"sinblog.cn/FunAnime-Server/router"
)

func main() {
	initHandler()
	err := router.NewRouter().Run(":8080")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
