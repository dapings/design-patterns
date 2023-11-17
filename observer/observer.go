package observer

import (
	"fmt"
)

// 观察者

type Observer interface {
	Update(*Subject)
}

type Reader struct {
	name string
}

func NewReader(name string) *Reader {
	return &Reader{
		name: name,
	}
}

func (r *Reader) Update(s *Subject) {
	fmt.Printf("%s receive %s\n", r.name, s.context)
}

// 被观察者

type Subject struct {
	observers []Observer
	context   string
}

func NewSubject() *Subject {
	return &Subject{
		observers: make([]Observer, 0),
	}
}

// Attach add observer, equal sub.
func (s *Subject) Attach(o Observer) {
	s.observers = append(s.observers, o)
}

// Detach remove observer.
func (s *Subject) Detach(o Observer) {
	for i, observer := range s.observers {
		if observer == o {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

func (s *Subject) notify() {
	for _, o := range s.observers {
		o.Update(s)
	}
}

// UpdateContext notify observers info, equal pub.
func (s *Subject) UpdateContext(context string) {
	s.context = context
	s.notify()
}
