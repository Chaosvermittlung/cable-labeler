package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/jung-kurt/gofpdf"
)

const (
	LabelWidth  = 60
	LabelHeight = 29
	PrintWidth  = 50
	PrintHeight = 26
)

func printFlagText(pdf *gofpdf.Fpdf, x, y float64, length string) {
	pdf.MoveTo(x, y)
	pdf.CellFormat(PrintWidth/2, PrintWidth/2, length, "", 1, "CT", false, 0, "")
	pdf.ImageOptions("chaosvermittlung_rev1_bw.png", x+PrintWidth/4-PrintHeight/8, y+PrintHeight/4, 0, PrintHeight/4, false, gofpdf.ImageOptions{ImageType: "PNG", ReadDpi: true}, 0, "")
}

func printFlag(length string) {
	pdf := gofpdf.NewCustom(&gofpdf.InitType{
		UnitStr:        "mm",
		Size:           gofpdf.SizeType{Wd: LabelWidth, Ht: LabelHeight},
		OrientationStr: "P",
	})
	pdf.RegisterImageOptions("chaosvermittlung_rev1_bw.png", gofpdf.ImageOptions{ImageType: "PNG", ReadDpi: true})
	pdf.SetAutoPageBreak(false, 0)
	pdf.SetFont("Arial", "", 15)
	pdf.AddPage()
	topMargin := (LabelHeight - PrintHeight) / float64(2)
	leftMargin := (LabelWidth - PrintWidth) / float64(2)
	pdf.Rect(leftMargin, topMargin, PrintWidth, PrintHeight, "D")
	pdf.SetDashPattern([]float64{1, 0.5}, 0)
	pdf.Line(LabelWidth/2, topMargin, LabelWidth/2, topMargin+PrintHeight)
	pdf.SetDashPattern([]float64{0.5, 0.5}, 0)
	pdf.Line(leftMargin, LabelHeight/2, leftMargin+PrintWidth, LabelHeight/2)
	pdf.SetDashPattern([]float64{0, 0}, 0)
	length = length + "m"
	printFlagText(pdf, leftMargin, topMargin, length)
	printFlagText(pdf, leftMargin+PrintWidth/2, topMargin, length)
	printFlagText(pdf, leftMargin, topMargin+PrintHeight/2, length)
	printFlagText(pdf, leftMargin+PrintWidth/2, topMargin+PrintHeight/2, length)
	pdfname := time.Now().Format("2006-01-02_15:04:05") + ".pdf"
	err := pdf.OutputFileAndClose(pdfname)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Please enter first length")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if text == "" {
			os.Exit(0)
		}
		start, err := strconv.Atoi(text)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Please enter second length")
		text, _ = reader.ReadString('\n')
		text = strings.TrimSpace(text)
		end := 0
		if text == "" {
			end = 0
		} else {
			end, err = strconv.Atoi(text)
			if err != nil {
				log.Fatal(err)
			}
		}
		cablelength := int(math.Abs(float64(end - start)))
		fmt.Println("Cable length: ", cablelength, "m \n")
		//printFlag(strconv.Itoa(cabellength))
	}
}
