package algorithm

import (
	"aueb.gr/cslabs/scheduler/model"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func MateSchedules(times []model.DayTime, schedule1 model.Schedule, schedule2 model.Schedule) model.Schedule {
	splitPoint := Generator.Intn(len(times))
	schedule := model.Schedule{
		Depth: min(schedule1.Depth + 1, schedule2.Depth + 1),
		Slots: make(map[string]map[string]int),
	}

	for index, time := range times {
		if index < splitPoint {
			schedule.Slots[time.String()] = schedule1.Slots[time.String()]
		} else {
			schedule.Slots[time.String()] = schedule2.Slots[time.String()]
		}
	}
	return schedule
}
