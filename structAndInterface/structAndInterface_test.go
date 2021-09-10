package structAndInterface

import "testing"

func TestPerimeter(t *testing.T) {
	perimeterTests := []struct{
		shape Shape
		want float64
	} {
		{shape: Rectangle{Height: 10.0,Width: 10.0}, want: 40.0},
	}
	for _, tt := range perimeterTests {
		got := tt.shape.Perimeter()
		if tt.want != got {
			t.Errorf("got %.2f , want %.2f", got, tt.want)
		}
	}
}

func TestArea(t *testing.T) {
	/*checkArea := func(t *testing.T, shape Shape,want float64) {
		t.Helper()
		got := shape.Area()
		if want != got {
			t.Errorf("got %.2f , want %.2f", got, want)
		}
	}
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{5.0, 4.0}
		want := 20.0
		checkArea(t, rectangle, want)
	})
	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		want := 314.1592653589793
		checkArea(t, circle, want)
	})*/
	// 表格驱动测试在我们要创建一系列相同测试方式的测试用例时很有用。
	/* 我们可以像遍历任何其他切片一样来遍历这个数组，进而用这个结构体的域来做我们的测试。
		你会看到开发人员能方便的引入一个新的几何形状，只需实现 Area 方法并把新的类型加到测试用例中。
		另外发现 Area 方法有错误，我们可以在修复这个错误之前非常容易的添加新的测试用例。
		列表驱动测试可以成为你工具箱中的得力武器。但是确保你在测试中真的需要使用它。
		如果你要测试一个接口的不同实现，或者传入函数的数据有很多不同的测试需求，这个武器将非常给力。
		*/
	areaTests := []struct {
		name string
		shape Shape
		hasArea float64
	} {
		{name: "Rectangle", shape: Rectangle{Width: 5.0,Height: 4.0}, hasArea: 20.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}
	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %.2f want %.2f.",tt.shape, got, tt.hasArea)
			}
		})
	}
}