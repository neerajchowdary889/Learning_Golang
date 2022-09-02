package main

import (
	"fmt"
	"log"
)

func main() {
	var input int
	fmt.Println("your fav int type karo: ")
	fmt.Scanln(&input)
	var example = func(input int) {
		defer cleanup()
		//In Golang, the defer keyword is used to delay the execution of a function
		//or a statement until the nearby function returns.
		if input >= 5 {
			log.Panicln("Ohh my god panicking right now.....")
		} else {
			log.SetFlags(log.Ldate & log.Ltime)
			log.Println("No issue what so ever.")
		}
	}
	example(input)
}
func cleanup() {
	r := recover()
	if r != nil {
		fmt.Println("Cleaned up using recover; Input here: ")
		var input2 int
		fmt.Scan(&input2)
		fmt.Println("Recovery done and your input is", input2)
	} else {
		log.Fatal("byee broo chill")
		// fatal is like panic but it dont have recover it just exit the programme.
		// fatal is also called stopper.(my own words)
	}
}
