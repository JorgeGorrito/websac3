package template

import (
	"bytes"
	"fmt"
	"websac3/app/domain/entity"
)

type ReportTemplate struct{ base }

func NewReportTemplate(templatePath string) *ReportTemplate {
	t := &base{}
	t.loadTemplate(templatePath)
	return &ReportTemplate{base: *t}
}

func (t *ReportTemplate) Render(ctx any) (string, error) {
	r, ok := ctx.(*entity.Report)
	if !ok {
		return "", fmt.Errorf("invalid context type, expected *entity.Report")
	}
	var buf bytes.Buffer
	if err := t.template.Execute(&buf, r); err != nil {
		return "", err
	}
	return buf.String(), nil
}
