package rules

import "testing"

func TestEmailIsValid(t *testing.T) {
	v := &EmailValidator{}

	valid := "hayoo@yahoo.com"
	invalid := "hey@yahoo"

	if !v.IsValid(valid) {
		t.Errorf("Should be valid: %s", valid)
	}

	if v.IsValid(invalid) {
		t.Errorf("Should be invalid: %s", invalid)
	}
}
