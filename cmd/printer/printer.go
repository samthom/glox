package main

import (
	"fmt"

	"github.com/samthom/glox/lib"
)


func main() {
    expression := lib.Binary{
        Left: &lib.Unary{
            Operator: lib.NewToken(lib.MINUS, "-", nil, 1),
            Right: &lib.Literal{
                Value: 123,
            },
        },
        Operator: lib.NewToken(lib.STAR, "*", nil, 1),
        Right: &lib.Grouping{
            Expression: &lib.Literal{
                Value: 45.67,
            },
        },
    }

    printer := lib.NewPrinter()
    str := printer.Print(&expression).(string)
    fmt.Print(str)
}
