package middleware

import (
	"os"
	"slices"
	"strings"

	"github.com/gin-gonic/gin"
)

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.GetHeader("Origin")

		// Si no hay variable de entorno, permitir todos los orígenes
		allowedOriginsEnv := os.Getenv("ALLOWED_ORIGINS")
		if allowedOriginsEnv == "" {
			// Permitir cualquier origen si no está configurado
			if origin != "" {
				c.Header("Access-Control-Allow-Origin", origin)
			} else {
				c.Header("Access-Control-Allow-Origin", "*")
			}
		} else {
			// Verificar contra la lista de orígenes permitidos
			allowedOrigins := strings.Split(allowedOriginsEnv, ",")
			// Limpiar espacios en blanco
			for i, o := range allowedOrigins {
				allowedOrigins[i] = strings.TrimSpace(o)
			}

			if slices.Contains(allowedOrigins, origin) {
				c.Header("Access-Control-Allow-Origin", origin)
				c.Header("Vary", "Origin")
			} else {
				// Si el origen no está permitido, no enviar el header
				c.Next()
				return
			}
		}

		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Accept, X-Requested-With")
		c.Header("Access-Control-Expose-Headers", "Content-Length")
		c.Header("Access-Control-Max-Age", "86400")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
