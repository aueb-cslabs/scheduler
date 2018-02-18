package output

import (
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"aueb.gr/cslabs/scheduler/model"
	"log"
	"strings"
)

func GeneratePDF(title string, schedule model.Schedule, admins []model.Admin, times []model.DayTime, dayLength int) error {
	pdfDoc, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}

	// Add one page from an URL
	pdfDoc.AddPage(wkhtmltopdf.NewPageReader(strings.NewReader(generateHtmlCode(title, schedule, admins, times, dayLength))))
	pdfDoc.Dpi.Set(600)
	pdfDoc.NoCollate.Set(false)
	pdfDoc.PageSize.Set(wkhtmltopdf.PageSizeA3)
	pdfDoc.Orientation.Set(wkhtmltopdf.OrientationLandscape)

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
	err = pdfDoc.WriteFile("out/" + title + ".pdf")
	if err != nil {
		return err
	}
	return nil
}
