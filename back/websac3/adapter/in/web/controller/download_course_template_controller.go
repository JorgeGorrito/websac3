package controller

import (
	"net/http"

	"websac3/app/port/out/message"
	"websac3/common/dependencies/container"

	"github.com/gin-gonic/gin"
)

type DownloadCourseTemplateController struct {
	Authenticable
}

func GetDownloadCourseTemplateController() *DownloadCourseTemplateController {
	return &DownloadCourseTemplateController{}
}

// DownloadCourseTemplate descarga una plantilla CSV para carga masiva de cursos
// @Summary Descargar Plantilla CSV de Cursos
// @Description Descarga un archivo CSV con las columnas requeridas y datos de ejemplo para la carga masiva de cursos.
// @Description La plantilla incluye todas las columnas necesarias: nombre, código, créditos, período, naturaleza, tipo y si es de ciberseguridad.
// @Tags Course
// @Accept json
// @Produce text/csv
// @Param lang path string true "Código de idioma" default(es) Enums(en, es)
// @Success 200 {file} file "Archivo CSV con plantilla"
// @Failure 401 {string} string "Usuario no autenticado"
// @Failure 403 {string} string "No tiene permisos para realizar esta acción"
// @Router /api/v1/{lang}/course/template [get]
// @Security BearerAuth
func (c *DownloadCourseTemplateController) DownloadTemplate(context *gin.Context) {
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
		if permission == "create" || permission == "all" {
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
	csvContent := `name;code;credits;period_number;nature_id;type_id;is_cybersecurity
Programación Avanzada;PROG301;3;3;1;1;no
Bases de Datos;BD201;4;2;1;1;no
Seguridad Informática;SEC401;3;4;1;2;si
Redes de Computadores;RED301;3;3;1;1;no
Desarrollo Web;WEB301;3;3;1;1;no`

	// Set headers for file download
	context.Header("Content-Type", "text/csv; charset=utf-8")
	context.Header("Content-Disposition", "attachment; filename=course_template.csv")
	context.Header("Content-Length", string(rune(len(csvContent))))

	// Write CSV content
	context.String(http.StatusOK, csvContent)
}


