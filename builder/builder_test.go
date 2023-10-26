package builder

import (
	"testing"
)

func TestBuilder1(t *testing.T) {
	b := &Builder1{}
	d := NewDirector(b)
	d.BuildX()
	res := b.Result()
	if res != "123" {
		t.Fatalf("builder1 fail, expect 123, but got %s", res)
	}
}

func TestBuilder2(t *testing.T) {
	b := &Builder2{}
	d := NewDirector(b)
	d.BuildX()
	res := b.Result()
	if res != 6 {
		t.Fatalf("builder2 fail, expect 6, but got %d", res)
	}
}
