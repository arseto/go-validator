package rules

import "testing"

func TestRequiredIsValid(t *testing.T) {

	valid := "AnyValue"
	invalid := ""

	v := MakeRequiredValidator("myField", valid)

	if !v.IsValid() {
		t.Errorf("Should be valid: %s", valid)
	}

	v = MakeRequiredValidator("myField", invalid)

	if v.IsValid() {
		t.Errorf("Should be invalid: %s", invalid)
	}

	actualMessage := v.Validate()
	if actualMessage == nil {
		t.Error("Should be invalid")
	}

	expectedMessage := "myField is required"
	if expectedMessage != actualMessage.Error() {
		t.Errorf("Expected message: %s, actual: %s",
			expectedMessage,
			actualMessage.Error(),
		)
	}

	validInt := 19
	invalidInt := 0

	v = &requiredValidator{"intField", validInt}

	if !v.IsValid() {
		t.Errorf("Should be valid: %d", validInt)
	}

	v = &requiredValidator{"intField", invalidInt}

	if v.IsValid() {
		t.Errorf("Should be invalid: %d", invalidInt)
	}
}
