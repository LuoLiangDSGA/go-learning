package main

import (
	"fmt"
	"math/rand"
	"math"
)

const c = 3.1

func main() {
    rand.Seed(100)
    fmt.Println("my favorite number is:", rand.Intn(10))
	fmt.Println("Hello,World!")
    fmt.Println(math.Pi)
    fmt.Println(c)
}
