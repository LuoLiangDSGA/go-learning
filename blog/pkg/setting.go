package setting

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

var (
	PageSize     int
	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
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
	loadApp()
	loadServer()
}

func loadServer() {
	HTTPPort = viper.GetInt("http.port")
	ReadTimeout = viper.GetDuration("http.read_timeout")
	WriteTimeout = viper.GetDuration("http.write_timeout")
}

func loadApp() {
	PageSize = viper.GetInt("app.pageSize")
}
