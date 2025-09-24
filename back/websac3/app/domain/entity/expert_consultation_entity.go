package entity

import (
	"time"
	"websac3/app/domain/constants"
)

type ExpertConsultation struct {
	ID uint

	// Usuario que solicita la asesoría
	RequesterID uint
	Requester   *User

	// Programa de grado asociado
	DegreeProgramID uint
	DegreeProgram   *DegreeProgram

	// Reporte asociado
	ReportID uint
	Report   *Report

	// Mensaje de solicitud (opcional)
	RequestMessage *string

	// Respuesta del experto
	ExpertResponse *string

	// Usuario experto que proporciona la respuesta
	ExpertID *uint
	Expert   *User

	// Estado de la consulta
	StatusID uint
	Status   *ExpertConsultationStatus

	// Fechas de auditoría
	CreatedAt  time.Time
	UpdatedAt  *time.Time
	AnsweredAt *time.Time
	ClosedAt   *time.Time
}

func NewExpertConsultation(
	requesterID uint,
	degreeProgramID uint,
	reportID uint,
	requestMessage *string,
) *ExpertConsultation {
	return &ExpertConsultation{
		RequesterID:     requesterID,
		DegreeProgramID: degreeProgramID,
		ReportID:        reportID,
		RequestMessage:  requestMessage,
		StatusID:        1, // ID for "pending" status
		CreatedAt:       time.Now(),
	}
}

func (ec *ExpertConsultation) IsRegistered() bool {
	return ec.ID != 0
}

func (ec *ExpertConsultation) IsPending() bool {
	return ec.Status != nil && ec.Status.Name == constants.ExpertConsultationStatusPending
}

func (ec *ExpertConsultation) IsRejected() bool {
	return ec.Status != nil && ec.Status.Name == constants.ExpertConsultationStatusRejected
}

func (ec *ExpertConsultation) IsAccepted() bool {
	return ec.Status != nil && ec.Status.Name == constants.ExpertConsultationStatusAccepted
}

func (ec *ExpertConsultation) AcceptConsultation(expertID uint, response string) {
	now := time.Now()
	ec.ExpertID = &expertID
	ec.ExpertResponse = &response
	ec.StatusID = 3 // ID for "accepted" status
	ec.AnsweredAt = &now
	ec.UpdatedAt = &now
}

func (ec *ExpertConsultation) RejectConsultation(expertID uint, rejectionReason *string) {
	now := time.Now()
	ec.ExpertID = &expertID
	if rejectionReason != nil {
		ec.ExpertResponse = rejectionReason
	}
	ec.StatusID = 2 // ID for "rejected" status
	ec.AnsweredAt = &now
	ec.UpdatedAt = &now
}

func (ec *ExpertConsultation) UpdateRequestMessage(message string) {
	now := time.Now()
	ec.RequestMessage = &message
	ec.UpdatedAt = &now
}
