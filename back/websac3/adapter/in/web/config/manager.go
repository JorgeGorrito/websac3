package config

import (
	"errors"
	"reflect"
	chandler "websac3/adapter/in/web/handler/command"
	qhandler "websac3/adapter/in/web/handler/query"
	"websac3/adapter/in/web/middleware"
	"websac3/app/port/in/dto/command"
	"websac3/app/port/in/dto/query"
	"websac3/app/port/in/usecase"
	"websac3/app/port/out/message"
	"websac3/common/dependencies/container"
	"websac3/common/jwt"
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
		chandler.NewCreateAccessRequestCommandHandler(
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
		chandler.NewValidateEmailCommandHandler(
			container.Inject[usecase.ValidateEmailUseCase](),
			container.Inject[validator.Validator](),
			container.Inject[message.Provider](),
			container.Inject[logging.Logger](),
		),
	); err != nil {
		*errorList = errors.Join(*errorList, err)
	}

	if err := iMediator.Register(
		reflect.TypeOf(command.LoginCommand{}),
		chandler.NewLoginHandler(
			container.Inject[usecase.LoginUseCase](),
			container.Inject[jwt.Generator](),
			container.Inject[message.Provider](),
			container.Inject[validator.Validator](),
			container.Inject[logging.Logger](),
		),
	); err != nil {
		*errorList = errors.Join(*errorList, err)
	}

	if err := iMediator.Register(
		reflect.TypeOf(command.RefreshTokenCommand{}),
		chandler.NewRefreshTokenHandler(
			container.Inject[usecase.GetUserByIDUseCase](),
			container.Inject[jwt.Generator](),
			container.Inject[message.Provider](),
			container.Inject[validator.Validator](),
			container.Inject[logging.Logger](),
		),
	); err != nil {
		*errorList = errors.Join(*errorList, err)
	}

	if err := iMediator.Register(
		reflect.TypeOf(query.ListAccessRequestQuery{}),
		qhandler.NewListAccessRequestQueryHandler(
			container.Inject[usecase.ListAccessRequestUseCase](),
			container.Inject[validator.Validator](),
			container.Inject[message.Provider](),
			container.Inject[logging.Logger](),
		),
	); err != nil {
		*errorList = errors.Join(*errorList, err)
	}

	if err := iMediator.Register(
		reflect.TypeOf(command.ApproveAccessRequestCommand{}),
		chandler.NewApproveAccessRequestCommandHandler(
			container.Inject[usecase.ApproveAccessRequestUseCase](),
			container.Inject[validator.Validator](),
			container.Inject[message.Provider](),
			container.Inject[logging.Logger](),
		),
	); err != nil {
		*errorList = errors.Join(*errorList, err)
	}

	if err := iMediator.Register(
		reflect.TypeOf(command.RejectAccessRequestCommand{}),
		chandler.NewRejectAccessRequestCommandHandler(
			container.Inject[usecase.RejectAccessRequestUseCase](),
			container.Inject[validator.Validator](),
			container.Inject[message.Provider](),
			container.Inject[logging.Logger](),
		),
	); err != nil {
		*errorList = errors.Join(*errorList, err)
	}

	if err := iMediator.Register(
		reflect.TypeOf(query.ListIdentificationTypeQuery{}),
		qhandler.NewListIdentificationTypeQueryHandler(
			container.Inject[usecase.ListIdentificationTypeUseCase](),
			container.Inject[validator.Validator](),
			container.Inject[message.Provider](),
			container.Inject[logging.Logger](),
		),
	); err != nil {
		*errorList = errors.Join(*errorList, err)
	}

	if err := iMediator.Register(
		reflect.TypeOf(query.ListHigherEducationInstitutionQuery{}),
		qhandler.NewListHigherEducationInstitutionQueryHandler(
			container.Inject[usecase.ListHigherEducationInstitutionUseCase](),
			container.Inject[message.Provider](),
			container.Inject[validator.Validator](),
			container.Inject[logging.Logger](),
		),
	); err != nil {
		*errorList = errors.Join(*errorList, err)
	}

	if err := iMediator.Register(
		reflect.TypeOf(query.ListTopicQuery{}),
		qhandler.NewListTopicQueryHandler(
			container.Inject[usecase.ListTopicUseCase](),
			container.Inject[message.Provider](),
			container.Inject[logging.Logger](),
			container.Inject[validator.Validator](),
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
