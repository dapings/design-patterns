package facade

import (
	"testing"
)

func TestFacader(t *testing.T) {
	f := NewFacade()
	result := f.Test()
	expect := "a sub system1 running\nb sub system1 running"
	if result != expect {
		t.Fatalf("expect %s, but got %s", expect, result)
	}
}
