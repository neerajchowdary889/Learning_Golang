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
	return 9, 12
}

func main() {
	Add := add(21, 41)
	fmt.Println(Add)
	Res1, Res2 := cal(31, 51)
	fmt.Println(Res1, Res2)
	a, b := vals()
	fmt.Println(a, b, "a and b")

}
