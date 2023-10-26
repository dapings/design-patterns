package builder

type Builder1 struct {
	// 可以是一个具体的产品，如房子
	result string
}

func (b *Builder1) Result() string {
	// 获取具体的产品信息
	return b.result
}

func (b *Builder1) Part1() {
	b.result += "1"
}

func (b *Builder1) Part2() {
	b.result += "2"
}

func (b *Builder1) Part3() {
	b.result += "3"
}

type Builder2 struct {
	result int
}

func (b *Builder2) Result() int {
	return b.result
}

func (b *Builder2) Part1() {
	b.result += 1
}

func (b *Builder2) Part2() {
	b.result += 2
}

func (b *Builder2) Part3() {
	b.result += 3
}
