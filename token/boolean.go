package token

import "strings"

// Boolean is a Token which contains a boolean Value.
type Boolean struct {
	// Value is a either "true" or "false".
	Value bool
}

// ID ...
func (*Boolean) ID() int {
	return IDBoolean
}

// Parse ...
func (*Boolean) Parse(v string) (Token, bool) {
	switch strings.ToLower(v) {
	case "false":
		return &Boolean{false}, true
	case "true":
		return &Boolean{true}, true
	}
	return nil, false
}

// init ...
func init() {
	add(&Boolean{})
}
