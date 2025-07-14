package dependencies

import (
	"github.com/JorgeGorrito/anise-dependency-injection/andi"
	"websac3/common/logging"

	"os"
	"strconv"
)

func (m *manager) registerLoggerDependencies() {
	m.binder.Bind(
		andi.GetAbstractType[logging.Logger](),
		func() any {
			var appName string = os.Getenv("APP_NAME")
			var logfilePath string = os.Getenv("LOGGER_PATH")
			bufferSize, err := strconv.Atoi(os.Getenv("LOGGER_BUFFER_SIZE"))
			if err != nil {
				panic(err)
			}
			maxEntries, err := strconv.Atoi(os.Getenv("LOGGER_MAX_ENTRIES"))
			if err != nil {
				panic(err)
			}
			logger, err := logging.NewLogger(
				appName,
				logfilePath,
				bufferSize,
				maxEntries,
			)
			if err != nil {
				panic("Failed to create logger: " + err.Error())
			}
			return logger
		},
	)
}
