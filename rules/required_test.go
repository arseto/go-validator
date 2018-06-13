package rules

import "testing"

func TestRequiredIsValid(t *testing.T) {
	v := &RequiredValidator{}

	valid := "AnyValue"
	invalid := ""

	if !v.IsValid(valid) {
		t.Errorf("Should be valid: %s", valid)
	}

	if v.IsValid(invalid) {
		t.Errorf("Should be invalid: %s", invalid)
	}

	validInt := 19
	invalidInt := 0

	if !v.IsValid(validInt) {
		t.Errorf("Should be valid: %d", validInt)
	}

	if v.IsValid(invalidInt) {
		t.Errorf("Should be invalid: %d", invalidInt)
	}
}
