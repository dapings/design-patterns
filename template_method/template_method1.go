package template_method

import (
	"fmt"
)

// AbstractClass 抽象类
type AbstractClass interface {
	TemplateMethod()
	PrimitiveOperation1()
	PrimitiveOperation2()
}

// ConcreteClassA 具体类A
type ConcreteClassA struct{}

func (c *ConcreteClassA) TemplateMethod() {
	fmt.Println("具体类A的模板方法")
	c.PrimitiveOperation1()
	c.PrimitiveOperation2()
}

func (c *ConcreteClassA) PrimitiveOperation1() {
	fmt.Println("具体类A的操作1")
}

func (c *ConcreteClassA) PrimitiveOperation2() {
	fmt.Println("具体类A的操作2")
}

// ConcreteClassB 具体类B
type ConcreteClassB struct{}

func (c *ConcreteClassB) TemplateMethod() {
	fmt.Println("具体类B的模板方法")
	c.PrimitiveOperation1()
	c.PrimitiveOperation2()
}

func (c *ConcreteClassB) PrimitiveOperation1() {
	fmt.Println("具体类B的操作1")
}

func (c *ConcreteClassB) PrimitiveOperation2() {
	fmt.Println("具体类B的操作2")
}
