package config

import (
	"errors"
	"reflect"
	handler "websac3/adapter/in/web/handler/command"
	"websac3/app/port/in/dto/command"
	"websac3/app/port/in/usecase"
	"websac3/app/port/out/logging"
	"websac3/common/dependencies/container"
	"websac3/common/mediator"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type manager struct{}

func NewConfigManager() *manager {
	return &manager{}
}

func (m *manager) loadEnviroment(errorList *error) {
	if err := godotenv.Load(); err != nil {
		*errorList = errors.Join(*errorList, err)
	}
}

func (m *manager) ConfigureMediator(errorList *error) {
	var iMediator mediator.IMediator = container.Inject[mediator.IMediator]()
	if iMediator == nil {
		err := errors.New("mediator is not configured")
		*errorList = errors.Join(*errorList, err)
	}

	if err := iMediator.Register(
		reflect.TypeOf(command.CreateAccessRequestCommand{}),
		handler.NewCreateAccessRequestCommandHandler(
			container.Inject[usecase.CreateAccessRequestUseCase](),
			container.Inject[logging.Logger](),
		),
	); err != nil {
		*errorList = errors.Join(*errorList, err)
	}
}

func (m *manager) ConfigureApplication() error {
	var errorList error
	m.loadEnviroment(&errorList)
	m.ConfigureMediator(&errorList)
	return errorList
}

func (m *manager) ConfigureEngine(engine *gin.Engine) error {
	var errorList error

	return errorList
}
