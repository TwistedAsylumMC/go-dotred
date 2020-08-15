package token

import (
	"regexp"
	"strconv"
)

// Number is a Token which represents an integer or a float.
type Number struct {
	// Value is the value of the number stored as a float64.
	Value float64
}

// ID ...
func (*Number) ID() int {
	return IDNumber
}

// Parse ...
func (*Number) Parse(v string) (Token, bool) {
	reg := regexp.MustCompile(`^\d*[.]?\d*$`)
	if reg.MatchString(v) {
		n, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return nil, false
		}
		return &Number{n}, true
	}
	return nil, false
}

// init ...
func init() {
	add(&Number{})
}
