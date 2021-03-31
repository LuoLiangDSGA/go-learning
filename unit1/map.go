package main

import (
	"fmt"
)

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"jay":     {1.0, 2.0},
	"fantasy": {11.0, 23.0},
}

func map1() {
	fmt.Println(m)
	hash := make(map[string]string)
	fmt.Println(hash)
	hash["ss"] = "ss"
	hash["aa"] = "ssd"
	hash["wer"] = "asdf"
	fmt.Println(hash)
	delete(hash, "ss")
	fmt.Println(hash)
	v, ok := hash["wer"]
	fmt.Println("the value:", v, "present?", ok)
}

func main() {
	map1()
}
