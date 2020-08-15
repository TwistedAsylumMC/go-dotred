package main

import (
	"bufio"
	"errors"
	"github.com/twistedasylummc/go-dotred/parser"
	"github.com/twistedasylummc/go-dotred/token"
	"os"
	"strings"
)

func main() {
	// Validate the arguments provided, the first argument is the path to the Go executable and the second
	// should be a path to a dotred file.
	if len(os.Args) < 2 {
		panic(errors.New("no file path provided"))
	}
	if !strings.HasSuffix(os.Args[1], ".red") {
		panic(errors.New("provided file does not contain the .red extension"))
	}

	// Open the file, Tokenize each line of the file individually.
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	var tokens []token.Token
	s := bufio.NewScanner(file)
	s.Split(bufio.ScanLines)
	for s.Scan() {
		line := s.Text()
		t, err := token.Tokenize(line)
		if err != nil {
			panic(err)
		}
		tokens = append(tokens, t...)
	}

	// Create a new Parser and Parse the Tokens previously obtained.
	p := parser.New()
	if err := p.Parse(tokens); err != nil {
		panic(err)
	}
}
