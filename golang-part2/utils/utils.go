package utils

import (
	"math/rand"
	"strconv"
	"strings"
)

/*
2001-02-13T05:16:14 INFO [Security] Successful login: IP address 137.221.132.152.201, username 'admin'
1974-01-01T05:34:59 INFO [Database] Executing SQL query: SELECT col1, col2 FROM table_name WHERE col2 IS NOT NULL
1997-06-11T11:27:18 INFO [SystemMonitor] CPU utilization: 30%, Memory usage: 60%, Disk space available: 50%
2010-08-26T02:54:02 ERROR [UserManagementService] Error processing user registration request: User email already exists
*/



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
	switch origin {
	case "[Security]":
		message = "Suspicious login attempt detected: IP address 56.75.109.2.235, username 'admin'"
	case "[Database]":
		message = "SQL query may be missing rows. Please double check your input data."
	case "[SystemMonitor]":
		message = "Suspicious login attempt detected: IP address 56.75.109.2.235, username 'admin'"
	case "[UserManagementService]":
		message = "Password may not contain enough variety. Proceed with caution."
	}
	return message
}

func ChooseInfoMessageFromOrigin(origin string) string {
	var message string
	switch origin {
	case "[Security]":
		message = "Successful login: IP address 190.95.66.237.224, username 'admin'"
	case "[Database]":
		message = "Executing SQL query: DELETE FROM table_name WHERE col2 IS NULL"
	case "[SystemMonitor]":
		message = "CPU utilization: 30%, Memory usage: 60%, Disk space available: 50%"
	case "[UserManagementService]":
		message = "Successfully created user account!"
	}
	return message
}

func ConstructFullRandomRequestMessage() string {
	ip := GenerateRandomIp()
	dateTime := GenerateRequestStyleRandomDate()
	request := ConstructHttpRequest()
	return ip + " - - " + dateTime + " " + request
}

func ConstructHttpRequest() string {
	messageType := ChooseRandomRequestType()
	path := " /index.html HHTP/1.1\" 200 "
	responseSize := strconv.Itoa(randomNumBetween(800, 1600))
	return "\"" + messageType + path + responseSize
}

func ConstructRequestDateTime(date, time string) string {
	return "[" + date + ":" + time + " +0000]"
}

func ConstructStandardDateTime(date, time string) string {
	components := []string{date, time}
	return strings.Join(components, "T")
}

func ChooseRandomMessageOrigin() string {
	origins := []string{"[Security]", "[Database]", "[SystemMonitor]", "[UserManagementService]"}
	return Sample(origins)
}

func GenerateRandomIp() string {
	var nums []string
	for i := 0; i < 5; i++ {
		rand := strconv.Itoa(randomNumUnder(256))
		nums = append(nums, rand)
	}
	return strings.Join(nums, ".")
}

func GenerateRandomTime() string {
	var times []string
	randHour := strconv.Itoa(randomNumBetween(1, 13))
	if len(randHour) < 2 {
		randHour = "0" + randHour
	}
	times = append(times, randHour)
	for i := 0; i < 2; i++ {
		rand := strconv.Itoa(randomNumUnder(60))
		if len(rand) < 2 {
			rand = "0" + rand
		}
		times = append(times, rand)
	}
	return strings.Join(times, ":")
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

func GenerateDateComponents() []string {
	year := strconv.Itoa(randomNumBetween(1970, 2025))
	month := strconv.Itoa(randomNumBetween(1, 13))
	if len(month) < 2 {
		month = "0" + month
	}
	day := strconv.Itoa(randomNumBetween(1, 32))
	if len(day) < 2 {
		day = "0" + day
	}
	return []string{year, month, day}
}

func GenerateStandardStyleRandomDate() string {
	times := GenerateDateComponents()
	return strings.Join(times, "-")
}

func GenerateRequestStyleRandomDate() string {
	times := GenerateDateComponents()
	year, m, day := times[0], times[1], times[2]
	var month string
	switch m {
	case "01":
		month = "Jan"
	case "02":
		month = "Feb"
	case "03":
		month = "Mar"
	case "04":
		month = "Apr"
	case "05":
		month = "May"
	case "06":
		month = "Jun"
	case "07":
		month = "Jul"
	case "08":
		month = "Aug"
	case "09":
		month = "Sep"
	case "10":
		month = "Oct"
	case "11":
		month = "Nov"
	case "12":
		month = "Dec"
	}
	dateComponents := []string{day, month, year}
	return strings.Join(dateComponents, "/")
}

func randomNumUnder(num int) int {
	return rand.Int() % num
}

func randomNumBetween(bottom, top int) int {
	return rand.Intn(top-bottom) + bottom
}

func Sample(coll []string) string {
	length := len(coll)
	num := randomNumUnder(length)
	return coll[num]
}
