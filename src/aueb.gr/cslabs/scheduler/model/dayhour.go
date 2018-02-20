package model

import "strconv"

type DayHour struct {
	Day     int  `json:"day"`
	Hour    int  `json:"hour"`
	Ignored bool `json:"-"`
}

func (dayTime DayHour) String() string {
	return strconv.Itoa(dayTime.Day) + "-" + strconv.Itoa(dayTime.Hour)
}

func (dayTime DayHour) GetPreviousHour() DayHour {
	return DayHour{Day: dayTime.Day, Hour: dayTime.Hour - 1}
}

func (dayTime DayHour) GetNextHour() DayHour {
	return DayHour{Day: dayTime.Day, Hour: dayTime.Hour + 1}
}

func (dayTime DayHour) IsStartOfDay() bool {
	return dayTime.Hour <= Config.ScheduleFirstHour
}

func (dayTime DayHour) IsEndOfDay() bool {
	return dayTime.Hour >= Config.ScheduleLastHour
}

func (dayTime DayHour) DayString() string {
	switch dayTime.Day {
	case 1:
		return "Δευτέρα"
	case 2:
		return "Τρίτη"
	case 3:
		return "Τετάρτη"
	case 4:
		return "Πέμπτη"
	case 5:
		return "Παρασκευή"
	}
	return ""
}

func (dayTime DayHour) TimeString() string {
	switch dayTime.Hour {
	case 1:
		return "9-11"
	case 2:
		return "11-1"
	case 3:
		return "1-3"
	case 4:
		return "3-5"
	case 5:
		return "5-7"
	case 6:
		return "7-9"
	}
	return ""
}
