package token

import "strings"

// If is a Token which indicates the start of an if statement.
type If struct {
	// Condition is a slice of Tokens which have been provided for the condition of the if statement.
	Condition []Token
	// Body is a slice of Tokens which have been provided for the body of the if statement.
	Body []Token
}

// ID ...
func (*If) ID() int {
	return IDIf
}

// Parse ...
func (*If) Parse(v string) (Token, bool) {
	if strings.ToLower(v) == "if" {
		return &If{}, true
	}
	return nil, false
}

// init ...
func init() {
	add(&If{})
}
