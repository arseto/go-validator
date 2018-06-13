package rules

import "reflect"

type RequiredValidator struct{}

func (r *RequiredValidator) IsValid(value interface{}) bool {
	zeroVal := reflect.Zero(reflect.TypeOf(value)).Interface()
	return value != zeroVal
}
