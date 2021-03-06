package forms

import (
	"fmt"
	"reflect"
)

// isSlice if given value is slice
func isSlice(value interface{}) bool {
	return reflect.TypeOf(value).Kind() == reflect.Slice
}

// valueInSlice if given string is in slice
func valueInSlice(s string, vs []string) bool {
	for _, v := range vs {
		if s == v {
			return true
		}
	}

	return false
}

// prepareAttributes prepares attributes to use in HTML tags
func prepareAttributes(attrs Attributes, noUse []string) string {
	attributes := ""

	for k, v := range attrs {
		if !valueInSlice(k, noUse) {
			attributes = attributes + fmt.Sprintf(" %s=\"%s\"", k, v)
		}
	}

	return attributes
}

// renderInput returns rendered input HTML tag
func renderInput(as Attributes, n, t string, noUse, vs []string) string {
	if as == nil {
		as = Attributes{}
	}

	if _, ok := as["id"]; !ok {
		as["id"] = fmt.Sprintf("f_%s", n)
	}

	attributes := prepareAttributes(as, noUse)

	if len(vs) > 0 && vs[0] != "" {
		attributes = attributes + fmt.Sprintf(" value=\"%s\"", vs[0])
	}

	return fmt.Sprintf("<input name=\"%s\" type=\"%s\"%s />", n, t, attributes)
}

func anyToString(v interface{}) (string, bool) {
	switch v.(type) {
	default:
		return fmt.Sprintf("%T", v), false
	case bool:
		if v.(bool) {
			return "1", true
		}
		return "0", true
	case int, int8, uint8, int16, uint16, int32, uint32, int64, uint64:
		return fmt.Sprintf("%d", v), true
	case float32, float64, complex64, complex128:
		return fmt.Sprintf("%g", v), true
	case string:
		return v.(string), true
	}
}
