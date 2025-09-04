package usecase

type CreateUserFromTokenUseCase interface {
	Execute(token string, password string, lang string) error
}
