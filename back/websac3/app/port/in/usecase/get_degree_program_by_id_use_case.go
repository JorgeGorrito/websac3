package usecase

import (
	"websac3/adapter/in/web/response"
)

// GetDegreeProgramByIDUseCase define el caso de uso para obtener el detalle de un programa de grado por ID
type GetDegreeProgramByIDUseCase interface {
	Execute(degreeProgramID uint, lang string) (response.ApiResponse[response.GetDegreeProgramByIDResponse], error)
}
