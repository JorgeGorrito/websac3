package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	_jwt "websac3/common/jwt"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func extractAccessTokenClaims(tokenStr string, secret string) (*_jwt.AccessTokenClaims, error) {
	// Parsear el token con claims personalizados
	parsedToken, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// Validar que el algoritmo sea HS256
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if !parsedToken.Valid {
		return nil, errors.New("invalid or expired token")
	}

	// Obtener las claims como mapa
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid claims format")
	}

	// Mapear claims manualmente
	subFloat, ok := claims["sub"].(float64)
	if !ok {
		return nil, errors.New("invalid sub claim")
	}

	username, _ := claims["username"].(string)
	role, _ := claims["role"].(string)

	// Mapear permissions (map[string][]string)
	permissions := make(map[string][]string)
	if rawPerms, ok := claims["permissions"].(map[string]interface{}); ok {
		for key, val := range rawPerms {
			if arr, ok := val.([]interface{}); ok {
				for _, item := range arr {
					if strItem, ok := item.(string); ok {
						permissions[key] = append(permissions[key], strItem)
					}
				}
			}
		}
	}

	return &_jwt.AccessTokenClaims{
		Sub:         uint(subFloat),
		Username:    username,
		Role:        role,
		Permissions: permissions,
	}, nil
}

func AuthMiddleware(jwtSecret []byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token faltante o inválido"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := extractAccessTokenClaims(tokenString, string(jwtSecret))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			c.Abort()
			return
		}

		c.Set("token", claims)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			c.Abort()
			return
		}

		c.Next()
	}
}
