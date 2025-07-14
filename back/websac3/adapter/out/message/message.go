package message

import (
	"fmt"
	imsg "websac3/app/port/out/message"
	"websac3/common/decoder"
)

// Implementación de Context
type context struct {
	messages map[string]map[string]string
}

func (c *context) GetMessage(namespace string, key string, params ...any) string {
	if ns, ok := c.messages[namespace]; ok {
		if msg, ok := ns[key]; ok {
			if len(params) > 0 {
				return fmt.Sprintf(msg, params...)
			}
			return msg
		}
	}
	return ""
}

func NewContext(messages map[string]map[string]string) *context {
	return &context{messages: messages}
}

// Implementación de Provider
type provider struct {
	messages map[string]map[string]map[string]string
}

func (p *provider) WithLang(lang string) imsg.Context {
	if messagesLang, ok := p.messages[lang]; ok {
		return NewContext(messagesLang)
	}
	return NewContext(make(map[string]map[string]string))
}

func NewProvider(messagesDirPath string, langs []string) imsg.Provider {
	decoder := decoder.Json()
	providerCreated := &provider{
		messages: make(map[string]map[string]map[string]string),
	}

	for _, lang := range langs {
		temp := make(map[string]map[string]string)
		path := fmt.Sprintf("%s%s/messages.json", messagesDirPath, lang)

		if err := decoder.Decode(path, &temp); err != nil {
			panic("failed to load messages for " + lang + ": " + err.Error())
		}

		providerCreated.messages[lang] = temp
	}

	return providerCreated
}
