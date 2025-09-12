package service

import (
	ptemplate "websac3/app/port/out/notification/template"
	ppdf "websac3/app/port/out/pdf"
)

type ReportPDFService struct {
	templates ptemplate.Provider
	pdf       ppdf.Converter
}

func NewReportPDFService(t ptemplate.Provider, p ppdf.Converter) *ReportPDFService {
	return &ReportPDFService{templates: t, pdf: p}
}

func (s *ReportPDFService) Generate(lang string, report any) ([]byte, error) {
	t := s.templates.GetByNameAndLang("report", lang)
	html, err := t.Render(report)
	if err != nil {
		return nil, err
	}
	return s.pdf.HTMLToPDF(html)
}
