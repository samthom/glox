package lib

type Expr interface {
	Accept(ExpressionVisitor) (any, error)
}

type ExpressionVisitor interface {
	VisitBinary(*Binary) (any, error)
	VisitGrouping(*Grouping) (any, error)
	VisitLiteral(*Literal) (any, error)
	VisitUnary(*Unary) (any, error)
}

type Binary struct {
	Left     Expr
	Operator Token
	Right    Expr
}

func (e *Binary) Accept(v ExpressionVisitor) (any, error) {
	return v.VisitBinary(e)
}

type Grouping struct {
	Expression Expr
}

func (e *Grouping) Accept(v ExpressionVisitor) (any, error) {
	return v.VisitGrouping(e)
}

type Literal struct {
	Value interface{}
}

func (e *Literal) Accept(v ExpressionVisitor) (any, error) {
	return v.VisitLiteral(e)
}

type Unary struct {
	Operator Token
	Right    Expr
}

func (e *Unary) Accept(v ExpressionVisitor) (any, error) {
	return v.VisitUnary(e)
}
