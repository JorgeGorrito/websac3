# WebSAC3 - Sistema Web de Asesoramiento en Ciberseguridad

<div align="center">

Sistema integral para la gestión de asesoramiento en ciberseguridad para programas académicos en instituciones de educación superior.

[![Go Version](https://img.shields.io/badge/Go-1.23.5-00ADD8?style=flat&logo=go)](https://golang.org/)
[![Next.js](https://img.shields.io/badge/Next.js-15.0-000000?style=flat&logo=next.js)](https://nextjs.org/)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-14-316192?style=flat&logo=postgresql)](https://www.postgresql.org/)
[![Docker](https://img.shields.io/badge/Docker-Ready-2496ED?style=flat&logo=docker)](https://www.docker.com/)

</div>

---

## 📋 Tabla de Contenidos

- [Descripción](#-descripción)
- [Arquitectura del Sistema](#️-arquitectura-del-sistema)
- [Inicio Rápido con Docker](#-inicio-rápido-con-docker)
- [Tecnologías](#️-tecnologías)
- [Estructura del Proyecto](#-estructura-del-proyecto)
- [Configuración](#️-configuración)
- [Ejecución en Desarrollo](#-ejecución-en-desarrollo)
- [Características Principales](#-características-principales)
- [Documentación Adicional](#-documentación-adicional)
- [Seguridad](#-seguridad)
- [Contacto](#-contacto)

---

## 📖 Descripción

**WebSAC3** es una plataforma completa diseñada para facilitar la gestión de asesoramiento en ciberseguridad para programas académicos. El sistema permite:

- Gestión de instituciones de educación superior y sus programas académicos
- Administración de cursos de ciberseguridad con taxonomía especializada
- Sistema de consultas de expertos con asignación y seguimiento
- Gestión de solicitudes de acceso con flujo de aprobación
- Generación de reportes y estadísticas
- Sistema de notificaciones por correo electrónico
- Autenticación segura con JWT
- Interfaz moderna y responsiva

El sistema está construido con arquitectura moderna, utilizando **arquitectura hexagonal** en el backend, **CQRS** para la separación de responsabilidades y **Next.js** para una experiencia de usuario fluida.

---

## 🏗️ Arquitectura del Sistema

El proyecto está compuesto por tres componentes principales:

```
┌─────────────────────────────────────────────────────────────┐
│                    WebSAC3 Frontend                         │
│                  (Next.js + React)                          │
│              https://localhost:3000                         │
└──────────────────────┬──────────────────────────────────────┘
                       │ HTTPS/JWT
                       ▼
┌─────────────────────────────────────────────────────────────┐
│                    WebSAC3 Backend                          │
│        (Go + Gin + Arquitectura Hexagonal + CQRS)          │
│              https://localhost:8110                         │
└──────────────────────┬──────────────────────────────────────┘
                       │ PostgreSQL
                       ▼
┌─────────────────────────────────────────────────────────────┐
│                    PostgreSQL 14                            │
│                  Base de Datos                              │
│              localhost:5432                                 │
└─────────────────────────────────────────────────────────────┘
```

### Backend (Go)
- **Arquitectura Hexagonal** (Ports & Adapters)
- **CQRS** (Command Query Responsibility Segregation)
- **Clean Architecture** con separación de capas
- API RESTful con documentación Swagger
- Autenticación JWT + OAuth2
- Comunicación HTTPS/TLS

### Frontend (Next.js)
- **Next.js 15** con App Router
- **React 18** con TypeScript
- **Redux Toolkit** para gestión de estado
- **RTK Query** para llamadas API
- **Tailwind CSS** + shadcn/ui para UI moderna
- Soporte completo para HTTPS

### Base de Datos
- **PostgreSQL 14**
- Migraciones automáticas con GORM
- Seeders para datos iniciales
- Zona horaria: America/Bogota

---

## 🚀 Inicio Rápido con Docker

La forma más rápida de ejecutar WebSAC3 es usando Docker Compose:

### Requisitos Previos
- Docker y Docker Compose instalados
- Puertos disponibles: 3000 (frontend), 8110 (backend), 5432 (database)

### Pasos

**1. Clonar el repositorio**
```bash
git clone <repository-url>
cd websac3
```

**2. Configurar credenciales de Gmail API** ⚠️ **IMPORTANTE**

Antes de ejecutar, debes colocar estos archivos en `back/websac3/`:
- `.mail-credentials.json` - Credenciales OAuth2 de Google Cloud Console
- `.mail-token.json` - Token de acceso OAuth2

> 📖 Ver [Configuración de Gmail API](#configuración-de-gmail-api) para más detalles.

**3. Ejecutar con Docker Compose**
```bash
docker-compose up -d
```

Esto iniciará:
- ✅ Base de datos PostgreSQL (puerto 5432)
- ✅ Backend API (https://localhost:8110)
- ✅ Frontend (https://localhost:3000)

**4. Acceder a la aplicación**
- **Frontend**: https://localhost:3000
- **Backend API**: https://localhost:8110
- **Swagger Docs**: https://localhost:8110/swagger/index.html

**5. Credenciales iniciales**

El sistema crea automáticamente usuarios por defecto:
- **Admin**: `admin@websac3.com` / `Admin123!`
- **Usuario**: `user@websac3.com` / `User123!`

> ⚠️ Cambia estas credenciales en producción.

### Comandos útiles
```bash
# Ver logs
docker-compose logs -f

# Detener servicios
docker-compose down

# Reiniciar servicios
docker-compose restart

# Ver estado de servicios
docker-compose ps
```

---

## 🛠️ Tecnologías

### Backend
| Tecnología | Versión | Propósito |
|-----------|---------|-----------|
| **Go** | 1.23.5 | Lenguaje principal |
| **Gin** | Latest | Framework web |
| **GORM** | Latest | ORM para PostgreSQL |
| **JWT** | - | Autenticación |
| **Swagger** | 1.16.4 | Documentación API |
| **Gmail API** | v1 | Envío de correos |
| **wkhtmltopdf** | - | Generación de PDFs |

### Frontend
| Tecnología | Versión | Propósito |
|-----------|---------|-----------|
| **Next.js** | 15.0 | Framework React |
| **React** | 18 | Librería UI |
| **TypeScript** | Latest | Lenguaje tipado |
| **Redux Toolkit** | Latest | Estado global |
| **RTK Query** | Latest | Fetching de datos |
| **Tailwind CSS** | Latest | Framework CSS |
| **shadcn/ui** | Latest | Componentes UI |

### Infraestructura
| Tecnología | Versión | Propósito |
|-----------|---------|-----------|
| **PostgreSQL** | 14 | Base de datos |
| **Docker** | Latest | Containerización |
| **Docker Compose** | Latest | Orquestación |

---

## 📁 Estructura del Proyecto

```
websac3/
├── back/                           # Backend (Go)
│   └── websac3/
│       ├── adapter/                # Capa de adaptadores
│       │   ├── in/web/            # Controladores HTTP
│       │   └── out/               # Persistencia, notificaciones, PDF
│       ├── app/                    # Capa de aplicación
│       │   ├── domain/            # Entidades, servicios, casos de uso
│       │   └── port/              # Puertos (Commands/Queries)
│       ├── common/                 # Utilidades compartidas
│       │   ├── dependencies/      # Inyección de dependencias
│       │   ├── mediator/          # Mediador CQRS
│       │   └── validator/         # Validadores
│       ├── logs/                   # Archivos de log
│       ├── Dockerfile
│       ├── go.mod
│       └── README.md              # 📖 Documentación del backend
│
├── front/                          # Frontend (Next.js)
│   └── websac3-front/
│       ├── src/
│       │   ├── app/               # Rutas (App Router)
│       │   │   ├── (no-session)/  # Páginas públicas
│       │   │   └── (with-session)/ # Páginas protegidas
│       │   ├── components/        # Componentes React
│       │   │   ├── ui/           # Componentes base
│       │   │   └── websac3/      # Componentes específicos
│       │   ├── services/          # API (RTK Query)
│       │   ├── store/             # Redux store
│       │   └── middleware.ts      # Middleware de Next.js
│       ├── public/                # Archivos estáticos
│       ├── certificates/          # Certificados SSL
│       ├── Dockerfile
│       ├── server.js              # Servidor HTTPS
│       └── README.md              # 📖 Documentación del frontend
│
├── docker-compose.yml             # Orquestación de servicios
└── README.md                      # 📖 Este archivo
```

---

## ⚙️ Configuración

### Configuración de Gmail API

⚠️ **OBLIGATORIO**: El sistema requiere credenciales de Gmail API para enviar notificaciones por correo.

**Pasos:**

1. **Crear proyecto en Google Cloud Console**
   - Ve a [Google Cloud Console](https://console.cloud.google.com/)
   - Crea un nuevo proyecto o selecciona uno existente

2. **Habilitar Gmail API**
   - En el proyecto, ve a **APIs & Services** → **Library**
   - Busca "Gmail API" y habilítala

3. **Crear credenciales OAuth 2.0**
   - Ve a **APIs & Services** → **Credentials**
   - Clic en **Create Credentials** → **OAuth 2.0 Client ID**
   - Tipo de aplicación: **Desktop app**
   - Descarga el JSON y guárdalo como `back/websac3/.mail-credentials.json`

4. **Generar token de acceso**
   - La primera vez que ejecutes el backend, se generará automáticamente
   - O genera manualmente usando las herramientas de Google OAuth
   - Debe tener el scope: `https://www.googleapis.com/auth/gmail.send`
   - Guárdalo como `back/websac3/.mail-token.json`

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

### Variables de Entorno

**Backend** (`back/websac3/.env`):
```env
# Base de datos
DB_HOST=websac3-db
DB_USER=websac3
DB_PASSWORD=12345678
DB_NAME=websac3
DB_PORT=5432

# Aplicación
APP_NAME=WEBSAC3
LOGGER_PATH=logs/websac3.log

# Gmail API (REQUERIDO)
MAIL_FROM=tu_email@gmail.com
MAIL_CREDENTIALS_PATH=.mail-credentials.json
MAIL_TOKEN_PATH=.mail-token.json

# JWT
JWT_SECRET_KEY=websac3_jwt_secret_key
JWT_REFRESH_SECRET_KEY=websac3_jwt_refresh_secret_key
JWT_EXPIRATION_TIME=1800

# CORS
ALLOWED_ORIGINS=https://localhost:3000
```

**Frontend** (`front/websac3-front/.env.local`):
```env
# URL del backend
NEXT_PUBLIC_API_BASE_URL=https://localhost:8110/api/v1
```

---

## 🏃 Ejecución en Desarrollo

### Backend

```bash
cd back/websac3

# Instalar dependencias
go mod download

# Generar certificados SSL
chmod +x ensure-certs.sh && ./ensure-certs.sh

# Generar documentación Swagger
swag init --dir . --output ./docs --parseDependency --parseInternal

# Ejecutar
go run main.go
```

El backend estará disponible en: `https://localhost:8110`

> 📖 Ver [back/websac3/README.md](back/websac3/README.md) para más detalles.

### Frontend

```bash
cd front/websac3-front

# Instalar dependencias
npm install

# Generar certificados SSL
npm run generate-certs

# Ejecutar en desarrollo con HTTPS
npm run dev:https
```

El frontend estará disponible en: `https://localhost:3000`

> 📖 Ver [front/websac3-front/README.md](front/websac3-front/README.md) para más detalles.

### Base de Datos

**Opción 1: Con Docker**
```bash
docker-compose up -d websac3-db
```

**Opción 2: PostgreSQL local**
```sql
CREATE DATABASE websac3;
CREATE USER websac3 WITH PASSWORD '12345678';
GRANT ALL PRIVILEGES ON DATABASE websac3 TO websac3;
```

---

## ✨ Características Principales

### 👥 Gestión de Usuarios
- ✅ Registro y activación de usuarios
- ✅ Autenticación con JWT (access + refresh tokens)
- ✅ Sistema de roles y permisos
- ✅ Cambio y recuperación de contraseña
- ✅ Activación/desactivación de usuarios

### 🎓 Gestión Académica
- ✅ Administración de instituciones de educación superior
- ✅ Gestión de programas de grado
- ✅ Catálogo de cursos de ciberseguridad
- ✅ Taxonomía de cursos (tipo, naturaleza, nivel)
- ✅ Unidades de duración

### 💼 Roles Profesionales
- ✅ 16 roles de ciberseguridad (SOC Analyst, Pentester, etc.)
- ✅ Áreas de conocimiento asociadas
- ✅ Gestión de competencias

### 📝 Solicitudes de Acceso
- ✅ Creación de solicitudes
- ✅ Flujo de aprobación/rechazo
- ✅ Estados de solicitud
- ✅ Notificaciones automáticas por email
- ✅ Historial de cambios

### 🎯 Consultas de Expertos
- ✅ Gestión de consultas
- ✅ Asignación de expertos
- ✅ Estados de consulta
- ✅ Seguimiento y resolución

### 📊 Reportes y Estadísticas
- ✅ Generación de reportes en PDF
- ✅ Estadísticas del sistema
- ✅ Exportación de datos
- ✅ Dashboard con métricas

### 📧 Sistema de Notificaciones
- ✅ Templates HTML personalizables
- ✅ Envío de correos con Gmail API
- ✅ Soporte para attachments
- ✅ Notificaciones de eventos importantes

### 🌐 Internacionalización
- ✅ Soporte multiidioma (español, inglés)
- ✅ Mensajes localizados
- ✅ Detección automática de idioma

### 🔒 Seguridad
- ✅ Comunicación HTTPS/TLS
- ✅ Autenticación JWT
- ✅ OAuth2 para Gmail
- ✅ Validación de datos de entrada
- ✅ Sanitización de queries SQL
- ✅ Gestión segura de contraseñas

---

## 📚 Documentación Adicional

### Backend
- 📖 [Backend README](back/websac3/README.md) - Documentación completa del backend
- 🔗 [API Swagger](https://localhost:8110/swagger/index.html) - Documentación interactiva de la API

### Frontend
- 📖 [Frontend README](front/websac3-front/README.md) - Documentación completa del frontend
- 📖 [HTTPS Setup](front/websac3-front/HTTPS-SETUP.md) - Configuración HTTPS
- 📖 [ENV Configuration](front/websac3-front/ENV-CONFIGURATION.md) - Variables de entorno

### Arquitectura
- **Backend**: Arquitectura Hexagonal + CQRS
  - `adapter/in`: Controladores HTTP
  - `adapter/out`: Persistencia, notificaciones, PDF
  - `app/domain`: Entidades y lógica de negocio
  - `app/port`: Interfaces (Commands/Queries)
  - `common`: Utilidades compartidas

- **Frontend**: Next.js App Router + Redux Toolkit
  - `app/(no-session)`: Páginas públicas
  - `app/(with-session)`: Páginas protegidas
  - `components/ui`: Componentes base
  - `components/websac3`: Componentes específicos
  - `services`: API con RTK Query

---

## 🔐 Seguridad

### Archivos Sensibles

⚠️ **IMPORTANTE**: Los siguientes archivos contienen información sensible y **NO deben ser versionados**:

```
back/websac3/
  ├── .env                      # Variables de entorno
  ├── .mail-credentials.json    # Credenciales OAuth2 Google
  ├── .mail-token.json          # Token OAuth2
  └── certs/                    # Certificados SSL/TLS

front/websac3-front/
  ├── .env.local                # Variables de entorno locales
  └── certificates/             # Certificados SSL
```

Asegúrate de que estos archivos estén en `.gitignore`.

### Buenas Prácticas
- ✅ Cambiar credenciales por defecto en producción
- ✅ Usar contraseñas fuertes para la base de datos
- ✅ Actualizar `JWT_SECRET_KEY` con un valor seguro
- ✅ Configurar `ALLOWED_ORIGINS` correctamente
- ✅ Usar certificados SSL válidos en producción
- ✅ Mantener las dependencias actualizadas

---

## 🗄️ Base de Datos

### Inicialización Automática

🚀 **La base de datos se inicializa automáticamente** la primera vez que ejecutas el backend.

**¿Qué incluye?**
- ✅ Migración de todas las tablas
- ✅ Roles y permisos por defecto
- ✅ Usuarios administrativos
- ✅ 16 roles profesionales de ciberseguridad
- ✅ Datos de enumeraciones (estados, tipos, etc.)

**Comandos CLI disponibles:**
```bash
# Inicialización manual (opcional)
go run main.go init

# Migrar todos los modelos
go run main.go migrate:all

# Ejecutar seeders
go run main.go seed:run --seed=essential_data

# Reset de base de datos (⚠️ Elimina todos los datos)
go run main.go migrate:reset
```

---

## 🧪 Testing

### Backend
```bash
cd back/websac3
go test ./... -v
```

### Frontend
```bash
cd front/websac3-front
npm run test
```

---

## 📦 Despliegue en Producción

### Con Docker Compose (Recomendado)

```bash
# 1. Configurar variables de entorno para producción
# Editar docker-compose.yml y archivos .env

# 2. Construir y ejecutar
docker-compose up -d --build

# 3. Ver logs
docker-compose logs -f
```

### Sin Docker

**Backend:**
```bash
cd back/websac3
go build -o websac3 .
./websac3
```

**Frontend:**
```bash
cd front/websac3-front
npm run build
npm run start:https
```

> 📖 En producción, considera usar Nginx o Apache como proxy inverso.

---

## 🤝 Contribución

Para contribuir al proyecto:

1. Fork el repositorio
2. Crea una rama para tu feature (`git checkout -b feature/AmazingFeature`)
3. Commit tus cambios (`git commit -m 'Add some AmazingFeature'`)
4. Push a la rama (`git push origin feature/AmazingFeature`)
5. Abre un Pull Request

---

## 📝 Licencia

Este proyecto es privado y propietario. Todos los derechos reservados.

---

## 👥 Contacto

- **Email de Soporte**: j0rg3.4b3ll4@gmail.com
- **Proyecto**: WebSAC3 - Sistema Web de Asesoramiento en Ciberseguridad

---

## 📋 Requisitos del Sistema

### Mínimos
- **CPU**: 2 cores
- **RAM**: 4 GB
- **Disco**: 10 GB libres
- **OS**: Windows 10, macOS, Linux

### Recomendados
- **CPU**: 4 cores
- **RAM**: 8 GB
- **Disco**: 20 GB SSD
- **OS**: Linux (Ubuntu 20.04+)

---

## 🎯 Roadmap

- [ ] Módulo de mensajería interna
- [ ] Dashboard de analytics avanzado
- [ ] Exportación a Excel
- [ ] Notificaciones en tiempo real (WebSocket)
- [ ] Modo oscuro
- [ ] App móvil (React Native)
- [ ] Integración con Microsoft Teams
- [ ] Sistema de backup automático

---

<div align="center">

**Desarrollado con ❤️ usando Go, Next.js y PostgreSQL**

⭐ Si te gusta este proyecto, considera darle una estrella

</div>

