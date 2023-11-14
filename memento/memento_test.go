package memento

import (
	"testing"
)

func ExampleGame() {
	game := &Game{
		hp: 10,
		mp: 10,
	}

	game.Status()
	progress := game.Save()

	game.Play(-2, -3)
	game.Status()

	game.Load(progress)
	game.Status()

	// Output:
	// Current HP:10, MP:10
	// Current HP:7, MP:8
	// Current HP:10, MP:10
}

func TestMemento(t *testing.T) {
	originator := NewOriginator()
	caretaker := NewCaretaker()

	originator.SetState("1")
	t.Logf("originator state: %s", originator.State())
	originator.SetState("2")
	t.Logf("originator state: %s", originator.State())
	caretaker.SetMemento(originator.Save())
	t.Logf("caretaker save originator state: %s", caretaker.Memento().State())
	originator.SetState("3")
	t.Logf("originator current state: %s", originator.State())
	originator.Restore(caretaker.Memento())
	t.Logf("originator restored state: %s", originator.State())
	if originator.State() != caretaker.Memento().State() {
		t.Fatalf("save and restore op, excepted and got not equal")
	}
}
