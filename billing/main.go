package main

import (
	"log"

	"github.com/jung-kurt/gofpdf"

	"github.com/mikerathbun/gopath_learning/billing/idrive"
)

const (
	customerName = "Network Enterprise Center"
	invoiceTitle = "Information Technology FY2017 Invoice"
	invoiceDates = "From 10/1/2016 To 9/30/2017"
)

var pdf = gofpdf.New("P", "mm", "A4", "")

// var pdf gofpdf

func RunCharges() {

	idrive.RunCharges()

}
func init() {
}
func main() {
	createPDFInvoiceHeader()
	// createPDFInvoice()

	err := pdf.OutputFileAndClose("hello.pdf")
	if err != nil {
		log.Fatal(err)
	}
	// RunCharges()
}

func createPDFInvoiceHeader() {

	// pdf := gofpdf.New("P", "mm", "A4", "")
	// width, _ := pdf.GetPageSize()
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 15)
	// pagew, pageh := pdf.GetPageSize()
	// mleft, mright, _, mbottom := pdf.GetMargins()

	pdf.Image("data/nec_logo.png", 10, 10, 50, 0, false, "", 0, "")
	pdf.CellFormat(15, 15, customerName, "0", 0, "TC", false, 0, "")
	pdf.CellFormat(15, 15, invoiceTitle, "0", 0, "TC", false, 0, "")
}

func createPDFInvoice() {
	// marginCell := 2. // Margin of top/bottom of cell
	pdf.AddPage()

	width, _ := pdf.GetPageSize()
	var rows = idrive.GetCharges()
	// CellFormat(w,h,text,borderStr,lineLocation (0-right1-nxtln-2below),alighStr,fill,link,linkstr)
	pdf.CellFormat(40, 7, "Personal Storage", "1", 0, "", false, 0, "")
	pdf.Ln(-1)

	for _, row := range rows {
		pdf.CellFormat(width-35, 6, row.String(), "1", 0, "", false, 0, "")
		pdf.CellFormat(20, 6, row.Cost(), "1", 0, "", false, 0, "")

		pdf.Ln(-1)
		// Add a new page if the height of the row doesn't fit on the page
		// if pdf.GetY()+height > pageh-mbottom {
		// 	pdf.AddPage()
		// 	y = pdf.GetY()
		// }

	}

	idrive.GetCharges()
}
