package facade

type ASubsystem interface {
	TestA() string
}

// NewASubsystem return new ASubsystem
func NewASubsystem() ASubsystem {
	return &ASubsystem1{}
}

type ASubsystem1 struct{}

func (a *ASubsystem1) TestA() string {
	return "a sub system1 running"
}

type BSubsystem interface {
	TestB() string
}

// NewBSubsystem return new BSubsystem
func NewBSubsystem() BSubsystem {
	return &BSubsystem1{}
}

type BSubsystem1 struct{}

func (b *BSubsystem1) TestB() string {
	return "b sub system1 running"
}
