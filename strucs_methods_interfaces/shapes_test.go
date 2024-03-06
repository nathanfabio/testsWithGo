package strucsmethodsinterfaces


import "testing"

func TestPerimeter(t *testing.T) {
	result:= Perimeter(10.0, 10.0)
	expected:= 40.0


	if result != expected {
		t.Errorf("got %.2f want %.2f", result, expected)
	}
}


func TestArea(t *testing.T) {
	result:= Area(20.0, 10.0)
	expected:= 200.0

	if result != expected {
		t.Errorf("got %.2f want %.2f", result, expected)
	}
}