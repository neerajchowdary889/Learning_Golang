package main

import "fmt"

func main() {
	// num := 2
	fmt.Print("Type out your input here: ")
	var inp int
	fmt.Scan(&inp)
	if inp <= 5 {
		fmt.Println("Hello")
	} else {
		fmt.Println("bye")
	}
}