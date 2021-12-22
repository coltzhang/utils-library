package pdf

import (
	"github.com/signintech/gopdf"
	"github.com/sirupsen/logrus"
	"log"
)

func GenPdf() {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.AddPage()
	err := pdf.AddTTFFont("truetype", "./simhei.ttf")
	if err != nil {
		logrus.Errorf(err.Error())
		return
	}

	err = pdf.SetFont("truetype", "", 14)
	if err != nil {
		log.Print(err.Error())
		return
	}
	pdf.Cell(nil, "您好")
	err = pdf.Text("我是某某某\n")

	pdf.WritePdf("hello.pdf")
}
