package hw09structvalidator

import (
	"strconv"
	"strings"
)

// intValidRule describes what the conditions target int field must satisfying to.
type intValidRule struct {
	Min int
	Max int
	In  []int
}

// validateNumber checks if the value satisfies the conditions.
func validateNumber[T int64 | float64](value T, condition string) error {
	rule := parseNumberCondition(condition)

	if rule.Min > 0 && value < T(rule.Min) {
		return ErrNotGreater
	}
	if rule.Max > 0 && value > T(rule.Max) {
		return ErrNotLesser
	}
	if len(rule.In) > 0 {
		found := false
		for _, v := range rule.In {
			if v == int(value) {
				found = true
				break
			}
		}
		if !found {
			return ErrNotInRange
		}
	}

	return nil
}

// parseNumberCondition parses the condition string and returns the rule to check the field value.
func parseNumberCondition(condition string) (rule intValidRule) {
	if condition == "" {
		return rule
	}

	for _, pattern := range strings.Fields(condition) {
		if pattern == "" {
			continue
		}

		parts := strings.SplitN(pattern, ":", 2)
		key, value := parts[0], parts[1]
		if value == "" {
			continue
		}

		switch key {
		case "min":
			min, _ := strconv.Atoi(value)
			rule.Min = min
		case "max":
			max, _ := strconv.Atoi(value)
			rule.Max = max
		case "in":
			for _, v := range strings.Split(value, ",") {
				x, err := strconv.Atoi(v)
				if err == nil {
					rule.In = append(rule.In, x)
				}
			}
		}
	}

	return rule
}
