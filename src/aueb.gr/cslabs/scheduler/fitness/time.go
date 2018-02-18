package fitness

import "aueb.gr/cslabs/scheduler/model"

func calculateFTime(schedule model.Schedule, admins []model.Admin, time model.DayHour, lab int) int {
	score := 0
	if len(schedule.Slots[time.String()]) == 0 {
		return 0
	}
	for _, admin := range admins {
		score += calculateFTimeAdmin(schedule, time, lab, admin)
	}
	return score
}

