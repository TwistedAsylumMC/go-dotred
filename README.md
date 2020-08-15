# go-dotred [![GoDoc](http://godoc.org/github.com/TwistedAsylumMC/go-dotred?status.svg)](https://godoc.org/github.com/TwistedAsylumMC/go-dotred)
A custom programming language interpreted in Go.

## What is the goal of this project?
The goal of go-dotred is to make a custom language than can be compiled into redstone, command blocks & functions for Minecraft: Bedrock Edition.

## Usage
``token.Tokenize()`` can be used to retreive a ``token.Token`` slice from the provided string. These Tokens can be used in ``(parer.Parser).Parse()`` or your own methods. A parser can be made using ``parser.New()``. An example of this in use can be found inside of [main.go](https://github.com/TwistedAsylumMC/go-dotred/blob/master/main.go).
> ``token.Tokenize()`` only works for single lines, and requires the user to handle splitting up the input into lines and tokenizing each line idividually.

## Examples
### Hello World
```
// "say()" is used to print to console.
say("Hello World!")
```
### If statements
```
// "if" is used to start an if statement.
// "true" in this case is just the condition for the if statement.
// "do" is used to end the condition and start the body.
if true do
    say("Hello World!")
// "fi" is used to close the body.
fi
```
These examples can be found in the [examples](https://github.com/TwistedAsylumMC/go-dotred/tree/master/examples) package.