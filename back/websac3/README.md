# Websac3 Backend

API backend para el sistema Websac3, construida con Go utilizando arquitectura hexagonal y patrones CQRS.

## 📋 Tabla de Contenidos

- [Descripción](#descripción)
- [Inicio Rápido](#inicio-rápido)
- [Arquitectura](#arquitectura)
- [Tecnologías](#tecnologías)
- [Requisitos Previos](#requisitos-previos)
- [Instalación](#instalación)
- [Configuración](#configuración)
- [Ejecución](#ejecución)
- [Comandos de Base de Datos](#-comandos-de-base-de-datos)
- [Documentación API](#documentación-api)
- [Estructura del Proyecto](#estructura-del-proyecto)
- [Características](#características)

## 📖 Descripción

Websac3 es un sistema backend robusto para la gestión de instituciones de educación superior, cursos, consultas de expertos y solicitudes de acceso. Implementa una arquitectura limpia con separación clara de responsabilidades y utiliza patrones modernos de desarrollo.

## ⚡ Inicio Rápido

Para poner en marcha la aplicación rápidamente:

```bash
# 1. Instalar dependencias
go mod download

# 2. Configurar variables de entorno
cp .env.example .env  # Edita el archivo .env con tus configuraciones

# 3. Generar certificados SSL
chmod +x ensure-certs.sh && ./ensure-certs.sh

# 4. Generar documentación Swagger
swag init --dir . --output ./docs --parseDependency --parseInternal

# 5. Ejecutar la aplicación
go run main.go
```

La aplicación estará disponible en: `https://localhost:8110`

**Nota:** 
- La base de datos se inicializa automáticamente la primera vez que ejecutas la aplicación
- Asegúrate de tener configuradas las credenciales de Gmail API (`.mail-credentials.json` y `.mail-token.json`) antes de ejecutar. Ver [Credenciales de Gmail API](#2-credenciales-de-gmail-api-requerido).

## 🏗️ Arquitectura

El proyecto sigue los principios de **Arquitectura Hexagonal** (Ports & Adapters) con:

- **Adapter Layer**: Manejadores de entrada (controladores web) y salida (persistencia, notificaciones, PDF)
- **Application Layer**: Lógica de dominio, entidades, servicios y casos de uso
- **Common Layer**: Utilidades compartidas, dependencias, validadores y mediadores

### Patrón CQRS

El proyecto implementa CQRS (Command Query Responsibility Segregation) utilizando un mediador para:
- Commands: Operaciones que modifican el estado
- Queries: Operaciones de lectura

## 🛠️ Tecnologías

### Framework y Lenguaje
- **Go 1.23.5**
- **Gin**: Framework web
- **Anise**: Framework personalizado para aplicaciones web con Gin

### Base de Datos
- **PostgreSQL**: Base de datos principal
- **GORM**: ORM para Go

### Autenticación y Seguridad
- **JWT**: Autenticación basada en tokens
- **OAuth2**: Integración con Google API
- **TLS/HTTPS**: Comunicación segura

### Otras Herramientas
- **Swagger**: Documentación API automática
- **wkhtmltopdf**: Generación de PDFs
- **Gmail API**: Envío de correos electrónicos
- **Docker**: Containerización

## 📦 Requisitos Previos

- Go 1.23.5 o superior
- PostgreSQL
- Docker y Docker Compose (opcional)
- wkhtmltopdf (para generación de PDFs)
- Certificados SSL/TLS (se generan automáticamente con `ensure-certs.sh`)
- **Credenciales de Gmail API** (OAuth2) - **REQUERIDO** para envío de correos

## 🚀 Instalación

### 1. Clonar el repositorio

```bash
git clone <repository-url>
cd websac3/back/websac3
```

### 2. Instalar dependencias

```bash
go mod download
```

### 3. Configurar credenciales de Gmail API

⚠️ **ANTES DE CONTINUAR**: Asegúrate de tener los archivos `.mail-credentials.json` y `.mail-token.json` en la raíz del proyecto. Ver sección [Credenciales de Gmail API](#2-credenciales-de-gmail-api-requerido) para más detalles.

### 4. Generar documentación Swagger

```bash
go install github.com/swaggo/swag/cmd/swag@v1.16.4
swag init --dir . --output ./docs --parseDependency --parseInternal
```

## ⚙️ Configuración

### 1. Variables de entorno

Crear un archivo `.env` en la raíz del proyecto con las siguientes variables:

```env
# Database Configuration
DB_HOST=websac3-db
DB_USER=websac3
DB_PASSWORD=12345678
DB_NAME=websac3
DB_PORT=5432
DB_SSLMODE=disable
DB_TIMEZONE=America/Bogota
DB_MAX_IDLE_CONNS=10
DB_MAX_OPEN_CONNS=100
DB_CONN_MAX_LIFETIME=3600

# Application Configuration
APP_NAME=WEBSAC3
LOGGER_PATH=logs/websac3.log
LOGGER_BUFFER_SIZE=1024
LOGGER_MAX_ENTRIES=1000

# Gmail API Configuration (REQUERIDO)
MAIL_FROM=tu_email@gmail.com
MAIL_CREDENTIALS_PATH=.mail-credentials.json
MAIL_TOKEN_PATH=.mail-token.json

# Internationalization
MESSAGES_PATH=adapter/out/message/lang/
MESSAGES_LANGUAGES="en,es"

# Email Templates
TEMPLATES_BASE_DIR=adapter/out/notification/template/templates/

# Cache Configuration
ENUM_TTL_MINUTES=15

# JWT Configuration
JWT_SECRET_KEY=websac3_jwt_secret_key
JWT_REFRESH_SECRET_KEY=websac3_jwt_refresh_secret_key
JWT_EXPIRATION_TIME=1800
JWT_REFRFESH_EXPIRATION_TIME=86400

# CORS Configuration
ALLOWED_ORIGINS=https://localhost:3000
```

### 2. Credenciales de Gmail API (REQUERIDO)

⚠️ **IMPORTANTE**: Estos archivos son **obligatorios** para que la aplicación funcione correctamente, ya que el sistema envía notificaciones por email.

Debes colocar estos archivos en la raíz del proyecto:

#### `.mail-credentials.json`
Credenciales OAuth2 de Google Cloud Console:
1. Ve a [Google Cloud Console](https://console.cloud.google.com/)
2. Crea un proyecto o selecciona uno existente
3. Habilita la **Gmail API**
4. Ve a **APIs & Services** → **Credentials**
5. Crea credenciales tipo **OAuth 2.0 Client ID**
6. Descarga el archivo JSON y guárdalo como `.mail-credentials.json`

#### `.mail-token.json`
Token de acceso OAuth2 generado:
1. La primera vez que ejecutes la aplicación, se generará automáticamente
2. O usa la herramienta de Google para generar el token con los scopes necesarios
3. El token debe tener el scope: `https://www.googleapis.com/auth/gmail.send`

**Estructura del `.mail-credentials.json`:**
```json
{
  "installed": {
    "client_id": "tu-client-id.apps.googleusercontent.com",
    "project_id": "tu-proyecto",
    "auth_uri": "https://accounts.google.com/o/oauth2/auth",
    "token_uri": "https://oauth2.googleapis.com/token",
    "client_secret": "tu-client-secret",
    "redirect_uris": ["http://localhost"]
  }
}
```

**Estructura del `.mail-token.json`:**
```json
{
  "access_token": "...",
  "token_type": "Bearer",
  "refresh_token": "...",
  "expiry": "2025-..."
}
```

### 3. Certificados SSL

Ejecuta el script para generar o asegurar los certificados:

```bash
chmod +x ensure-certs.sh
./ensure-certs.sh
```

## 🏃 Ejecución

### Modo Desarrollo

```bash
go run main.go
```

### Con Docker

```bash
docker-compose up --build
```

### Compilar

```bash
go build -o websac3.exe .
./websac3.exe
```

La aplicación estará disponible en: `https://localhost:8110`

## 🗄️ Comandos de Base de Datos

El proyecto incluye comandos CLI para gestionar la base de datos:

### Inicialización Automática

🚀 **La aplicación se inicializa automáticamente al arrancar por primera vez.**

Cada vez que ejecutas `go run main.go`, la aplicación verifica si la base de datos ha sido inicializada. Si es la primera vez, ejecuta automáticamente:

**¿Qué hace?**
- Crea la tabla de control de migraciones
- Verifica si la base de datos ya fue inicializada
- Si es la primera vez:
  - Migra todas las tablas (usuarios, roles, permisos, cursos, instituciones, etc.)
  - Ejecuta los seeders esenciales (roles, permisos, usuarios por defecto, etc.)
  - Ejecuta el seeder de roles profesionales (16 roles de ciberseguridad con sus áreas de conocimiento)
  - Registra la migración como completada con timestamp
- Si ya fue inicializada, muestra un mensaje informativo y continúa normalmente

**Características:**
- ✅ **Automático**: No necesitas ejecutar ningún comando adicional
- ✅ **Idempotente**: Se ejecuta cada vez que arranca la app, pero solo inicializa una vez
- ✅ **Seguro**: Verifica antes de hacer cambios
- ✅ **Transaccional**: Todos los cambios se hacen en una transacción (si algo falla, se revierte todo)

### Comando init manual (opcional)

También puedes ejecutar la inicialización manualmente si lo prefieres:

```bash
go run main.go init
```

### Otros comandos disponibles

```bash
# Migrar todos los modelos
go run main.go migrate:all

# Migrar un modelo específico
go run main.go migrate:model --model=users

# Resetear la base de datos (⚠️ Elimina todos los datos)
go run main.go migrate:reset

# Ejecutar un seeder específico
go run main.go seed:run --seed=essential_data
```

## 📚 Documentación API

La documentación Swagger está disponible en:

```
https://localhost:8110/swagger/index.html
```

### Información de la API
- **Versión**: 1.0
- **Host**: localhost:8110
- **Esquema**: HTTPS
- **Autenticación**: Bearer Token (JWT)

## 📁 Estructura del Proyecto

```
websac3/
├── adapter/                  # Capa de adaptadores
│   ├── in/                  # Adaptadores de entrada
│   │   └── web/             # Controladores, handlers, routing
│   └── out/                 # Adaptadores de salida
│       ├── message/         # Mensajes i18n
│       ├── notification/    # Envío de emails
│       ├── pdf/             # Generación de PDFs
│       └── persistence/     # Repositorios PostgreSQL
├── app/                     # Capa de aplicación
│   ├── domain/              # Dominio de la aplicación
│   │   ├── constants/       # Constantes del dominio
│   │   ├── entity/          # Entidades del dominio
│   │   ├── errs/            # Errores personalizados
│   │   ├── notification/    # Templates de notificaciones
│   │   ├── service/         # Servicios de dominio
│   │   └── usecase/         # Casos de uso
│   └── port/                # Puertos (interfaces)
│       ├── in/              # Puertos de entrada (Commands/Queries)
│       └── out/             # Puertos de salida (Repositories)
├── common/                  # Utilidades comunes
│   ├── decoder/             # Decodificadores
│   ├── dependencies/        # Inyección de dependencias
│   ├── filter/              # Filtros
│   ├── jwt/                 # Utilidades JWT
│   ├── logging/             # Logger
│   ├── mail/                # Utilidades de email
│   ├── mapper/              # Mappers
│   ├── mediator/            # Mediador CQRS
│   ├── paginator/           # Paginación
│   ├── uuid/                # Generador de UUIDs
│   └── validator/           # Validadores
├── docs/                    # Documentación Swagger (generada)
├── logs/                    # Archivos de log
├── .env                     # Variables de entorno
├── Dockerfile               # Configuración Docker
├── ensure-certs.sh          # Script para certificados
├── go.mod                   # Dependencias Go
├── go.sum                   # Checksums de dependencias
└── main.go                  # Punto de entrada
```

## ✨ Características

### Gestión de Usuarios
- Registro y activación de usuarios
- Autenticación con JWT
- Cambio de contraseña
- Desactivación de usuarios
- Gestión de roles y permisos

### Gestión Académica
- Administración de cursos
- Programas de grado
- Instituciones de educación superior
- Niveles de formación
- Tipos de curso y naturaleza

### Solicitudes de Acceso
- Creación de solicitudes
- Aprobación/rechazo de solicitudes
- Estados de solicitud
- Notificaciones por email

### Consultas de Expertos
- Gestión de consultas
- Estados de consultas
- Asignación de expertos

### Reportes y Estadísticas
- Generación de reportes en PDF
- Estadísticas del sistema
- Exportación de datos

### Notificaciones
- Sistema de templates para emails
- Envío de correos con Gmail API
- Templates HTML personalizables
- Soporte para attachments

### Internacionalización
- Soporte multiidioma
- Mensajes en español e inglés

## 🔐 Seguridad

- Autenticación mediante JWT
- Comunicación HTTPS/TLS
- Validación de datos de entrada
- Sanitización de queries SQL con GORM
- Gestión segura de contraseñas
- OAuth2 para Gmail API

### ⚠️ Archivos Sensibles

Los siguientes archivos contienen información sensible y **NO deben ser versionados** en Git:
- `.env` - Variables de entorno
- `.mail-credentials.json` - Credenciales OAuth2 de Google
- `.mail-token.json` - Token de acceso OAuth2
- `certs/` - Certificados SSL/TLS

Asegúrate de que estos archivos estén listados en `.gitignore`.

## 📝 Notas Adicionales

### Zona Horaria
El proyecto está configurado para la zona horaria `America/Bogota`.

### Seeders
Los datos iniciales (seeds) se encuentran en:
```
adapter/out/persistence/postgresql/seeders/seeds/
```

### Templates de Email
Los templates HTML para emails están en:
```
adapter/out/notification/template/templates/
```

## 👥 Contacto

- **Email de Soporte**: j0rg3.4b3ll4@gmail.com

---

**Desarrollado con ❤️ usando Go y Arquitectura Hexagonal**

