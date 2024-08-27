package hw09structvalidator

import (
	"fmt"
	"reflect"
	"strings"
)

// ValidationError describes a validation error.
type ValidationError struct {
	Field string
	Err   error
}

// ValidationErrors is a list of validation errors.
type ValidationErrors []ValidationError

// Error returns the error message.
func (v ValidationErrors) Error() string {
	sb := strings.Builder{}
	for _, err := range v {
		sb.WriteString(fmt.Sprintf("field: %s, error: %v\n", err.Field, err.Err))
	}

	return sb.String()
}

// Validate validates the struct fields according to its tags.
func Validate(v interface{}) error {
	errs := make(ValidationErrors, 0)

	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := val.Type().Field(i)
		tag := fieldType.Tag.Get("validate")

		if tag == "" {
			continue
		}

		if err := validateField(field, tag); err != nil {
			errs = append(errs, ValidationError{
				Field: fieldType.Name,
				Err:   err,
			})
		}
	}

	if len(errs) == 0 {
		return nil
	}
	return errs
}

// validateField validates the field according to the tag.
func validateField(field reflect.Value, tag string) error {
	conditions := strings.Split(tag, "|")
	for _, condition := range conditions {
		if err := checkCondition(field, condition); err != nil {
			return err
		}
	}

	return nil
}

// checkCondition checks if the field satisfies the condition.
func checkCondition(field reflect.Value, condition string) error {
	switch field.Kind() {
	case reflect.String:
		return validateString(field.String(), condition)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return validateNumber(int(field.Int()), condition)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return validateNumber(uint(field.Uint()), condition)
	case reflect.Float32, reflect.Float64:
		return validateNumber(field.Float(), condition)
	case reflect.Slice:
		for i := 0; i < field.Len(); i++ {
			elem := field.Index(i)
			if err := checkCondition(elem, condition); err != nil {
				return fmt.Errorf("elem %d: %v", i, err)
			}
		}
	}

	return nil
}
