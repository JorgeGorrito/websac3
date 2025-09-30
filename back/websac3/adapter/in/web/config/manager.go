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
		reflect.TypeOf(command.EvaluateDegreeProgramCommand{}),
		chandler.NewEvaluateDegreeProgramCommandHandler(
			container.Inject[validator.Validator](),
			container.Inject[logging.Logger](),
			container.Inject[usecase.EvaluateDegreeProgramUseCase](),
			container.Inject[message.Provider](),
		),
	); err != nil {
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
		reflect.TypeOf(query.ListUserAccessRequestsQuery{}),
		qhandler.NewListUserAccessRequestsQueryHandler(
			container.Inject[validator.Validator](),
			container.Inject[logging.Logger](),
			container.Inject[usecase.ListUserAccessRequestsUseCase](),
			container.Inject[message.Provider](),
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

	if err := iMediator.Register(
		reflect.TypeOf(query.ListDurationUnitQuery{}),
		qhandler.NewListDurationUnitQueryHandler(
			container.Inject[usecase.ListDurationUnitUseCase](),
			container.Inject[message.Provider](),
			container.Inject[logging.Logger](),
			container.Inject[validator.Validator](),
		),
	); err != nil {
		*errorList = errors.Join(*errorList, err)
	}

	if err := iMediator.Register(
		reflect.TypeOf(query.ListFormationLevelQuery{}),
		qhandler.NewListFormationLevelQueryHandler(
			container.Inject[usecase.ListFormationLevelUseCase](),
			container.Inject[message.Provider](),
			container.Inject[logging.Logger](),
			container.Inject[validator.Validator](),
		),
	); err != nil {
		*errorList = errors.Join(*errorList, err)
	}

	if err := iMediator.Register(
		reflect.TypeOf(command.CreateDegreeProgramCommand{}),
		chandler.NewCreateDegreeProgramCommandHandler(
			container.Inject[usecase.CreateDegreeProgramUseCase](),
			container.Inject[validator.Validator](),
			container.Inject[message.Provider](),
			container.Inject[logging.Logger](),
		),
	); err != nil {
		*errorList = errors.Join(*errorList, err)
	}

	if err := iMediator.Register(
		reflect.TypeOf(query.ListDegreeProgramQuery{}),
		qhandler.NewListDegreeProgramQueryHandler(
			container.Inject[usecase.ListDegreeProgramUseCase](),
			container.Inject[message.Provider](),
			container.Inject[logging.Logger](),
			container.Inject[validator.Validator](),
		),
	); err != nil {
		*errorList = errors.Join(*errorList, err)
	}

	if err := iMediator.Register(
		reflect.TypeOf(command.CreateCourseCommand{}),
		chandler.NewCreateCourseCommandHandler(
			container.Inject[usecase.CreateCourseUseCase](),
			container.Inject[validator.Validator](),
			container.Inject[message.Provider](),
			container.Inject[logging.Logger](),
		),
	); err != nil {
		*errorList = errors.Join(*errorList, err)
	}

	if err := iMediator.Register(
		reflect.TypeOf(command.DeleteCourseCommand{}),
		chandler.NewDeleteCourseCommandHandler(
			container.Inject[usecase.DeleteCourseUseCase](),
			container.Inject[message.Provider](),
			container.Inject[logging.Logger](),
			container.Inject[validator.Validator](),
		),
	); err != nil {
		*errorList = errors.Join(*errorList, err)
	}

	if err := iMediator.Register(
		reflect.TypeOf(command.UpdateCourseCommand{}),
		chandler.NewUpdateCourseCommandHandler(
			container.Inject[usecase.UpdateCourseUseCase](),
			container.Inject[message.Provider](),
			container.Inject[logging.Logger](),
			container.Inject[validator.Validator](),
		),
	); err != nil {
		*errorList = errors.Join(*errorList, err)
	}

	if err := iMediator.Register(
		reflect.TypeOf(command.CreateUserFromTokenCommand{}),
		chandler.NewCreateUserFromTokenCommandHandler(
			container.Inject[usecase.CreateUserFromTokenUseCase](),
			container.Inject[validator.Validator](),
			container.Inject[message.Provider](),
			container.Inject[logging.Logger](),
		),
	); err != nil {
		*errorList = errors.Join(*errorList, err)
	}

	if err := iMediator.Register(
		reflect.TypeOf(query.ListReportsByDegreeProgramQuery{}),
		qhandler.NewListReportsByDegreeProgramQueryHandler(
			container.Inject[usecase.ListReportsByDegreeProgramUseCase](),
			container.Inject[message.Provider](),
			container.Inject[logging.Logger](),
			container.Inject[validator.Validator](),
		),
	); err != nil {
		*errorList = errors.Join(*errorList, err)
	}

	if err := iMediator.Register(
		reflect.TypeOf(query.GetReportByIDQuery{}),
		qhandler.NewGetReportByIDQueryHandler(
			container.Inject[usecase.GetReportByIDUseCase](),
			container.Inject[message.Provider](),
			container.Inject[logging.Logger](),
			container.Inject[validator.Validator](),
		),
	); err != nil {
		*errorList = errors.Join(*errorList, err)
	}

	if err := iMediator.Register(
		reflect.TypeOf(query.ListProfessionalRoleQuery{}),
		qhandler.NewListProfessionalRoleQueryHandler(
			container.Inject[usecase.ListProfessionalRoleUseCase](),
			container.Inject[validator.Validator](),
			container.Inject[message.Provider](),
			container.Inject[logging.Logger](),
		),
	); err != nil {
		*errorList = errors.Join(*errorList, err)
	}

	if err := iMediator.Register(
		reflect.TypeOf(query.ListCourseByDegreeProgramQuery{}),
		qhandler.NewListCourseByDegreeProgramQueryHandler(
			container.Inject[usecase.ListCourseModelsByDegreeProgramUseCase](),
			container.Inject[message.Provider](),
			container.Inject[logging.Logger](),
			container.Inject[validator.Validator](),
		),
	); err != nil {
		*errorList = errors.Join(*errorList, err)
	}

	if err := iMediator.Register(
		reflect.TypeOf(query.GetCourseByIDQuery{}),
		qhandler.NewGetCourseByIDQueryHandler(
			container.Inject[usecase.GetCourseByIDUseCase](),
			container.Inject[message.Provider](),
			container.Inject[logging.Logger](),
			container.Inject[validator.Validator](),
		),
	); err != nil {
		*errorList = errors.Join(*errorList, err)
	}

	if err := iMediator.Register(
		reflect.TypeOf(query.GetCourseTopicsByCourseIDQuery{}),
		qhandler.NewGetCourseTopicsByCourseIDQueryHandler(
			container.Inject[usecase.GetCourseTopicsByCourseIDUseCase](),
			container.Inject[message.Provider](),
			container.Inject[logging.Logger](),
			container.Inject[validator.Validator](),
		),
	); err != nil {
		*errorList = errors.Join(*errorList, err)
	}

	if err := iMediator.Register(
		reflect.TypeOf(query.ListCourseTypesQuery{}),
		qhandler.NewListCourseTypesQueryHandler(
			container.Inject[usecase.ListCourseTypesUseCase](),
			container.Inject[message.Provider](),
			container.Inject[logging.Logger](),
			container.Inject[validator.Validator](),
		),
	); err != nil {
		*errorList = errors.Join(*errorList, err)
	}

	if err := iMediator.Register(
		reflect.TypeOf(query.ListCourseNaturesQuery{}),
		qhandler.NewListCourseNaturesQueryHandler(
			container.Inject[usecase.ListCourseNaturesUseCase](),
			container.Inject[message.Provider](),
			container.Inject[logging.Logger](),
			container.Inject[validator.Validator](),
		),
	); err != nil {
		*errorList = errors.Join(*errorList, err)
	}

	if err := iMediator.Register(
		reflect.TypeOf(command.CreateReportFeedbackCommand{}),
		chandler.NewCreateReportFeedbackCommandHandler(
			container.Inject[validator.Validator](),
			container.Inject[logging.Logger](),
			container.Inject[usecase.CreateReportFeedbackUseCase](),
			container.Inject[message.Provider](),
		),
	); err != nil {
		*errorList = errors.Join(*errorList, err)
	}

	if err := iMediator.Register(
		reflect.TypeOf(query.GetReportFeedbackQuery{}),
		qhandler.NewGetReportFeedbackQueryHandler(
			container.Inject[validator.Validator](),
			container.Inject[logging.Logger](),
			container.Inject[usecase.GetReportFeedbackUseCase](),
			container.Inject[message.Provider](),
		),
	); err != nil {
		*errorList = errors.Join(*errorList, err)
	}

	if err := iMediator.Register(
		reflect.TypeOf(query.ListReportsPendingFeedbackQuery{}),
		qhandler.NewListReportsPendingFeedbackQueryHandler(
			container.Inject[validator.Validator](),
			container.Inject[logging.Logger](),
			container.Inject[usecase.ListReportsPendingFeedbackUseCase](),
			container.Inject[message.Provider](),
		),
	); err != nil {
		*errorList = errors.Join(*errorList, err)
	}

	if err := iMediator.Register(
		reflect.TypeOf(query.ListReportFeedbacksQuery{}),
		qhandler.NewListReportFeedbacksQueryHandler(
			container.Inject[validator.Validator](),
			container.Inject[logging.Logger](),
			container.Inject[usecase.ListReportFeedbacksUseCase](),
			container.Inject[message.Provider](),
		),
	); err != nil {
		*errorList = errors.Join(*errorList, err)
	}

	if err := iMediator.Register(
		reflect.TypeOf(query.ListRoleQuery{}),
		qhandler.NewListRoleQueryHandler(
			container.Inject[usecase.ListRoleUseCase](),
			container.Inject[validator.Validator](),
			container.Inject[message.Provider](),
			container.Inject[logging.Logger](),
		),
	); err != nil {
		*errorList = errors.Join(*errorList, err)
	}

	if err := iMediator.Register(
		reflect.TypeOf(query.ListApprovedAccessRequestQuery{}),
		qhandler.NewListApprovedAccessRequestQueryHandler(
			container.Inject[usecase.ListApprovedAccessRequestUseCase](),
			container.Inject[validator.Validator](),
			container.Inject[message.Provider](),
			container.Inject[logging.Logger](),
		),
	); err != nil {
		*errorList = errors.Join(*errorList, err)
	}

	if err := iMediator.Register(
		reflect.TypeOf(query.ListRejectedAccessRequestQuery{}),
		qhandler.NewListRejectedAccessRequestQueryHandler(
			container.Inject[usecase.ListRejectedAccessRequestUseCase](),
			container.Inject[validator.Validator](),
			container.Inject[message.Provider](),
			container.Inject[logging.Logger](),
		),
	); err != nil {
		*errorList = errors.Join(*errorList, err)
	}

	if err := iMediator.Register(
		reflect.TypeOf(query.ListUsersQuery{}),
		qhandler.NewListUsersQueryHandler(
			container.Inject[usecase.ListUsersUseCase](),
			container.Inject[message.Provider](),
			container.Inject[logging.Logger](),
		),
	); err != nil {
		*errorList = errors.Join(*errorList, err)
	}

	if err := iMediator.Register(
		reflect.TypeOf(command.DeactivateUserCommand{}),
		chandler.NewDeactivateUserCommandHandler(
			container.Inject[usecase.DeactivateUserUseCase](),
			container.Inject[message.Provider](),
			container.Inject[logging.Logger](),
			container.Inject[validator.Validator](),
		),
	); err != nil {
		*errorList = errors.Join(*errorList, err)
	}

	if err := iMediator.Register(
		reflect.TypeOf(command.ActivateUserCommand{}),
		chandler.NewActivateUserCommandHandler(
			container.Inject[usecase.ActivateUserUseCase](),
			container.Inject[message.Provider](),
			container.Inject[logging.Logger](),
			container.Inject[validator.Validator](),
		),
	); err != nil {
		*errorList = errors.Join(*errorList, err)
	}

	if err := iMediator.Register(
		reflect.TypeOf(query.ListUserExpertConsultationsQuery{}),
		qhandler.NewListUserExpertConsultationsQueryHandler(
			container.Inject[validator.Validator](),
			container.Inject[logging.Logger](),
			container.Inject[usecase.ListUserExpertConsultationsUseCase](),
			container.Inject[message.Provider](),
		),
	); err != nil {
		*errorList = errors.Join(*errorList, err)
	}

	if err := iMediator.Register(
		reflect.TypeOf(command.CreateExpertConsultationCommand{}),
		chandler.NewCreateExpertConsultationCommandHandler(
			container.Inject[usecase.CreateExpertConsultationUseCase](),
			container.Inject[validator.Validator](),
			container.Inject[message.Provider](),
			container.Inject[logging.Logger](),
		),
	); err != nil {
		*errorList = errors.Join(*errorList, err)
	}

	if err := iMediator.Register(
		reflect.TypeOf(query.ListPendingExpertConsultationsQuery{}),
		qhandler.NewListPendingExpertConsultationsQueryHandler(
			container.Inject[validator.Validator](),
			container.Inject[logging.Logger](),
			container.Inject[usecase.ListPendingExpertConsultationsUseCase](),
			container.Inject[message.Provider](),
		),
	); err != nil {
		*errorList = errors.Join(*errorList, err)
	}

	if err := iMediator.Register(
		reflect.TypeOf(command.AcceptExpertConsultationCommand{}),
		chandler.NewAcceptExpertConsultationCommandHandler(
			container.Inject[validator.Validator](),
			container.Inject[logging.Logger](),
			container.Inject[usecase.AcceptExpertConsultationUseCase](),
			container.Inject[message.Provider](),
		),
	); err != nil {
		*errorList = errors.Join(*errorList, err)
	}

	if err := iMediator.Register(
		reflect.TypeOf(command.RejectExpertConsultationCommand{}),
		chandler.NewRejectExpertConsultationCommandHandler(
			container.Inject[validator.Validator](),
			container.Inject[logging.Logger](),
			container.Inject[usecase.RejectExpertConsultationUseCase](),
			container.Inject[message.Provider](),
		),
	); err != nil {
		*errorList = errors.Join(*errorList, err)
	}

	if err := iMediator.Register(
		reflect.TypeOf(query.ListExpertConsultationStatusQuery{}),
		qhandler.NewListExpertConsultationStatusQueryHandler(
			container.Inject[validator.Validator](),
			container.Inject[logging.Logger](),
			container.Inject[usecase.ListExpertConsultationStatusUseCase](),
			container.Inject[message.Provider](),
		),
	); err != nil {
		*errorList = errors.Join(*errorList, err)
	}
}

func (m *manager) ConfigureMappers() {
	mapper.RegisterMappers()
}

func (m *manager) ConfigureApplication() error {
	var errorList error
	m.ConfigureMediator(&errorList)
	m.ConfigureMappers()
	return errorList
}

func (m *manager) ConfigureMiddleware(engine *gin.Engine) error {
	var errorList error
	// Aplicar CORS primero
	engine.Use(middleware.CorsMiddleware())
	// Luego aplicar el middleware de idioma
	engine.Use(middleware.LangMiddleware())

	return errorList
}

func (m *manager) ConfigureEngine(engine *gin.Engine) error {
	var errorList error
	m.ConfigureMiddleware(engine)

	return errorList
}
