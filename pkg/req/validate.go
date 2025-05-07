package req

import (
	"github.com/go-playground/validator/v10"
)

func IsValid[T any](body T) error {
	validate := validator.New()
	err := validate.Struct(body)
	return err
}

