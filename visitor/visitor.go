package visitor

import (
	"fmt"
)

// Visitor 访问者接口
type Visitor interface {
	Visit(Element)
}

// Element 元素接口
type Element interface {
	Accept(Visitor)
}

// EnterpriseCustomer 具体元素类型
type EnterpriseCustomer struct {
	name string
}

func NewEnterpriseCustomer(name string) *EnterpriseCustomer {
	return &EnterpriseCustomer{name: name}
}

func (c *EnterpriseCustomer) Accept(visitor Visitor) {
	visitor.Visit(c)
}

// IndividualCustomer 具体的元素类型
type IndividualCustomer struct {
	name string
}

func NewIndividualCustomer(name string) *IndividualCustomer {
	return &IndividualCustomer{name: name}
}

func (c *IndividualCustomer) Accept(visitor Visitor) {
	visitor.Visit(c)
}

// ServiceRequestVisitor 具体的访问者类型
type ServiceRequestVisitor struct{}

func (*ServiceRequestVisitor) Visit(customer Element) {
	switch c := customer.(type) {
	case *EnterpriseCustomer:
		fmt.Printf("serving enterprise customer %s\n", c.name)
	case *IndividualCustomer:
		fmt.Printf("serving individual customer %s\n", c.name)
	}
}

// AnalysisVisitor only for enterprise
type AnalysisVisitor struct{}

func (*AnalysisVisitor) Visit(customer Element) {
	switch c := customer.(type) {
	case *EnterpriseCustomer:
		fmt.Printf("analysis enterprise customer %s\n", c.name)
	}
}

// ConcreteCustomer 具体对象结构
type ConcreteCustomer struct {
	customers []Element
}

func (c *ConcreteCustomer) Accept(visitor Visitor) {
	for _, customer := range c.customers {
		customer.Accept(visitor)
	}
}

func (c *ConcreteCustomer) Attach(customer Element) {
	c.customers = append(c.customers, customer)
}

func (c *ConcreteCustomer) Detach(customer Element) {
	for i, v := range c.customers {
		if v == customer {
			c.customers = append(c.customers[:i], c.customers[i+1:]...)
			break
		}
	}
}
