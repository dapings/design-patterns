package prototype

import (
	"testing"
)

var manager *PrototypeManager

func init() {
	manager = NewPrototypeManager()
	obj1 := &Obj1{
		name: "obj1",
	}
	manager.SetPrototypeObj("obj1", obj1)
}

func TestCloneFromManager(t *testing.T) {
	co := manager.PrototypeObj("obj1").Clone()

	obj1 := co.(*Obj1)
	if obj1.name != "obj1" {
		t.Fatal("error")
	}
}

func TestClone(t *testing.T) {
	obj1 := manager.PrototypeObj("obj1")
	obj1Clone := obj1.Clone()

	if obj1 == obj1Clone {
		t.Fatal("error, expect the different obj, but got same")
	}
}
