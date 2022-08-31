package main

import "fmt"

var Name string
var Age int
var Id int32
var email string
var Midmarks float64
var Endmarks float64
var marks float64

func init() {
	fmt.Println("Enter your Name: ")
	fmt.Scan(&Name)

	fmt.Println("Enter your Age:")
	fmt.Scan(&Age)

	fmt.Println("Enter your ID Number:")
	fmt.Scan(&Id)

	fmt.Println("Enter your Email:")
	fmt.Scan(&email)

	fmt.Println("Enter your Midmarks:")
	fmt.Scan(&Midmarks)

	fmt.Println("Enter your Endmarks:")
	fmt.Scan(&Endmarks)

	marks = (Midmarks + Endmarks) / 2
}
func main() {
	fmt.Println("___________________________")
	fmt.Printf("Confirm details:\n%v\n%v\n%v\n%v\nMids: %v\nEnd: %v\nAverage: %v", Name, Age, Id, email, Midmarks, Endmarks, marks)
}
