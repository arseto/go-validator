package validator

import (
	"reflect"
	"testing"
)

type TestObject struct {
	ID      int64  `validate:"required"`
	Name    string `validate:"required"`
	Email   string `validate:"required|email"`
	Phone   string `validate:"numeric"`
	Remarks string
}

func TestReadValidationRules(t *testing.T) {
	obj := &TestObject{
		10,
		"Test Name",
		"admin@123.com",
		"99188301",
		"",
	}

	rules := ReadValidationRules(obj)

	idRule := rules["ID"]
	expectedIDValue := obj.ID
	expectedIDRule := []string{"required"}

	if expectedIDValue != idRule.Value {
		t.Errorf("Expected value: %d, actual: %d", expectedIDValue, idRule.Value)
	}
	if !reflect.DeepEqual(expectedIDRule, idRule.Rules) {
		t.Errorf("Expected: %s, actual: %s", expectedIDRule, idRule.Rules)
	}

	nameRule := rules["Name"]
	expectedNameValue := obj.Name
	expectedNameRule := []string{"required"}

	if expectedNameValue != nameRule.Value {
		t.Errorf("Expected value: %s, actual: %s", expectedNameValue, nameRule.Value)
	}
	if !reflect.DeepEqual(expectedNameRule, nameRule.Rules) {
		t.Errorf("Expected: %s, actual: %s", expectedNameRule, nameRule.Rules)
	}

	emailRule := rules["Email"]
	expectedEmailValue := obj.Email
	expectedEmailRule := []string{"required", "email"}

	if expectedEmailValue != emailRule.Value {
		t.Errorf("Expected value: %s, actual: %s", expectedEmailValue, emailRule.Value)
	}
	if !reflect.DeepEqual(expectedEmailRule, emailRule.Rules) {
		t.Errorf("Expected: %s, actual: %s", expectedEmailRule, emailRule.Rules)
	}

	phoneRule := rules["Phone"]
	expectedPhoneValue := obj.Phone
	expectedPhoneRule := []string{"numeric"}

	if expectedPhoneValue != phoneRule.Value {
		t.Errorf("Expected value: %s, actual: %s", expectedPhoneValue, phoneRule.Value)
	}
	if !reflect.DeepEqual(expectedPhoneRule, phoneRule.Rules) {
		t.Errorf("Expected: %s, actual: %s", expectedPhoneRule, phoneRule.Rules)
	}

	remarksRule := rules["Remarks"]
	expectedRemarksValue := obj.Remarks
	expectedRemarksRule := []string{}

	if expectedRemarksValue != remarksRule.Value {
		t.Errorf("Expected value: %s, actual: %s", expectedRemarksValue, remarksRule.Value)
	}
	if !reflect.DeepEqual(expectedRemarksRule, remarksRule.Rules) {
		t.Errorf("Expected: %s, actual: %s", expectedRemarksRule, remarksRule.Rules)
	}
}
