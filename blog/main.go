package main

import (
	"blog/model"
	setting "blog/pkg"
	"blog/router"
)

func init() {
	setting.Init()
	model.Init()
}

func main() {
	r := router.InitRouter()
	r.Run()
}
