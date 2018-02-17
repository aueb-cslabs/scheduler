package generator

import (
	"aueb.gr/cslabs/scheduler/model"
	"strconv"
)

func GenerateHtml(schedule model.Schedule, admins []model.Admin, times []model.DayTime) string {
	html := "<html><head><meta charset=\"UTF-8\"><link rel=\"stylesheet\" href=\"https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0/css/bootstrap.min.css\">"
	html += "<style>td {text-align: center;}</style>"
	html += "</head><body><table class=table><tr><td></td>"
	for _, time := range times {
		html += "<td>" + strconv.Itoa(time.Day) + "-" + strconv.Itoa(time.Time) + "</td>"
	}
	html += "</tr>"
	for _, admin := range admins {
		html += "<tr><td>" + admin.Name + "</td>"
		for _, time := range times {
			slot, ok := schedule.Slots[time.String()][admin.String()]
			if ok && slot > 0 {
				html += "<td style=\"background: red; color: white\">" + strconv.Itoa(slot) + "</td>"
			} else {
				html += "<td>" + string(admin.Preferences[time.String()]) + "</td>"
			}
		}
		html += "</tr>"
	}
	html += "</table></body></html>"
	return html
}
