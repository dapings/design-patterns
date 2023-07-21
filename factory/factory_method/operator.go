package factory_method

// Operator 被封装的实际类接口。
type Operator interface {
	SetA(int)
	SetB(int)
	Result() int
}

// OperatorFactory 工厂接口。
type OperatorFactory interface {
	Create() Operator
}

// BaseOperator 封装公用的方法，是 Operator 实例的基类。
type BaseOperator struct {
	a, b int
}

func (o *BaseOperator) SetA(n int) {
	o.a = n
}

func (o *BaseOperator) SetB(n int) {
	o.b = n
}

// PlusOperator 加法实现
type PlusOperator struct {
	*BaseOperator
}

func (o PlusOperator) Result() int {
	return o.a + o.b
}

// PlusOperatorFactory is PlusOperator 的工厂类。
type PlusOperatorFactory struct{}

func (PlusOperatorFactory) Create() Operator {
	return &PlusOperator{&BaseOperator{}}
}

type MinusOperator struct {
	*BaseOperator
}

func (o MinusOperator) Result() int {
	return o.a - o.b
}

type MinusOperatorFactory struct{}

func (MinusOperatorFactory) Create() Operator {
	return MinusOperator{&BaseOperator{}}
}
