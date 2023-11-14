package memento

// Originator 发起人，需要被保存或恢复状态的对象。
type Originator struct {
	state string
}

func NewOriginator() *Originator {
	return &Originator{}
}

func (o *Originator) State() string {
	return o.state
}

func (o *Originator) SetState(state string) {
	o.state = state
}

func (o *Originator) Save() *Memento1 {
	return &Memento1{state: o.State()}
}

func (o *Originator) Restore(m *Memento1) {
	o.state = m.State()
}

// Memento1 备忘录，用于存储发起人对象的状态。
type Memento1 struct {
	state string
}

func (m *Memento1) State() string {
	return m.state
}

// Caretaker 负责人，管理备忘录对象，存储或获取备忘录对象。
type Caretaker struct {
	memento *Memento1
}

func NewCaretaker() *Caretaker {
	return &Caretaker{}
}

func (c *Caretaker) Memento() *Memento1 {
	return c.memento
}

func (c *Caretaker) SetMemento(memento *Memento1) {
	c.memento = memento
}
