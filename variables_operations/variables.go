package main

import "fmt"

func main() {
	variables()
	operations()
	Forloops()
}

func variables() {
	var Name = "Neeraj"
	var Age = 18
	fmt.Println("Name: " + Name)
	fmt.Println("Age: ", Age)
	fmt.Println("_________________end of Variables")
}

func operations() {
	num1, num2, num3 := 21, 31, 41
	const num = 15
	sum := num1 + num2 + num3
	sub := sum - num
	mult := num1 * num2 * num3
	division := sum / num
	Mod := sum % num
	fmt.Println("Sum", sum)
	fmt.Println("sub", sub)
	fmt.Println("Multiply", mult)
	fmt.Println("Division", division)
	fmt.Println("Mod", Mod)
	fmt.Println("_________________end of Operations")

}
