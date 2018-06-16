package validator

type ValidationError interface {
	error
	Messages() MessageBag
}

type validationError struct {
	err      string
	messages MessageBag
}

func NewValidationErrorShort(err string) ValidationError {
	return &validationError{
		err,
		&messageBag{make([]string, 0)},
	}
}

func NewValidationError(err string, messages MessageBag) ValidationError {
	return &validationError{
		err,
		messages,
	}
}

func (e *validationError) Error() string {
	return e.err
}

func (e *validationError) Messages() MessageBag {
	return e.messages
}
