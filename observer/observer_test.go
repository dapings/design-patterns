package observer

func ExampleObserver() {
	subject := NewSubject()
	reader1 := NewReader("reader1")
	reader2 := NewReader("reader2")
	reader3 := NewReader("reader3")

	subject.Attach(reader1)
	subject.Attach(reader2)
	subject.Attach(reader3)

	subject.UpdateContext("observer mode")

	// Output:
	// reader1 receive observer mode
	// reader2 receive observer mode
	// reader3 receive observer mode
}

func ExampleEventListener() {
	observer1 := &ConcreteObserver{"Observer1"}
	observer2 := &ConcreteObserver{"Observer2"}

	subject := &ConcreteSubject{}
	subject.AddObserver(observer1)
	subject.AddObserver(observer2)

	// 修改被观察者对象的状态
	subject.state = 100

	// 通知观察者对象
	subject.NotifyObservers()

	// Output:
	// Observer Observer1 got state 100
	// Observer Observer2 got state 100
}
