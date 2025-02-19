# 🚀 Backend - OAuth Google Demo

Backend desarrollado en Go para implementar autenticación OAuth 2.0 con Google.

## 🛠️ Prerrequisitos

- Go 1.22 o superior
- MySQL/MariaDB

## 📦 Instalación

1. Clona el repositorio
2. Instala las dependencias:

```bash
go mod tidy
```

## 🗄️ Base de Datos

1. Crea una base de datos en MySQL/MariaDB
2. Crea la tabla `usuarios`:
```sql
CREATE TABLE usuarios (
    id INT AUTO_INCREMENT PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    rol VARCHAR(50) NOT NULL
);
```

## 🔐 Configuración de Variables de Entorno

1. Crea un archivo `.env` en la raíz del proyecto backend:
```env
# Puerto del servidor
PORT=8080

# Configuración de la base de datos
DB_HOST=localhost
DB_PORT=3306
DB_USER=tu_usuario
DB_PASSWORD=tu_password
DB_NAME=nombre_de_tu_bd

# Configuración JWT
JWT_ENCRYPTION=tu_clave_secreta_muy_larga
JWT_EXPIRATION=1440

# Configuración OAuth Google
GOOGLE_CLIENT_ID=tu_client_id
GOOGLE_CLIENT_SECRET=tu_client_secret
GOOGLE_REDIRECT_URL=http://localhost:5173/callback
```

### 🔑 Obtener Credenciales de Google

Sigue estos pasos para obtener las credenciales OAuth 2.0:

1. Ve a [Google Cloud Console](https://console.cloud.google.com/)

2. Crea un nuevo proyecto o selecciona uno existente
   - En la barra superior, junto al logo de Google Cloud
   - Click en "Selecciona un proyecto" → "Nuevo Proyecto"
   - Dale un nombre y click en "Crear"

3. Configura la pantalla de consentimiento:
   - En el menú lateral, ve a "APIs & Servicios" → "Pantalla de consentimiento"
   - Selecciona "Externo" como Tipo de usuario
   - Rellena la información requerida:
     - Nombre aplicación: Nombre de tu aplicación
     - Email de soporte: Tu email
     - Información de contacto: Tu email

4. Crea las credenciales:
   - En el menú lateral, ve a "APIs & Servicios" → "Credenciales"
   - Click en "Crear credenciales" → "ID de cliente OAuth"
   - Selecciona "Aplicación Web"
   - Configura:
     ```
     Name: [Nombre de tu aplicación]
     Orígenes autorizados de JavaScript:
       - http://localhost:5173
     URI de redireccionamiento autorizados:
       - http://localhost:5173/callback
     ```
   - Click en "Crear"

5. Guarda las credenciales:
   - Se te mostrarán el Client ID y Client Secret
   - Cópialos a tu archivo `.env`:
   ```env
   GOOGLE_CLIENT_ID=tu_client_id
   GOOGLE_CLIENT_SECRET=tu_client_secret
   GOOGLE_REDIRECT_URL=http://localhost:5173/callback
   ```

⚠️ **Importante**: 
- Sigue los pasos en el orden exacto indicado
- No saltes la configuración de la pantalla de consentimiento
- Para desarrollo, asegúrate de añadir tu email como usuario de prueba
- En producción, deberás cambiar los dominios `localhost:5173` por tu dominio real

## 🚀 Ejecución

```bash
go run main.go
```

## 🛣️ Rutas Disponibles

- `GET /auth/google/url` - Obtiene la URL para iniciar el flujo de OAuth
- `POST /auth/google/callback` - Callback de Google OAuth
- `GET /protegida` - Ruta protegida (requiere token JWT)

## 🔒 Autenticación

El sistema utiliza dos niveles de autenticación:
1. OAuth de Google para el login inicial
2. JWT para proteger las rutas de la API

Los tokens JWT incluyen:
- Email del usuario
- Rol del usuario
- ID del usuario en la base de datos