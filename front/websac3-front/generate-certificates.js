const { execSync } = require('child_process');
const fs = require('fs');
const path = require('path');

const certDir = path.join(__dirname, 'certificates');

// Create certificates directory if it doesn't exist
if (!fs.existsSync(certDir)) {
  fs.mkdirSync(certDir, { recursive: true });
  console.log('✓ Created certificates directory');
}

// Check if certificates already exist
const keyPath = path.join(certDir, 'localhost-key.pem');
const certPath = path.join(certDir, 'localhost.pem');

if (fs.existsSync(keyPath) && fs.existsSync(certPath)) {
  console.log('✓ SSL certificates already exist');
  process.exit(0);
}

console.log('Generating self-signed SSL certificates...');

try {
  // Generate self-signed certificate using OpenSSL
  // This works on Windows (with Git Bash), macOS, and Linux
  execSync(
    `openssl req -x509 -nodes -days 365 -newkey rsa:2048 ` +
    `-keyout "${keyPath}" ` +
    `-out "${certPath}" ` +
    `-subj "/C=CO/ST=Meta/L=Villavicencio/O=WebSAC3/CN=localhost"`,
    { stdio: 'inherit' }
  );
  
  console.log('✓ SSL certificates generated successfully!');
  console.log(`  Key: ${keyPath}`);
  console.log(`  Cert: ${certPath}`);
  console.log('\nNote: These are self-signed certificates for development only.');
  console.log('Your browser will show a security warning - this is normal.');
} catch (error) {
  console.error('Error generating certificates:', error.message);
  console.log('\nAlternative: Use mkcert (recommended)');
  console.log('1. Install mkcert: https://github.com/FiloSottile/mkcert');
  console.log('2. Run: mkcert -install');
  console.log('3. Run: mkcert -key-file certificates/localhost-key.pem -cert-file certificates/localhost.pem localhost');
  process.exit(1);
}


