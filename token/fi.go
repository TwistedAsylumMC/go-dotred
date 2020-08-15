package token

import (
	"strings"
)

// Fi is a Token which indicates the end of an if statement.
type Fi struct{}

// ID ...
func (*Fi) ID() int {
	return IDFi
}

// Parse ...
func (*Fi) Parse(v string) (Token, bool) {
	if strings.ToLower(v) == "fi" {
		return &Fi{}, true
	}
	return nil, false
}

// init ...
func init() {
	add(&Fi{})
}
