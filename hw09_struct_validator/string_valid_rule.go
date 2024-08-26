package hw09structvalidator

import (
	"regexp"
	"strconv"
	"strings"
)

// stringValidRule describes what the conditions target string field must satisfying to.
type stringValidRule struct {
	Len    int
	Regexp string
	In     []string
}

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
			len, _ := strconv.Atoi(value)
			rule.Len = len
		case "regexp":
			rule.Regexp = value
		case "in":
			rule.In = strings.Split(value, ",")
		}
	}

	return rule
}
