package interpreter

// Expression 表达式接口
type Expression interface {
	Interpret() bool
}

// TerminalExpression 终结符表达式
type TerminalExpression struct {
	data string
}

func NewTerminalExpression(data string) *TerminalExpression {
	return &TerminalExpression{data: data}
}

func (t *TerminalExpression) Interpret() bool {
	return t.data == "true"
}

// OrExpression 非终结符表达式
type OrExpression struct {
	expr1 Expression
	expr2 Expression
}

func NewOrExpression(expr1 Expression, expr2 Expression) *OrExpression {
	return &OrExpression{expr1: expr1, expr2: expr2}
}

func (o *OrExpression) Interpret() bool {
	return o.expr1.Interpret() || o.expr2.Interpret()
}

// Interpreter 解释器
type Interpreter struct {
	expression Expression
}

func NewInterpreter(expression Expression) *Interpreter {
	return &Interpreter{expression: expression}
}

func (i *Interpreter) Interpret() bool {
	return i.expression.Interpret()
}
