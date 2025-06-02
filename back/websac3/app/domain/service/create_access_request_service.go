package service

import (
	"errors"
	"websac3/app/domain/entity"
	"websac3/app/domain/errs"
	"websac3/app/port/out/persistence"
)

type CreateAccessRequestService struct {
	createAccessRequestPort persistence.CreateAccessRequestPort
	createPersonPort        persistence.CreatePersonPort
	createUserPort          persistence.CreateUserPort
	getUserPort             persistence.GetUserPort
	updateUserPort          persistence.UpdateUserPort
	updatePersonPort        persistence.UpdatePersonPort
	getPersonPort           persistence.GetPersonPort
	getAccessRequestPort    persistence.GetAccessRequestPort
	getStatusPort           persistence.GetStatusPort
	statusPending           *entity.Status
	txManager               persistence.TransactionManager
}

func (c *CreateAccessRequestService) initStatusPending(tx persistence.Transaction) error {
	var err error
	var statusPendingFound entity.Status

	if statusPendingFound, err = c.getStatusPort.GetByName("pending", tx); err != nil {
		if errors.Is(err, errs.NotFoundError) {
			err = errors.Join(err, errors.New("status pending not found, please verify the essentials data"))
		}
		return err
	}
	c.statusPending = &statusPendingFound
	return nil
}

func NewCreateAccessRequestService(
	createAccessRequestPort persistence.CreateAccessRequestPort,
	createPersonPort persistence.CreatePersonPort,
	createUserPort persistence.CreateUserPort,
	updateUserPort persistence.UpdateUserPort,
	updatePersonPort persistence.UpdatePersonPort,
	getUserPort persistence.GetUserPort,
	getPersonPort persistence.GetPersonPort,
	getAccessRequestPort persistence.GetAccessRequestPort,
	getStatusPort persistence.GetStatusPort,
	txManager persistence.TransactionManager,
) *CreateAccessRequestService {
	c := &CreateAccessRequestService{
		createAccessRequestPort: createAccessRequestPort,
		createPersonPort:        createPersonPort,
		createUserPort:          createUserPort,
		getUserPort:             getUserPort,
		getPersonPort:           getPersonPort,
		getAccessRequestPort:    getAccessRequestPort,
		getStatusPort:           getStatusPort,
		updateUserPort:          updateUserPort,
		updatePersonPort:        updatePersonPort,
		statusPending:           nil,
		txManager:               txManager,
	}

	if err := txManager.ExecuteInTransaction(func(tx persistence.Transaction) error {
		c.initStatusPending(tx)
		return nil
	}); err != nil {
		panic(err)
	}

	return c
}

func (c *CreateAccessRequestService) CreateAccessRequest(requestToCreate *entity.AccessRequest, tx persistence.Transaction) error {
	requestToCreate.Status = c.statusPending
	if err := c.createUserPort.Create(requestToCreate.Person.User, tx); err != nil {
		return err
	}
	if err := c.createPersonPort.Create(requestToCreate.Person, tx); err != nil {
		return err
	}
	if err := c.createAccessRequestPort.Create(requestToCreate, tx); err != nil {
		return err
	}
	return nil
}

func (c *CreateAccessRequestService) Execute(
	requestToCreate entity.AccessRequest,
) error {
	return c.txManager.ExecuteInTransaction(
		func(tx persistence.Transaction) error {
			var err error
			var requestFound entity.AccessRequest

			requestFound, err = c.getAccessRequestPort.GetLastCreatedPersonIdentificationNumber(requestToCreate.Person.IdentificationNumber, tx)
			if err != nil {
				return err
			}

			if requestFound.CanRegisterForFirstTime() {
				return c.CreateAccessRequest(&requestToCreate, tx)
			}

			return err
		},
	)
}
