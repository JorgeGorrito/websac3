/**
 * Utility functions for translating role names from English to Spanish
 */

export const ROLE_TRANSLATIONS: Record<string, string> = {
  "program lead": "Director de programa",
  "guest": "Invitado", 
  "admin": "Administrador",
  "cybersecurity auditor": "Auditor Ciberseguridad",
};

/**
 * Translates a role name from English to Spanish
 * @param roleName - The role name in English
 * @returns The translated role name in Spanish, or the original name if no translation exists
 */
export function translateRole(roleName: string): string {
  return ROLE_TRANSLATIONS[roleName.toLowerCase()] || roleName;
}

/**
 * Gets all available role translations
 * @returns Object with English role names as keys and Spanish translations as values
 */
export function getRoleTranslations(): Record<string, string> {
  return ROLE_TRANSLATIONS;
}
