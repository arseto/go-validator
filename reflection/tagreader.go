package reflection

import (
	"reflect"
	"strings"
)

type ValidationRule struct {
	Value interface{}
	Rules []string
}

func ReadValidationRules(obj interface{}) map[string]*ValidationRule {

	rules := make(map[string]*ValidationRule)

	val := reflect.ValueOf(obj).Elem()

	for i := 0; i < val.NumField(); i++ {

		typeField := val.Type().Field(i)
		valueField := val.Field(i)

		name := typeField.Name
		tag := typeField.Tag
		tagStr := tag.Get("validate")
		validateTags := make([]string, 0)
		if len(tagStr) > 0 {
			validateTags = strings.Split(tagStr, ",")
		}
		rule := &ValidationRule{
			valueField.Interface(),
			validateTags,
		}

		rules[name] = rule
	}

	return rules
}
