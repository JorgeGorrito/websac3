package template

import (
	"bytes"
	"fmt"
	"websac3/app/domain/notification/template/context"
)

type AccessRequestRejectedTemplate struct{ base }

func NewAccessRequestRejectedTemplate(templatePath string) *AccessRequestRejectedTemplate {
	t := &base{}
	t.loadTemplate(templatePath)
	return &AccessRequestRejectedTemplate{
		base: *t,
	}
}

func (t *AccessRequestRejectedTemplate) Render(ctx any) (string, error) {
	c, ok := ctx.(*context.AccessRequestRejected)
	if !ok {
		return "", fmt.Errorf("invalid context type, expected *AccessRequestRejectedContext")
	}

	var buf bytes.Buffer
	err := t.template.Execute(&buf, c)
	if err != nil {
		return "", err
	}

	render := buf.String()
	return render, nil
}
