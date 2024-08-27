package hw09structvalidator

import (
	"regexp"
	"strconv"
	"strings"
)

// stringValidRule describes what the conditions target string field must satisfying to.
type stringValidRule struct {
	Len    int      // The length that the field must be equal to.
	Regexp string   // The regular expression that the field must match.
	In     []string // The list of values that the field must be equal to.
}

// validateString checks if the value satisfies the conditions.
func validateString(value string, condition string) error {
	rule := parseStringCondition(condition)

	if rule.Len > 0 && len(value) != rule.Len {
		return ErrInvalidLength
	}
	if rule.Regexp != "" && !regexp.MustCompile(rule.Regexp).MatchString(value) {
		return ErrNotMatchRegexp
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
			return ErrNotInRange
		}
	}

	return nil
}

// parseStringCondition parses the condition string and returns the rule to check the field value.
func parseStringCondition(condition string) (rule stringValidRule) {
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
		case "len":
			l, _ := strconv.Atoi(value)
			rule.Len = l
		case "regexp":
			rule.Regexp = value
		case "in":
			rule.In = strings.Split(value, ",")
		}
	}

	return rule
}
