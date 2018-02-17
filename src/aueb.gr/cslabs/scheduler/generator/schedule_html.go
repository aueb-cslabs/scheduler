package generator

import (
	"aueb.gr/cslabs/scheduler/model"
	"strconv"
)

func GenerateHtml(schedule model.Schedule, admins []model.Admin, times []model.DayTime, dayLength int) string {
	html := "<html><head><meta charset=\"UTF-8\"><link rel=\"stylesheet\" href=\"https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0/css/bootstrap.min.css\">"
	html += "<style>td {text-align: center;} @page { margin: 0; }</style>"
	html += "</head><body><div class=\"p-4\"><table class=\"table table-bordered\"><tr><td></td>"

	printedDay := 0
	for _, time := range times {
		if printedDay != time.Day {
			printedDay = time.Day
			html += "<td colspan=\"" + strconv.Itoa(dayLength) + "\">" + time.DayString() + "</td>"
		}
	}
	html += "</tr><tr><td></td>"
	for _, time := range times {
		html += "<td>" + time.TimeString() + "</td>"
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
	html += "</table></div></body></html>"
	return html
}
