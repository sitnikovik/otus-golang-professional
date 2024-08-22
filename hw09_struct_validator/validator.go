package hw09structvalidator

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type ValidationError struct {
	Field string
	Err   error
}

type ValidationErrors []ValidationError

type stringValidRule struct {
	Len    int
	Regexp string
	In     []string
}

type intValidRule struct {
	Min int
	Max int
	In  [2]int
}

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

		if !isTypeOk(field.Type) {
			continue
		}

		validTag := val.Type().Field(i).Tag.Get("validate")
		if validTag == "" {
			continue
		}

		fieldValue := val.Field(i)
		name := field.Name
		switch field.Type.Kind() {
		case reflect.String:
			rules := parseStringValidRules(validTag)
			s := fieldValue.String()
			if err := validateString(s, rules); err != nil {
				errors = append(errors, ValidationError{name, err})
			}
		case reflect.Int:
			rules := parseIntValidRules(validTag)
			num := int(fieldValue.Int())
			if err := validateInt(num, rules); err != nil {
				errors = append(errors, ValidationError{name, err})
			}
		default:
			continue
		}

	}

	return errors
}

func isTypeOk(t reflect.Type) bool {
	return isString(t) || isInt(t)
}

func isString(t reflect.Type) bool {
	return t.Kind() == reflect.String
}

func isInt(t reflect.Type) bool {
	return t.Kind() == reflect.Int
}

func parseStringValidRules(tag string) []stringValidRule {
	rules := []stringValidRule{}

	for _, cond := range strings.Split(tag, "|") {
		for _, pattern := range strings.Fields(cond) {
			if pattern == "" {
				continue
			}

			parts := strings.SplitN(pattern, ":", 2)
			v := parts[1]
			if v == "" {
				continue
			}

			rule := stringValidRule{}
			switch parts[0] {
			case "len":
				len, _ := strconv.Atoi(v)
				rule.Len = len
			case "regexp":
				rule.Regexp = v
			case "in":
				rule.In = strings.Split(v, ",")
			}
			rules = append(rules, rule)
		}
	}

	return rules
}

func parseIntValidRules(tag string) []intValidRule {
	rules := []intValidRule{}

	for _, cond := range strings.Split(tag, "|") {
		for _, pattern := range strings.Fields(cond) {
			if pattern == "" {
				continue
			}

			parts := strings.SplitN(pattern, ":", 2)
			v := parts[1]
			if v == "" {
				continue
			}

			rule := intValidRule{}
			switch parts[0] {
			case "min":
				min, _ := strconv.Atoi(v)
				rule.Min = min
			case "max":
				max, _ := strconv.Atoi(v)
				rule.Max = max
			case "in":
				in := strings.Split(v, ",")
				if len(in) != 2 {
					continue
				}
				min, _ := strconv.Atoi(in[0])
				max, _ := strconv.Atoi(in[1])
				rule.In = [2]int{min, max}
			}
			rules = append(rules, rule)
		}

	}

	return rules
}

func validateString(value string, rules []stringValidRule) error {
	for _, rule := range rules {
		if rule.Len > 0 && len(value) != rule.Len {
			return fmt.Errorf("string length must be %d", rule.Len)
		}
		if rule.Regexp != "" && !regexp.MustCompile(rule.Regexp).MatchString(value) {
			return fmt.Errorf("string must match the pattern %s", rule.Regexp)
		}
		if len(rule.In) > 0 {
			found := false
			for _, v := range rule.In {
				if v == value {
					found = true
					break
				}
			}
			if !found {
				return fmt.Errorf("string must be one of %v", rule.In)
			}
		}
	}

	return nil
}

func validateInt(value int, rules []intValidRule) error {
	for _, rule := range rules {
		if rule.Min > 0 && value < rule.Min {
			return fmt.Errorf("value must be greater than %d", rule.Min)
		}
		if rule.Max > 0 && value > rule.Max {
			return fmt.Errorf("value must be less than %d", rule.Max)
		}
		if rule.In != [2]int{} && (value < rule.In[0] || value > rule.In[1]) {
			return fmt.Errorf("value must be in range %d-%d", rule.In[0], rule.In[1])
		}
	}

	return nil
}
