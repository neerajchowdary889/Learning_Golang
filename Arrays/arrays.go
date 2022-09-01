package main

import (
	"fmt"
)

func arrays() {
	my_slice := []string{
		"Winnie",
		"Rick",
		"Anna",
		"Elsa",
		"Olaf",
	}
	twoslice := []string{"Wine", "rickey", "sugar"}
	var arr [4]string
	arr[0] = "Tirupati"
	arr[1] = "Kollam"
	arr[2] = "Kayankulam"
	my_slice = append(my_slice, "foss")
	my_slice = append(my_slice, twoslice...)
	fmt.Println(arr)
	fmt.Println(my_slice)
}

func main() {
	arrays()
	fmt.Println("__________________arrays")
	maps()
	fmt.Println("__________________maps")
	mapsstruct()
	fmt.Println("__________________structmaps")
}
