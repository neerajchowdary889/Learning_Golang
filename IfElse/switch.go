package main

import "fmt"

func switches() {
	fmt.Println("Enter between 1 to 4")
	var inp int
	fmt.Scan(&inp)

	switch inp {
	case 1:
		fmt.Println(1, "Switch case")
	case 2:
		fmt.Println(2, "Switch case")
	case 3:
		fmt.Println(3, "Switch case")
	case 4:
		fmt.Println(4, "Switch case")
	default:
		fmt.Println(inp, "Wrong Input")
	}

}
