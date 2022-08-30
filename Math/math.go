package main

import (
	"fmt"
	"math"
)

func cal() {
	// Square root doesn't take int as a input, it only take float
	var Num float64 = 49
	sqroot := math.Sqrt(Num)
	var round float64 = math.Round(3.1238249884834804)
	ceil := math.Ceil(3.8899)
	floor := math.Floor(3.8899)
	fmt.Println("Square root of 49:", sqroot)
	fmt.Printf("Round of 3.1238249884834804 is %.2f \n", round)
	fmt.Printf("Ceil of 3.8899 is %.2f \n", ceil)
	fmt.Printf("Floor of 3.8899 is %.2f \n", floor)

}

func main() {
	cal()
}
