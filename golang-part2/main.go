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
	fmt.Println("Random message type:\t\t\t" + utils.ChooseRandomMessageType())
	fmt.Println("Random message origin:\t\t\t" + utils.ChooseRandomMessageOrigin())
	fmt.Println(utils.ConstructFullRandomRequestMessage())
	fmt.Println(utils.ConstructFullRandomRequestMessage())
	fmt.Println(utils.ConstructFullRandomRequestMessage())
	fmt.Println(utils.ConstructFullRandomRequestMessage())
	fmt.Println(utils.ConstructFullRandomRequestMessage())
	fmt.Println()
}
