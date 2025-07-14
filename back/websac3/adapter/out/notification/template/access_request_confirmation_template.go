package template

import (
	"bytes"
	"fmt"
	"websac3/app/domain/notification/template/context"
)

type AccessRequestConfirmationTemplate struct {
	base
}

func NewAccessRequestConfirmationTemplate(templatePath string) *AccessRequestConfirmationTemplate {
	t := &base{}
	t.loadTemplate(templatePath)
	return &AccessRequestConfirmationTemplate{
		base: *t,
	}
}

func (t *AccessRequestConfirmationTemplate) Render(ctx any) (string, error) {
	c, ok := ctx.(*context.AccessRequestConfirmation)
	if !ok {
		return "", fmt.Errorf("invalid context type, expected *AccessRequestConfirmationContext")
	}

	var buf bytes.Buffer
	err := t.template.Execute(&buf, c)
	if err != nil {
		return "", err
	}

	render := buf.String()
	return render, nil
}
