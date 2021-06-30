package main

import "blog/router"

func main() {
	r := router.InitRouter()
	r.Run()
}
