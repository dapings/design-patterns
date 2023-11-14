package memento

import (
	"fmt"
)

type Memento interface{}

// Game 发起人，需要被保存和恢复状态的对象。
type Game struct {
	hp, mp int
}

// 备忘录，用于存储发起人对象的状态。
type gameMemento struct {
	hp, mp int
}

func (g *Game) Play(mpDelta, hpDelta int) {
	g.mp += mpDelta
	g.hp += hpDelta
}

// Save 保存对象，形成一份快照。
func (g *Game) Save() Memento {
	return &gameMemento{
		hp: g.hp,
		mp: g.mp,
	}
}

// Load 恢复对象，恢复快照中的数据
func (g *Game) Load(m Memento) {
	gm := m.(*gameMemento)
	g.mp = gm.mp
	g.hp = gm.hp
}

func (g *Game) Status() {
	fmt.Printf("Current HP:%d, MP:%d\n", g.hp, g.mp)
}
