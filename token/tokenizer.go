package token

import (
	"fmt"
	"regexp"
	"strings"
)

// knownTokens is a map of all of the Tokens added from the add() function.
var knownTokens = make(map[int]Token)

// Token is an interface that holds the information for a Token.
type Token interface {
	ID() int
	Parse(v string) (Token, bool)
}

// add adds a Token to the knownTokens map so it can be obtained in FromID().
func add(t Token) {
	knownTokens[t.ID()] = t
}

// FromID attempts to find a Token from the provided ID.
func FromID(id int) (Token, bool) {
	t, ok := knownTokens[id]
	return t, ok
}

// Tokenize splits a line of dotred code into Tokens which can then be parsed by the Parser. This function
// currently only works with single lines, so Tokenize must be called seperately for each line.
func Tokenize(v string) ([]Token, error) {
	var tokens []Token
	reg := regexp.MustCompile(`^//.*$|"((?:\\.|[^"])*)"|(\S+)\(([^)]+)\)|\w+`)
	words := reg.FindAllString(v, -1)
	skip := 0
	for i, word := range words {
		if skip > 0 {
			skip--
			continue
		}
		valid := false
		for _, t := range knownTokens {
			if token, ok := t.Parse(word); ok {
				if stmt, ok := token.(*If); ok {
					rem := strings.Join(words[i+1:], " ")
					cond := rem[:strings.Index(strings.ToLower(rem), "do")]
					condTokens, err := Tokenize(cond)
					if err != nil {
						return tokens, err
					}
					skip = len(condTokens)
					stmt.Condition = condTokens
				}
				if cmt, ok := token.(*Comment); ok {
					cmt.Message = strings.Trim(strings.Join(words[i:], " ")[2:], " ")
				}
				tokens = append(tokens, token)
				valid = true
				continue
			}
		}
		if !valid {
			return tokens, fmt.Errorf("unexpected \"%s\"", word)
		}
	}
	return tokens, nil
}
