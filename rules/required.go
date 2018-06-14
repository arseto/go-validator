package rules

import (
	"fmt"
	"reflect"
)

type requiredValidator struct {
	fieldName string
	value     interface{}
}

func MakeRequiredValidator(fieldName string, value interface{}) *requiredValidator {
	return &requiredValidator{
		fieldName,
		value,
	}
}

func (r *requiredValidator) IsValid() bool {
	zeroVal := reflect.Zero(reflect.TypeOf(r.value)).Interface()
	return r.value != zeroVal
}

func (r *requiredValidator) Validate() error {
	if !r.IsValid() {
		return fmt.Errorf("%s is required", r.fieldName)
	}

	return nil
}
