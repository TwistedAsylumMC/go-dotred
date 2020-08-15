package token

import "strings"

// Comment is a Token which has no effect on the code and is just used to document or leave notes.
type Comment struct {
	// Message is everything inside of the comment
	Message string
}

// ID ...
func (*Comment) ID() int {
	return IDComment
}

// Parse ...
func (*Comment) Parse(v string) (Token, bool) {
	if strings.HasPrefix(v, "//") {
		return &Comment{}, true
	}
	return nil, false
}

// init ...
func init() {
	add(&Comment{})
}
