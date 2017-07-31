package main

import (
	"fmt"
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
	pdf.AddPage()
	pdf.SetFont("Arial", "", 11)
	pagew, pageh := pdf.GetPageSize()
	mleft, mright, _, mbottom := pdf.GetMargins()

	pdf.Image("data/nec_logo.png", 10, 10, 50, 0, false, "", 0, "")
	pdf.Cell(99, 10, customerName)
	pdf.AddPage()
	cols := []float64{60, 100, pagew - mleft - mright - 100 - 60}
	fmt.Println(cols)
	// rows := [][]string{}
	var rows = idrive.GetCharges()

	for i, row := range rows {
		curx, y := pdf.GetXY()
		x := curx
		fmt.Println(x)
		height := 0.
		_, lineHt := pdf.GetFontSize()
		// fmt.Println(row)

		// Add a new page if the height of the row doesn't fit on the page
		if pdf.GetY()+height > pageh-mbottom {
			pdf.AddPage()
			y = pdf.GetY()
		}
		// for i, txt := range row {
		// 	width := cols[i]
		// 	pdf.Rect(x, y, width, height, "")
		// 	pdf.MultiCell(width, lineHt+marginCell, txt, "", "", false)
		// 	x += width
		// 	pdf.SetXY(x, y)
		// }
		pdf.SetXY(curx, y+height)

	}
	err := pdf.OutputFileAndClose("hello.pdf")
	if err != nil {
		log.Fatal(err)
	}

	idrive.GetCharges()
}
