package rules

import "testing"

func TestNumericIsValid(t *testing.T) {
	v := &NumericValidator{}

	valid := "09299301"
	invalid := "a940.999"

	if !v.IsValid(valid) {
		t.Errorf("Should be valid: %s", valid)
	}

	if v.IsValid(invalid) {
		t.Errorf("Should be invalid: %s", invalid)
	}
}
