package usecase

type RejectExpertConsultationUseCase interface {
	Execute(consultationID uint, expertResponse string, expertID uint, lang string) error
}
