package factory_method

import (
	"testing"
)

func compute(f OperatorFactory, a, b int) int {
	op := f.Create()
	op.SetA(a)
	op.SetB(b)
	return op.Result()
}

func TestOperator(t *testing.T) {
	var f OperatorFactory
	f = PlusOperatorFactory{}
	if compute(f, 1, 2) != 3 {
		t.Fatal("error with factory method pattern")
	}
	
	f = MinusOperatorFactory{}
	if compute(f, 4, 2) != 2 {
		t.Fatal("error with factory method pattern")
	}
}
