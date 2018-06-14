package rules

import "testing"

func TestNumericIsValid(t *testing.T) {

	valid := "09299301"
	invalid := "a940.999"

	v := MakeNumericValidator("numericField", valid)

	if !v.IsValid() {
		t.Errorf("Should be valid: %s", valid)
	}

	v = MakeNumericValidator("numericField", invalid)

	if v.IsValid() {
		t.Errorf("Should be invalid: %s", invalid)
	}

	actual := v.Validate()
	if actual == nil {
		t.Error("Should give error message")
	}

	expectedMessage := "numericField must be a valid numeric string"
	if expectedMessage != actual.Error() {
		t.Errorf("Expected message: %s, actual: %s",
			expectedMessage,
			actual.Error(),
		)
	}
}
