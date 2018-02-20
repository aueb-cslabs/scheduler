package fitness

import (
	"aueb.gr/cslabs/scheduler/model"
	"math"
)

var HoursPerAdmin = 0.0

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func calculateHours(schedule model.Schedule, admins []model.Admin, times []model.DayHour) int {
	fit := 0
	for _, admin := range admins {
		hours := 0

		dayCount := 0
		day := 0
		startOfDay := 0
		currentDay := 0
		for _, time := range times {
			slot, ok := schedule.Slots[time.String()][admin.String()]
			if currentDay != time.Day {
				if day != 0 {
					dayCount++
				}
				currentDay = time.Day
				day = 0
				startOfDay = 0
			}
			if ok && slot > 0 {
				hours++
				day++
				if startOfDay == 0 {
					startOfDay = day
				}
				if day-startOfDay >= 3 {
					fit += fitnessLongStay
				}
				if day-startOfDay >= 5 {
					fit += fitnessDayLongStay
				}
			}
			if day == 3 {
				fit += fitnessMoreThan2Hours
			}
			if day > 3 {
				fit += fitnessMoreThan3Hours
			}
		}
		fit += dayCount * fitnessDayInLabs
		fit += int(math.Max(0, HoursPerAdmin-float64(hours)) * fitnessLessHoursPerHour)
		if float64(hours)-HoursPerAdmin > 2 {
			fit += int(math.Max(0, float64(hours)-HoursPerAdmin) * fitnessMoreHoursPerHourAfter2)
		} else {
			fit += int(math.Max(0, float64(hours)-HoursPerAdmin) * fitnessMoreHoursPerHour)
		}
	}

	for _, time := range times {
		if len(schedule.Slots[time.String()]) < 2 {
			fit += fitnessEmptyHours
		}
	}
	return fit
}
