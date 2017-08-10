package main

import (
	"fmt"
	"log"

	"github.com/jung-kurt/gofpdf"

	"github.com/mikerathbun/gopath_learning/billing/idrive"
)

const (
	customerName = "502 Force Support Squadron"
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
	// createPDFInvoiceLines()
	myTestList()

	err := pdf.OutputFileAndClose("hello.pdf")
	if err != nil {
		log.Fatal(err)
	}
	// RunCharges()
}
func myTestList() {
	pdf.AddPage()

	width, _ := pdf.GetPageSize()
	pdf.SetFont("Helvetica", "", 8)
	fmt.Printf("The size is %f\n", width)
	currx, curry := pdf.GetXY()
	fmt.Printf("1: x:%f y:%f\n", currx, curry)
	pdf.CellFormat(width-35, 6, "Mike.rathbun I drive size 54345.MB", "", 0, "", false, 0, "")
	currx, curry = pdf.GetXY()
	fmt.Printf("2: x:%f y:%f\n", currx, curry)
	pdf.CellFormat(20, 6, "8/30/2017", "", 0, "", false, 0, "")
	currx, curry = pdf.GetXY()
	fmt.Printf("3: x:%f y:%f\n", currx, curry)
	pdf.CellFormat(20, 6, "Gigabyte(s)", "", 0, "", false, 0, "")
	currx, curry = pdf.GetXY()
	fmt.Printf("4: x:%f y:%f\n", currx, curry)
	pdf.CellFormat(20, 6, "$5.00", "", 0, "", false, 0, "")
	currx, curry = pdf.GetXY()
	fmt.Printf("5: x:%f y:%f\n", currx, curry)
	pdf.CellFormat(20, 6, "$100.00", "", 0, "", false, 0, "")

	pdf.Ln(-1)

}
func createPDFInvoiceHeader() {

	// pdf := gofpdf.New("P", "mm", "A4", "")
	// width, _ := pdf.GetPageSize()
	pdf.AddPage()
	pdf.SetFont("Helvetica", "", 8)
	// pagew, pageh := pdf.GetPageSize()
	// mleft, mright, _, mbottom := pdf.GetMargins()

	pdf.Image("data/nec_logo.png", 10, 10, 50, 0, false, "", 0, "")
	pdf.CellFormat(15, 15, customerName, "0", 0, "TC", false, 0, "")
	pdf.CellFormat(15, 15, invoiceTitle, "0", 0, "TC", false, 0, "")
}

func createPDFInvoiceLines() {
	// marginCell := 2. // Margin of top/bottom of cell
	pdf.AddPage()

	width, _ := pdf.GetPageSize()
	var rows = idrive.GetCharges()
	// CellFormat(w,h,text,borderStr,lineLocation (0-right1-nxtln-2below),alighStr,fill,link,linkstr)
	pdf.CellFormat(40, 7, "Personal Storage", "", 0, "", false, 0, "")
	pdf.Ln(-1)

	for _, row := range rows {
		pdf.CellFormat(width-35, 6, row.String(), "", 0, "", false, 0, "")
		pdf.CellFormat(20, 6, row.Cost(), "", 0, "", false, 0, "")
		pdf.CellFormat(20, 6, row.Cost(), "", 0, "", false, 0, "")

		pdf.Ln(-1)
		// Add a new page if the height of the row doesn't fit on the page
		// if pdf.GetY()+height > pageh-mbottom {
		// 	pdf.AddPage()
		// 	y = pdf.GetY()
		// }

	}

	// idrive.GetCharges()
}
