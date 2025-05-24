package service

import (
	"errors"
	"websac3/app/port/in/dto/command"
)

type Service struct{}

func (*Service) validateInputData(dtos []command.Validator) error {
	var errorList error
	for _, dto := range dtos {
		if err := dto.Validate(); err != nil {
			errorList = errors.Join(errorList, err)
		}
	}
	return errorList
}
