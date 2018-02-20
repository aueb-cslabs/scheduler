package output

import (
	"aueb.gr/cslabs/scheduler/model"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"log"
	"strings"
)

func GenerateOfficialPDF(schedule model.Schedule, admins []model.Admin, times []model.DayHour, dayLength int) error {
	pdfDoc, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}

	// Add one page from an URL
	pdfDoc.AddPage(wkhtmltopdf.NewPageReader(strings.NewReader(generateOfficialHtmlCode(schedule, admins, times))))
	pdfDoc.Dpi.Set(600)
	pdfDoc.NoCollate.Set(false)
	pdfDoc.PageSize.Set(wkhtmltopdf.PageSizeA4)
	pdfDoc.Orientation.Set(wkhtmltopdf.OrientationPortrait)

	pdfDoc.MarginTop.Set(12)
	pdfDoc.MarginLeft.Set(12)
	pdfDoc.MarginRight.Set(12)
	pdfDoc.MarginBottom.Set(12)

	// Create PDF document in internal buffer
	err = pdfDoc.Create()
	if err != nil {
		return err
	}
	// Write buffer contents to file on disk
	prepareOutDir()
	err = pdfDoc.WriteFile(getOutputFile("official", "pdf"))
	if err != nil {
		return err
	}
	return nil
}
