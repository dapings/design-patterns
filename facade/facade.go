package facade

import (
	"fmt"
)

type Facader interface {
	Test() string
}

func NewFacade() Facader {
	return &Facade{
		a: NewASubsystem(),
		b: NewBSubsystem(),
	}
}

// Facade 门面类
type Facade struct {
	a ASubsystem
	b BSubsystem
}

func (f *Facade) Test() string {
	aResult := f.a.TestA()
	bResult := f.b.TestB()
	return fmt.Sprintf("%s\n%s", aResult, bResult)
}
