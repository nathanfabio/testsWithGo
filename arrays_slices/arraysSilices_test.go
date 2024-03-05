package arraysslices

import "testing"

func TestSum(t *testing.T) {

	t.Run("collections of any size", func (t *testing.T)  {
		numbers:= []int{1, 2, 3, 4}

		result:= Sum(numbers)
		expected:= 10

		if result != expected {
			t.Errorf("got %d want %d given, %v", result, expected, numbers)
			
		}
	})
}