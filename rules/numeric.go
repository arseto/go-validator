package rules

import "strconv"

type NumericValidator struct{}

func (r *NumericValidator) IsValid(value interface{}) bool {
	str, ok := value.(string)
	if !ok {
		return false
	}
	_, err := strconv.ParseInt(str, 10, 64)
	return err == nil
}
