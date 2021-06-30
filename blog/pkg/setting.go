package setting

import (
	"fmt"
	"github.com/spf13/viper"
)

func Init() {
	viper.SetConfigFile("./conf/app.yaml")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	Cfg := viper.Get("db")
	fmt.Println(Cfg)
	// 监控配置文件变化
	viper.WatchConfig()
}
