package message

type Context interface {
	GetMessage(namespace string, key string, params ...any) string
}

type Provider interface {
	WithLang(lang string) Context
}
