package adapter

// Target 适配的目标接口
type Target interface {
	Request() string
}

// Adaptee 适配者接口，表示被适配的目标接口
type Adaptee interface {
	SpecificRequest() string
}

// NewAdaptee 适配者接口的工厂函数
func NewAdaptee() Adaptee {
	return &adapteeImpl{}
}

// 适配者接口实现实例
type adapteeImpl struct{}

func (a *adapteeImpl) SpecificRequest() string {
	return "adaptee method"
}

// 适配器实现，转换Adaptee为Target接口的适配器
type adapter struct {
	Adaptee
}

// Request 适配器实现 Target 接口
func (a *adapter) Request() string {
	return a.SpecificRequest()
}

// NewAdapter Adapter 适配器的工厂函数
func NewAdapter(adaptee Adaptee) Target {
	return &adapter{
		Adaptee: adaptee,
	}
}
