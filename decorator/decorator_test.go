package decorator

import (
	"fmt"
)

func ExampleDecorator() {
	var c Component = &ConcreteComponent{}
	c = WrapAddDecorator(c, 10)
	c = WrapMulDecorator(c, 8)
	c = WrapPlusDecorator(c, 5)
	result := c.Calc()

	fmt.Printf("result %d\n", result)
	// Output:
	// result 85
}
