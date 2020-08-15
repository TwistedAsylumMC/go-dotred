package token

import (
	"regexp"
	"strings"
)

// Function is a Token which links to a built-in function.
type Function struct {
	// Name is the name of the function provided in the dotred code.
	Name string
	// Parameters is a slice of Tokens which have been provided as the function parameters.
	Parameters []Token
}

// ID ...
func (*Function) ID() int {
	return IDFunction
}

// Parse ...
func (*Function) Parse(v string) (Token, bool) {
	reg := regexp.MustCompile(`^\w+\((.*?)\)$`)
	if reg.MatchString(v) {
		reg = regexp.MustCompile("^[^(]*")
		name := reg.FindString(v)

		reg = regexp.MustCompile(`\((.*?)\)`)
		rep := strings.NewReplacer("(", "", ")", "")
		params, err := Tokenize(rep.Replace(reg.FindString(v)))
		if err != nil {
			panic(err)
		}

		return &Function{name, params}, true
	}
	return nil, false
}

// init ...
func init() {
	add(&Function{})
}
