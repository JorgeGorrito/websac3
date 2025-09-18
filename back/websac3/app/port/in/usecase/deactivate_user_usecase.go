package usecase

type DeactivateUserUseCase interface {
	Execute(userID uint, lang string) error
}
