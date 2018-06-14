package main

import "testing"

type mockMessage struct{}

func (m *mockMessage) AddMessage(message string) {}

func (m *mockMessage) GetMessages() []string { return nil }

func (m *mockMessage) Empty() bool { return false }

func TestValidationError(t *testing.T) {
	messages := &mockMessage{}
	e := &validationError{
		"test error",
		messages,
	}

	if e.Error() != "test error" {
		t.Error("Mismatch error message")
	}

	if e.Messages() != messages {
		t.Error("Mismatch error messages")
	}

}
