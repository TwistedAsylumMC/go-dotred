package function

import (
	"fmt"
	"github.com/twistedasylummc/go-dotred/token"
)

// Say is a function that (currently) prints the first parameter provided. In the future this function will
// compile into a command block which runs the "/say" command in-game.
type Say struct{}

// Name ...
func (Say) Name() string {
	return "say"
}

// Run ...
func (s Say) Run(v []token.Token) interface{} {
	if len(v) != 1 {
		invalidParamCount(s, 1, len(v))
	}
	m, ok := v[0].(*token.String)
	if !ok {
		invalidParamType(s, 1, &token.String{}, v[0])
	}
	fmt.Println(m.Value)
	return nil
}

// init ...
func init() {
	add(&Say{})
}
