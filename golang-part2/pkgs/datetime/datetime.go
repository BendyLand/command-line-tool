package datetime

import (
	"golang-part2/pkgs/utils"
	"strconv"
	"strings"
)

func GenerateDateComponents() []string {
	year := strconv.Itoa(utils.RandomNumBetween(1970, 2025))
	month := strconv.Itoa(utils.RandomNumBetween(1, 13))
	if len(month) < 2 {
		month = "0" + month
	}
	day := strconv.Itoa(utils.RandomNumBetween(1, 32))
	if len(day) < 2 {
		day = "0" + day
	}
	return []string{year, month, day}
}

func GenerateRandomTime() string {
	var times []string
	randHour := strconv.Itoa(utils.RandomNumBetween(1, 13))
	if len(randHour) < 2 {
		randHour = "0" + randHour
	}
	times = append(times, randHour)
	for i := 0; i < 2; i++ {
		rand := strconv.Itoa(utils.RandomNumUnder(60))
		if len(rand) < 2 {
			rand = "0" + rand
		}
		times = append(times, rand)
	}
	return strings.Join(times, ":")
}

func GenerateStandardStyleRandomDate() string {
	components := GenerateDateComponents()
	return strings.Join(components, "-")
}

func ConstructStandardDateTime(date, time string) string {
	components := []string{date, time}
	return strings.Join(components, "T")
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

func ConstructRequestDateTime(date, time string) string {
	return "[" + date + ":" + time + " +0000]"
}
