package adapter

import (
	"testing"
)

func TestAdapter(t *testing.T) {
	expect := "adaptee method"
	adaptee := NewAdaptee()
	target := NewAdapter(adaptee)
	result := target.Request()
	if result != expect {
		t.Fatalf("expect: %s, got %s", expect, result)
	}
}
