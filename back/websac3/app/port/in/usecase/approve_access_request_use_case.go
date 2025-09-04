package usecase

type ApproveAccessRequestUseCase interface {
	Execute(accessRequestToApproveID uint, roleID uint, lang string) error
}
