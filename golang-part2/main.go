package main

import (
	"fmt"
	"golang-part2/utils"
)

func main() {
	for i := 0; i < 100; i++ {
		displayAll()
	}
}

func displayAll() {
	fmt.Println()
	fmt.Println("Welcome to the Golang Random Log Message Generator!\n")
	fmt.Println("Random IP address:\t\t" + utils.GenerateRandomIp())
	fmt.Println("Random time:\t\t\t" + utils.GenerateRandomTime())
	fmt.Println("Random message type:\t\t" + utils.ChooseRandomMessageType())
	fmt.Println("Random standard style date:\t" + utils.GenerateStandardStyleRandomDate())
	fmt.Println("Random request style date:\t" + utils.GenerateRequestStyleRandomDate())
	fmt.Println()
}
