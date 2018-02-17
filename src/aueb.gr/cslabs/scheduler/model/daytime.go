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