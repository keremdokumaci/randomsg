package validator

import (
	"github.com/go-playground/validator/v10"
)

type customValidator struct {
	validator *validator.Validate
}

var CustomValidator customValidator

func NewValidator() {
	CustomValidator.validator = validator.New()
}

func (c customValidator) ValidateStruct(model interface{}) (bool, []string) {
	errFields := []string{}
	err := c.validator.Struct(model)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errFields = append(errFields, err.Field())
		}
		return true, errFields
	}
	return false, []string{}
}
