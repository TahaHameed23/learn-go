package main

import (
	"fmt"
	"math"
)

var c, python, java bool

func main() {
	fmt.Println("Hello, Go")
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))

	var i int
	fmt.Println(i, c, python, java)
}
