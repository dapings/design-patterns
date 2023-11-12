package visitor

func ExampleRequestVisitor() {
	c := &ConcreteCustomer{}
	c.Attach(NewEnterpriseCustomer("Tom"))
	c.Attach(NewEnterpriseCustomer("Bob"))
	c.Attach(NewIndividualCustomer("Ra"))
	c.Accept(&ServiceRequestVisitor{})
	// Output:
	// serving enterprise customer Tom
	// serving enterprise customer Bob
	// serving individual customer Ra
}

func ExampleAnalysisVisitor() {
	c := &ConcreteCustomer{}
	c.Attach(NewEnterpriseCustomer("Tom A"))
	c.Attach(NewIndividualCustomer("Ra"))
	c.Attach(NewEnterpriseCustomer("Bob B"))
	c.Accept(&AnalysisVisitor{})
	// Output:
	// analysis enterprise customer Tom A
	// analysis enterprise customer Bob B
}
