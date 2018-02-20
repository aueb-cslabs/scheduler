package model

var Config = GeneratorConfiguration{
	ScheduleFirstDay:     1,
	ScheduleLastDay:      5,
	ScheduleFirstHour:    2,
	ScheduleLastHour:     6,
	PreferencesDays:      5,
	PreferencesDayLength: 5,

	IgnoreDays:     []int{},
	IgnoreHours:    []int{},
	IgnoreDayTimes: []DayHour{},
}

type GeneratorConfiguration struct {
	ScheduleFirstDay  int `json:"schedule_first_day"`
	ScheduleLastDay   int `json:"schedule_last_day"`
	ScheduleFirstHour int `json:"schedule_first_hour"`
	ScheduleLastHour  int `json:"schedule_last_hour"`

	PreferencesDays      int `json:"preferences_days"`
	PreferencesDayLength int `json:"preferences_day_length"`

	IgnoreDays     []int     `json:"ignore_days"`
	IgnoreHours    []int     `json:"ignore_hours"`
	IgnoreDayTimes []DayHour `json:"ignore_day_times"`
}

func (configuration GeneratorConfiguration) ScheduleDayLength() int {
	return configuration.ScheduleLastHour - configuration.ScheduleFirstHour + 1
}

func (configuration GeneratorConfiguration) ScheduleDays() int {
	return configuration.ScheduleLastDay - configuration.ScheduleFirstDay + 1
}
