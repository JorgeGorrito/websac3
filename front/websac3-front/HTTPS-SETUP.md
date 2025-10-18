# Configuración HTTPS para WebSAC3 Frontend

Este documento explica cómo implementar HTTPS en tu proyecto Next.js en diferentes entornos.

## ⚠️ Importante: Backend en HTTPS

**El frontend ahora está configurado para usar HTTPS en todas las conexiones con el backend.**

URL por defecto: `https://localhost:8110/api/v1/{language}`

Para configurar una URL diferente, consulta el archivo **[ENV-CONFIGURATION.md](./ENV-CONFIGURATION.md)**

## 📋 Tabla de Contenidos

- [Configuración del Backend](#configuración-del-backend)
- [Desarrollo Local](#desarrollo-local)
- [Producción con Docker + Nginx](#producción-con-docker--nginx)
- [Producción con Let's Encrypt](#producción-con-lets-encrypt)

---

## 🔌 Configuración del Backend

### Paso 1: Asegúrate de que tu Backend esté en HTTPS

El frontend ahora envía todas las peticiones a `https://localhost:8110/api/v1/{language}` por defecto.

**Opciones:**

1. **Backend ya tiene HTTPS:** No necesitas hacer nada más en esta sección.

2. **Backend en HTTP:** Tienes dos opciones:
   - **Opción A (Recomendada):** Configura HTTPS en tu backend
   - **Opción B (Solo Desarrollo):** Crear un archivo `.env.local`:
     ```bash
     NEXT_PUBLIC_API_BASE_URL=http://localhost:8110/api/v1
     ```

3. **Backend en diferente URL/Puerto:** Crear archivo `.env.local`:
   ```bash
   NEXT_PUBLIC_API_BASE_URL=https://tu-backend-url:puerto/api/v1
   ```

### Paso 2: Verificar Certificados SSL del Backend

Si el backend usa certificados autofirmados en desarrollo:

1. Visita directamente la URL del backend en tu navegador: `https://localhost:8110`
2. Acepta la advertencia de seguridad del certificado
3. Ahora el frontend podrá comunicarse con el backend

📖 **Más información:** Lee [ENV-CONFIGURATION.md](./ENV-CONFIGURATION.md) para configuración detallada.

---

## 🔧 Desarrollo Local

### Opción 1: Certificados Autofirmados (Rápido)

1. **Generar certificados SSL locales:**
   ```bash
   npm run generate-certs
   ```

2. **Iniciar el servidor en modo desarrollo con HTTPS:**
   ```bash
   npm run dev:https
   ```

3. **Acceder a la aplicación:**
   ```
   https://localhost:3000
   ```

   ⚠️ **Nota:** Tu navegador mostrará una advertencia de seguridad porque el certificado es autofirmado. Esto es normal en desarrollo.
   - En Chrome/Edge: Haz clic en "Avanzado" → "Continuar a localhost"
   - En Firefox: Haz clic en "Avanzado" → "Aceptar el riesgo y continuar"

### Opción 2: mkcert (Recomendado para Desarrollo)

`mkcert` crea certificados SSL que el navegador reconoce como válidos.

1. **Instalar mkcert:**

   **Windows (con Chocolatey):**
   ```bash
   choco install mkcert
   ```

   **macOS:**
   ```bash
   brew install mkcert
   ```

   **Linux:**
   ```bash
   wget https://github.com/FiloSottile/mkcert/releases/download/v1.4.4/mkcert-v1.4.4-linux-amd64
   chmod +x mkcert-v1.4.4-linux-amd64
   sudo mv mkcert-v1.4.4-linux-amd64 /usr/local/bin/mkcert
   ```

2. **Instalar la CA local:**
   ```bash
   mkcert -install
   ```

3. **Crear directorio para certificados:**
   ```bash
   mkdir certificates
   ```

4. **Generar certificados para localhost:**
   ```bash
   mkcert -key-file certificates/localhost-key.pem -cert-file certificates/localhost.pem localhost 127.0.0.1 ::1
   ```

5. **Iniciar el servidor:**
   ```bash
   npm run dev:https
   ```

6. **Acceder sin advertencias:**
   ```
   https://localhost:3000
   ```

---

## 🐳 Producción con Docker + Nginx

Esta configuración usa Nginx como reverse proxy con soporte SSL.

### 1. Preparar Certificados SSL

**Opción A: Certificados Autofirmados (Testing)**
```bash
# Generar certificados
npm run generate-certs
```

**Opción B: Certificados de Producción**
- Obtén certificados SSL de tu proveedor (GoDaddy, Namecheap, etc.)
- O usa Let's Encrypt (ver sección siguiente)
- Coloca los archivos en la carpeta `certificates/`:
  - `cert.pem` (certificado)
  - `key.pem` (clave privada)

### 2. Actualizar nginx.conf

Si tus certificados tienen nombres diferentes, edita `nginx.conf`:
```nginx
ssl_certificate /etc/nginx/ssl/tu-certificado.crt;
ssl_certificate_key /etc/nginx/ssl/tu-clave-privada.key;
```

### 3. Construir y Ejecutar con Docker Compose

```bash
# Construir las imágenes
docker-compose -f docker-compose.https.yml build

# Iniciar los servicios
docker-compose -f docker-compose.https.yml up -d

# Ver logs
docker-compose -f docker-compose.https.yml logs -f
```

### 4. Acceder a la Aplicación

```
https://localhost
https://tu-dominio.com
```

### 5. Detener los Servicios

```bash
docker-compose -f docker-compose.https.yml down
```

---

## 🔐 Producción con Let's Encrypt (Certificados Gratuitos)

Let's Encrypt proporciona certificados SSL gratuitos y renovables automáticamente.

### Requisitos:
- Un dominio público (ej: `websac3.example.com`)
- El servidor debe ser accesible desde internet en el puerto 80

### Configuración con Certbot

1. **Instalar Certbot:**

   **Ubuntu/Debian:**
   ```bash
   sudo apt update
   sudo apt install certbot python3-certbot-nginx
   ```

   **CentOS/RHEL:**
   ```bash
   sudo yum install certbot python3-certbot-nginx
   ```

2. **Obtener certificado:**
   ```bash
   sudo certbot certonly --standalone -d websac3.example.com
   ```

   Los certificados se guardarán en:
   - Certificado: `/etc/letsencrypt/live/websac3.example.com/fullchain.pem`
   - Clave: `/etc/letsencrypt/live/websac3.example.com/privkey.pem`

3. **Actualizar docker-compose.https.yml:**
   ```yaml
   nginx:
     volumes:
       - ./nginx.conf:/etc/nginx/conf.d/default.conf
       - /etc/letsencrypt:/etc/letsencrypt:ro
   ```

4. **Actualizar nginx.conf:**
   ```nginx
   ssl_certificate /etc/letsencrypt/live/websac3.example.com/fullchain.pem;
   ssl_certificate_key /etc/letsencrypt/live/websac3.example.com/privkey.pem;
   ```

5. **Renovación automática:**
   ```bash
   # Probar renovación
   sudo certbot renew --dry-run

   # Configurar cron para renovación automática (se ejecuta 2 veces al día)
   sudo crontab -e
   ```

   Agregar:
   ```
   0 0,12 * * * certbot renew --quiet --post-hook "docker-compose -f /ruta/a/tu/proyecto/docker-compose.https.yml restart nginx"
   ```

---

## 🔍 Verificar Configuración SSL

### Herramientas Online:
- [SSL Labs SSL Test](https://www.ssllabs.com/ssltest/)
- [SSL Checker](https://www.sslshopper.com/ssl-checker.html)

### Comandos de Verificación:

```bash
# Ver información del certificado
openssl s_client -connect localhost:443 -servername localhost

# Verificar fecha de expiración
openssl s_client -connect tu-dominio.com:443 2>/dev/null | openssl x509 -noout -dates

# Probar conexión HTTPS
curl -k https://localhost:3000
```

---

## 🛠️ Solución de Problemas

### Error: "ENOENT: no such file or directory, open 'certificates/localhost-key.pem'"
**Solución:** Ejecuta `npm run generate-certs` para generar los certificados.

### Error: "Address already in use"
**Solución:** 
```bash
# Windows
netstat -ano | findstr :3000
taskkill /PID <PID> /F

# Linux/macOS
lsof -ti:3000 | xargs kill -9
```

### Navegador muestra advertencia de seguridad
**Para desarrollo:** Esto es normal con certificados autofirmados. Puedes continuar de forma segura.
**Para producción:** Usa certificados válidos de Let's Encrypt o un proveedor comercial.

### Docker: Nginx no puede acceder a los certificados
**Solución:** Verifica los permisos de los archivos de certificados:
```bash
chmod 644 certificates/*.pem
```

---

## 📝 Archivos Importantes

- `server.js` - Servidor HTTPS personalizado para Next.js
- `generate-certificates.js` - Script para generar certificados SSL
- `nginx.conf` - Configuración de Nginx con SSL
- `docker-compose.https.yml` - Docker Compose con Nginx
- `certificates/` - Directorio para certificados SSL (git-ignored)

---

## 🚀 Scripts Disponibles

| Script | Descripción |
|--------|-------------|
| `npm run dev` | Desarrollo HTTP (puerto 3000) |
| `npm run dev:https` | Desarrollo HTTPS (puerto 3000) |
| `npm run build` | Construir para producción |
| `npm run start` | Producción HTTP |
| `npm run start:https` | Producción HTTPS |
| `npm run generate-certs` | Generar certificados SSL autofirmados |

---

## 🔒 Mejores Prácticas de Seguridad

1. **Nunca hagas commit de certificados privados:** Los archivos `.pem` y `.key` están en `.gitignore`
2. **Usa HTTPS en producción:** Siempre usa SSL/TLS en entornos de producción
3. **Mantén actualizados tus certificados:** Configura renovación automática
4. **Usa TLS 1.2 o superior:** Desactiva protocolos antiguos (SSLv3, TLS 1.0, TLS 1.1)
5. **Headers de seguridad:** Ya incluidos en la configuración de Nginx

---

## 📚 Referencias

- [Next.js Custom Server](https://nextjs.org/docs/advanced-features/custom-server)
- [mkcert](https://github.com/FiloSottile/mkcert)
- [Let's Encrypt](https://letsencrypt.org/)
- [Nginx SSL Configuration](https://nginx.org/en/docs/http/configuring_https_servers.html)
- [Mozilla SSL Configuration Generator](https://ssl-config.mozilla.org/)

