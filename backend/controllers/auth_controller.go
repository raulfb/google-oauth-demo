package controllers

import (
	"net/http"

	"oauth/backend/models"
	"oauth/backend/services"

	"github.com/gin-gonic/gin"
)

func GetGoogleAuthURL(c *gin.Context) {
	url := services.GetGoogleAuthURL()
	c.JSON(http.StatusOK, gin.H{
		"url": url,
	})
}

func GoogleCallback(c *gin.Context) {
	var request struct {
		Code string `json:"code"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Código no recibido correctamente"})
		return
	}

	email, err := services.HandleGoogleCallback(request.Code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Buscar usuario por email
	var usuario models.Usuario
	result := models.DB.Where("email = ?", email).First(&usuario)

	if result.Error != nil {
		// Si no existe, crear nuevo usuario
		usuario = models.Usuario{
			Email: email,
			Rol:   "Usuario",
		}
		if err := models.DB.Create(&usuario).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear usuario"})
			return
		}
	}

	// Generar token con los datos del usuario
	token, err := services.GenerarToken(usuario.Email, usuario.Rol, usuario.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al generar el token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"user":  usuario,
	})
}
