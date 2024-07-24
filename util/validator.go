package util

import (
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

func Validate(v *validator.Validate, log *logrus.Logger, u interface{}) (bool, map[string]string) {
	if err := v.Struct(&u); err != nil {
		log.WithError(err).Error("Validation failed")
		validationErrors := err.(validator.ValidationErrors)
		errorMessages := make(map[string]string)
		for _, fieldError := range validationErrors {
			errorMessages[fieldError.Field()] = fieldError.Error()
		}
		return false, errorMessages
	}
	return true, nil
}
