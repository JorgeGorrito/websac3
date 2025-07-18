package usecase

type ValidateEmailUseCase interface {
	Execute(validationToken string, lang string) error
}
