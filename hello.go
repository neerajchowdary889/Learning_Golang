package main

import (
	"Learning_Golang/testpackage"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Hello World")
	fmt.Print("HelloWorld\n")
	strings := testpackage.Strings
	fmt.Println(strings)
	testpackage.Packagess()
	//convert int to string
	g := strconv.FormatInt(int64(78), 10)
	p, _ := strconv.ParseInt("123", 10, 64)
	fmt.Printf("%T\n", g)
	fmt.Printf("%T", p)
}
