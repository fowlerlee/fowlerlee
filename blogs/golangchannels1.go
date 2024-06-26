package main

import "fmt"

func main() {
	fmt.Print("hello ")
	done := make(chan bool)
	readData := make(chan string, 2)

	work := func(done <-chan bool, readData <- ){

	}
}