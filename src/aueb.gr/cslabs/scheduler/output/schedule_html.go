package output

import (
	"aueb.gr/cslabs/scheduler/model"
	"strconv"
	"os"
	"bufio"
)

func GenerateHtml(title string, schedule model.Schedule, admins []model.Admin, times []model.DayTime, dayLength int) error {
	prepareOutDir()
	html := generateHtmlCode(title, schedule, admins, times, dayLength)
	f, err := os.Create("out/" + title + ".html")
	if err != nil {
		return err
	}
	w := bufio.NewWriter(f)
	_, err = w.WriteString(html)
	if err != nil {
		return err
	}
	w.Flush()
	return nil
}

func generateHtmlCode(title string, schedule model.Schedule, admins []model.Admin, times []model.DayTime, dayLength int) string {
	html := "<html><head><meta charset=\"UTF-8\"><link rel=\"stylesheet\" href=\"https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0/css/bootstrap.min.css\">"
	html += "<style>td {text-align: center;} .active {background: red; color: white;} @page { margin: 0; } @media print { td {-webkit-print-color-adjust: exact;} .active {background: red; color: white;} }</style>"
	html += "</head><body><div class=\"p-4\">"
	html += "<h1 class='mb-4'>" + title + "</h1>"
	html += "<table class=\"table table-bordered table-striped\"><tr><td></td>"

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
				html += "<td class='active lead'>" + strconv.Itoa(slot) + "</td>"
			} else {
				html += "<td class='lead'>" + string(admin.Preferences[time.String()]) + "</td>"
			}
		}
		html += "</tr>"
	}
	html += "</table></div></body></html>"
	return html
}
