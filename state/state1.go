package state

import (
	"fmt"
)

// 状态类

type State interface {
	Handle()
}

type ConcreteStateA struct{}

func (s *ConcreteStateA) Handle() {
	fmt.Println("当前状态为A，执行A状态的处理逻辑")
}

type ConcreteStateB struct{}

func (s *ConcreteStateB) Handle() {
	fmt.Println("当前状态为B，执行B状态的处理逻辑")
}

// 上下文类

type ConcreteStateCtx struct {
	// 持有当前状态的引用，并将具体的行为委托给当前状态对象处理
	state State
}

func NewConcreteStateCtx() *ConcreteStateCtx {
	return &ConcreteStateCtx{}
}

func (ctx *ConcreteStateCtx) SetState(state State) {
	ctx.state = state
}

func (ctx *ConcreteStateCtx) ThirdHandle() {
	// 委托状态对象处理请求
	ctx.state.Handle()
}
