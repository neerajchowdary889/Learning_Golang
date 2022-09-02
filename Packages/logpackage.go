package main

import (
	"log"
	"os"
	"time"
)

func main() {
	file, _ := os.Create("File.log")

	log.SetFlags(log.Ltime | log.Ldate)

	log.Println("Hello World")

	log.SetOutput(file)

	log.Println("Hello Bro, I'm inside the Log file...")
	log.Println("Cooooolllllll brooo...")

	for i := 1; i <= 5; i++ {
		log.Println("Wowwwww", i)
		time.Sleep(time.Millisecond * 300)
	}

	log.SetOutput(os.Stdout)
	log.Println("nahhhh, I'm outside Log file...")
	log.Println("Be like a good boyyyy...")

	flags := log.Ldate | log.Ltime
	infoLogger := log.New(file, "Info: ", flags)
	infoLogger.Println("Im Back broo")

	warnLogger := log.New(os.Stdout, "Warn: ", flags)
	warnLogger.Println("Sup Mannn")

	file.Close()
}
