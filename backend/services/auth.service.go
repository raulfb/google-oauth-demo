package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// SecretKey se obtiene desde la variable de entorno
var SecretKey = []byte(os.Getenv("JWT_ENCRYPTION"))

// Claims define los datos que estarán dentro del token
type Claims struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	ID    int    `json:"id"`
	jwt.RegisteredClaims
}

// Configuración de Google OAuth
var googleOAuthConfig *oauth2.Config

func InitGoogleOAuth() {
	googleOAuthConfig = &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URL"),
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
		},
		Endpoint: google.Endpoint,
	}
}

func GetGoogleAuthURL() string {
	return googleOAuthConfig.AuthCodeURL("state-token")
}

type GoogleUserInfo struct {
	Email string `json:"email"`
}

func HandleGoogleCallback(code string) (string, error) {
	token, err := googleOAuthConfig.Exchange(context.Background(), code)
	if err != nil {
		return "", fmt.Errorf("error al intercambiar código: %v", err)
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return "", fmt.Errorf("error al obtener información del usuario: %v", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error al leer respuesta: %v", err)
	}

	var userInfo GoogleUserInfo
	if err := json.Unmarshal(data, &userInfo); err != nil {
		return "", fmt.Errorf("error al parsear respuesta: %v", err)
	}

	return userInfo.Email, nil
}

// GenerarToken genera un JWT con los datos del usuario
func GenerarToken(email string, role string, id int) (string, error) {
	// Obtener la duración de expiración desde la variable de entorno
	expirationMinutes, err := strconv.Atoi(os.Getenv("JWT_EXPIRATION"))
	if err != nil {
		log.Fatal("JWT_EXPIRATION no está definida o no es válida")
	}

	// Crear los claims
	claims := &Claims{
		Email: email,
		Role:  role,
		ID:    id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expirationMinutes) * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// Crear el token firmado
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(SecretKey)
	if err != nil {
		return "", err
	}

	return "Bearer " + signedToken, nil
}

// ValidarToken verifica que un token sea válido y devuelve sus claims
func ValidarToken(tokenString string) (*Claims, error) {
	// Parsear el token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Validar el método de firma
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return SecretKey, nil
	})

	// Verificar si hubo un error o si el token no es válido
	if err != nil {
		return nil, err
	}

	// Extraer los claims si el token es válido
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrTokenInvalidClaims
}
