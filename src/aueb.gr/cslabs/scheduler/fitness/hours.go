package scorer

import "aueb.gr/cslabs/scheduler/model"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func calculateHours(schedule model.Schedule, admins []model.Admin, times []model.DayTime) int {
	fit := 0
	hoursPerAdmin := len(times) * 2 / len(admins)

	for _, admin := range admins {
		hours := 0
		day := 0
		currentDay := 0
		for _, time := range times {
			slot, ok := schedule.Slots[time.String()][admin.String()]
			if ok && slot > 0 {
				hours++
				if currentDay != time.Day {
					currentDay = time.Day
					day = 0
				}
				day++
			}
			if day == 3 {
				fit += FITNESS_MORE_THAN_2_HOURS
			}
			if day > 3 {
				fit += FITNESS_MORE_THAN_3_HOURS
			}
		}
		fit += max(0, hoursPerAdmin - hours) * FITNESS_LESS_HOURS_PER_HOUR
		if hours - hoursPerAdmin > 2 {
			fit += max(0, hours - hoursPerAdmin) * FITNESS_MORE_HOURS_PER_HOUR_AF2
		} else {
			fit += max(0, hours - hoursPerAdmin) * FITNESS_MORE_HOURS_PER_HOUR
		}
	}

	for _, time := range times {
		if len(schedule.Slots[time.String()]) < 2 {
			fit += FITNESS_EMPTY_HOURS
		}
	}

	return fit
}
