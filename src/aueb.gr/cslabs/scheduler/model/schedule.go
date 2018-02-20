package model

import (
	"fmt"
	"strconv"
)

var CustomBlockRule func(admin Admin, time DayHour, lab int) bool

type Schedule struct {
	Title string `json:"title"`
	//Date and Time/Administrator/Lab(0 for none)
	Slots map[string]map[string]int `json:"slots"`

	Index   int `json:"-"`
	Fitness int `json:"fitness"`
}

func (schedule *Schedule) AvailableAdminsAt(allAdmins []Admin, time DayHour, lab int) []Admin {
	var admins []Admin
	hasRequested := false
	for _, admin := range allAdmins {
		available, req := schedule.IsAdminAvailableAt(admin, time, lab)
		if req {
			hasRequested = true
			admins = []Admin{}
			break
		}
		if available {
			admins = append(admins, admin)
		}
	}
	if !hasRequested {
		return admins
	}
	for _, admin := range allAdmins {
		available, req := schedule.IsAdminAvailableAt(admin, time, lab)
		if available && req {
			admins = append(admins, admin)
		}
	}
	return admins
}

func (schedule *Schedule) IsAdminAvailableAt(admin Admin, time DayHour, lab int) (bool, bool) {
	val, ok := admin.Preferences[time.String()]
	if !ok {
		return false, false
	}
	if CustomBlockRule != nil && CustomBlockRule(admin, time, lab) {
		return false, false
	}

	slot, ok := schedule.Slots[time.String()][admin.String()]
	if ok && slot > 0 {
		return false, false
	}
	if (val == Request1 && lab == 1) || (val == Request2 && lab == 2) {
		return true, true
	}
	return val == Able || val == AbleIfNoneAvailable || val == AbleNotPreferable || val == LessonAble ||
		(val == In1NotPreferable && lab == 1) ||
		(val == In2NotPreferable && lab == 2), false
}

func (schedule Schedule) Print(times []DayHour) {
	for _, time := range times {
		fmt.Println(strconv.Itoa(time.Day) + ": " + strconv.Itoa(time.Hour))
		for admin, lab := range schedule.Slots[time.String()] {
			fmt.Println("\t" + admin + ": " + strconv.Itoa(lab))
		}
	}
}
