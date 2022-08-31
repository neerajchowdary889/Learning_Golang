package main

import "fmt"

func Forloops() {
	j := 0
	for i := 0; i <= 5; i++ {
		// For Loop
		if i == 4 {
			fmt.Println("Skip 4")
			continue
		}
		fmt.Println("For loop", i)
	}
	fmt.Println("________________End of For loop")
	for j <= 4 {
		fmt.Println("While loop", j)
		j++
	}
	fmt.Println("________________End of While loop")
	nums := []int{1, 2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum = sum + num
		fmt.Println(sum)
	}

}
