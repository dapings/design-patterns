package proxy

import (
	"testing"
)

func TestProxy(t *testing.T) {
	// 定义了Real,Proxy的共同行为
	var sub Subject
	sub = &Proxy{}
	result := sub.Do()
	expect := "before:real:after"
	if result != expect {
		t.Fatalf("expect: %s, but got %s", expect, result)
	}
}
