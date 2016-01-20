package forms

import (
	"fmt"
	"regexp"
)

const (
	EMAIL_PATTERN = `(?i)\b[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,}\b`
)

func patternMatched(pattern, value string) bool {
	matched, _ := regexp.MatchString(pattern, value)
	return matched
}

type Validator interface {
	IsValid(values []string) (bool, string)
}

type Required struct{}

func (r *Required) IsValid(values []string) (bool, string) {
	if len(values) > 0 && len(values[0]) > 0 {
		return true, ""
	}

	return false, translations["REQUIRED"]
}

type Regexp struct {
	Pattern string
}

func (r *Regexp) IsValid(values []string) (bool, string) {
	if r.Pattern == "" {
		return false, fmt.Sprintf(translations["NO_MATCH_PATTERN"], r.Pattern)
	}

	m := patternMatched(r.Pattern, values[0])
	if !m {
		return false, fmt.Sprintf(translations["NO_MATCH_PATTERN"], r.Pattern)
	}

	return m, ""
}

type Email struct{}

func (v *Email) IsValid(values []string) (bool, string) {
	m := patternMatched(EMAIL_PATTERN, values[0])
	if m {
		return m, ""
	}

	return m, translations["INCORRECT_EMAIL"]
}

type MinLength struct {
	Min int
}

func (v *MinLength) IsValid(values []string) (bool, string) {
	if len(values[0]) >= v.Min {
		return true, ""
	}

	return false, fmt.Sprintf(translations["INCORRECT_MIN_LENGTH"], v.Min)
}

type MaxLength struct {
	Max int
}

func (v *MaxLength) IsValid(values []string) (bool, string) {
	if len(values[0]) <= v.Max {
		return true, ""
	}

	return false, fmt.Sprintf(translations["INCORRECT_MAX_LENGTH"], v.Max)
}

type InSlice struct {
	Values []string
}

func (v *InSlice) IsValid(values []string) (bool, string) {
	for _, value := range values {
		if !ValueInSlice(value, v.Values) {
			return false, fmt.Sprintf(translations["VALUE_NOT_FOUND"], value)
		}
	}

	return true, ""
}
