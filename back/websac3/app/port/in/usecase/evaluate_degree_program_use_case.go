package usecase

type EvaluateDegreeProgramUseCase interface {
	Execute(degreeProgramID uint, professionalRoleID uint, userID uint, lang string) error
}
