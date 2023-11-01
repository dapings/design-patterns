package strategy

func ExamplePayByCash() {
	payment := NewPayment("Foo", "", 123, &Cash{})
	payment.Pay()
	// Output:
	// pay $123 to Foo by cash
}

func ExamplePayByBank() {
	payment := NewPayment("Bar", "0001", 888, &Bank{})
	payment.Pay()
	// Output:
	// pay $888 to Bar by bank account 0001
}
