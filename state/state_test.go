package state

func ExampleWeek() {
	ctx := NewDayCtx()
	todayAndNext := func() {
		ctx.Today()
		ctx.Next()
	}

	for i := 0; i < 8; i++ {
		todayAndNext()
	}
	// Output:
	// Sunday
	// Monday
	// Tuesday
	// Wednesday
	// Thursday
	// Friday
	// Saturday
	// Sunday
}

func ExampleConcreteState() {
	ctx := NewConcreteStateCtx()
	a := &ConcreteStateA{}
	b := &ConcreteStateB{}

	ctx.SetState(a)
	ctx.ThirdHandle()

	ctx.SetState(b)
	ctx.ThirdHandle()

	// Output:
	// 当前状态为A，执行A状态的处理逻辑
	// 当前状态为B，执行B状态的处理逻辑
}
