package hw09structvalidator

import (
	"fmt"
	"strconv"
	"strings"
)

// intValidRule describes what the conditions target int field must satisfying to.
type intValidRule struct {
	Min int
	Max int
	In  [2]int
}

// parseIntValidRules creates list of rules to check the int struct field by its tag.
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

// Validate validates the string for satisfying the rule
func (r intValidRule) Validate(n int) error {
	if r.Min > 0 && n < r.Min {
		return fmt.Errorf("value must be greater than %d", r.Min)
	}
	if r.Max > 0 && n > r.Max {
		return fmt.Errorf("value must be less than %d", r.Max)
	}
	if r.In != [2]int{} && (n < r.In[0] || n > r.In[1]) {
		return fmt.Errorf("value must be in range %d-%d", r.In[0], r.In[1])
	}

	return nil
}
