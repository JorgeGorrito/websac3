package usecase

import (
	"websac3/app/domain/entity"
)

type CreateDegreeProgramUseCase interface {
	Execute(degreeProgram entity.DegreeProgram, lang string) error
}
