package controller

import (
	"fmt"
	"net/http"
	"websac3/adapter/in/web/request"
	"websac3/adapter/in/web/response"
	"websac3/app/port/in/dto/command"
	"websac3/common/mapper"
	"websac3/common/mediator"

	"github.com/gin-gonic/gin"
)

type CreateCourseController struct{ Authenticable }

var instanceCreateCourseController *CreateCourseController = nil

func GetCreateCourseController() *CreateCourseController {
	if instanceCreateCourseController == nil {
		instanceCreateCourseController = &CreateCourseController{}
	}
	return instanceCreateCourseController
}

// CreateCourse manejador para crear un curso
// @Summary Crear Curso
// @Description Crea un nuevo curso con toda la información académica necesaria.
// @Tags Course
// @Accept json
// @Produce json
// @Param request body request.CreateCourseRequest true "CreateCourseRequest"
// @Param lang path string true "Código de idioma" default(en) English Enums(en, es)
// @Success 201 {object} response.ApiResponse[string] "Curso creado exitosamente"
// @Failure 400 {object} response.ApiResponse[string] "Formato de solicitud inválido (ejemplo: JSON mal formado)"
// @Failure 409 {object} response.ApiResponse[string] "Curso ya existe"
// @Failure 500 {object} response.ApiResponse[string] "Error interno del servidor"
// @Router /api/v1/{lang}/course [post]
// @Security BearerAuth
func (c *CreateCourseController) CreateCourse(context *gin.Context) {
	var err error
	var createCourseRequest request.CreateCourseRequest
	if err = context.ShouldBindJSON(&createCourseRequest); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"result": err.Error()})
		return
	}
	fmt.Printf("createCourseRequest: %+v\n", createCourseRequest)
	var createCourseCommand command.CreateCourseCommand
	createCourseCommand, err = mapper.Map[request.CreateCourseRequest, command.CreateCourseCommand](&createCourseRequest)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"result": err.Error()})
		return
	}
	fmt.Printf("createCourseCommand: %+v\n", createCourseCommand)

	token := c.GetToken(context)
	lang := context.Param("lang")
	createCourseCommand.CreatedBy = token.Sub
	createCourseCommand.Permissions = token.Permissions["courses"]

	result, _ := mediator.Send[command.CreateCourseCommand, response.ApiResponse[string]](createCourseCommand, lang)
	context.JSON(result.HttpStatusCode, result.ToResponseFormat())
}
