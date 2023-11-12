package interpreter

import (
	"fmt"
)

func ExampleInterpret() {
	expr1 := NewTerminalExpression("true")
	expr2 := NewTerminalExpression("false")
	expr3 := NewOrExpression(expr1, expr2)

	i := NewInterpreter(expr3)
	fmt.Println(i.Interpret())
	// Output:
	// true
}
