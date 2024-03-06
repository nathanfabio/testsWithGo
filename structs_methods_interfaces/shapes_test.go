package structsmethodsinterfaces

import "testing"


func TestPerimeter(t *testing.T) {
	rectangle:= Rectangle{10.0, 10.0}
	result:= Perimeter(rectangle)
	expected:= 40.0

	if result != expected {
		t.Errorf("got %.2f want %.2f", result, expected)
	}
}


func TestArea(t *testing.T) {
	checkArea:= func(t *testing.T, shape Shape, expected float64) {
		t.Helper()
		result:= shape.Area()

		if result != expected {
			t.Errorf("result %.2f, expected %.2f", result, expected)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle:= Rectangle{20.0, 10.0}
		checkArea(t, rectangle, 200.0)
	})

	t.Run("circles", func(t *testing.T) {
		circle:= Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})
}