package utils

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

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
