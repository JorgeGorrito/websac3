package usecase

type ApproveAccessRequestUseCase interface {
	Execute(accessRequestToApproveID uint, lang string) error
}
