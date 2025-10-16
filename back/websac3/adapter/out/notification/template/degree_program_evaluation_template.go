package template

import (
	"bytes"
	"fmt"
	"websac3/app/domain/notification/template/context"
)

type DegreeProgramEvaluationTemplate struct {
	base
}

func NewDegreeProgramEvaluationTemplate(templatePath string) *DegreeProgramEvaluationTemplate {
	t := &base{}
	t.loadTemplate(templatePath)
	return &DegreeProgramEvaluationTemplate{
		base: *t,
	}
}

func (t *DegreeProgramEvaluationTemplate) Render(ctx any) (string, error) {
	c, ok := ctx.(*context.DegreeProgramEvaluation)
	if !ok {
		return "", fmt.Errorf("invalid context type, expected *DegreeProgramEvaluation")
	}

	var buf bytes.Buffer
	err := t.template.Execute(&buf, c)
	if err != nil {
		return "", err
	}

	render := buf.String()
	return render, nil
}



















