package utils

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func Cleanup(numLogFiles int) {
	success := CombineTextFiles(numLogFiles)
	if success {
		fmt.Println("Successfully combined files to 'result_logs.txt'!")
	} else {
		fmt.Println("Unable to combine files.")
	}
	EmptyDirectory("logs")
}

func EmptyDirectory(path string) {
	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("Error getting files:", err)
		os.Exit(1)
	}
	for i, file := range files {
		err := os.Remove(path + "/" + file.Name())
		if err != nil {
			fmt.Println("Problem removing file", i+1)
		}
	}
	fmt.Println("Directories removed successfully!")
}

func Exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return false
}

func CreateDirectory(path string) bool {
	err := os.Mkdir(path, 0750)
	if err != nil {
		fmt.Println("Error creating directory")
		return false
	}
	fmt.Println("Successfully created directory 'logs'")
	return true
}

func CombineTextFiles(numLogFiles int) bool {
	var filenames []string
	for i := 0; i < numLogFiles; i++ {
		filename := fmt.Sprintf("logs/log%d.txt", i+1)
		filenames = append(filenames, filename)
	}
	resultFile, err := os.Create("result_logs.txt")
	if err != nil {
		fmt.Println("Error creating output file: ", err)
		return false
	}
	defer resultFile.Close()

	for _, filename := range filenames {
		inputFile, err := os.Open(filename)
		if err != nil {
			fmt.Println("Error opening file: ", err)
			return false
		}
		defer inputFile.Close()

		success := ScanInputFileToResult(inputFile, resultFile)
		if !success {
			return false
		}
	}
	return true
}

func ScanInputFileToResult(inputFile, resultFile *os.File) bool {
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		_, err := resultFile.WriteString(scanner.Text() + "\n")
		if err != nil {
			fmt.Println("Error writing to result file: ", err)
			return false
		}
	}
	if scanErr := scanner.Err(); scanErr != nil {
		fmt.Println("Error reading input file: ", scanErr)
		return false
	}
	return true
}

func WriteMessagesToFile(logChan chan string, numLogFiles, logsPerFile int) {
	// Create each new file based on iteration #
	fmt.Println("Generating data...")
	for i := 0; i < numLogFiles; i++ {
		filename := fmt.Sprintf("logs/log%d.txt", i+1)
		file, err := os.Create(filename)
		if err != nil {
			fmt.Println("Error creating file:", err)
			continue
		}
		defer file.Close()

		// Write each log to the file with a newline
		for j := 0; j < logsPerFile; j++ {
			logMessage := <-logChan
			_, err := file.WriteString(logMessage + "\n")
			if err != nil {
				fmt.Println("Error writing to file:", err)
				continue
			}
		}
	}
}

func ChooseErrorMessageFromOrigin(origin string) string {
	var message string
	switch origin {
	case "[Security]":
		message = "Permission Denied: You do not have access to this content."
	case "[Database]":
		message = "There was a problem executing the SQL Query."
	case "[SystemMonitor]":
		message = "Not enough disk space to perform task."
	case "[UserManagementService]":
		message = "Error processing user registration request: User email already exists."
	}
	return message
}

func ChooseWarnMessageFromOrigin(origin string) string {
	var message string
	ip := GenerateRandomIp()
	switch origin {
	case "[Security]":
		message = fmt.Sprintf("Suspicious login attempt detected: IP address %s, username 'admin'", ip)
	case "[Database]":
		message = "SQL query may be missing rows. Please double check your input data."
	case "[SystemMonitor]":
		message = "Disk space utilization at 80%. Consider adding more storage or removing unneeded files."
	case "[UserManagementService]":
		message = "Password may not contain enough variety. Proceed with caution."
	}
	return message
}

func ChooseInfoMessageFromOrigin(origin string) string {
	var message string
	ip := GenerateRandomIp()
	switch origin {
	case "[Security]":
		message = fmt.Sprintf("Successful login: IP address %s, username 'admin'", ip)
	case "[Database]":
		message = "Executing SQL query: DELETE FROM table_name WHERE col2 IS NULL"
	case "[SystemMonitor]":
		message = "CPU utilization: 30%, Memory usage: 60%, Disk space available: 50%"
	case "[UserManagementService]":
		message = "Successfully created user account!"
	}
	return message
}

func ChooseRandomMessageOrigin() string {
	origins := []string{"[Security]", "[Database]", "[SystemMonitor]", "[UserManagementService]"}
	return Sample(origins)
}

func GenerateRandomIp() string {
	var nums []string
	for i := 0; i < 5; i++ {
		rand := strconv.Itoa(RandomNumUnder(256))
		nums = append(nums, rand)
	}
	return strings.Join(nums, ".")
}

func ChooseRandomMessageType() string {
	messageTypes := []string{"INFO", "WARN", "ERROR"}
	return Sample(messageTypes)
}

func ChooseRandomRequestType() string {
	requestTypes := []string{"GET", "POST", "PUT", "DELETE"}
	return Sample(requestTypes)
}

func ChooseRandomQueryType() string {
	queryTypes := []string{
		"SELECT col1, col2 FROM table_name WHERE col2 IS NOT NULL",
		"DELETE FROM table_name WHERE col2 IS NULL",
		"INSERT INTO table_name (col1, col2) VALUES (val1, val2)",
		"UPDATE table_name SET col1 = val1 WHERE col2 IS NOT NULL"}
	return Sample(queryTypes)
}

func RandomNumUnder(num int) int {
	return rand.Int() % num
}

func RandomNumBetween(bottom, top int) int {
	return rand.Intn(top-bottom) + bottom
}

func Sample(coll []string) string {
	length := len(coll)
	num := RandomNumUnder(length)
	return coll[num]
}
