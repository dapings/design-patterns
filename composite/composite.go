package composite

import (
	"fmt"
)

// Component 组件接口
// 通过统一的接口  Component ，以一致的方式处理单个对象和对象组合。
type Component interface {
	Parent() Component
	SetParent(Component)
	Name() string
	SetName(string)
	AddChild(Component)
	Print(string)
	Remove(Component)
}

const (
	LeafNode = iota
	CompositeNode
)

func NewComponent(kind int, name string) Component {
	var c Component
	switch kind {
	case LeafNode:
		c = NewLeaf()
	case CompositeNode:
		c = NewComposite()
	}

	c.SetName(name)
	return c
}

type component struct {
	name   string
	parent Component
}

func (c *component) Parent() Component {
	return c.parent
}

func (c *component) SetParent(parent Component) {
	c.parent = parent
}

func (c *component) Name() string {
	return c.name
}

func (c *component) SetName(name string) {
	c.name = name
}

func (c *component) AddChild(Component) {}

func (c *component) Print(string) {}

func (c *component) Remove(Component) {}

// Leaf 叶节点，没有子节点
type Leaf struct {
	component
}

func NewLeaf() *Leaf {
	return &Leaf{}
}

func (l *Leaf) AddChild(Component) {
	fmt.Println("cannot add to a leaf")
}

func (l *Leaf) Print(pre string) {
	fmt.Printf("%s-%s\n", pre, l.Name())
}

func (l *Leaf) Remove(Component) {
	fmt.Println("cannot remove from a leaf")
}

// Composite 组合节点，分支节点，可以包含其他节点，如叶子节点、组合节点
type Composite struct {
	component
	childs []Component
}

func NewComposite() *Composite {
	return &Composite{
		childs: make([]Component, 0),
	}
}

func (c *Composite) AddChild(child Component) {
	c.childs = append(c.childs, child)
}

func (c *Composite) Print(pre string) {
	fmt.Printf("%s+%s\n", pre, c.Name())
	pre += " "
	for _, child := range c.childs {
		child.Print(pre)
	}
}

func (c *Composite) Remove(child Component) {
	for i, comp := range c.childs {
		if child == comp {
			c.childs = append(c.childs[:i], c.childs[i+1:]...)
			break
		}
	}
}
