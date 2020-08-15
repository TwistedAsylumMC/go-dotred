package parser

import (
	"fmt"
	"github.com/twistedasylummc/go-dotred/function"
	"github.com/twistedasylummc/go-dotred/token"
)

// Parser contains an array of expected Token IDs so the parser can validate the order of tokens provided.
type Parser struct {
	expecting []int
}

// New creates a new Parser which can be used to parse Tokens.
func New() Parser {
	return Parser{}
}

// Parse parses a slice of Tokens, validates their order and executes the code. In the future, the Parse
// method (or something similar) will be used to compile the dotred code into usable redstone circuits.
func (p *Parser) Parse(tokens []token.Token) error {
	skip := 0
	bypassSkip := false
	for i, t := range tokens {
		if _, ok := t.(*token.Comment); ok {
			continue
		}
		if skip > 0 && !bypassSkip {
			skip--
			continue
		}
		if bypassSkip {
			bypassSkip = false
		}
		if !p.expects(t) {
			return fmt.Errorf("unexpected token %#v", t)
		}
		switch t := t.(type) {
		case *token.Function:
			f, ok := function.FromName(t.Name)
			if !ok {
				panic(fmt.Errorf("unknown function %s", t.Name))
			}
			f.Run(t.Parameters)
		case *token.If:
			met := false
			for _, cond := range t.Condition {
				switch cond := cond.(type) {
				case *token.Boolean:
					met = cond.Value
				case *token.Function:
					f, ok := function.FromName(cond.Name)
					if !ok {
						panic(fmt.Errorf("unknown function %s", cond.Name))
					}
					res := f.Run(cond.Parameters)
					if b, ok := res.(bool); ok {
						met = b
					} else {
						panic(fmt.Errorf("condition for if statement must result in a boolean, %T given", res))
					}
				}
			}
			for _, body := range tokens[i+len(t.Condition)+1:] {
				if body.ID() != token.IDComment && body.ID() != token.IDFi {
					t.Body = append(t.Body, body)
					skip++
				} else {
					break
				}
			}
			if met {
				stmt := Parser{}
				if err := stmt.Parse(t.Body); err != nil {
					panic(err)
				}
			}
			bypassSkip = true
			p.expect(token.IDDo)
		case *token.Do:
			p.clearExpectations()
			p.expect(token.IDFi)
		case *token.Fi:
			p.clearExpectations()
		}
	}
	if len(p.expecting) > 0 {
		t, ok := token.FromID(p.expecting[0])
		if ok {
			return fmt.Errorf("expecting %#v", t)
		} else {
			return fmt.Errorf("expecting unknown token %v", p.expecting[0])
		}
	}
	return nil
}

// expect makes the parser expect a specific token ID as the next token.
func (p *Parser) expect(id int) {
	p.expecting = append(p.expecting, id)
}

// expects checks whether or not the Parser is expecting the provided token. If there are no expected tokens,
// the parser will allow tokens such as a function or an if statement.
func (p *Parser) expects(t token.Token) bool {
	if len(p.expecting) == 0 {
		if t.ID() == token.IDFunction || t.ID() == token.IDIf {
			return true
		}
	}
	for _, id := range p.expecting {
		if id == t.ID() {
			return true
		}
	}
	return false
}

// clearExpectations removes all expectations from the Parser until it calls expect() again.
func (p *Parser) clearExpectations() {
	p.expecting = []int{}
}
