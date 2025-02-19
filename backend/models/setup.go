package models

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // MySQL driver
	"github.com/joho/godotenv"
)

// DB define la conexión a la base de datos
var DB *gorm.DB

// DBConfig representa la configuración de la base de datos
type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

// BuildDBConfig genera la configuración de la base de datos
func BuildDBConfig() *DBConfig {
	err := godotenv.Load() // Cargar el archivo .env
	if err != nil {
		log.Println("Can't load .env file. Let's try environment variables...")
	}

	// Convertir el puerto de string a int
	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatal("Undefined environment variables!")
	}

	host := os.Getenv("DB_HOST")
	if host == "" {
		log.Fatal("Error: La variable de entorno 'DB_HOST' no está definida.")
	}
	port := os.Getenv("DB_PORT")
	if port == "" {
		log.Fatal("Error: La variable de entorno 'DB_PORT' no está definida.")
	}
	user := os.Getenv("DB_USER")
	if user == "" {
		log.Fatal("Error: La variable de entorno 'DB_USER' no está definida.")
	}
	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		log.Fatal("Error: La variable de entorno 'DB_PASSWORD' no está definida.")
	}
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		log.Fatal("Error: La variable de entorno 'DB_NAME' no está definida.")
	}

	dbConfig := DBConfig{
		Host:     host,
		Port:     dbPort,
		User:     user,
		Password: password,
		DBName:   dbName,
	}

	return &dbConfig
}

// DbURL genera la URL de conexión para la base de datos
func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}
