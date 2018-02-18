package scorer

import "aueb.gr/cslabs/scheduler/model"

func CalculateFitness(schedule model.Schedule, admins []model.Admin, times []model.DayHour) int {
	fit := 0
	for lab := 1; lab <= 2; lab++ {
		for _, time := range times {
			fit += calculateFTime(schedule, admins, time, lab)
		}
	}
	fit += calculateHours(schedule, admins, times)

	return fit
}
