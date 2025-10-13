/**
 * Utility functions for role handling and normalization
 */

export type UserRole = 'admin' | 'guest' | 'program lead' | 'cybersecurity_auditor';

/**
 * Normalizes a role string by converting to lowercase and trimming whitespace
 * Also handles common role variations and typos
 * @param role - The role string to normalize
 * @returns The normalized role string
 */
export function normalizeRole(role: string | undefined | null): string {
  if (!role) return '';
  
  const normalized = role.toLowerCase().trim();
  
  // Handle common variations and typos
  const roleMappings: Record<string, string> = {
    'guess': 'guest', // Common typo
    'guests': 'guest', // Plural form
    'program_lead': 'program lead', // Underscore variant
    'programlead': 'program lead', // No space variant
    'cybersecurity_auditor': 'cybersecurity_auditor', // Keep as is
    'cybersecurityauditor': 'cybersecurity_auditor', // No underscore variant
  };
  
  return roleMappings[normalized] || normalized;
}

/**
 * Checks if a user role matches any of the required roles
 * @param userRole - The user's role
 * @param requiredRoles - The required roles (can be string or array)
 * @returns True if the user role matches any required role
 */
export function hasRequiredRole(userRole: string | undefined | null, requiredRoles: string | string[]): boolean {
  const normalizedUserRole = normalizeRole(userRole);
  const allowedRoles = Array.isArray(requiredRoles) ? requiredRoles : [requiredRoles];
  
  return allowedRoles.some(role => normalizeRole(role) === normalizedUserRole);
}

/**
 * Gets the appropriate dashboard route for a given role
 * @param role - The user's role
 * @returns The dashboard route path
 */
export function getDashboardRoute(role: string | undefined | null): string {
  const normalizedRole = normalizeRole(role);
  
  console.log(`🎯 getDashboardRoute: Original role "${role}" -> Normalized "${normalizedRole}"`);
  
  switch (normalizedRole) {
    case 'admin':
      return "/admin/dashboard";
    case 'guest':
    case 'program lead':
      return "/director/dashboard";
    case 'cybersecurity_auditor':
    case 'cybersecurity auditor':
      return "/experto/dashboard";
    default:
      console.warn(`Unknown role: "${role}" (normalized: "${normalizedRole}"). Redirecting to admin dashboard.`);
      return "/admin/dashboard";
  }
}

/**
 * Validates if a role is a valid UserRole
 * @param role - The role to validate
 * @returns True if the role is valid
 */
export function isValidRole(role: string | undefined | null): role is UserRole {
  const normalizedRole = normalizeRole(role);
  const validRoles: UserRole[] = ['admin', 'guest', 'program lead', 'cybersecurity_auditor'];
  return validRoles.includes(normalizedRole as UserRole);
}
