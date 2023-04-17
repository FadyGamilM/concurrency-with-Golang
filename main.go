package main

import (
	"fmt"
)

func main() {
	go getData()
	fmt.Println("processing is finished")
}

func getData() {
	fmt.Println("data ..")
}
 