package main

import (
	"fmt"
	"time"
)

var d = "Maggie"

func main() {
	go count("SigmaCynide")
	count("DarkKnight")
}

func count(data string) {
	for i := 1; i < 15; i++ {
		fmt.Println("Nerd", data, i)
		time.Sleep(time.Millisecond * 500)
	}
}
