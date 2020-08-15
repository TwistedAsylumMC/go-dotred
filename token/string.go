package token

import (
	"regexp"
)

// String is a Token which represents a string literal.
type String struct {
	// Value is the string provided, excluding the quotes.
	Value string
}

// ID ...
func (*String) ID() int {
	return IDString
}

// Parse ...
func (*String) Parse(v string) (Token, bool) {
	reg := regexp.MustCompile("^\"[^\"]*\"$")
	if reg.MatchString(v) {
		return &String{v[1 : len(v)-1]}, true
	}
	return nil, false
}

// init ...
func init() {
	add(&String{})
}
