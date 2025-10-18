#!/bin/bash
set -euo pipefail

CERT_DIR="${CERT_DIR:-/app/certificates}"
CERT_CN="${CERT_CN:-localhost}"
KEY_BITS="${KEY_BITS:-2048}"
CERT_DAYS="${CERT_DAYS:-365}"
CERT_KEY="${CERT_KEY:-$CERT_DIR/localhost-key.pem}"
CERT_CRT="${CERT_CRT:-$CERT_DIR/localhost.pem}"

mkdir -p "$CERT_DIR"

# Verificar si los certificados ya existen y son válidos
if [[ -f "$CERT_KEY" && -f "$CERT_CRT" ]]; then
  echo "[certs] Verificando certificados existentes..."
  
  # Verificar si el certificado está próximo a expirar (30 días)
  if openssl x509 -checkend 2592000 -noout -in "$CERT_CRT" 2>/dev/null; then
    echo "[certs] Usando certificados existentes válidos en $CERT_DIR"
    exec "$@"
  else
    echo "[certs] Certificado próximo a expirar o inválido, regenerando..."
  fi
fi

echo "[certs] Generando certificado SSL autofirmado..."

# Crear configuración temporal para Subject Alternative Names (SAN)
SAN_CFG="$(mktemp)"
cat > "$SAN_CFG" << EOF
[ req ]
default_bits       = $KEY_BITS
distinguished_name = req_distinguished_name
req_extensions     = req_ext
prompt             = no

[ req_distinguished_name ]
C  = CO
ST = Meta
L  = Villavicencio
O  = WebSAC3
OU = Frontend
CN = $CERT_CN

[ req_ext ]
subjectAltName = @alt_names

[ alt_names ]
DNS.1 = localhost
DNS.2 = *.localhost
IP.1  = 127.0.0.1
IP.2  = ::1
EOF

# Generar clave privada
openssl genrsa -out "$CERT_KEY" "$KEY_BITS" 2>/dev/null

# Generar certificado autofirmado con SAN
openssl req -new -x509 -key "$CERT_KEY" -out "$CERT_CRT" \
  -days "$CERT_DAYS" -extensions req_ext -config "$SAN_CFG" 2>/dev/null

# Limpiar archivo temporal
rm -f "$SAN_CFG"

# Establecer permisos seguros
chmod 600 "$CERT_KEY"
chmod 644 "$CERT_CRT"

echo "[certs] ✓ Certificado SSL generado exitosamente"
echo "[certs]   - Clave privada: $CERT_KEY"
echo "[certs]   - Certificado: $CERT_CRT"
echo "[certs]   - Validez: $CERT_DAYS días"
echo "[certs]   - CN: $CERT_CN"
echo ""
echo "[certs] ⚠️  ADVERTENCIA: Este es un certificado autofirmado para desarrollo"
echo "[certs]    Tu navegador mostrará una advertencia de seguridad (esto es normal)"

# Ejecutar el comando pasado como parámetros al script
exec "$@"

