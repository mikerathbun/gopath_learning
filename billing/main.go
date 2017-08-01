package main

import (
	"log"

	"github.com/jung-kurt/gofpdf"

	"github.com/mikerathbun/gopath_learning/billing/idrive"
)

const customerName = "NEC"

func RunCharges() {

	idrive.RunCharges()

}

func main() {
	createPDFInvoice()

	// RunCharges()
}

func createPDFInvoice() {
	// marginCell := 2. // Margin of top/bottom of cell
	pdf := gofpdf.New("P", "mm", "A4", "")
	var width, _ = pdf.GetPageSize()
	pdf.AddPage()
	pdf.SetFont("Arial", "", 11)
	// pagew, pageh := pdf.GetPageSize()
	// mleft, mright, _, mbottom := pdf.GetMargins()

	pdf.Image("data/nec_logo.png", 10, 10, 50, 0, false, "", 0, "")
	pdf.Cell(99, 10, customerName)
	pdf.AddPage()

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
	err := pdf.OutputFileAndClose("hello.pdf")
	if err != nil {
		log.Fatal(err)
	}

	idrive.GetCharges()
}
