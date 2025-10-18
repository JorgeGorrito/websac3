# Configuración de Variables de Entorno

## 📋 Configuración del API Backend

El frontend ahora está configurado para usar **HTTPS** en lugar de HTTP para todas las conexiones con el backend.

### URL por Defecto

Si no se configura ninguna variable de entorno, la aplicación usará:
```
https://localhost:8110/api/v1/{language}
```

---

## 🔧 Configurar Variables de Entorno

### Para Desarrollo Local

1. **Crear archivo `.env.local` en la raíz del proyecto:**

```bash
# .env.local
NEXT_PUBLIC_API_BASE_URL=https://localhost:8110/api/v1
```

2. **Reiniciar el servidor de desarrollo:**
```bash
npm run dev
# o
npm run dev:https
```

### Para Producción

1. **Crear archivo `.env.production` en la raíz del proyecto:**

```bash
# .env.production
NEXT_PUBLIC_API_BASE_URL=https://api.websac3.example.com/api/v1
```

2. **Construir la aplicación:**
```bash
npm run build
npm run start
```

---

## 🌐 Ejemplos de Configuración

### Backend en Localhost con HTTPS (Puerto 8110)
```bash
NEXT_PUBLIC_API_BASE_URL=https://localhost:8110/api/v1
```

### Backend en Docker con HTTPS
```bash
NEXT_PUBLIC_API_BASE_URL=https://backend:8110/api/v1
```

### Backend Remoto en Desarrollo
```bash
NEXT_PUBLIC_API_BASE_URL=https://dev-api.websac3.unillanos.edu.co/api/v1
```

### Backend en Producción
```bash
NEXT_PUBLIC_API_BASE_URL=https://api.websac3.unillanos.edu.co/api/v1
```

### Backend en HTTP (No recomendado para producción)
```bash
NEXT_PUBLIC_API_BASE_URL=http://localhost:8110/api/v1
```

---

## 📝 Notas Importantes

### 1. **Prefijo NEXT_PUBLIC_**
- Las variables deben empezar con `NEXT_PUBLIC_` para estar disponibles en el navegador
- Sin este prefijo, Next.js no expondrá la variable al cliente

### 2. **No incluir el idioma en la URL**
- ❌ Incorrecto: `https://localhost:8110/api/v1/es`
- ✅ Correcto: `https://localhost:8110/api/v1`
- El idioma (`es` o `en`) se agrega automáticamente

### 3. **No incluir barra final**
- ❌ Incorrecto: `https://localhost:8110/api/v1/`
- ✅ Correcto: `https://localhost:8110/api/v1`
- La aplicación elimina automáticamente las barras finales

### 4. **Certificados SSL**
- Si tu backend usa HTTPS con certificados autofirmados en desarrollo, el navegador puede mostrar advertencias
- Para desarrollo: acepta el certificado en el navegador
- Para producción: usa certificados válidos (Let's Encrypt, etc.)

---

## 🔍 Verificar Configuración

### Ver la URL que está usando la aplicación:

Abre las DevTools del navegador (F12) y ejecuta:
```javascript
console.log('API Base URL:', process.env.NEXT_PUBLIC_API_BASE_URL);
```

O revisa las llamadas de red en la pestaña "Network" para ver las URLs completas.

---

## 🐳 Configuración con Docker

### docker-compose.yml

```yaml
services:
  frontend:
    build: .
    environment:
      - NEXT_PUBLIC_API_BASE_URL=https://backend:8110/api/v1
    ports:
      - "3000:3000"
    depends_on:
      - backend

  backend:
    image: tu-backend-image
    ports:
      - "8110:8110"
```

### Dockerfile con variables de entorno en build time

```dockerfile
FROM node:18-alpine AS builder

WORKDIR /app

COPY package*.json ./
RUN npm install --frozen-lockfile

COPY . .

# Define la variable de entorno para build time
ARG NEXT_PUBLIC_API_BASE_URL
ENV NEXT_PUBLIC_API_BASE_URL=$NEXT_PUBLIC_API_BASE_URL

RUN npm run build

FROM node:18-alpine AS production

WORKDIR /app

COPY --from=builder /app/package*.json ./
COPY --from=builder /app/.next ./.next
COPY --from=builder /app/public ./public
COPY --from=builder /app/node_modules ./node_modules

ENV NODE_ENV=production
ENV PORT=3000

EXPOSE 3000
CMD ["npm", "start"]
```

Construir con:
```bash
docker build --build-arg NEXT_PUBLIC_API_BASE_URL=https://api.websac3.com/api/v1 -t websac3-front .
```

---

## 🛠️ Solución de Problemas

### Error: "Failed to fetch" o "CORS error"

**Causa:** El backend no está accesible o no permite conexiones HTTPS desde el frontend.

**Solución:**
1. Verifica que el backend esté corriendo en HTTPS
2. Verifica los certificados SSL
3. Revisa la configuración CORS del backend

### Error: "NET::ERR_CERT_INVALID"

**Causa:** Certificado SSL no válido o autofirmado.

**Solución para Desarrollo:**
1. Acepta el certificado en el navegador visitando directamente la URL del API
2. O usa `mkcert` para generar certificados confiables localmente

**Solución para Producción:**
1. Usa certificados válidos de Let's Encrypt o un proveedor comercial
2. Verifica que el certificado no haya expirado

### La URL no cambia después de modificar .env.local

**Solución:**
1. Reinicia completamente el servidor de desarrollo (Ctrl+C y `npm run dev` de nuevo)
2. Las variables `NEXT_PUBLIC_*` se incrustan en el bundle durante el build
3. Para producción, debes hacer `npm run build` de nuevo

---

## 📚 Referencias

- [Next.js Environment Variables](https://nextjs.org/docs/basic-features/environment-variables)
- [Next.js API Routes](https://nextjs.org/docs/api-routes/introduction)
- [Docker Environment Variables](https://docs.docker.com/compose/environment-variables/)

---

## ✅ Checklist de Implementación

- [x] Cambiar URL por defecto a HTTPS en `src/services/api.ts`
- [ ] Crear archivo `.env.local` con tu configuración local
- [ ] Verificar que el backend esté corriendo en HTTPS
- [ ] Probar las llamadas al API desde el frontend
- [ ] Para producción: Crear `.env.production` con la URL correcta
- [ ] Configurar certificados SSL válidos para producción


