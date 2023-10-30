package flyweight

import (
	"fmt"
)

// Flyweight 接口
type Flyweight interface {
	Operation(extrinsicState int)
}

// FlyweightFactory 工厂实例，维护享元对象的池子
type FlyweightFactory struct {
	flyweights map[string]Flyweight
}

// GetFlyweight 获取享元对象
// 如果池子中已经存在对应的享元对象，则直接返回；否则，创建新的享元对象并加入池子。
func (f *FlyweightFactory) GetFlyweight(key string) Flyweight {
	if fw, ok := f.flyweights[key]; ok {
		return fw
	}

	fw := NewImageConcreteFlyweight(key)
	f.flyweights[key] = fw
	return fw
}

var imageFactory *FlyweightFactory

// GetImageFlyweightFactory 获取六图像享元工厂实例
func GetImageFlyweightFactory() *FlyweightFactory {
	if imageFactory == nil {
		imageFactory = &FlyweightFactory{flyweights: make(map[string]Flyweight)}
	}
	return imageFactory
}

// ImageConcreteFlyweight 具体的享元对象
type ImageConcreteFlyweight struct {
	intrinsicStat int
	data          string
}

func (i *ImageConcreteFlyweight) Data() string {
	return i.data
}

func (i *ImageConcreteFlyweight) Operation(extrinsicState int) {
	fmt.Printf("具体享元对象，内部状态： %d, 外部状态：%d\n", i.intrinsicStat, extrinsicState)
}

func NewImageConcreteFlyweight(fileName string) *ImageConcreteFlyweight {
	return &ImageConcreteFlyweight{data: fmt.Sprintf("image data %s", fileName)}
}

type ImageViewer struct {
	*ImageConcreteFlyweight
}

func NewImageViewer(fileName string) *ImageViewer {
	image := GetImageFlyweightFactory().GetFlyweight(fileName)
	return &ImageViewer{ImageConcreteFlyweight: image.(*ImageConcreteFlyweight)}
}

func (v *ImageViewer) Display(state int) {
	v.Operation(state)
	fmt.Printf("Display: %s\n", v.Data())
}
