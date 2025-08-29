package service

import (
	"websac3/app/domain/entity"
	"websac3/app/port/out/message"
	"websac3/app/port/out/persistence"
	"websac3/app/port/out/persistence/db"
)

type CreateDegreeProgramService struct {
	createDegreeProgramPort persistence.CreateDegreeProgramPort
	persistenceManager      db.Manager
	msgProvider             message.Provider
}

func NewCreateDegreeProgramService(
	createDegreeProgramPort persistence.CreateDegreeProgramPort,
	persistenceManager db.Manager,
	msgProvider message.Provider,
) *CreateDegreeProgramService {
	return &CreateDegreeProgramService{
		createDegreeProgramPort: createDegreeProgramPort,
		persistenceManager:      persistenceManager,
		msgProvider:             msgProvider,
	}
}

func (s *CreateDegreeProgramService) Execute(
	degreeProgram entity.DegreeProgram,
	lang string,
) error {
	return s.persistenceManager.ExecuteInTransaction(
		func(ctx db.Context) error {
			if err := s.createDegreeProgramPort.Create(&degreeProgram, ctx); err != nil {
				return err
			}
			return nil
		},
	)
}
