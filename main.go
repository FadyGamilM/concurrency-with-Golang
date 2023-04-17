package main

import (
	"fmt"
)

func main() {
	// define the communication channel
	communicationChannel := make(chan string)

	// call the routine
	go getData(communicationChannel)

	// wait for joining all forked routines
	response := <-communicationChannel

	fmt.Println("response is => ", response)

	fmt.Println("processing is finished")
}

func getData(communicationChannel chan string) {
	fmt.Println("Doing some work ...")

	// join the main routine
	communicationChannel <- "done"
}
