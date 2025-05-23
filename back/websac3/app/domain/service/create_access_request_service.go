package service

import (
	"errors"
	"fmt"
	"time"
	"websac3/app/domain/builder"
	"websac3/app/domain/entity"
	"websac3/app/domain/errs"
	"websac3/app/port/in/dto"
	"websac3/app/port/in/dto/command"
	"websac3/app/port/out/persistence"
)

type CreateAccessRequestService struct {
	Service
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
	return &CreateAccessRequestService{
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
}

func (c *CreateAccessRequestService) validateInputData(validators []dto.Validator) error {
	for _, validator := range validators {
		if err := validator.Validate(); err != nil {
			return err
		}
	}
	return nil
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

func (c *CreateAccessRequestService) buildNewUser(userBuilder builder.UserBuilder, userCommand *command.CreateUserCommand) *entity.User {
	return userBuilder.
		WithEmail(userCommand.Email).
		WithRole(nil).
		Build()
}

func (c *CreateAccessRequestService) buildNewPerson(personBuilder builder.PersonBuilder, personCommand *command.CreatePersonCommand, user *entity.User) *entity.Person {
	return personBuilder.
		WithIdentificationNumber(personCommand.IdentificationNumber).
		WithName(personCommand.Name).
		WithLastname(personCommand.Lastname).
		WithIdentificationTypeID(personCommand.IdentificationTypeID).
		WithIdentificationNumber(personCommand.IdentificationNumber).
		WithHigherEducationInstitutionSnies(personCommand.HigherEducationInstitutionSnies).
		WithJobPosition(personCommand.JobPosition).
		WithUser(user).
		Build()
}

func (c *CreateAccessRequestService) buildAccessRequest(
	accessRequestBuilder builder.AccessRequestBuilder,
	person *entity.Person,
) *entity.AccessRequest {
	return accessRequestBuilder.
		WithPerson(person).
		WithCreatedAt(time.Now()).
		WithUpdatedAt(time.Now()).
		WithStatus(c.statusPending).
		Build()
}

func (c *CreateAccessRequestService) createFirstAccessRequest(
	createAccessRequest *entity.AccessRequest,
	tx persistence.Transaction,
) (err error) {
	fmt.Println("Ingreso a create first")
	if err = c.createUserPort.Create(createAccessRequest.Person.User, tx); err != nil {
		return err
	}
	fmt.Printf("User entity: %+v\n", *createAccessRequest.Person.User)
	if err = c.createPersonPort.Create(createAccessRequest.Person, tx); err != nil {
		return err
	}
	fmt.Printf("Person entity: %+v\n", *createAccessRequest.Person)
	if err = c.createAccessRequestPort.Create(createAccessRequest, tx); err != nil {
		return err
	}
	fmt.Printf("create access request entity: %+v\n", *createAccessRequest)
	fmt.Println("No se rompio")
	return nil
}

func (c *CreateAccessRequestService) createNewAccessRequest(
	newAccessRequest *entity.AccessRequest,
	accessRequestFound *entity.AccessRequest,
	tx persistence.Transaction,
) (err error) {
	if newAccessRequest.Person.User.Email != accessRequestFound.Person.User.Email {
		if err = c.updateUserPort.UpdateById(newAccessRequest.Person.User, accessRequestFound.ID, tx); err != nil {
			return err
		}
	}

	if newAccessRequest.Person.JobPosition != accessRequestFound.Person.JobPosition ||
		newAccessRequest.Person.HigherEducationInstitutionSnies != accessRequestFound.Person.HigherEducationInstitutionSnies ||
		newAccessRequest.Person.Name != accessRequestFound.Person.Name ||
		newAccessRequest.Person.Lastname != accessRequestFound.Person.Lastname {
		if err = c.updatePersonPort.UpdateById(newAccessRequest.Person, accessRequestFound.ID, tx); err != nil {
			return err
		}
	}
	if err = c.createAccessRequestPort.Create(newAccessRequest, tx); err != nil {
		return err
	}
	return nil
}

func (c *CreateAccessRequestService) CreateAccessRequest(
	createAccessRequestCommand command.CreateAccessRequestCommand,
) error {
	if err := c.validateInputData([]dto.Validator{&createAccessRequestCommand}); err != nil {
		return err
	}

	return c.txManager.ExecuteInTransaction(
		func(tx persistence.Transaction) error {
			var err error
			var canRegisterForFirstTime, canRegisterAnother bool

			var accessRequestFound entity.AccessRequest
			var user *entity.User = nil
			var person *entity.Person = nil
			var newAccessRequest *entity.AccessRequest = nil

			var accessRequestBuilder builder.AccessRequestBuilder = builder.NewAccessRequestBuilder()
			var personBuilder builder.PersonBuilder = builder.NewPersonBuilder()
			var userBuilder builder.UserBuilder = builder.NewUserBuilder()

			if c.statusPending == nil {
				if err = c.initStatusPending(tx); err != nil {
					return err
				}
			}
			accessRequestFound, err = c.getAccessRequestPort.GetLastCreatedPersonIdentificationNumber(
				createAccessRequestCommand.Person.IdentificationNumber, tx,
			)
			if err != nil {
				return err
			}

			user = c.buildNewUser(userBuilder, &createAccessRequestCommand.Person.User)
			person = c.buildNewPerson(personBuilder, &createAccessRequestCommand.Person, user)
			newAccessRequest = c.buildAccessRequest(accessRequestBuilder, person)

			if canRegisterForFirstTime = !accessRequestFound.IsRegistered(); canRegisterForFirstTime {
				return c.createFirstAccessRequest(newAccessRequest, tx)
			}

			canRegisterAnother, err = accessRequestFound.CanRegisterAnother()
			if err != nil {
				return err
			}
			if canRegisterAnother {
				return c.createNewAccessRequest(newAccessRequest, &accessRequestFound, tx)
			}

			return err
		},
	)
}
