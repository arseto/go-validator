package main

type Validator interface {
	IsValid(value interface{}) bool
}
