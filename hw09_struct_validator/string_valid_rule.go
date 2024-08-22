package hw09structvalidator

import (
	"fmt"
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

// parseStringValidRules creates list of rules to check the string struct field by its tag.
func parseStringValidRules(tag string) []stringValidRule {
	rules := []stringValidRule{}

	for _, condition := range strings.Split(tag, "|") {
		for _, pattern := range strings.Fields(condition) {
			if pattern == "" {
				continue
			}

			parts := strings.SplitN(pattern, ":", 2)
			key, value := parts[0], parts[1]
			if value == "" {
				continue
			}

			rule := stringValidRule{}
			switch key {
			case "len":
				len, _ := strconv.Atoi(value)
				rule.Len = len
			case "regexp":
				rule.Regexp = value
			case "in":
				rule.In = strings.Split(value, ",")
			}
			rules = append(rules, rule)
		}
	}

	return rules
}

// Validate validates the string for satisfying the rule
func (r stringValidRule) Validate(s string) error {
	if r.Len > 0 && len(s) != r.Len {
		return fmt.Errorf("string length must be %d", r.Len)
	}
	if r.Regexp != "" && !regexp.MustCompile(r.Regexp).MatchString(s) {
		return fmt.Errorf("string must match the pattern %s", r.Regexp)
	}
	if len(r.In) > 0 {
		found := false
		for _, v := range r.In {
			if v == s {
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("string must be one of %v", r.In)
		}
	}

	return nil
}
