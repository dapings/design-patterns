package bridge

func ExampleConcreteSms() {
	m := NewConcreteAbstractionMessage(UseSMS())
	m.SendMessage("have a drink?", "rob")
	// Output:
	// send have a drink? to rob by sms
}

func ExampleConcreteEmail() {
	m := NewConcreteAbstractionMessage(UseEmail())
	m.SendMessage("have a drink?", "rob")
	// Output:
	// send have a drink? to rob by email
}

func ExampleUrgencySms() {
	m := NewUrgencyAbstractionMessage(UseSMS())
	m.SendMessage("have a drink?", "rob")
	// Output:
	// send [Urgency] have a drink? to rob by sms
}

func ExampleUrgencyEmail() {
	m := NewUrgencyAbstractionMessage(UseEmail())
	m.SendMessage("have a drink?", "rob")
	// Output:
	// send [Urgency] have a drink? to rob by email
}
