# Etapa 1: Builder - Construir la aplicación
FROM node:20.11.0-bookworm-slim AS builder

# Instalar OpenSSL para generación de certificados
RUN apt-get update && \
    apt-get install -y --no-install-recommends openssl ca-certificates && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /app

# Copiar archivos de configuración de dependencias
COPY package*.json ./

# Instalar TODAS las dependencias (incluidas devDependencies para compilar)
RUN npm ci && \
    npm cache clean --force

# Copiar el resto del código fuente
COPY . .

# Construir la aplicación Next.js
RUN npm run build

# Eliminar devDependencies después de compilar para reducir tamaño
RUN npm prune --production

# Etapa 2: Runner - Ejecutar la aplicación
FROM node:20.11.0-bookworm-slim

# Instalar OpenSSL y CA certificates
RUN apt-get update && \
    apt-get install -y --no-install-recommends openssl ca-certificates && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /app

# Copiar node_modules y archivos construidos desde builder
COPY --from=builder /app/node_modules ./node_modules
COPY --from=builder /app/.next ./.next
COPY --from=builder /app/public ./public
COPY --from=builder /app/package*.json ./
COPY --from=builder /app/next.config.ts ./
COPY --from=builder /app/server.js ./

# Copiar script de generación de certificados
COPY ensure-certs-front.sh /app/ensure-certs-front.sh
RUN chmod +x /app/ensure-certs-front.sh

# Crear directorio para certificados
RUN mkdir -p /app/certificates

# Exponer puerto HTTPS
EXPOSE 3000

# Variables de entorno por defecto
ENV NODE_ENV=production
ENV CERT_DIR=/app/certificates
ENV CERT_CN=localhost
ENV KEY_BITS=2048
ENV CERT_DAYS=365

# Ejecutar script de certificados y luego iniciar el servidor
ENTRYPOINT ["/app/ensure-certs-front.sh"]
CMD ["npm", "run", "start:https"]
