package template

import (
	"html/template"
	ptemplate "websac3/app/port/out/notification/template"
)

type Provider struct {
	registry map[string]ptemplate.Template
}

func NewProvider(content map[string]ptemplate.Template) *Provider {
	return &Provider{
		registry: content,
	}
}

func (p *Provider) GetByName(name string) ptemplate.Template {
	if template, ok := p.registry[name]; ok {
		return template
	}
	return nil
}

func (p *Provider) GetByNameAndLang(name string, lang string) ptemplate.Template {
	if t, ok := p.registry[name+":"+lang]; ok {
		return t
	}
	if t, ok := p.registry[name+":es"]; ok {
		return t
	}
	return p.GetByName(name)
}

type base struct {
	template *template.Template
}

func (b *base) loadTemplate(templatePath string) {
	template, err := template.ParseFiles(templatePath)
	if err != nil {
		return
	}
	b.template = template
}
