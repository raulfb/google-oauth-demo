package middlewares

import (
	"net/http"
	"strings"

	"oauth/backend/services"

	"github.com/gin-gonic/gin"
)

// Middleware para autenticar JWT
func Authenticate(c *gin.Context) {
	// Obtener el token del encabezado Authorization
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token no proporcionado"})
		c.Abort() // Detiene el flujo y no pasa al siguiente controlador
		return
	}

	// El token debería estar en el formato "Bearer <token>"
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Formato de token inválido"})
		c.Abort()
		return
	}

	// Validar el token
	tokenString := parts[1]
	claims, err := services.ValidarToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
		c.Abort()
		return
	}

	// Obtener el rol del claim
	rol := claims.Role
	email := claims.Email
	id := claims.ID
	// Almacenar el rol y en el contexto
	c.Set("rol", rol)
	c.Set("email", email)
	c.Set("id", id)
	// Si el token es válido, pasar a la siguiente ruta
	c.Next()
}
