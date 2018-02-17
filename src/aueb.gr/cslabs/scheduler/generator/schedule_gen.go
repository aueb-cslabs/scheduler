package generator

import (
	"aueb.gr/cslabs/scheduler/model"
)

func GenerateRandomSchedule(admins []model.Admin, times []model.DayTime) model.Schedule {
	schedule := model.Schedule{
		Depth: 0,
		Slots: make(map[string]map[string]int),
	}

	for _, time := range times {
		timeSlice := make(map[string]int)
		for lab := 1; lab <= 2; lab++ {
			availableAdmins := schedule.AvailableAdminsAt(admins, time, lab)
			if len(availableAdmins) != 0 {
				randAdmin := availableAdmins[Generator.Intn(len(availableAdmins))]
				timeSlice[randAdmin.String()] = lab
			}
		}
		schedule.Slots[time.String()] = timeSlice
	}
	return schedule
}
