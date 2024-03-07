package utils

import (
	"strconv"
	"strings"
	"math/rand"
)

/*  
2001-02-13T05:16:14 INFO [Security] Successful login: IP address 137.221.132.152.201, username 'admin'
1974-01-01T05:34:59 INFO [Database] Executing SQL query: SELECT col1, col2 FROM table_name WHERE col2 IS NOT NULL
1997-06-11T11:27:18 INFO [SystemMonitor] CPU utilization: 30%, Memory usage: 60%, Disk space available: 50%
2010-08-26T02:54:02 ERROR [UserManagementService] Error processing user registration request: User email already exists
247.102.65.159.61 - - [30/Sep/2022:09:38:29 +0000] "POST /index.html HTTP/1.1" 200 1301
*/

func ConstructRequestDateTime(date, time string) string {
	return "[" + date + ":" + time + " +0000]"
}

func ConstructStandardDateTime(date, time string) string {
	components := []string{date, time}
	return strings.Join(components, "T")
}

func ChooseRandomMessageOrigin() string {
	origins := []string{"[Security]", "[Database]", "[SystemMonitor]", "[UserManagementService]"}
	num := randomNumUnder(4)
	return origins[num]
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
	num := randomNumUnder(3)
	return messageTypes[num]
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