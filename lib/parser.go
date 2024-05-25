package lib

import "errors"

type Parser interface {
	Parse() Expr
}

type parser struct {
	current int
	tokens  []Token
	glx     Glox
}

func NewParser(tokens []Token, glx Glox) Parser {
	return &parser{0, tokens, glx}
}

func (p *parser) Parse() Expr {
	expr, _ := p.expression()
	return expr
}

func (p *parser) expression() (Expr, error) {
	return p.equality()
}

func (p *parser) equality() (Expr, error) {
	expr, err := p.comparison()
	if err != nil {
		return nil, err
	}
	for p.match(BANG_EQUAL, EQUAL_EQUAL) {
		operator := p.previous()
		right, err := p.comparison()
		if err != nil {
			return nil, err
		}
		expr = &Binary{expr, operator, right}
	}
	return expr, nil
}

func (p *parser) comparison() (Expr, error) {
	expr, err := p.term()
	if err != nil {
		return nil, err
	}
	for p.match(GREATER, GREATER_EQUAL, LESS, LESS_EQUAL) {
		operator := p.previous()
		right, err := p.term()
		if err != nil {
			return nil, err
		}
		expr = &Binary{expr, operator, right}
	}
	return expr, nil
}

func (p *parser) term() (Expr, error) {
	expr, err := p.factor()
	if err != nil {
		return nil, err
	}
	for p.match(MINUS, PLUS) {
		operator := p.previous()
		right, err := p.factor()
		if err != nil {
			return nil, err
		}
		expr = &Binary{expr, operator, right}
	}
	return expr, nil
}

func (p *parser) factor() (Expr, error) {
	expr, err := p.unary()
	if err != nil {
		return nil, err
	}
	for p.match(SLASH, STAR) {
		operator := p.previous()
		right, err := p.unary()
		if err != nil {
			return nil, err
		}
		expr = &Binary{expr, operator, right}
	}
	return expr, nil
}

func (p *parser) unary() (Expr, error) {
	if p.match(BANG, MINUS) {
		operator := p.previous()
		right, err := p.unary()
		if err != nil {
			return nil, err
		}
		return &Unary{operator, right}, nil
	}

	return p.primary()
}

func (p *parser) primary() (Expr, error) {
	if p.match(FALSE) {
		return &Literal{false}, nil
	}
	if p.match(TRUE) {
		return &Literal{false}, nil
	}
	if p.match(NIL) {
		return &Literal{false}, nil
	}

	if p.match(NUMBER, STRING) {
		return &Literal{p.previous().Literal()}, nil
	}

	if p.match(LEFT_PAREN) {
		expr, err := p.expression()
		if err != nil {
			return nil, err
		}
		p.consume(RIGHT_PAREN, "Expect ')' after expression")

		return &Grouping{expr}, nil
	}

	return nil, errors.New("Invalid syntx at " + p.tokens[p.current].ToString())
}

func (p *parser) match(types ...TokenType) bool {
	for _, v := range types {
		if p.check(v) {
			p.advance()
			return true
		}
	}

	return false
}

func (p *parser) check(t TokenType) bool {
	if p.isAtEnd() {
		return false
	}

	return p.peek().Type() == t
}

func (p *parser) advance() Token {
	if !p.isAtEnd() {
		p.current++
	}
	return p.previous()
}

func (p *parser) isAtEnd() bool {
	return p.peek().Type() == EOF
}

func (p *parser) peek() Token {
	return p.tokens[p.current]
}

func (p *parser) previous() Token {
	return p.tokens[p.current-1]
}

func (p *parser) consume(t TokenType, message string) (Token, error) {
	if p.check(t) {
		return p.advance(), nil
	}

	return nil, errors.New(p.peek().ToString() + " " + message)
}

type ParseError interface {
}

type parseError struct {
	token   Token
	message string
}

func (p *parser) parseError(token Token, message string) ParseError {
	p.glx.ParseError(token, message)
	return &parseError{token, message}
}

func (p *parser) synchronize() {
	p.advance()

	for !p.isAtEnd() {
		if p.previous().Type() == SEMICOLON {
			return
		}
		switch p.peek().Type() {
		case CLASS:
		case FUN:
		case VAR:
		case FOR:
		case IF:
		case WHILE:
		case PRINT:
		case RETURN:
			return
		}
		p.advance()
	}

}
