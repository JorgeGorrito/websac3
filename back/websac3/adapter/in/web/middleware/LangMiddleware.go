package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func isConfiguredLang(lang string, allowed []string) bool {
	for _, l := range allowed {
		if strings.TrimSpace(l) == lang {
			return true
		}
	}
	return false
}

func isSwaggerRoute(path string) bool {
	return strings.Contains(path, "/swagger")
}

func LangMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if isSwaggerRoute(c.Request.URL.Path) {
			c.Next()
			return
		}

		configuredLangs := strings.Split(os.Getenv("MESSAGES_LANGUAGES"), ",")
		lang := ""

		if pathLang := c.Param("lang"); pathLang != "" {
			lang = pathLang
		} else if headerLang := c.GetHeader("Accept-Language"); headerLang != "" {
			lang = headerLang
		}

		if lang == "" || !isConfiguredLang(lang, configuredLangs) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{})
			return
		}

		c.Set("lang", lang)
		c.Next()
	}
}
