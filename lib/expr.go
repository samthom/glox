package lib

type Expr interface {
    Accept(ExpressionVisitor) any
}

type ExpressionVisitor interface {
    VisitBinary(*Binary) any
    VisitGrouping(*Grouping) any
    VisitLiteral(*Literal) any
    VisitUnary(*Unary) any
}

type Binary struct {
    Left  Expr
    Operator  Token
    Right  Expr
}

func (e *Binary) Accept(v ExpressionVisitor) any {
    return v.VisitBinary(e)
}

type Grouping struct {
    Expression  Expr
}

func (e *Grouping) Accept(v ExpressionVisitor) any {
    return v.VisitGrouping(e)
}

type Literal struct {
    Value  interface{}
}

func (e *Literal) Accept(v ExpressionVisitor) any {
    return v.VisitLiteral(e)
}

type Unary struct {
    Operator  Token
    Right  Expr
}

func (e *Unary) Accept(v ExpressionVisitor) any {
    return v.VisitUnary(e)
}

