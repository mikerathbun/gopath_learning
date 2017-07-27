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
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Hello PDF World!!")
	err := pdf.OutputFileAndClose("hello.pdf")
	if err != nil {
		log.Fatal(err)
	}
}
