package conf

import (
	"github.com/spf13/viper"
	"sinblog.cn/FunAnime-Server/util/logger"
)

func Init(runType string) {
	viper.SetConfigType("toml")
	viper.SetConfigName(runType)
	viper.AddConfigPath("./conf/")

	err := viper.ReadInConfig()
	if err != nil {
		logger.Panic("Read Config File Failed", logger.Fields{"err": err})
	}
}
