package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fmt.Println("Welcome welcome....")
	data := "Welcome to the file bro, have a nice day...Bye"
	file, err := os.Create("MyFile.txt")

	if err != nil {
		panic("Gofer is panicking...")
	}
	io.WriteString(file, data)
	//fmt.Println(writer)
	//readfile(file)
	readfile()
	file.Close()
}

func readfile() {
	File, _ := os.Create("NewFile.txt")
	log.SetOutput(File)
	log.Println("Hello world")
	log.Println("Hello bro I'm in new File")
	log.SetOutput(os.Stdout)
	log.Println("Nothing much")
}
