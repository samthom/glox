package main

import (
	"log"
	"os"
	"strings"
)

// package [arg0]

// import (
// 	"github.com/samthom/glox/lib"
// )
//
// type Expr interface{}
//
// type Binary struct {
// 	left     Expr
// 	operator lib.Token
// 	right    Expr
// }

var asList = []string{
	"Binary : left Expr, operator Token, right Expr",
	"Grouping: expression Expr",
	"Literal: value interface{}",
	"Unary: operator Token, right Expr",
}

func main() {
    args := os.Args;

    if len(args) != 3 {
        log.Fatal("Usage: gen_ast <output directory> <package name>")
    }

    outputDir := args[1]
    outputPackage := args[2]
    err := defineAst(outputDir, "expr", outputPackage)
    if err != nil {
        log.Panic(err)
        os.Exit(1)
    }

    return
}

func defineAst(outputDir, basename, outputPackage string) error {
    filePath := outputDir + "/" + basename + ".go"

    file, err := os.Create(filePath)
    if err != nil {
        return err
    }
    defer file.Close();

    file.Write([]byte("package " + outputPackage + "\n\n"))
//     file.Write([]byte(`import (
//    "github.com/samthom/glox/lib"
// )` + "\n\n"))    

    file.Write([]byte("type Expr interface{}\n\n"))

    for _, expression := range asList {
        elems := strings.Split(expression, ":")
        exprName := elems[0];
        file.Write([]byte("type " + exprName +" struct {\n"))
        fields := strings.Split(elems[1], ",")
        for _, field := range fields {
            fieldStr := strings.TrimSpace(field)
            val := strings.Split(fieldStr, " ")
            file.Write([]byte("    " + val[0] + "  " + val[1] + "\n"));
        }
        file.Write([]byte("}\n\n"))
    }

    return nil
}
