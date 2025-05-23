package service

import (
	"errors"
	"websac3/app/port/in/dto"
)

type Service struct{}

func (*Service) validateInputData(dtos []dto.Validator) error {
	var errorList error
	for _, dto := range dtos {
		if err := dto.Validate(); err != nil {
			errorList = errors.Join(errorList, err)
		}
	}
	return errorList
}
