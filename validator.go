package validator

import (
	"github.com/arseto/validator/reflection"
	"github.com/arseto/validator/rules"
)

type FieldValidator interface {
	IsValid() bool
	Validate() error
}

type StructValidator interface {
	Fail() bool
	Success() bool
	Messages() MessageBag
	Validate() ValidationError
}

type structValidator struct {
	messages MessageBag
}

func MakeStructValidator(obj interface{}) StructValidator {
	validationRules := reflection.ReadValidationRules(obj)

	mb := &messageBag{make([]string, 0)}

	for key, rule := range validationRules {
		for _, fieldRule := range rule.Rules {
			var validator FieldValidator

			switch fieldRule {
			case "email":
				validator = rules.MakeEmailValidator(key, rule.Value)
			case "required":
				validator = rules.MakeRequiredValidator(key, rule.Value)
			case "numeric":
				validator = rules.MakeNumericValidator(key, rule.Value)
			default:
				validator = nil
			}

			if validator != nil {
				err := validator.Validate()
				if err != nil {
					mb.AddMessage(err.Error())
				}
			}
		}
	}

	return &structValidator{mb}
}

func (v *structValidator) Fail() bool {
	return !v.messages.Empty()
}

func (v *structValidator) Success() bool {
	return !v.Fail()
}

func (v *structValidator) Messages() MessageBag {
	return v.messages
}

func (v *structValidator) Validate() ValidationError {
	if v.Fail() {
		return NewValidationError(
			"Validation Failed",
			v.messages,
		)
	}
	return nil
}
