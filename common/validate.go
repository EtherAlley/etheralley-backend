package common

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateStruct(s interface{}) error {
	if s == nil {
		return errors.New("validation failed on nil struct")
	}

	err := validate.Struct(s)

	if err != nil {
		return buildValidationError(err)
	}

	return nil
}

func ValidateField(f interface{}, tag string) error {
	err := validate.Var(f, tag)

	if err != nil {
		return buildValidationError(err)
	}

	return nil
}

func buildValidationError(err error) error {
	msg := ""
	for i, err := range err.(validator.ValidationErrors) {
		if i != 0 {
			msg += " and "
		}
		msg = msg + fmt.Sprintf("validation failed on field %v", err.Field())
	}
	return errors.New(msg)
}
