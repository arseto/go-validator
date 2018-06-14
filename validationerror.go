package main

type ValidationError error

type validationError struct {
	message  string
	messages MessageBag
}

func (e *validationError) Error() string {
	return e.message
}

func (e *validationError) Messages() MessageBag {
	return e.messages
}
