package conf

import (
	"FunAnime-Server/cache"
	"FunAnime-Server/model"
	"FunAnime-Server/util"
	"os"

	"github.com/joho/godotenv"
)

// Init 初始化配置项
func Init() {
	// 从本地读取环境变量
	godotenv.Load()

	// 读取翻译文件
	if err := LoadLocales("conf/locales/zh-cn.yaml"); err != nil {
		util.Errnof(util.LogClassic, -100003, "can_not_load_translate_file")
	}

	// 连接数据库
	model.Database(os.Getenv("MYSQL_DSN"))
	cache.Redis()
}
