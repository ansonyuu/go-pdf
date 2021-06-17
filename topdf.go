package main

import (
	_ "fmt"
	"io/ioutil"
	"log"
	"github.com/jung-kurt/gofpdf"
)

func main() {
	file := "test.txt"
	// reading in file to converter
	content, err := ioutil.ReadFile(file)

	//error checking
	if err != nil {
		log.Fatalf("%s file not found", file)
	}

	// parameters are (orientation, unit, size, font directory)
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	// parameters are (font, style, size)
	pdf.SetFont("Arial", "B", 16)

	// parameters are (width, line height, string, border, alignment)
	pdf.MultiCell(190, 5, string(content), "0", "0", false)

	_ = pdf.OutputFileAndClose("test.pdf")
}
