package lib

import (
	"fmt"
	"reflect"
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
	t := reflect.TypeOf(e.Value)
	typ := t.String()
	switch typ {
	case "int":
		return fmt.Sprintf("%d", e.Value)
	case "float64":
		return fmt.Sprintf("%f", e.Value)
	default:
		return e.Value
	}

}

func (p *Printer) VisitUnary(e *Unary) any {
	return p.parenthesize(e.Operator.Lexeme(), []Expr{e.Right})
}

func (p *Printer) parenthesize(name string, exprs []Expr) string {
	res := "("
	res += name
	for _, expr := range exprs {
		res += " "
		res += expr.Accept(p).(string)
	}
	res += ")"
	return res
}

func NewPrinter() *Printer {
	return &Printer{}
}
