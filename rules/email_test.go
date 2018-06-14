package rules

import "testing"

func TestEmailIsValid(t *testing.T) {

	valid := "hayoo@yahoo.com"
	invalid := "hey@yahoo"

	v := MakeEmailValidator("emailField", valid)

	if !v.IsValid() {
		t.Errorf("Should be valid: %s", valid)
	}

	v = MakeEmailValidator("emailField", invalid)

	if v.IsValid() {
		t.Errorf("Should be invalid: %s", invalid)
	}

	actual := v.Validate()
	if actual == nil {
		t.Error("Should give error message")
	}

	expectedMessage := "emailField must be a valid email address"
	if expectedMessage != actual.Error() {
		t.Errorf("Expected message: %s, actual: %s",
			expectedMessage,
			actual.Error(),
		)
	}
}
