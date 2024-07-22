package lib

import (
	"fmt"
	"reflect"
)

type Printer struct {
	result string
}

func (p *Printer) Print(expr Expr) (any, error) {
	// return expr.Accept(p).(string)
	v, err := expr.Accept(p)
	return v.(string), err
}

func (p *Printer) VisitBinary(e *Binary) (any, error) {
	return p.parenthesize(e.Operator.Lexeme(), []Expr{e.Left, e.Right})
}

func (p *Printer) VisitGrouping(e *Grouping) (any, error) {
	return p.parenthesize("group", []Expr{e.Expression})
}

func (p *Printer) VisitLiteral(e *Literal) (any, error) {
	if e.Value == nil {
		return "nil", nil
	}
	t := reflect.TypeOf(e.Value)
	typ := t.String()
	switch typ {
	case "int":
		return fmt.Sprintf("%d", e.Value), nil
	case "float64":
		return fmt.Sprintf("%f", e.Value), nil
	default:
		return e.Value, nil
	}

}

func (p *Printer) VisitUnary(e *Unary) (any, error) {
	return p.parenthesize(e.Operator.Lexeme(), []Expr{e.Right})
}

func (p *Printer) parenthesize(name string, exprs []Expr) (string, error) {
	res := "("
	res += name
	for _, expr := range exprs {
		res += " "
		r, err := expr.Accept(p)
		if err != nil {
			return "", err
		}
		res += r.(string)
		// res += expr.Accept(p).(string)
	}
	res += ")"
	return res, nil
}

func NewPrinter() *Printer {
	return &Printer{}
}
