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
	numLogFiles := 10
	logsPerFile := 100
	numGoroutines := 10
	wg := sync.WaitGroup{}

	logChan := make(chan string, numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go GenerateLogMessages(logChan, logsPerFile, &wg)
	}

	WriteMessagesToFile(logChan, numLogFiles, logsPerFile)

	wg.Wait()
}

func WriteMessagesToFile(logChan chan string, numLogFiles, logsPerFile int) {
	for i := 0; i < numLogFiles; i++ {
		filename := fmt.Sprintf("logs/log%d.txt", i+1)
		file, err := os.Create(filename)
		if err != nil {
			fmt.Println("Error creating file:", err)
			continue
		}
		defer file.Close()

		for j := 0; j < logsPerFile; j++ {
			logMessage := <-logChan

			_, err := file.WriteString(logMessage + "\n")
			if err != nil {
				fmt.Println("Error writing to file:", err)
				continue
			}
		}
		fmt.Println("Log messages written to:", filename)
	}
}

// Generates multiple random log messages and sends them to the provided channel
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

func displayAll() {
	fmt.Println()
	fmt.Println(ConstructFullRandomLogMessage())
	fmt.Println()
	fmt.Println(ConstructFullRandomRequestMessage())
	fmt.Println()
}
