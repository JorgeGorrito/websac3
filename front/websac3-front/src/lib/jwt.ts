// Utility functions for JWT token handling

export interface JWTPayload {
  sub: number; // user ID
  username: string;
  email?: string; // Make email optional
  role: string;
  permissions?: Record<string, string[]>;
  iat: number; // issued at
  exp: number; // expires at
}

/**
 * Decode JWT token without verification (client-side only)
 * Note: This is for display purposes only. Never trust client-side decoded data for security decisions.
 */
export function decodeJWT(token: string): JWTPayload | null {
  try {
    const base64Url = token.split('.')[1];
    const base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/');
    const jsonPayload = decodeURIComponent(
      atob(base64)
        .split('')
        .map(c => '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2))
        .join('')
    );
    return JSON.parse(jsonPayload);
  } catch (error) {
    console.error('Error decoding JWT:', error);
    return null;
  }
}

/**
 * Check if JWT token is expired
 */
export function isTokenExpired(token: string): boolean {
  const payload = decodeJWT(token);
  if (!payload) return true;
  
  const currentTime = Math.floor(Date.now() / 1000);
  return payload.exp < currentTime;
}

/**
 * Get time until token expires (in seconds)
 */
export function getTokenTimeUntilExpiry(token: string): number {
  const payload = decodeJWT(token);
  if (!payload) return 0;
  
  const currentTime = Math.floor(Date.now() / 1000);
  return Math.max(0, payload.exp - currentTime);
}
