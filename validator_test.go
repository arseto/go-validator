package validator

import "testing"

type testObject struct {
	ID      int64  `validate:"required"`
	Name    string `validate:"required"`
	Email   string `validate:"required,email"`
	Phone   string `validate:"numeric"`
	Remarks string
}

func TestSuccessValidation(t *testing.T) {
	obj := &testObject{
		10,
		"Test Name",
		"admin@123.com",
		"99188301",
		"",
	}

	sv := MakeStructValidator(obj)

	if !sv.Success() {
		t.Error("Expected success validation")
	}

	err := sv.Validate()

	if err != nil {
		t.Error("Not expecting error")
	}
}

func TestFailedValidation(t *testing.T) {
	obj := &testObject{
		0,
		"Default Name",
		"waton@email",
		"oo9934",
		"Notes here",
	}

	sv := MakeStructValidator(obj)

	if sv.Success() {
		t.Error("Expected failed validation")
	}

	messages := sv.Messages()

	if len(messages.GetMessages()) != 3 {
		t.Error("Expected exactly 3 errors")
	}

	err := sv.Validate()

	if err == nil {
		t.Error("Expected error")
	}
}
