package usecase

type AcceptExpertConsultationUseCase interface {
	Execute(consultationID uint, expertResponse string, expertID uint, lang string) error
}
