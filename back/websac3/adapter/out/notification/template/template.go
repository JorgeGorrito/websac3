package template

import (
	"fmt"
	"html/template"
	"path/filepath"
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
	template, ok := p.registry[name+":"+lang]
	if !ok {
		return nil
	}
	return template
}

type base struct {
	template *template.Template
}

func (b *base) loadTemplate(templatePath string) {
	funcMap := template.FuncMap{
		"mul": func(a, b interface{}) float64 {
			var fa, fb float64
			switch v := a.(type) {
			case float32:
				fa = float64(v)
			case float64:
				fa = v
			case int:
				fa = float64(v)
			}
			switch v := b.(type) {
			case float32:
				fb = float64(v)
			case float64:
				fb = v
			case int:
				fb = float64(v)
			}
			return fa * fb
		},
		"div": func(a, b interface{}) float64 {
			var fa, fb float64
			switch v := a.(type) {
			case float32:
				fa = float64(v)
			case float64:
				fa = v
			case int:
				fa = float64(v)
			}
			switch v := b.(type) {
			case float32:
				fb = float64(v)
			case float64:
				fb = v
			case int:
				fb = float64(v)
			}
			if fb == 0 {
				return 0
			}
			return fa / fb
		},
		"printf": func(format string, args ...interface{}) string {
			return fmt.Sprintf(format, args...)
		},
		"ge": func(a, b interface{}) bool {
			var fa, fb float64
			switch v := a.(type) {
			case float32:
				fa = float64(v)
			case float64:
				fa = v
			case int:
				fa = float64(v)
			}
			switch v := b.(type) {
			case float32:
				fb = float64(v)
			case float64:
				fb = v
			case int:
				fb = float64(v)
			}
			return fa >= fb
		},
	}

	// Crear template con funciones personalizadas y parsear el archivo
	tmpl, err := template.New("").Funcs(funcMap).ParseFiles(templatePath)
	if err != nil {
		return
	}

	// Obtener el nombre del archivo sin la ruta
	fileName := filepath.Base(templatePath)

	// El template principal será el nombre del archivo
	b.template = tmpl.Lookup(fileName)
	if b.template == nil {
		// Si no encuentra el template por nombre, usar el primer template disponible
		b.template = tmpl.Templates()[0]
	}
}
