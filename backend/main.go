package main

import (
	"log"
	"os"

	"oauth/backend/controllers"
	"oauth/backend/middlewares"
	"oauth/backend/models"
	"oauth/backend/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

func main() {
	// Carga las variables del archivo .env

	err := godotenv.Load()
	if err != nil {
		log.Println("Error al cargar el archivo .env. Se usarán las variables de entorno del sistema.")
	}

	// Lee el puerto desde las variables de entorno
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Error: La variable de entorno 'PORT' no está definida.")
	}

	// GOOGLE_CLIENT_ID
	googleClientID := os.Getenv("GOOGLE_CLIENT_ID")
	if googleClientID == "" {
		log.Fatal("Error: La variable de entorno 'GOOGLE_CLIENT_ID' no está definida.")
	}

	// GOOGLE_CLIENT_SECRET
	googleClientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	if googleClientSecret == "" {
		log.Fatal("Error: La variable de entorno 'GOOGLE_CLIENT_SECRET' no está definida.")
	}

	// GOOGLE_REDIRECT_URL
	googleRedirectURL := os.Getenv("GOOGLE_REDIRECT_URL")
	if googleRedirectURL == "" {
		log.Fatal("Error: La variable de entorno 'GOOGLE_REDIRECT_URL' no está definida.")
	}

	models.DB, err = gorm.Open("mysql", models.DbURL(models.BuildDBConfig()))
	if err != nil {
		panic(err)
	}
	defer models.DB.Close()

	// Configuración del servidor Gin
	router := gin.Default()

	// cors
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true           // acceder desde calquer dominio
	config.AllowCredentials = true          // aceptar tokens
	config.AddAllowHeaders("authorization") // para poder usar o token
	config.AddAllowMethods("OPTIONS")       // para o vuejs
	router.Use(cors.New(config))            //si se usa despois de definir as rutas non funciona

	// Inicializar Google OAuth
	services.InitGoogleOAuth()

	// Rutas
	router.GET("/auth/google/url", controllers.GetGoogleAuthURL)
	router.POST("/auth/google/callback", controllers.GoogleCallback)

	router.GET("/protegida", middlewares.Authenticate, controllers.GetProtectedInfo)

	router.Run(":" + port)
}
