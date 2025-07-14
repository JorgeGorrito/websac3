package dependencies

import (
	"github.com/JorgeGorrito/anise-dependency-injection/andi"
	"os"
	"strings"
	"websac3/adapter/out/message"
	imsg "websac3/app/port/out/message"
)

func (m *manager) registerMessageProviderDependencies() {
	var configuredLangs string = os.Getenv("MESSAGES_LANGUAGES")
	var messagesLanguages []string = strings.Split(configuredLangs, ",")
	m.binder.Bind(
		andi.GetAbstractType[imsg.Provider](),
		func() any {
			return message.NewProvider(
				os.Getenv("MESSAGES_PATH"),
				messagesLanguages,
			)
		},
	)
}
