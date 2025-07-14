package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
)

func isConfiguredLang(lang string, allowed []string) bool {
	for _, l := range allowed {
		if strings.TrimSpace(l) == lang {
			return true
		}
	}
	return false
}

func LangMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
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
