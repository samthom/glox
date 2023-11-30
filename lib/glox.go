package lib

import (
    "os"
    "fmt"
    "log"
    "bufio"
)

type glox struct {
    hadError bool
}

type Glox interface {
    RunFile(string)
    RunPrompt()
    Run(string)
    Error(int, string)
    Report(int, string, string)
}

func NewGlox() Glox {
    return new(glox)
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
    scanner := NewScanner(source)
    scanner.ScanTokens()
}

func (g *glox) Error(line int, message string) {
    g.Report(line, "", message)
}

func (g *glox) Report(line int, where, message string) {
    err := fmt.Errorf("[line %d] Error%s: %s", line, where, message)
    fmt.Println(err)
    g.hadError = true
}
