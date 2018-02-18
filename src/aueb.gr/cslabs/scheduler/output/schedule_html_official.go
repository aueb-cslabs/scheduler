package output

import (
	"aueb.gr/cslabs/scheduler/model"
	"strconv"
	"os"
	"bufio"
)

func GenerateOfficialHtml(schedule model.Schedule, admins []model.Admin, times []model.DayHour) error {
	prepareOutDir()
	html := generateOfficialHtmlCode(schedule, admins, times)
	f, err := os.Create(getOutputFile("official", "html"))
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

func generateOfficialHtmlCode(schedule model.Schedule, admins []model.Admin, times []model.DayHour) string {

	colorCollection := []string{"bg-primary", "bg-warning", "bg-success"}

	html := "<html><head><meta charset=\"UTF-8\"><link rel=\"stylesheet\" href=\"https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0/css/bootstrap.min.css\">"
	html += "<style>td, th {text-align: center;}</style>"
	html += "</head><body><div class=\"p-4\">"
	html += "<h2 class='mb-4'>" + schedule.Title + "</h2>"

	for lab := 1; lab <= 2; lab++ {
		html += "<table class='table table-bordered table-striped mb-4'><tr><th class='lead " + colorCollection[lab - 1] + "' colspan='" + strconv.Itoa(model.Config.ScheduleDays() + 1) + "'> CSLab " + strconv.Itoa(lab) + "</th></tr>"

		printedDay := 0
		html += "<tr><td></td>"
		for _, time := range times {

			if printedDay != time.Day {
				printedDay = time.Day
				html += "<td>" + time.DayString() + "</td>"
			}
		}
		html += "</tr>"

		for hour := model.Config.ScheduleFirstHour; hour <= model.Config.ScheduleLastHour; hour++ {
			html += "<tr><td>" + model.DayHour{Day:0, Time: hour}.TimeString() + "</td>"
			for day := model.Config.ScheduleFirstDay; day <= model.Config.ScheduleLastDay; day++ {
				dayHour := model.DayHour{Day: day, Time: hour}
				admin := "-"
				for _, adm := range admins {
					slot, ok := schedule.Slots[dayHour.String()][adm.String()]
					if ok && slot == lab {
						admin = adm.Name
						break
					}
				}
				html += "<td>" + admin + "</td>"
			}
			html += "</tr>"
		}
		html += "</table>"
	}

	return html
}
