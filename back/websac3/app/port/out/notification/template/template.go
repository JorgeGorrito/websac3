package template

type Template interface {
	Render(context any) (string, error)
}

type Provider interface {
	GetByName(name string) Template
}
