package validator

import (
	"fmt"
	"strconv"
)

type numericValidator struct {
	fieldName string
	value     interface{}
}

func MakeNumericValidator(fieldName string, value interface{}) *numericValidator {
	return &numericValidator{fieldName, value}
}

func (r *numericValidator) IsValid() bool {
	str, ok := r.value.(string)
	if !ok {
		return false
	}
	_, err := strconv.ParseInt(str, 10, 64)
	return err == nil
}

func (r *numericValidator) Validate() error {
	if !r.IsValid() {
		return fmt.Errorf("%s must be a valid numeric string", r.fieldName)
	}

	return nil
}
