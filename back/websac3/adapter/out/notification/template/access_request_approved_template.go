package template

import (
	"bytes"
	"fmt"
	"websac3/app/domain/notification/template/context"
)

type AccessRequestApprovedTemplate struct{ base }

func NewAccessRequestApprovedTemplate(templatePath string) *AccessRequestApprovedTemplate {
	t := &base{}
	t.loadTemplate(templatePath)
	return &AccessRequestApprovedTemplate{
		base: *t,
	}
}

func (t *AccessRequestApprovedTemplate) Render(ctx any) (string, error) {
	c, ok := ctx.(*context.AccessRequestApproved)
	if !ok {
		return "", fmt.Errorf("invalid context type, expected *AccessRequestApprovedContext")
	}

	var buf bytes.Buffer
	err := t.template.Execute(&buf, c)
	if err != nil {
		return "", err
	}

	render := buf.String()
	return render, nil
}
