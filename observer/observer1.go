package observer

import (
	"fmt"
)

// 抽象被观察者类

type EventEmitter interface {
	AddObserver(EventListener)
	RemoveObserver(EventListener)
	NotifyObservers()
}

// 具体被观察者类

type ConcreteSubject struct {
	state     int
	observers []EventListener
}

func (c *ConcreteSubject) AddObserver(observer EventListener) {
	c.observers = append(c.observers, observer)
}

func (c *ConcreteSubject) RemoveObserver(observer EventListener) {
	for i, o := range c.observers {
		if o == observer {
			c.observers = append(c.observers[:i], c.observers[i+1:]...)
			break
		}
	}
}

func (c *ConcreteSubject) NotifyObservers() {
	for _, observer := range c.observers {
		observer.Update(c.state)
	}
}

// 抽象观察者类

type EventListener interface {
	Update(state int)
}

// 具体观察者类

type ConcreteObserver struct {
	name string
}

func (c *ConcreteObserver) Update(state int) {
	fmt.Printf("Observer %s got state %d\n", c.name, state)
}
