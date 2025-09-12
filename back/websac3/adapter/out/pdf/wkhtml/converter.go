package wkhtml

import (
	"strings"

	wkhtml "github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

type Converter struct{}

func New() *Converter { return &Converter{} }

func (c *Converter) HTMLToPDF(html string) ([]byte, error) {
	pdfg, err := wkhtml.NewPDFGenerator()
	if err != nil {
		return nil, err
	}
	page := wkhtml.NewPageReader(strings.NewReader(html))
	page.EnableLocalFileAccess.Set(true)
	pdfg.AddPage(page)
	if err := pdfg.Create(); err != nil {
		return nil, err
	}
	return pdfg.Bytes(), nil
}
