package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func authMiddleware(c *gin.Context) {
	authorizationHeader := c.GetHeader("Authorization")
	if authorizationHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token de autenticação não fornecido"})
		c.Abort()
		return
	}

	tokenParts := strings.Split(authorizationHeader, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token de autenticação inválido"})
		c.Abort()
		return
	}

	tokenString := tokenParts[1]

	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return getSecretKey(), nil
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token de autenticação inválido"})
		c.Abort()
		return
	}

	c.Next()
}
