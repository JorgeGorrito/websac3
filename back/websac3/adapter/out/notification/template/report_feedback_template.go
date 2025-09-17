package template

import (
	"bytes"
	"fmt"
	"websac3/app/domain/notification/template/context"
)

type ReportFeedbackTemplate struct{ base }

func NewReportFeedbackTemplate(templatePath string) *ReportFeedbackTemplate {
	t := &base{}
	t.loadTemplate(templatePath)
	return &ReportFeedbackTemplate{
		base: *t,
	}
}

func (t *ReportFeedbackTemplate) Render(data any) (string, error) {
	ctx, ok := data.(*context.ReportFeedbackContext)
	if !ok {
		return "", fmt.Errorf("invalid context type, expected *ReportFeedbackContext")
	}

	var buf bytes.Buffer
	err := t.template.Execute(&buf, ctx)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
