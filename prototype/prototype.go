package prototype

// Prototype 原理接口
type Prototype interface {
	Clone() Prototype
}

// PrototypeManager 具体原型类型
type PrototypeManager struct {
	prototypes map[string]Prototype
}

func NewPrototypeManager() *PrototypeManager {
	return &PrototypeManager{prototypes: make(map[string]Prototype)}
}

func (p *PrototypeManager) PrototypeObj(name string) Prototype {
	return p.prototypes[name].Clone()
}

func (p *PrototypeManager) SetPrototypeObj(name string, prototype Prototype) {
	p.prototypes[name] = prototype
}
