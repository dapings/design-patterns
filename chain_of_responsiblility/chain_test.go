package chain

func ExampleChain() {
	pm := NewProjectManager()
	dm := NewDepManager()
	gm := NewGeneralManager()

	pm.SetSuccessor(dm)
	dm.SetSuccessor(gm)

	var m Manager = pm
	m.HandleFeeRequest("bob", 400)
	m.HandleFeeRequest("tom", 1400)
	m.HandleFeeRequest("ra", 10000)
	m.HandleFeeRequest("floor", 400)
	m.HandleFeeRequest("header", 400)
	// Output:
	// project manager permit bob 400 fee request
	// dep manager permit tom 1400 fee request
	// general manager permit ra 10000 fee request
	// project manager don't permit floor 400 fee request
	// project manager don't permit header 400 fee request
}
