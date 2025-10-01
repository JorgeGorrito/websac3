package controller

import (
	"net/http"

	"websac3/app/port/out/message"
	"websac3/common/dependencies/container"

	"github.com/gin-gonic/gin"
)

type DownloadDegreeProgramTemplateController struct {
	Authenticable
}

func GetDownloadDegreeProgramTemplateController() *DownloadDegreeProgramTemplateController {
	return &DownloadDegreeProgramTemplateController{}
}

// DownloadDegreeProgramTemplate descarga una plantilla CSV para carga masiva de programas de grado
// @Summary Descargar Plantilla CSV
// @Description Descarga un archivo CSV con las columnas requeridas y datos de ejemplo para la carga masiva de programas de grado.
// @Description La plantilla incluye todas las columnas necesarias: SNIES, nombre, créditos totales, duración, nivel de formación,
// @Description roles profesionales, enfoque del programa y perfiles de entrada, egreso y profesional.
// @Tags DegreeProgram
// @Accept json
// @Produce text/csv
// @Param lang path string true "Código de idioma" default(es) Enums(en, es)
// @Success 200 {file} file "Archivo CSV con plantilla"
// @Failure 401 {string} string "Usuario no autenticado"
// @Failure 403 {string} string "No tiene permisos para realizar esta acción"
// @Router /api/v1/{lang}/degree-program/template [get]
// @Security BearerAuth
func (c *DownloadDegreeProgramTemplateController) DownloadTemplate(context *gin.Context) {
	// Validar token de autenticación
	token := c.GetToken(context)
	if token == nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
		return
	}

	// Validar permisos para descargar plantilla
	permissions := token.Permissions["degree-programs"]
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
	csvContent := `snies;name;total_credits;duration_value;duration_unit_id;formation_level_id;professional_role_ids;program_focus;entry_profile;graduate_profile;professional_profile
12345;Ingeniería de Sistemas;160;5;1;3;1,2,3;Desarrollo de software y sistemas informáticos;Bachiller con conocimientos básicos en matemáticas y física;Ingeniero capaz de desarrollar sistemas informáticos complejos;Desarrollador de software - analista de sistemas - arquitecto de software
67890;Ingeniería de Software;150;4;1;3;2,4,5;Desarrollo de aplicaciones y sistemas de software;Bachiller con aptitudes lógicas y matemáticas;Ingeniero especializado en desarrollo de software;Desarrollador full-stack - ingeniero de DevOps - líder técnico
11111;Ingeniería Informática;170;6;1;3;1,3,6;Tecnologías de la información y comunicaciones;Bachiller con interés en tecnología;Ingeniero experto en infraestructura tecnológica;Administrador de sistemas - consultor IT - especialista en redes`

	// Set headers for file download
	context.Header("Content-Type", "text/csv")
	context.Header("Content-Disposition", "attachment; filename=degree_program_template.csv")
	context.Header("Content-Length", string(rune(len(csvContent))))

	// Write CSV content
	context.String(http.StatusOK, csvContent)
}
