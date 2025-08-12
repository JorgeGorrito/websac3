package usecase

type RejectAccessRequestUseCase interface {
	Execute(accessRequestToRejectID uint, lang string) error
}
