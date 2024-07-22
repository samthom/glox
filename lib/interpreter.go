package lib

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

type interpreter struct {
	value int
}

func NewInterpreter() ExpressionVisitor {
	return &interpreter{}
}

func (i *interpreter) Interpret(e Expr) error {
	value, err := i.evaluate(e)
	if err != nil {
		// report runtime error to Lox object
		return err
	}
	fmt.Println(">", stringify(value))
	return nil
}

func (i *interpreter) VisitLiteral(e *Literal) (any, error) {
	return e.Value, nil
}

func (i *interpreter) VisitGrouping(e *Grouping) (any, error) {
	return i.evaluate(e)
}

func (i *interpreter) VisitUnary(e *Unary) (any, error) {
	right, err := i.evaluate(e.Right)
	if err != nil {
		return nil, err
	}

	switch e.Operator.Type() {
	case MINUS:
		if err := checkNumberAndOperand(e.Operator, right); err != nil {
			return nil, err
		}
		v := right.(float32)
		return -v, nil
	case BANG:
		return !isTruthy(right), nil
	}

	return nil, nil
}

func (i *interpreter) VisitBinary(e *Binary) (any, error) {
	right, err := i.evaluate(e.Right)
	if err != nil {
		return nil, err
	}
	left, err := i.evaluate(e.Left)
	if err != nil {
		return nil, err
	}

	leftType := reflect.TypeOf(left).String()
	rightType := reflect.TypeOf(right).String()
	switch e.Operator.Type() {
	case PLUS:
		if (leftType == "string" || rightType == "string") && (leftType == "float64" || rightType == "float64") {
			leftString := fmt.Sprintf("%v", left)
			rightString := fmt.Sprintf("%v", right)
			return leftString + rightString, nil
		} else if leftType == "float64" && rightType == "float64" {
			return (left.(float64) + right.(float64)), nil
		} else {
			return nil, RuntimeError{e.Operator, errors.New("Operands must be two numbers or strings")}
		}
	case MINUS:
		if err := checkNumberAndOperands(e.Operator, left, right); err != nil {
			return nil, err
		}
		return (left.(float64) - right.(float64)), nil
	case SLASH:
		if err := checkNumberAndOperands(e.Operator, left, right); err != nil {
			return nil, err
		}
		return (left.(float64) / right.(float64)), nil
	case STAR:
		if err := checkNumberAndOperands(e.Operator, left, right); err != nil {
			return nil, err
		}
		return (left.(float64) * right.(float64)), nil
	case GREATER:
		if err := checkNumberAndOperands(e.Operator, left, right); err != nil {
			return nil, err
		}
		return (left.(float64) > right.(float64)), nil
	case GREATER_EQUAL:
		if err := checkNumberAndOperands(e.Operator, left, right); err != nil {
			return nil, err
		}
		return (left.(float64) >= right.(float64)), nil
	case LESS:
		if err := checkNumberAndOperands(e.Operator, left, right); err != nil {
			return nil, err
		}
		return (left.(float64) < right.(float64)), nil
	case LESS_EQUAL:
		if err := checkNumberAndOperands(e.Operator, left, right); err != nil {
			return nil, err
		}
		return (left.(float64) <= right.(float64)), nil
	case BANG_EQUAL:
		return !isEqual(left, right), nil
	case EQUAL_EQUAL:
		return isEqual(left, right), nil
	}

	return nil, nil
}

func (i *interpreter) evaluate(e Expr) (any, error) {
	return e.Accept(i)
}

func isTruthy(object interface{}) bool {
	if reflect.TypeOf(object).String() == "bool" {
		return object.(bool)
	}
	return false
}

func isEqual(a interface{}, b interface{}) bool {
	if a == nil && b == nil {
		return true
	} else if a == nil {
		return false
	}

	return a == b
}

func checkNumberAndOperand(operator Token, operand interface{}) error {
	if reflect.TypeOf(operand).String() == "float64" {
		return nil
	}
	return RuntimeError{operator, errors.New("Operand must be a number")}
}

func checkNumberAndOperands(operator Token, left interface{}, right interface{}) error {
	if reflect.TypeOf(left).String() == "float64" && reflect.TypeOf(right).String() == "float64" {
		return nil
	}
	return RuntimeError{operator, errors.New("Operands must be a number")}
}

func stringify(object interface{}) string {
	if object == nil {
		return "nil"
	}

	objectType := reflect.TypeOf(object).String()

	if objectType == "float64" {
		return strconv.FormatFloat(object.(float64), 'g', -1, 32)
	} else if objectType == "bool" {
		if object.(bool) {
			return "true"
		} else {
			return "false"
		}
	}

	return object.(string)
}
