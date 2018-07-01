package validator

import (
	"fmt"
	"regexp"
)

type emailValidator struct {
	fieldName string
	value     interface{}
}

func MakeEmailValidator(fieldName string, value interface{}) *emailValidator {
	return &emailValidator{fieldName, value}
}

func (r *emailValidator) IsValid() bool {
	re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	email := r.value.(string)

	return re.MatchString(email)
}

func (r *emailValidator) Validate() error {
	if !r.IsValid() {
		return fmt.Errorf("%s must be a valid email address", r.fieldName)
	}

	return nil
}
