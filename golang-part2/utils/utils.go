package utils

import (
	"strconv"
	"strings"
	"math/rand"
)

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