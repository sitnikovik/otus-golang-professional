package hw09structvalidator

import (
	"fmt"
	"reflect"
	"strings"
)

type ValidationError struct {
	Field string
	Err   error
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	sb := strings.Builder{}
	for _, err := range v {
		sb.WriteString(fmt.Sprintf("field: %s, error: %v\n", err.Field, err.Err))
	}

	return sb.String()
}

func Validate(v interface{}) error {
	errors := ValidationErrors{}

	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	reftype := reflect.TypeOf(v)
	for i := 0; i < val.NumField(); i++ {
		field := reftype.Field(i)

		validTag := val.Type().Field(i).Tag.Get("validate")
		if validTag == "" {
			continue
		}

		fieldName := field.Name
		fieldValue := val.Field(i)
		switch field.Type.Kind() {
		case reflect.String:
			rules := parseStringValidRules(validTag)
			s := fieldValue.String()
			if err := validateString(s, rules); err != nil {
				errors = append(errors, ValidationError{Field: fieldName, Err: err})
			}
		case reflect.Int:
			rules := parseIntValidRules(validTag)
			num := int(fieldValue.Int())
			if err := validateInt(num, rules); err != nil {
				errors = append(errors, ValidationError{Field: fieldName, Err: err})
			}
		}
	}

	if len(errors) > 0 {
		return errors
	}
	return nil
}

func validateString(value string, rules []stringValidRule) error {
	for _, r := range rules {
		if err := r.Validate(value); err != nil {
			return err
		}
	}

	return nil
}

func validateInt(value int, rules []intValidRule) error {
	for _, r := range rules {
		if err := r.Validate(value); err != nil {
			return err
		}
	}

	return nil
}
