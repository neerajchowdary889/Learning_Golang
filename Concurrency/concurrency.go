package main

import (
	"fmt"
	"time"
)

var d = "Maggie"

func main() {
	go count("SigmaCynide")
	count("DarkKnight")

	c := make(chan string)

	go channel("Valorent", c)

	//for {
	//	msg, open := <-c
	//	if open {
	//		fmt.Println(msg)
	//	} else {
	//		break
	//	}
	//}
	for msg := range c {
		fmt.Println(msg)
	}
	Ch1 := make(chan string)
	Ch2 := make(chan string)
	go func() {
		for {
			Ch1 <- "Every 5 seconds"
			time.Sleep(time.Millisecond * 500)
		}
	}()
	go func() {
		for {
			Ch2 <- "Every 2 seconds"
			time.Sleep(time.Millisecond * 200)
		}
	}()

	for {
		select {
		case msg1 := <-Ch1:
			fmt.Println(msg1)
		case msg2 := <-Ch2:
			fmt.Println(msg2)
		}
	}
	// if we dont use select then Ch2 will wait until 5secs done for Ch1 then
	// code become that both Ch1 and Ch2 wait for 5secs. when we use "select" then
	// it automatically select which case is ready and print those, if Ch2 is ready
	// then it will print Ch2 by not waiting for Ch1 to finish its 5secs...
}
func count(data string) {
	for i := 1; i < 15; i++ {
		fmt.Println("Nerd", data, i)
		time.Sleep(time.Millisecond * 300)
	}
}
