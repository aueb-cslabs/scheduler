package parser

import (
	"aueb.gr/cslabs/scheduler/model"
	"encoding/csv"
	"io"
)

func ParsePreferenceCSV(r io.Reader, days, dayLength int) []model.Admin {

	csvReader := csv.NewReader(r)
	for {
		fields, _ := csvReader.Read()
		if len(fields) > 0 && fields[0] == "S" {
			break
		}
	}

	var admins []model.Admin
	for {
		fields, _ := csvReader.Read()
		if len(fields) < 0 || fields[0] == "E" {
			break
		}

		admin := model.Admin{Name: fields[0], Preferences: make(map[string]model.Preference)}
		slot := 1
		for day := 1; day <= days; day++ {
			for hour := 1; hour <= dayLength; hour++ {
				if len(fields[slot]) > 0 {
					admin.Preferences[model.DayHour{Day: day, Hour: hour}.String()] = model.Preference(fields[slot])
				}
				slot++
			}
		}
		admins = append(admins, admin)
	}
	return admins
}
