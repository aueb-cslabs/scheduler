package scorer

import (
	"aueb.gr/cslabs/scheduler/model"
)


func calculateFTimeAdmin(schedule model.Schedule, time model.DayHour, lab int, admin model.Admin) int {
	fit := 0
	slot, ok := schedule.Slots[time.String()][admin.String()]
	if !ok || slot == 0 {
		return 0
	}

	//Calculate is already was in uni
	if !time.IsStartOfDay() {
		avail, _ := admin.Preferences[time.GetPreviousHour().String()]
		slot, ok := schedule.Slots[time.GetPreviousHour().String()][admin.String()]
		if ok && slot == lab {
			fit += FITNESS_SAME_LAB
			goto ScoreAvailability
		} else if ok && slot > 0 && slot != lab {
			fit += FITNESS_DIFFERENT_LAB
		} else if (ok && slot > 0) || avail == model.LESSON || avail == model.LESSON_ABLE ||
			avail == model.LAB_IN_1 || avail == model.LAB_IN_2 ||
			avail == model.LAB_IN_1_NO_PREF || avail == model.LAB_IN_2_NO_PREF {
			fit += FITNESS_IN_AUEB + FITNESS_IN_AUEB_MULT * admin.Distance
			goto ScoreAvailability
		} else if avail == model.UNABLE {
			fit += FITNESS_WILL_BE_UNAVAILABLE
			goto ScoreAvailability
		}
	}
	//Calculate if will stay in uni
	if !time.IsEndOfDay() {
		avail, _ := admin.Preferences[time.GetNextHour().String()]
		slot, ok := schedule.Slots[time.GetNextHour().String()][admin.String()]
		if ok && slot == lab {
			fit += FITNESS_SAME_LAB
		} else if ok && slot > 0 && slot != lab {
			fit += FITNESS_DIFFERENT_LAB
		} else if (ok && slot > 0) || avail == model.LESSON || avail == model.LESSON_ABLE ||
			avail == model.LAB_IN_1 || avail == model.LAB_IN_2 ||
			avail == model.LAB_IN_1_NO_PREF || avail == model.LAB_IN_2_NO_PREF {
			fit += FITNESS_IN_AUEB + FITNESS_IN_AUEB_MULT * admin.Distance
		} else if avail == model.UNABLE {
			fit += FITNESS_WILL_BE_UNAVAILABLE
		}
	}

ScoreAvailability:
	avail := admin.Preferences[time.String()]

	//Calculate based on availability
	switch avail {
	case model.ABLE:
		fit += FITNESS_ABLE
		break
	case model.ABLE_NOT_PREF:
		fit += FITNESS_ABLE_NOT_PREF
		break
	case model.ABLE_IF_NONE:
		fit += FITNESS_ABLE_IF_NONE
		break
	case model.LAB_IN_1:
		if lab == 1 {
			fit += FITNESS_REQUESTED
		}
		break
	case model.LAB_IN_2:
		if lab == 2 {
			fit += FITNESS_REQUESTED
		}
		break
	case model.LAB_IN_1_NO_PREF:
		if lab == 1 {
			fit += FITNESS_LAB_NOT_REQUESTED
		}
		break
	case model.LAB_IN_2_NO_PREF:
		if lab == 2 {
			fit += FITNESS_LAB_NOT_REQUESTED
		}
		break
	case model.LESSON_ABLE:
		fit += FITNESS_LESSON_NOT_REQUESTED
	}
	return fit
}
