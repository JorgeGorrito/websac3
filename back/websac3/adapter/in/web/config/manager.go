package config

import (
	"errors"
	"reflect"
	handler "websac3/adapter/in/web/handler/command"
	"websac3/adapter/in/web/middleware"
	"websac3/app/port/in/dto/command"
	"websac3/app/port/in/usecase"
	"websac3/app/port/out/message"
	"websac3/common/dependencies/container"
	"websac3/common/logging"
	"websac3/common/mapper"
	"websac3/common/mediator"
	"websac3/common/validator"

	"github.com/gin-gonic/gin"
)

type manager struct{}

func NewConfigManager() *manager {
	return &manager{}
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
			container.Inject[validator.Validator](),
			container.Inject[message.Provider](),
			container.Inject[logging.Logger](),
		),
	); err != nil {
		*errorList = errors.Join(*errorList, err)
	}

	if err := iMediator.Register(
		reflect.TypeOf(command.ValidateEmailCommand{}),
		handler.NewValidateEmailCommandHandler(
			container.Inject[usecase.ValidateEmailUseCase](),
			container.Inject[validator.Validator](),
			container.Inject[message.Provider](),
			container.Inject[logging.Logger](),
		),
	); err != nil {
		*errorList = errors.Join(*errorList, err)
	}
}

func (m *manager) ConfigureMappers() {
	mapper.RegisterMapFunctions()
}

func (m *manager) ConfigureApplication() error {
	var errorList error
	m.ConfigureMediator(&errorList)
	m.ConfigureMappers()
	return errorList
}

func (m *manager) ConfigureMiddleware(engine *gin.Engine) error {
	var errorList error
	engine.Use(middleware.LangMiddleware())

	return errorList
}

func (m *manager) ConfigureEngine(engine *gin.Engine) error {
	var errorList error
	m.ConfigureMiddleware(engine)

	return errorList
}
