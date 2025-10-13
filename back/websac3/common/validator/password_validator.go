package validator

import (
	"regexp"
	"websac3/app/domain/errs"
	"websac3/app/port/out/message"
)

// ValidatePasswordComplexity valida que la contraseña cumpla con los requisitos de complejidad:
// - Al menos una letra (mayúscula o minúscula)
// - Al menos un número
// - Al menos un símbolo especial
func ValidatePasswordComplexity(password string, msgProvider message.Provider, lang string) error {
	// Validar que contenga al menos una letra
	hasLetter := regexp.MustCompile(`[a-zA-Z]`).MatchString(password)
	if !hasLetter {
		return errs.NewValidationError(
			msgProvider.
				WithLang(lang).
				GetMessage("validator", "password_must_contain_letter"),
		)
	}

	// Validar que contenga al menos un número
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)
	if !hasNumber {
		return errs.NewValidationError(
			msgProvider.
				WithLang(lang).
				GetMessage("validator", "password_must_contain_number"),
		)
	}

	// Validar que contenga al menos un símbolo especial
	hasSymbol := regexp.MustCompile(`[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?~` + "`" + `]`).MatchString(password)
	if !hasSymbol {
		return errs.NewValidationError(
			msgProvider.
				WithLang(lang).
				GetMessage("validator", "password_must_contain_symbol"),
		)
	}

	return nil
}

