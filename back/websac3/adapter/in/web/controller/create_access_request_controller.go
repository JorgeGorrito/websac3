package controller

import (
	"errors"
	"net/http"
	"websac3/adapter/in/web/request"
	"websac3/app/domain/errs"
	"websac3/app/port/in/dto/command"
	"websac3/app/port/in/usecase"
	"websac3/common/mapper"

	"github.com/gin-gonic/gin"
)

type CreateAccessRequestController struct {
	createAccessRequestUseCase usecase.CreateAccessRequestUseCase
}

var createAccessRequestControllerInstance *CreateAccessRequestController = nil

func InitCreateAccessRequestController(
	createAccessRequestUseCase usecase.CreateAccessRequestUseCase,
) *CreateAccessRequestController {
	if createAccessRequestControllerInstance == nil {
		createAccessRequestControllerInstance = &CreateAccessRequestController{
			createAccessRequestUseCase: createAccessRequestUseCase,
		}
	}
	return createAccessRequestControllerInstance
}

func (c *CreateAccessRequestController) CreateAccessRequest(context *gin.Context) {
	var createAccessRequestRequest request.CreateAccessRequestRequest
	if err := context.ShouldBindJSON(&createAccessRequestRequest); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var createAccessRequestCommand command.CreateAccessRequestCommand
	if err := mapper.Map(&createAccessRequestRequest, &createAccessRequestCommand); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	switch err := c.createAccessRequestUseCase.CreateAccessRequest(createAccessRequestCommand); {
	case err == nil:
		context.JSON(http.StatusCreated, gin.H{"message": "Access request created successfully"})
	case errors.Is(err, errs.ValidationError):
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	default:
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}
