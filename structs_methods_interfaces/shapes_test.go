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
	t.Run("rectangles", func(t *testing.T) {
		rectangle:= Rectangle{20.0, 10.0}
		result:= rectangle.Area()
		expected:= 200.0

		if result != expected {
			t.Errorf("got %.2f want %.2f", result, expected)
		}
	})

	t.Run("circles", func(t *testing.T)  {
		circle:= Circle{5}
		result:= circle.Area()
		expected:= 78.53981633974483

		if result != expected {
			t.Errorf("got %.2f want %.2f", result, expected)
		}
	})
}