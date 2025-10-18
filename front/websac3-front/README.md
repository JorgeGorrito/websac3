# WebSAC3 Frontend

Sistema Web de Asesoramiento en Ciberseguridad para Programas Académicos - Frontend

Este es un proyecto [Next.js](https://nextjs.org) que proporciona la interfaz de usuario para WebSAC3.

## 🚀 Inicio Rápido

### Requisitos Previos

- Node.js >= 20.11.0
- npm o yarn
- Backend de WebSAC3 corriendo en HTTPS (por defecto: `https://localhost:8110`)

### Instalación

```bash
# Instalar dependencias
npm install
```

### Desarrollo

**Opción 1: HTTP (por defecto)**
```bash
npm run dev
```
Abre [http://localhost:3000](http://localhost:3000) en tu navegador.

**Opción 2: HTTPS (recomendado)**
```bash
# Generar certificados SSL
npm run generate-certs

# Iniciar servidor con HTTPS
npm run dev:https
```
Abre [https://localhost:3000](https://localhost:3000) en tu navegador.

## 🔒 Configuración HTTPS

**⚠️ IMPORTANTE:** El frontend está configurado para comunicarse con el backend usando **HTTPS** por defecto.

### URL del Backend por Defecto
```
https://localhost:8110/api/v1/{language}
```

### Configurar URL Personalizada

Si tu backend está en una URL diferente, crea un archivo `.env.local` en la raíz del proyecto:

```bash
# .env.local
NEXT_PUBLIC_API_BASE_URL=https://tu-backend-url:puerto/api/v1
```

📖 **Documentación completa:**
- **[HTTPS-SETUP.md](./HTTPS-SETUP.md)** - Configuración HTTPS para el frontend
- **[ENV-CONFIGURATION.md](./ENV-CONFIGURATION.md)** - Configuración de variables de entorno y backend

## 📦 Scripts Disponibles

| Script | Descripción |
|--------|-------------|
| `npm run dev` | Servidor de desarrollo HTTP (puerto 3000) |
| `npm run dev:https` | Servidor de desarrollo HTTPS (puerto 3000) |
| `npm run build` | Construir para producción |
| `npm run start` | Servidor de producción HTTP |
| `npm run start:https` | Servidor de producción HTTPS |
| `npm run lint` | Ejecutar linter |
| `npm run generate-certs` | Generar certificados SSL autofirmados |

## 🏗️ Estructura del Proyecto

```
websac3-front/
├── src/
│   ├── app/              # Rutas y páginas (App Router)
│   │   ├── (no-session)/ # Páginas sin sesión (login, registro)
│   │   └── (with-session)/ # Páginas con sesión (dashboard, admin, etc.)
│   ├── components/       # Componentes React
│   │   ├── ui/          # Componentes UI base
│   │   └── websac3/     # Componentes específicos de WebSAC3
│   ├── services/        # Servicios API (RTK Query)
│   ├── store/           # Redux store
│   ├── hooks/           # Custom hooks
│   ├── lib/             # Utilidades y helpers
│   └── middleware.ts    # Middleware de Next.js
├── public/              # Archivos estáticos
├── certificates/        # Certificados SSL (generados, git-ignored)
├── server.js            # Servidor HTTPS personalizado
└── docker-compose.https.yml # Docker Compose con Nginx SSL
```

## 🐳 Despliegue con Docker

El Dockerfile incluye generación automática de certificados SSL y compilación de la aplicación:

```bash
# Construir y ejecutar (compila y genera certificados automáticamente)
docker-compose up -d

# Ver logs
docker-compose logs -f websac3-front

# Detener
docker-compose down
```

Accede a la aplicación en [https://localhost:3000](https://localhost:3000)

**Características:**
- ✓ Compilación automática dentro del contenedor
- ✓ Generación automática de certificados SSL
- ✓ Verificación y renovación automática (si expiran en 30 días)
- ✓ Certificados persistidos en volumen Docker
- ✓ Build multi-etapa optimizado

### Variables de Entorno

Personaliza el comportamiento en `docker-compose.yml`:

```yaml
environment:
  - CERT_CN=mi-dominio.com           # Nombre del certificado
  - CERT_DAYS=730                    # Validez del certificado (días)
  - NEXT_PUBLIC_API_BASE_URL=https://mi-backend:8110/api/v1
```

📖 **Guía rápida:** Ver [DOCKER-QUICKSTART.md](./DOCKER-QUICKSTART.md)  
📖 **Configuración HTTPS:** Ver [HTTPS-SETUP.md](./HTTPS-SETUP.md)

## Learn More

To learn more about Next.js, take a look at the following resources:

- [Next.js Documentation](https://nextjs.org/docs) - learn about Next.js features and API.
- [Learn Next.js](https://nextjs.org/learn) - an interactive Next.js tutorial.

You can check out [the Next.js GitHub repository](https://github.com/vercel/next.js) - your feedback and contributions are welcome!

## Deploy on Vercel

The easiest way to deploy your Next.js app is to use the [Vercel Platform](https://vercel.com/new?utm_medium=default-template&filter=next.js&utm_source=create-next-app&utm_campaign=create-next-app-readme) from the creators of Next.js.

Check out our [Next.js deployment documentation](https://nextjs.org/docs/app/building-your-application/deploying) for more details.
