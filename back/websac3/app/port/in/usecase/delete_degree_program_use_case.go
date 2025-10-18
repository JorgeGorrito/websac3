package usecase

type DeleteDegreeProgramUseCase interface {
	Execute(degreeProgramID uint, userID uint, lang string) error
}






