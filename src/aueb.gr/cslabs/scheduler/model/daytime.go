package model

import "strconv"

var FirstDay = 1
var LastDay = 5
var FirstHour = 2
var LastHour = 6

type DayTime struct {
	Day int
	Time int
}

func (dayTime DayTime) String() string {
	return strconv.Itoa(dayTime.Day) + strconv.Itoa(dayTime.Time)
}

func (dayTime DayTime) GetPreviousHour() DayTime {
	return DayTime{Day: dayTime.Day, Time: dayTime.Time - 1}
}

func (dayTime DayTime) GetNextHour() DayTime {
	return DayTime{Day: dayTime.Day, Time: dayTime.Time + 1}
}

func (dayTime DayTime) IsStartOfDay() bool {
	return dayTime.Time <= FirstHour
}

func (dayTime DayTime) IsEndOfDay() bool {
	return dayTime.Time >= LastHour
}

func (dayTime DayTime) DayString() string {
	switch dayTime.Day {
	case 1: return "Δευτέρα"
	case 2: return "Τρίτη"
	case 3: return "Τετάρτη"
	case 4: return "Πέμπτη"
	case 5: return "Παρασκευή"
	}
	return ""
}

func (dayTime DayTime) TimeString() string {
	switch dayTime.Time {
	case 1: return "9-11"
	case 2: return "11-1"
	case 3: return "1-3"
	case 4: return "3-5"
	case 5: return "5-7"
	case 6: return "7-9"
	}
	return ""
}