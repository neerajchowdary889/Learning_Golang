package main

import "fmt"

func add(Num1 int, Num2 int) int {
	sum := Num1 + Num2
	return sum
}

func cal(Num3 int, Num4 int) (int, int) {
	Mul := Num3 * Num4
	Mod := Num3 % Num4
	return Mul, Mod
}

func vals() (int, int) {
	// Multiple return function
	return 9, 12
}

func SumAll(num ...int) {
	// variadic function
	// dynamic in nature
	// num is a array
	fmt.Println(num, " ")
	temp := 0
	for i := 0; i < len(num); i++ {
		temp += num[i]
	}
	fmt.Println(temp, "SumAll")
}

var Name string

func init() {
	// First init will be initilized at any cost...
	fmt.Println("init function...")
	Name = "Neeraj Chowadary"
}

func main() {
	Add := add(21, 41)
	fmt.Println(Add)
	Res1, Res2 := cal(31, 51)
	fmt.Println(Res1, Res2)
	a, b := vals()
	fmt.Println(a, b, "a and b")
	SumAll(1, 2, 3, 4, 5)
	// %T allow us to print the type of variable.
	// %v allow us to print the Int value of variable and even String value.
	fmt.Printf("init function testing-- %s\n", Name)
}
