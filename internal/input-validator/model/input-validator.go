package model

import (
	"strings"
	"fmt"
)

type Validator interface {
	Validate(input string) bool
}

type EmailValidator struct {}
type PasswordValidator struct {}

func (e *EmailValidator) Validate(input string) bool {
	if strings.Contains(input,"@") {
		return true
	} else {
		return false
	}
}

func (e *PasswordValidator) Validate(input string) bool {
	if len(input) >= 8 {
		return true
	} else {
		return false
	}
}

type RegistrationService struct {
	validators []Validator
}

func NewRegistrationService(validators []Validator) (*RegistrationService){
	return &RegistrationService{validators: validators}
}

func (r *RegistrationService) Register(input string){
	for _,validator := range r.validators {
		if !validator.Validate(input) {
			fmt.Println("FAILED")
			return
		}
	}
    fmt.Println("PASSED")
}

