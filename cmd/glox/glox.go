package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
    "github.com/samthom/glox/lib"
)

type glox struct {
    hadError bool
}

type Glox interface {
    runFile(string)
    runPrompt()
    run(string)
    error(int, string)
    report(int, string, string)
}

func NewGlox() *glox {
    return new(glox)
}

var glx Glox

func main() {
	args := os.Args

    glx = NewGlox()

	if len(args) > 2 {
		log.Println("Usage: glox [script]")
		os.Exit(64)
	} else if len(args) == 2 {
		glx.runFile(args[1])
	} else {
		glx.runPrompt()
	}
}

func (g *glox) runFile(path string) {
    byteFormat, err := os.ReadFile(path)
    if err != nil {
        log.Fatal(err.Error())
    }

    g.run(string(byteFormat))
    if g.hadError {
        os.Exit(65)
    }
}

func (g *glox) runPrompt() {
    reader := bufio.NewReader(os.Stdin)

    for {
        fmt.Print("> ")
        line, err := reader.ReadString('\n')
        if err != nil || len(line) == 0 {
            break
        }
        g.run(line)
        g.hadError = false
    }
}

func (g *glox) run(source string) {
    // scanner
    scanner := lib.NewScanner(source)
    scanner.ScanTokens()
}

func (g *glox) error(line int, message string) {
    g.report(line, "", message)
}

func (g *glox) report(line int, where, message string) {
    err := fmt.Errorf("[line %d] Error%s: %s", line, where, message)
    fmt.Println(err)
    g.hadError = true
}
