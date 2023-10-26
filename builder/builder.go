package builder

// Builder 生成器接口
type Builder interface {
	Part1()
	Part2()
	Part3()
}

// Director 指挥者
type Director struct {
	// 需要建造什么，交给专业的建造者
	builder Builder
}

func NewDirector(builder Builder) *Director {
	return &Director{
		builder: builder,
	}
}

// BuildX 专业的建造者建造具体的事物
func (d *Director) BuildX() {
	d.builder.Part1()
	d.builder.Part2()
	d.builder.Part3()
}
