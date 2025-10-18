#!/bin/bash
set -euo pipefail

CERT_DIR="${CERT_DIR:-/app/certs}"
CERT_CN="${CERT_CN:-localhost}"
KEY_BITS="${KEY_BITS:-2048}"
CERT_DAYS="${CERT_DAYS:-365}"
RENEW_BEFORE_SECS="${RENEW_BEFORE_SECS:-2592000}"
CERT_KEY="${CERT_KEY:-$CERT_DIR/localhost.key}"
CERT_CRT="${CERT_CRT:-$CERT_DIR/localhost.crt}"
CERT_CSR="${CERT_CSR:-$CERT_DIR/localhost.csr}"

mkdir -p $CERT_DIR

if [[ ! -f "$CERT_KEY" || ! -f "$CERT_CRT" || ! -f "$CERT_CSR" ]]; then
    echo "Generando certificado SSL..."
    ts="$(date +%Y%m%d%H%M%S)"
    openssl genrsa -out "$CERT_KEY" "$KEY_BITS"
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
O  = Dev
OU = Dev
CN = $CERT_CN

[ req_ext ]
subjectAltName = @alt_names

[ alt_names ]
DNS.1 = localhost
IP.1  = 127.0.0.1
EOF
  openssl req -new -key "$CERT_KEY" -out "$CERT_CSR" -config "$SAN_CFG"

  # Cert autofirmado con SAN
  openssl x509 -req -in "$CERT_CSR" -signkey "$CERT_KEY" -out "$CERT_CRT" \
    -days "$CERT_DAYS" -extensions req_ext -extfile "$SAN_CFG"

  rm -f "$SAN_CFG"

  chmod 600 "$CERT_KEY"
  echo "[certs] Certificado nuevo creado en $CERT_DIR (validez: $CERT_DAYS días)"
else
  echo "[certs] Usando certificados existentes en $CERT_DIR"
fi

# Ejecutar el comando pasado como parámetros al script
exec "$@"