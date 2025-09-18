package usecase

type ActivateUserUseCase interface {
	Execute(userID uint, lang string) error
}
