package flyweight

import (
	"testing"
)

func ExampleFlyweight() {
	viewer := NewImageViewer("image1.png")
	viewer.Display(1)
	// Output:
	// 具体享元对象，内部状态： 0, 外部状态：1
	// Display: image data image1.png
}

func TestFlyweight(t *testing.T) {
	viewer1 := NewImageViewer("image1.png")
	viewer2 := NewImageViewer("image1.png")

	if viewer2.ImageConcreteFlyweight != viewer1.ImageConcreteFlyweight {
		t.Fail()
	}
}
