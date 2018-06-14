package main

import (
	"reflect"
	"testing"
)

func TestMessageBag(t *testing.T) {
	mb := &messageBag{make([]string, 0)}

	mb.AddMessage("message 1")

	expected := []string{"message 1"}

	if !reflect.DeepEqual(expected, mb.GetMessages()) {
		t.Error("Messages should be equal")
	}

	mb.AddMessage("message 2")
	expected = []string{"message 1", "message 2"}

	if !reflect.DeepEqual(expected, mb.GetMessages()) {
		t.Error("Messages should be equal")
	}
}
