package lib

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type glox struct {
	hadRuntimeError bool
	hadError        bool
	interpreter     *interpreter
}

type Glox interface {
	RunFile(string)
	RunPrompt()
	Run(string)
	Error(int, string)
	Report(int, string, string)
	ParseError(Token, string)
}

func NewGlox() Glox {
	return &glox{false, false, new(interpreter)}
}

func (g *glox) RunFile(path string) {
	byteFormat, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err.Error())
	}

	g.Run(string(byteFormat))
	if g.hadError {
		os.Exit(65)
	}
	if g.hadRuntimeError {
		os.Exit(70)
	}
}

func (g *glox) RunPrompt() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		line, err := reader.ReadString('\n')
		if err != nil || len(line) == 0 {
			break
		}
		g.Run(line)
		g.hadError = false
	}
}

func (g *glox) Run(source string) {
	// scanner
	scanner := NewScanner(source, g)
	tokens := scanner.ScanTokens()

	parser := NewParser(tokens, g)
	expression := parser.Parse()
	if err := g.interpreter.Interpret(expression); err != nil {
		g.RuntimeError(err.(RuntimeError))
	}

	if g.hadError {
		return
	}

	// printer := NewPrinter()
	// fmt.Println(printer.Print(expression))
}

func (g *glox) Error(line int, message string) {
	g.Report(line, "", message)
}

func (g *glox) ParseError(token Token, message string) {
	if token.Type() == EOF {
		g.Report(token.Line(), " at end", message)
	} else {
		g.Report(token.Line(), "at '"+token.Lexeme()+"'", message)
	}
}

func (g *glox) Report(line int, where, message string) {
	err := fmt.Errorf("[line %d] Error%s: %s", line, where, message)
	fmt.Println(err)
	g.hadError = true
}

func (g *glox) RuntimeError(err RuntimeError) {
	fmt.Printf("%v\n[line %d]", err.Error(), err.token.Line())
	g.hadRuntimeError = true
}
