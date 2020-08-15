package function

import (
	"fmt"
	"github.com/twistedasylummc/go-dotred/token"
	"strings"
)

// knownFunctions is a map of all of the functions added from the add() function.
var knownFunctions = make(map[string]Function)

// Function is an interface that holds information about a function to be used in the parser.
type Function interface {
	// Name is the name of the function that will be used in the dotred code.
	Name() string
	// Run is called by the parser when it parses a function call for this function. It can return values to
	// be used in the dotred code.
	Run(v []token.Token) interface{}
}

// add adds a Function to the knownFunctions map so it can be obtained in FromName().
func add(f Function) {
	knownFunctions[strings.ToLower(f.Name())] = f
}

// FromName attempts to find a Function from the provided name. This method is not case sensitive.
func FromName(name string) (Function, bool) {
	f, ok := knownFunctions[strings.ToLower(name)]
	return f, ok
}

// invalidParamCount panics a new error if a function expected more parameters than provided.
func invalidParamCount(f Function, e, g int) {
	panic(fmt.Errorf("%s expects %v parameters, %v given", f.Name(), e, g))
}

// invalidParamType panics a new error if a parameter does not match the expected type in the function.
func invalidParamType(f Function, n int, e, g token.Token) {
	panic(fmt.Errorf("%s expects parameter %v to be a %T, %T given", f.Name(), n, e, g))
}
