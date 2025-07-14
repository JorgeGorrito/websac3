package controller

import (
	"net/http"
	"websac3/adapter/in/web/request"
	"websac3/adapter/in/web/response"
	"websac3/app/port/in/dto/command"
	"websac3/common/mapper"
	"websac3/common/mediator"

	"github.com/gin-gonic/gin"
)

type CreateAccessRequestController struct{}

var createAccessRequestControllerInstance *CreateAccessRequestController = nil

func InitCreateAccessRequestController() *CreateAccessRequestController {
	if createAccessRequestControllerInstance == nil {
		createAccessRequestControllerInstance = &CreateAccessRequestController{}
	}
	return createAccessRequestControllerInstance
}

func (c *CreateAccessRequestController) CreateAccessRequest(context *gin.Context) {
	var err error
	var createAccessRequestRequest request.CreateAccessRequestRequest
	if err = context.ShouldBindJSON(&createAccessRequestRequest); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"result": err.Error()})
		return
	}

	var createAccessRequestCommand command.CreateAccessRequestCommand
	createAccessRequestCommand, err = mapper.Map[request.CreateAccessRequestRequest, command.CreateAccessRequestCommand](&createAccessRequestRequest)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"result": err.Error()})
		return
	}

	result, _ := mediator.Send[command.CreateAccessRequestCommand, response.ApiResponse[string]](createAccessRequestCommand, context.GetString("lang"))
	context.JSON(result.HttpStatusCode, result.ToResponseFormat())
}
