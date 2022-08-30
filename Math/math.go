package main

import (
	"fmt"
	"math"
)

func cal() {
	// Square root doesn't take int as a input, it only take float
	var Num float64 = 49
	sqroot := math.Sqrt(Num)
	fmt.Println("Square root of 49:", sqroot)
}

func main() {
	cal()
}
