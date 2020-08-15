package token

import "strings"

// Do is a Token which indicates the end of the condition for an if statement.
type Do struct{}

// ID ...
func (*Do) ID() int {
	return IDDo
}

// Parse ...
func (*Do) Parse(v string) (Token, bool) {
	if strings.ToLower(v) == "do" {
		return &Do{}, true
	}
	return nil, false
}

// init ...
func init() {
	add(&Do{})
}
