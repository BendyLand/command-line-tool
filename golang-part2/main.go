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
	fmt.Println("Random IP address:\t\t\t" + utils.GenerateRandomIp())
	fmt.Println("Random message type:\t\t\t" + utils.ChooseRandomMessageType())
	fmt.Println("Random message origin:\t\t\t" + utils.ChooseRandomMessageOrigin())
	stdDate := utils.GenerateStandardStyleRandomDate()
	rqstDate := utils.GenerateRequestStyleRandomDate()
	time := utils.GenerateRandomTime()
	stdDateTime := utils.ConstructStandardDateTime(stdDate, time)
	rqstDateTime := utils.ConstructRequestDateTime(rqstDate, time)
	fmt.Println("Randomly created standard datetime:\t" + stdDateTime)
	fmt.Println("Randomly created request datetime:\t" + rqstDateTime)
	fmt.Println()
}
