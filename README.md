# Overview

Simple helper for struct validation using attributes.

# Usage

```go

type testObject struct {
	ID      int64  `validate:"required"`
	Name    string `validate:"required"`
	Email   string `validate:"required,email"`
	Phone   string `validate:"numeric"`
	Remarks string
}

func ProcessObject(obj *testObject) {
    sv := validator.MakeStructValidator(obj)

    //check if validation failed
    fail := sv.Fail()

    //check if validation success
    //opposite of Fail()
    success := sv.Success()

    //returns error message with detailed field error message
    err := sv.Validate()

    //global error message
    errorMessage := err.Error()

    //field error messages
    fieldErrorMessages := err.Messages()

    //do whatever with error message
    //for http request, return 422 error message
    //etc

    ...
}

```

# Contributing

Please feel free to submit issue or pull request :).
