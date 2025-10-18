package controller

import (
	"net/http"

	"websac3/app/port/out/message"
	"websac3/common/dependencies/container"

	"github.com/gin-gonic/gin"
)

type DownloadCourseTopicTemplateController struct {
	Authenticable
}

func GetDownloadCourseTopicTemplateController() *DownloadCourseTopicTemplateController {
	return &DownloadCourseTopicTemplateController{}
}

// DownloadCourseTopicTemplate descarga una plantilla CSV para carga masiva de tópicos de curso
// @Summary Descargar Plantilla CSV de Tópicos de Curso
// @Description Descarga un archivo CSV con las columnas requeridas y datos de ejemplo para la carga masiva de tópicos de curso.
// @Description La plantilla incluye: ID del tópico y horas de estudio.
// @Tags Course
// @Accept json
// @Produce text/csv
// @Param lang path string true "Código de idioma" default(es) Enums(en, es)
// @Success 200 {file} file "Archivo CSV con plantilla"
// @Failure 401 {string} string "Usuario no autenticado"
// @Failure 403 {string} string "No tiene permisos para realizar esta acción"
// @Router /api/v1/{lang}/course/topic/template [get]
// @Security BearerAuth
func (c *DownloadCourseTopicTemplateController) DownloadTemplate(context *gin.Context) {
	// Validar token de autenticación
	token := c.GetToken(context)
	if token == nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
		return
	}

	// Validar permisos para descargar plantilla
	permissions := token.Permissions["courses"]
	hasPermission := false
	for _, permission := range permissions {
		if permission == "create" || permission == "update" || permission == "all" {
			hasPermission = true
			break
		}
	}

	if !hasPermission {
		msgProvider := container.Inject[message.Provider]()
		lang := context.Param("lang")
		context.JSON(http.StatusForbidden, gin.H{
			"error": msgProvider.WithLang(lang).GetMessage("base_error", "forbidden"),
		})
		return
	}

	// CSV template content with semicolon separator for Excel compatibility
	csvContent := `topic_id;study_hours
1;20
2;15
3;25
4;10
5;30`

	// Set headers for file download
	context.Header("Content-Type", "text/csv; charset=utf-8")
	context.Header("Content-Disposition", "attachment; filename=course_topic_template.csv")
	context.Header("Content-Length", string(rune(len(csvContent))))

	// Write CSV content
	context.String(http.StatusOK, csvContent)
}





