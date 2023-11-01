package decorator

// Component 抽象组件
type Component interface {
	Calc() int
}

// ConcreteComponent 具体组件
type ConcreteComponent struct{}

func (*ConcreteComponent) Calc() int {
	return 0
}

// Decorator 装饰器
type Decorator struct {
	Component
}

func (d *Decorator) Calc() int {
	if d.Component != nil {
		return d.Component.Calc()
	}
	return 0
}

// PlusDecorator plus 装饰器
type PlusDecorator struct {
	Decorator
	num int
}

func (d *PlusDecorator) Calc() int {
	return d.Component.Calc() + d.num
}

func WrapPlusDecorator(c Component, num int) Component {
	d := &PlusDecorator{num: num}
	d.Component = c
	return d
}

// MulDecorator mul 装饰器
type MulDecorator struct {
	Component
	num int
}

func (d *MulDecorator) Calc() int {
	return d.Component.Calc() * d.num
}

func WrapMulDecorator(c Component, num int) Component {
	return &MulDecorator{
		Component: c,
		num:       num,
	}
}

// AddDecorator add 装饰器
type AddDecorator struct {
	Component
	num int
}

func (d *AddDecorator) Calc() int {
	return d.Component.Calc() + d.num
}

func WrapAddDecorator(c Component, num int) Component {
	return &AddDecorator{
		Component: c,
		num:       num,
	}
}
