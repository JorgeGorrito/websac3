package template

import (
	"bytes"
	"fmt"
	"websac3/app/domain/notification/template/context"
)

type ExpertConsultationNotificationTemplate struct {
	base
}

func NewExpertConsultationNotificationTemplate(templatePath string) *ExpertConsultationNotificationTemplate {
	t := &base{}
	t.loadTemplate(templatePath)
	return &ExpertConsultationNotificationTemplate{
		base: *t,
	}
}

func (t *ExpertConsultationNotificationTemplate) Render(ctx any) (string, error) {
	c, ok := ctx.(context.ExpertConsultationNotificationContext)
	if !ok {
		// Intentar con puntero
		cPtr, okPtr := ctx.(*context.ExpertConsultationNotificationContext)
		if !okPtr {
			return "", fmt.Errorf("invalid context type, expected ExpertConsultationNotificationContext or *ExpertConsultationNotificationContext")
		}
		c = *cPtr
	}

	var buf bytes.Buffer
	err := t.template.Execute(&buf, c)
	if err != nil {
		return "", err
	}

	render := buf.String()
	return render, nil
}
