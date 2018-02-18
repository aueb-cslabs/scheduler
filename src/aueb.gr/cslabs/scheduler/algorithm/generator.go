package algorithm

import (
	"aueb.gr/cslabs/scheduler/model"
)

func GenerateRandomSchedule(admins []model.Admin, times []model.DayHour) model.Schedule {
	schedule := model.Schedule{
		Slots: make(map[string]map[string]int),
	}

	for _, time := range times {
		timeSlice := make(map[string]int)
		if !time.Ignored {
			for lab := 1; lab <= 2; lab++ {
				availableAdmins := schedule.AvailableAdminsAt(admins, time, lab)
				if len(availableAdmins) != 0 {
					randAdmin := availableAdmins[Generator.Intn(len(availableAdmins))]
					timeSlice[randAdmin.String()] = lab
				}
			}
		}
		schedule.Slots[time.String()] = timeSlice
	}
	return schedule
}
