package utils

import "github.com/go-playground/validator/v10"

func ValidationFormatter(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}
