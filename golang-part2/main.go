package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to the Golang Random Log Message Generator!")
	fmt.Println(generateRandomIp())
	fmt.Println(generateRandomTime())
	fmt.Println(chooseRandomMessageType())
	fmt.Println(generateStandardStyleRandomDate())
	fmt.Println(generateRequestStyleRandomDate())
}

/*
2001-02-13T05:16:14 INFO [Security] Successful login: IP address 137.221.132.152.201, username 'admin'
1974-01-01T05:34:59 INFO [Database] Executing SQL query: SELECT col1, col2 FROM table_name WHERE col2 IS NOT NULL
1997-06-11T11:27:18 INFO [SystemMonitor] CPU utilization: 30%, Memory usage: 60%, Disk space available: 50%
2010-08-26T02:54:02 ERROR [UserManagementService] Error processing user registration request: User email already exists
247.102.65.159.61 - - [30/Sep/2022:09:38:29 +0000] "POST /index.html HTTP/1.1" 200 1301
2013-08-18T05:54:53 WARN [Database] SQL query may be missing rows. Please double check your input data.
1982-10-11T04:07:39 INFO [UserManagementService] Successfully created user account!
1997-05-25T05:04:55 ERROR [SystemMonitor] Not enough disk space to perform task.
2016-01-23T09:25:05 WARN [Security] Suspicious login attempt detected: IP address 191.116.246.170.88, username 'admin'
2009-12-16T07:36:28 INFO [Database] Executing SQL query: SELECT col1, col2 FROM table_name WHERE col2 IS NOT NULL
2017-04-14T12:40:32 ERROR [UserManagementService] Error processing user registration request: User email already exists
1998-09-06T03:31:43 INFO [Database] Executing SQL query: SELECT col1, col2 FROM table_name WHERE col2 IS NOT NULL
229.44.95.187.240 - - [08/Jul/1977:08:10:07 +0000] "POST /index.html HTTP/1.1" 200 1370
2006-07-09T02:27:10 INFO [UserManagementService] Successfully created user account!
1972-09-30T09:38:40 INFO [Security] Successful login: IP address 173.17.22.107.95, username 'admin'
13.172.83.232.64 - - [03/Oct/1984:04:08:07 +0000] "PUT /index.html HTTP/1.1" 200 1540
2012-09-13T09:29:51 ERROR [UserManagementService] Error processing user registration request: User email already exists
2006-07-07T11:37:05 WARN [Security] Suspicious login attempt detected: IP address 62.44.36.76.46, username 'admin'
2011-07-21T06:31:26 ERROR [UserManagementService] Error processing user registration request: User email already exists
1983-03-24T06:17:38 ERROR [SystemMonitor] Not enough disk space to perform task.
219.185.255.193.120 - - [11/Apr/1983:03:15:44 +0000] "POST /index.html HTTP/1.1" 200 1491
1997-12-17T06:32:30 ERROR [UserManagementService] Error processing user registration request: User email already exists
1971-05-14T09:34:07 INFO [SystemMonitor] CPU utilization: 30%, Memory usage: 60%, Disk space available: 50%
163.85.115.80.181 - - [11/Feb/2021:07:50:29 +0000] "DELETE /index.html HTTP/1.1" 200 1083
*/

func generateRandomIp() string {
	var nums []string
	for i := 0; i < 5; i++ {
		rand := strconv.Itoa(rand.Int() % 255 + 1)
		nums = append(nums, rand)
	}
	return strings.Join(nums, ".")
}

func generateRandomTime() string {
	var times []string
	randHour := strconv.Itoa(rand.Int() % 12 + 1)
	if len(randHour) < 2 {
		randHour = "0" + randHour
	}
	times = append(times, randHour)
	for i := 0; i < 2; i++ {
		rand := strconv.Itoa(rand.Int() % 60)
		if len(rand) < 2 {
			rand = "0" + rand
		}
		times = append(times, rand)
	}
	return strings.Join(times, ":")
}

func chooseRandomMessageType() string {
	messageTypes := []string{"INFO", "WARN", "ERROR"}
	num := rand.Int() % 3
	return messageTypes[num]
}

func generateDateComponents() []string {
	lowerBound := 1970
	upperBound := 2025
	year := strconv.Itoa(rand.Intn(upperBound-lowerBound) + lowerBound)
	month := strconv.Itoa(rand.Int() % 12 + 1)
	if len(month) < 2 {
		month = "0" + month
	}
	day := strconv.Itoa(rand.Int() % 30 + 1)
	if len(day) < 2 {
		day = "0" + day
	}
	return []string{year, month, day}
}

func generateStandardStyleRandomDate() string {
	times := generateDateComponents()
	return strings.Join(times, "-")
}

func generateRequestStyleRandomDate() string {
	times := generateDateComponents()
	year, m, day  := times[0], times[1], times[2]
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