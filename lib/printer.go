package lib

import (
	"fmt"
)

type Printer struct {
    result string
}

func (p *Printer) Print(expr Expr) any {
    return expr.Accept(p).(string)
}

func (p *Printer) VisitBinary(e *Binary) any {
    return p.parenthesize(e.Operator.Lexeme(), []Expr{e.Left, e.Right})
}

func (p *Printer) VisitGrouping(e *Grouping) any {
    return p.parenthesize("group", []Expr{e.Expression})
}

func (p *Printer) VisitLiteral(e *Literal) any {
    if e.Value == nil {
        return "nil"
    }
    s := fmt.Sprintf("%f", e.Value)
    return s
}

func (p *Printer) VisitUnary(e *Unary) any {
    return p.parenthesize(e.Operator.Lexeme(), []Expr{e.Right})
}

func (p *Printer) parenthesize(name string, exprs []Expr) string {
    res := "("
    res += name
    for _, expr := range exprs {
        fmt.Printf("%#v\n", expr)
        res += " "
        res += expr.Accept(p).(string)
        fmt.Println(res)
    }
    res += ")"
    return res
}

func NewPrinter() *Printer {
    return &Printer{}
}
