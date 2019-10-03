package conf

import (
	"fmt"
	"github.com/spf13/viper"
)

func Init(runType string) {
	viper.SetConfigType("toml")
	viper.SetConfigName(runType)
	viper.AddConfigPath("./conf/")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Read Config File Failed, Error: %s\n", err.Error()))
	}
}
