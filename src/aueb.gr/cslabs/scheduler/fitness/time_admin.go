package fitness

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
			fit += fitnessSameLab
			goto ScoreAvailability
		} else if ok && slot > 0 && slot != lab {
			fit += fitnessDifferentLab
		} else if (ok && slot > 0) || avail == model.Lesson || avail == model.LessonAble ||
			avail == model.Request1 || avail == model.Request2 ||
			avail == model.In1NotPreferable || avail == model.In2NotPreferable {
			fit += fitnessInPremises + fitnessInPremisesDistanceMult*admin.Distance
			goto ScoreAvailability
		} else if avail == model.NotAble {
			fit += fitnessWillBeUnavailable
			goto ScoreAvailability
		}
	}
	//Calculate if will stay in uni
	if !time.IsEndOfDay() {
		avail, _ := admin.Preferences[time.GetNextHour().String()]
		slot, ok := schedule.Slots[time.GetNextHour().String()][admin.String()]
		if ok && slot == lab {
			fit += fitnessSameLab
		} else if ok && slot > 0 && slot != lab {
			fit += fitnessDifferentLab
		} else if (ok && slot > 0) || avail == model.Lesson || avail == model.LessonAble ||
			avail == model.Request1 || avail == model.Request2 ||
			avail == model.In1NotPreferable || avail == model.In2NotPreferable {
			fit += fitnessInPremises + fitnessInPremisesDistanceMult*admin.Distance
		} else if avail == model.NotAble {
			fit += fitnessWillBeUnavailable
		}
	}

ScoreAvailability:
	avail := admin.Preferences[time.String()]

	//Calculate based on availability
	switch avail {
	case model.Able:
		fit += fitnessAble
		break
	case model.AbleNotPreferable:
		fit += fitnessAbleNotPref
		break
	case model.AbleIfNoneAvailable:
		fit += fitnessAbleIfNone
		break
	case model.Request1:
		if lab == 1 {
			fit += fitnessRequested
		}
		break
	case model.Request2:
		if lab == 2 {
			fit += fitnessRequested
		}
		break
	case model.In1NotPreferable:
		if lab == 1 {
			fit += fitnessLabNotRequested
		}
		break
	case model.In2NotPreferable:
		if lab == 2 {
			fit += fitnessLabNotRequested
		}
		break
	case model.LessonAble:
		fit += fitnessLessonNotRequested
	}
	return fit
}
