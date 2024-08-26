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
	errs := make(ValidationErrors, 0)

	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)
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

func validateField(field reflect.Value, tag string) error {
	conditions := strings.Split(tag, "|")
	for _, condition := range conditions {
		if err := checkCondition(field, condition); err != nil {
			return err
		}
	}

	return nil
}

func checkCondition(field reflect.Value, condition string) error {
	switch field.Kind() {
	case reflect.String:
		return validateString(field.String(), condition)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return validateNumber(field.Int(), condition)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return validateNumber(int64(field.Uint()), condition)
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
