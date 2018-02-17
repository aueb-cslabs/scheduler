package model

import (
	"fmt"
	"strconv"
)

type Schedule struct {
	Depth int
	//Date and Time/Administrator/Lab(0 for none)
	Slots map[string]map[string]int

	Index   int
	Fitness int
}

func (schedule *Schedule) AvailableAdminsAt(allAdmins []Admin, time DayTime, lab int) []Admin {
	var admins []Admin
	for _, admin := range allAdmins {
		if schedule.IsAdminAvailableAt(admin, time, lab) {
			admins = append(admins, admin)
		}
	}
	return admins
}

func (schedule *Schedule) IsAdminAvailableAt(admin Admin, time DayTime, lab int) bool {
	val, ok := admin.Preferences[time.String()]
	if !ok {
		return false
	}
	slot, ok := schedule.Slots[time.String()][admin.String()]
	if ok && slot > 0 {
		return false
	}
	return val == ABLE || val == ABLE_IF_NONE || val == ABLE_NOT_PREF || val == LESSON_ABLE ||
			((val == LAB_IN_1 || val == LAB_IN_1_NO_PREF) && lab == 1) ||
			((val == LAB_IN_2 || val == LAB_IN_2_NO_PREF) && lab == 2)
}

func (schedule Schedule) Print(times []DayTime) {
	for _, time := range times {
		fmt.Println(strconv.Itoa(time.Day) + ": " + strconv.Itoa(time.Time))
		for admin, lab := range schedule.Slots[time.String()] {
			fmt.Println("\t" + admin + ": " + strconv.Itoa(lab))
		}
	}
}