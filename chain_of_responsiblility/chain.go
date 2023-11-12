package chain

import (
	"fmt"
)

type Manager interface {
	HaveRight(money int) bool
	HandleFeeRequest(name string, money int) bool
}

type RequestChain struct {
	Manager
	successor *RequestChain
}

func (r *RequestChain) SetSuccessor(m *RequestChain) {
	r.successor = m
}

func (r *RequestChain) HandleFeeRequest(name string, money int) bool {
	if r.HaveRight(money) {
		return r.Manager.HandleFeeRequest(name, money)
	}
	if r.successor != nil {
		r.successor.HandleFeeRequest(name, money)
	}
	return false
}

type ProjectManager struct{}

func (*ProjectManager) HaveRight(money int) bool {
	return money < 500
}

func (*ProjectManager) HandleFeeRequest(name string, money int) bool {
	if name == "bob" {
		fmt.Printf("project manager permit %s %d fee request\n", name, money)
		return true
	}
	fmt.Printf("project manager don't permit %s %d fee request\n", name, money)
	return false
}

func NewProjectManager() *RequestChain {
	return &RequestChain{
		Manager: &ProjectManager{},
	}
}

type DepManager struct{}

func (*DepManager) HaveRight(money int) bool {
	return money < 5000
}

func (*DepManager) HandleFeeRequest(name string, money int) bool {
	if name == "tom" {
		fmt.Printf("dep manager permit %s %d fee request\n", name, money)
		return true
	}
	fmt.Printf("dep manager don't permit %s %d fee request\n", name, money)
	return false
}

func NewDepManager() *RequestChain {
	return &RequestChain{
		Manager: &DepManager{}}
}

type GeneralManager struct{}

func (*GeneralManager) HaveRight(money int) bool {
	return true
}

func (*GeneralManager) HandleFeeRequest(name string, money int) bool {
	if name == "ra" {
		fmt.Printf("general manager permit %s %d fee request\n", name, money)
		return true
	}
	fmt.Printf("general manager don't permit %s %d fee request\n", name, money)
	return false
}

func NewGeneralManager() *RequestChain {
	return &RequestChain{
		Manager: &GeneralManager{}}
}
