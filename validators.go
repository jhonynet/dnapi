package main

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	mutant "github.com/jhonynet/dna"
)

func squareMatrix(fl validator.FieldLevel) bool {
	return mutant.IsSquareMatrix(fl.Field().Interface().([]string))
}

func validateDnaCharacters(fl validator.FieldLevel) bool {
	return !mutant.HasInvalidCharacters(fl.Field().Interface().([]string))
}

// Register custom validations in validator engine
func registerCustomValidators() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("squareMatrix", squareMatrix)
		_ = v.RegisterValidation("validDnaCharacters", validateDnaCharacters)
	}
}
