package main

import (
	"fmt"
	"time"
)

func channel(data string, c chan string) {
	for j := 1; j <= 5; j++ {
		c <- (data + "-- Channel")
		time.Sleep(time.Millisecond * 500)
	}
	close(c)
	fmt.Println("Closed channel....")
}
