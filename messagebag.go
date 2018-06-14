package main

type MessageBag interface {
	AddMessage(message string)
	GetMessages() []string
	Empty() bool
}

type messageBag struct {
	messages []string
}

func (mb *messageBag) AddMessage(message string) {
	mb.messages = append(mb.messages, message)
}

func (mb *messageBag) GetMessages() []string {
	return mb.messages
}

func (mb *messageBag) Empty() bool {
	return len(mb.messages) == 0
}
