package main

import (
	"fmt"
	"golang-part2/pkgs/datetime"
	"golang-part2/pkgs/utils"
	"os"
	"strconv"
	"sync"
)

func main() {
	dirExists, _ := utils.Exists("logs")
	if !dirExists {
		os.Mkdir("logs", 0750)
	}
	numLogFiles, logsPerFile := Greet()
	var wg sync.WaitGroup

	logChan := make(chan string, numLogFiles)
	defer close(logChan)
	for i := 0; i < numLogFiles; i++ {
		wg.Add(1)
		go GenerateLogMessages(logChan, logsPerFile, &wg)
	}
	utils.WriteMessagesToFile(logChan, numLogFiles, logsPerFile)
	wg.Wait()

	utils.Cleanup(numLogFiles)
}

func Greet() (int, int) {
	fmt.Println(
		"Welcome to the Go Random Log Generator!",
		"\nHow many log messages would you like to generate?",
	)
	var input string
	fmt.Scan(&input)
	var logsToGenerate, numLogFiles, logsPerFile int
	logsToGenerate, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid input")
		numLogFiles, logsPerFile = Greet()
	}
	numLogFiles = int(logsToGenerate / 100)
	if numLogFiles < 1 {
		numLogFiles = 1
	}
	logsPerFile = int(logsToGenerate / numLogFiles)
	return numLogFiles, logsPerFile
}

// Generates a given number of random logs and sends them to the provided channel
func GenerateLogMessages(ch chan<- string, logsPerFile int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < logsPerFile; i++ {
		logMessage := ConstructFullRandomLogMessage()
		ch <- logMessage
	}
}

// Functions to call the 'ConstructFullMessage' functions and send results to a channel
func GenerateRequestMessage(ch chan<- string) {
	ch <- ConstructFullRandomRequestMessage()
}

func GenerateLogMessage(ch chan<- string) {
	ch <- ConstructFullRandomLogMessage()
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

func ConstructHttpRequest() string {
	messageType := utils.ChooseRandomRequestType()
	path := " /index.html HTTP/1.1\" 200 "
	responseSize := strconv.Itoa(utils.RandomNumBetween(800, 1600))
	return "\"" + messageType + path + responseSize
}
