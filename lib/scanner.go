package lib

import (
	"strconv"
)

var keywords = map[string]TokenType{
	"and":    AND,
	"class":  CLASS,
	"else":   ELSE,
	"false":  FALSE,
	"for":    FOR,
	"fun":    FUN,
	"if":     IF,
	"nil":    NIL,
	"or":     OR,
	"print":  PRINT,
	"return": RETURN,
	"super":  SUPER,
	"this":   THIS,
	"true":   TRUE,
	"var":    VAR,
	"while":  WHILE,
}

type scanner struct {
	source  string
	tokens  []Token
	start   int
	current int
	line    int
}

type Scanner interface {
	ScanTokens()
	isAtEnd() bool
}

func NewScanner(source string) Scanner {
	list := []Token{}
	return &scanner{source, list, 0, 0, 1}
}

func (s *scanner) ScanTokens() {
	for !s.isAtEnd() {
		s.start = s.current
		s.scanToken()
	}

	t := NewToken(EOF, "", nil, s.line)
	s.tokens = append(s.tokens, t)
}

func (s *scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

func (s *scanner) scanToken() {
	c := s.advance()
	switch c {
	case string(LEFT_PAREN):
		s.add(LEFT_PAREN)
		break
	case string(RIGHT_PAREN):
		s.add(RIGHT_PAREN)
		break
	case string(LEFT_BRACE):
		s.add(LEFT_BRACE)
		break
	case string(RIGHT_BRACE):
		s.add(RIGHT_BRACE)
		break
	case string(COMMA):
		s.add(COMMA)
		break
	case string(DOT):
		s.add(DOT)
		break
	case string(MINUS):
		s.add(MINUS)
		break
	case string(PLUS):
		s.add(PLUS)
		break
	case string(SEMICOLON):
		s.add(SEMICOLON)
		break
	case string(STAR):
		s.add(STAR)
		break
	case string(BANG):
		if s.match(string(EQUAL)) {
			s.add(BANG_EQUAL)
		} else {
			s.add(BANG)
		}
		break
	case string(EQUAL):
		if s.match(string(EQUAL)) {
			s.add(EQUAL_EQUAL)
		} else {
			s.add(EQUAL)
		}
		break
	case string(LESS):
		if s.match(string(EQUAL)) {
			s.add(LESS_EQUAL)
		} else {
			s.add(LESS)
		}
		break
	case string(GREATER):
		if s.match(string(EQUAL)) {
			s.add(GREATER_EQUAL)
		} else {
			s.add(GREATER)
		}
		break
	case string(SLASH):
		if s.match(string(SLASH)) {
			for s.peek() != "\n" && !s.isAtEnd() {
				s.advance()
			}
		} else {
			s.add(SLASH)
		}
		break
	case " ", "\r", "\t":
		break
	case "\n":
		s.line++
		break
	case string(STRING):
		s.string()
		break
	default:
		if isDigit(c) {
			s.number()
		} else if isAlpha(c) {
			s.identifier()
		} else {
			glx.error(s.line, "Unexpected character.")
		}
		break
	}
}

func (s *scanner) advance() string {
    v := s.source[s.current];
	s.current++
	return string(v) // going out on a limb and only expecting ASCII characters
}

func (s *scanner) add(tknType TokenType) {
	s.addToken(tknType, nil)
}

func (s *scanner) addToken(tknType TokenType, literal interface{}) {
	text := s.source[s.start:s.current]
	s.tokens = append(s.tokens, NewToken(tknType, text, literal, s.line))
}

func (s *scanner) match(expected string) bool {
	if s.isAtEnd() {
		return false
	}
	if string(s.source[s.current]) != expected {
		return false
	}
	s.current++
	return true
}

func (s *scanner) peek() string {
	if s.isAtEnd() {
		return "\\0"
	}
	return string(s.source[s.current])
}

func (s *scanner) peekNext() string {
	if s.current+1 >= len(s.source) {
		return "\\0"
	}
	return string(s.source[s.current+1])
}

func (s *scanner) string() {
	for s.peek() != string(STRING) && !s.isAtEnd() {
		if s.peek() == "\n" {
			s.line++
		}
		s.advance()
	}

	if s.isAtEnd() {
		glx.error(s.line, "Unterminated string.")
	} else {
		s.advance()
		value := s.source[s.start+1 : s.current-1]
		s.addToken(STRING, value)
	}
}

func (s *scanner) number() {
	if isDigit(s.peek()) {
		s.advance()
	}

	// Look for a fractional part
	if s.peek() == "." && isDigit(s.peekNext()) {
		// consume the "."
		s.advance()
		for isDigit(s.peek()) {
			s.advance()
		}
	}
	f, _ := strconv.ParseFloat(s.source[s.start:s.current], 64)
	s.addToken(NUMBER, f)
}

func (s *scanner) identifier() {
	if isAlphaNumeric(s.peek()) {
		s.advance()
	}
	text := s.source[s.start:s.current]
	keyword := keywords[text]
	s.add(keyword)
}

func isDigit(c string) bool {
	return c >= "0" && c <= "9"
}

func isAlpha(c string) bool {
	return (c >= "a" && c <= "z") || (c >= "A" && c <= "Z") || c == "_"
}

func isAlphaNumeric(c string) bool {
	return isAlpha(c) || isDigit(c)
}
