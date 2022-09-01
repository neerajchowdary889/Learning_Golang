package main

import "fmt"

type student struct {
	Name   string
	Rollno int32
	Course string
}

func maps() {
	var Map_1 = map[int]string{
		1: "Cat",
		2: "Dog",
		3: "Cow",
		4: "Tiger",
		5: "Liger",
	}
	Map_1[6] = "Lion"
	Map_1[0] = "Bug"

	fmt.Println("Map_1: \n", Map_1)

	delete(Map_1, 0)
	Map_1[1] = "CatsGoneWrong"

	fmt.Println("Map_1: \n", Map_1)

	fmt.Println(Map_1[5])

	for i := 1; i <= len(Map_1); i++ {
		temp := Map_1[i]
		fmt.Println(">>>", temp)
	}

	var Map_2 = make(map[int]string)

	Map_2[18] = "Virat"
	Map_2[45] = "Rohit"
	Map_2[1] = "KLRahul"
	Map_2[63] = "SKY"
	Map_2[21] = "DK"

	fmt.Println("Map_2: \n", Map_2)

	for jersey, player := range Map_2 {
		fmt.Println(">>>", jersey, player)
	}
}
func mapsstruct() {
	a1 := student{"Neeraj", 21067, "AIE"}
	a2 := student{"Ayyappa", 21084, "AIE"}
	a3 := student{"Subhash", 21036, "AIE"}
	a4 := student{"Manasa", 21057, "AIE"}
	a5 := student{"Akshaya", 21051, "AIE"}

	Studentmap := map[int]student{
		1: a1,
		2: a2,
		3: a3,
		4: a4,
		5: a5,
	}
	fmt.Println(Studentmap)
}
