package main

import (
	"fmt"
	"math"
)

func flow() {
	sum := 1
	//The init and post statements are optional.
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
func pow(x, n, lim float64) float64 {
	//Variables declared by the statement are only in scope until the end of the `if`
	if v := math.Pow(x, n); v < lim {
		return v
	}

	// can't use v here
	//return lim.v
	return lim
}
