# 🔐 Demo de OAuth con Google

Este proyecto muestra una implementación completa de autenticación OAuth 2.0 con Google, utilizando Vue 3 para el frontend y Go para el backend.

## 📝 Descripción general

La aplicación permite a los usuarios:
- Iniciar sesión usando sus cuentas de Google
- Acceder a rutas protegidas mediante tokens JWT
- Mantener sesiones seguras y persistentes

## 🏗️ Arquitectura

El proyecto está dividido en dos partes:

### 🎨 Frontend (Vue 3)
- Maneja la interfaz de usuario y la experiencia de autenticación
- Gestiona el estado de la sesión y los tokens
- Implementa las redirecciones necesarias para OAuth

### 📦 Backend (Go) 
- Procesa las solicitudes de autenticación
- Verifica y gestiona tokens
- Maneja la persistencia de datos de usuarios
- Protege rutas mediante middleware

## 📚 Documentación detallada

Para instrucciones específicas de configuración y ejecución:

- [Documentación del Frontend](./frontend/README.md)
- [Documentación del Backend](./backend/README.md)


## 🔄 Flujo de autenticación OAuth

El proceso de autenticación OAuth con Google sigue los siguientes pasos:

1. **Inicio del login con Google**
   - El usuario hace clic en "Login con Google" en el frontend
   - El frontend hace una petición a `/auth/google/url` del backend
   - El backend devuelve la URL de autenticación de Google

2. **Autenticación en Google**
   - El usuario es redirigido a la página de login de Google
   - El usuario autoriza el acceso a su información
   - Google devuelve un código de autorización al frontend

3. **Procesamiento en el backend**
   - El frontend envía el código a `/auth/google/callback`
   - El backend verifica el código con Google y obtiene el email del usuario
   - Busca al usuario en la base de datos, si no existe lo crea
   - Genera un JWT (JSON Web Token) con los datos del usuario
   - Devuelve el token JWT y la información del usuario al frontend

4. **Acceso a rutas protegidas**
   - Para cada petición a rutas protegidas, el frontend incluye el token JWT en el header: `Authorization: Bearer <token>`
   - El middleware `Authenticate` del backend valida el token
   - Si el token es válido, permite el acceso y proporciona los datos del usuario al controlador
   - Si el token no es válido o ha expirado, devuelve un error 401

Esta demo tiene como objetivo proporcionar un ejemplo completo de cómo implementar OAuth con Google en una aplicación web, cubriendo tanto los aspectos frontend como backend.


