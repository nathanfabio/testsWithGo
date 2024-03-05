package arraysslices

import (
	"reflect"
	"testing"
)

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

// func TestSumAll(t *testing.T) {
// 	result:= SumAll([]int{2, 3}, []int{3, 3})
// 	expected:= []int{5, 6}

// 	if !reflect.DeepEqual(result, expected) {
// 		t.Errorf("got %v want %v", result, expected)
// 	}
// }

func TestSumAllTheRest(t *testing.T) {
	checkSums:= func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}


	t.Run("sum up some slices", func(t *testing.T){
		result:= SumAllTheRest([]int{1, 4}, []int{0, 7})
		expected:= []int{4, 7}

		checkSums(t, result, expected)
	})


	t.Run("sum up empty slices", func(t *testing.T){
		result:= SumAllTheRest([]int{}, []int{3, 8, 2})
		expected:= []int{0, 10}

		checkSums(t, result, expected)
	})
}