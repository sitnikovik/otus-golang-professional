package hw09structvalidator

import (
	"strconv"
	"strings"
)

// intValidRule describes what the conditions target int field must satisfying to.
type intValidRule struct {
	Min int
	Max int
	In  [2]int
}

func validateNumber[T int64 | float64](value T, condition string) error {
	rule := parseNumberCondition(condition)

	if rule.Min > 0 && value < T(rule.Min) {
		return ErrNotGreater
	}
	if rule.Max > 0 && value > T(rule.Max) {
		return ErrNotLesser
	}
	if rule.In[0] > 0 && rule.In[1] > 0 && (value < T(rule.In[0]) || value > T(rule.In[1])) {
		return ErrNotInRange
	}

	return nil
}

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
			in := strings.Split(value, ",")
			if len(in) != 2 {
				continue
			}
			min, _ := strconv.Atoi(in[0])
			max, _ := strconv.Atoi(in[1])
			rule.In = [2]int{min, max}
		}
	}

	return rule
}
