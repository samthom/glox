package main

import (
	"log"
	"os"
	"strings"
)

// package [arg0]

// type Expr interface {
// 	Accept(Visitor) any
// }
//
// type Visitor interface {
// 	VisitBinaryExpr(*Binary) any
// }
//
// type Binary struct {
// 	left Expr
// 	right Expr
// 	operator Token
// }
//
// func (b *Binary) accept(visitor Visitor) any {
// 	return visitor.VisitBinaryExpr(b)
// }
//
// type AstPrinter struct{}
//
// func (p *AstPrinter) visitBinaryExpr(b *Binary) any {
//     return "";
// }

var asList = []string{
	"Binary: Left Expr, Operator Token, Right Expr",
	"Grouping: Expression Expr",
	"Literal: Value interface{}",
	"Unary: Operator Token, Right Expr",
}

func main() {

	args := os.Args
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
	defer file.Close()

	file.Write([]byte("package " + outputPackage + "\n\n"))

	file.Write([]byte(`type Expr interface {
    Accept(ExpressionVisitor) any
}`))

	file.Write([]byte("\n\n"))

	writeVisitorInterface(file)

	for _, expression := range asList {
		elems := strings.Split(expression, ":")
		writeExpression(file, elems)
	}

	return nil
}

func writeExpression(file *os.File, elems []string) {
	writeStruct(file, elems[0], elems[1])
	writeAcceptMethod(file, elems[0])
}

func writeAcceptMethod(file *os.File, exprName string) {
	file.Write([]byte("func (e *" + exprName + ") Accept(v ExpressionVisitor) any {\n"))
	file.Write([]byte("    return v.Visit" + exprName + "(e)\n"))
	file.Write([]byte("}\n\n"))
}

func writeStruct(file *os.File, exprName string, fieldStr string) {
	file.Write([]byte("type " + exprName + " struct {\n"))
	fields := strings.Split(fieldStr, ",")
	for _, field := range fields {
		fieldStr := strings.TrimSpace(field)
		val := strings.Split(fieldStr, " ")
		file.Write([]byte("    " + val[0] + "  " + val[1] + "\n"))
	}
	file.Write([]byte("}\n\n"))
}

func writeVisitorInterface(file *os.File) {
	file.Write([]byte(`type ExpressionVisitor interface {`))
	file.Write([]byte("\n"))
	for _, expression := range asList {
		elems := strings.Split(expression, ":")
		exprName := elems[0]
		file.Write([]byte("    Visit" + exprName + "(*" + exprName + ") any"))
		file.Write([]byte("\n"))
	}

	file.Write([]byte(`}`))
	file.Write([]byte("\n\n"))
}
