package rules

import "regexp"

type EmailValidator struct{}

func (r *EmailValidator) IsValid(value interface{}) bool {
	re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	email := value.(string)

	return re.MatchString(email)
}
