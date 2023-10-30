package proxy

// Subject 接口定义 RealSubject, Proxy 的共同行为
type Subject interface {
	Do() string
}

// RealSubject 真实对象，实现了 Subject 接口
type RealSubject struct{}

func (r *RealSubject) Do() string {
	return "real"
}

// Proxy 代理对象，持有一个RealSubject实例的引用
type Proxy struct {
	real RealSubject
}

func (p *Proxy) Do() string {
	// 	在调用RealSubject之前或之后可以执行一些额外的操作
	var result string

	// 在调用真实对象之前的工作，检查缓存，判断权限，实例化真实对象等
	result += "before:"
	// 调用真实对象
	result += p.real.Do()
	// 调用之后的操作，如缓存结果，对结果进行处理等
	result += ":after"
	return result
}
