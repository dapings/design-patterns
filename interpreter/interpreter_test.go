package interpreter

import (
	"fmt"
	"testing"
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

func TestInterpreter(t *testing.T) {
	p := &Parser{}
	p.Parse("1 + 2 + 3 - 4 + 5 - 6")
	got := p.Result().Interpret()
	expected := 1
	if got != expected {
		t.Fatalf("expected %d, but got %d", expected, got)
	}
}
