package main

import (
	"fmt"
	"golang-part2/pkgs/datetime"
	"golang-part2/pkgs/utils"
	"strconv"
)

func main() {
	for i := 0; i < 100; i++ {
		displayAll()
	}
}

func ConstructHttpRequest() string {
	messageType := utils.ChooseRandomRequestType()
	path := " /index.html HTTP/1.1\" 200 "
	responseSize := strconv.Itoa(utils.RandomNumBetween(800, 1600))
	return "\"" + messageType + path + responseSize
}

func ConstructFullRandomRequestMessage() string {
	ip := utils.GenerateRandomIp()
	date := datetime.GenerateRequestStyleRandomDate()
	time := datetime.GenerateRandomTime()
	dateTime := datetime.ConstructRequestDateTime(date, time)
	request := ConstructHttpRequest()
	return ip + " - - " + dateTime + " " + request
}

func ConstructFullRandomLogMessage() string {
	date := datetime.GenerateStandardStyleRandomDate()
	time := datetime.GenerateRandomTime()
	dateTime := datetime.ConstructStandardDateTime(date, time)
	messageType := utils.ChooseRandomMessageType()
	origin := utils.ChooseRandomMessageOrigin()
	var message string
	switch messageType {
	case "INFO":
		message = utils.ChooseInfoMessageFromOrigin(origin)
	case "WARN":
		message = utils.ChooseWarnMessageFromOrigin(origin)
	case "ERROR":
		message = utils.ChooseErrorMessageFromOrigin(origin)
	}
	return fmt.Sprintf("%s %s %s %s", dateTime, messageType, origin, message)
}

func displayAll() {
	fmt.Println()
	fmt.Println(ConstructFullRandomLogMessage())
	fmt.Println(ConstructFullRandomLogMessage())
	fmt.Println(ConstructFullRandomLogMessage())
	fmt.Println(ConstructFullRandomLogMessage())
	fmt.Println(ConstructFullRandomLogMessage())
	fmt.Println()
	fmt.Println(ConstructFullRandomRequestMessage())
	fmt.Println(ConstructFullRandomRequestMessage())
	fmt.Println(ConstructFullRandomRequestMessage())
	fmt.Println(ConstructFullRandomRequestMessage())
	fmt.Println(ConstructFullRandomRequestMessage())
	fmt.Println()
}
