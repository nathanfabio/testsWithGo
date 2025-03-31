package reflection

import "testing"

func TestBrowse(t *testing.T) {
	expected := "Nathan"
	var result []string

	x := struct {
		Name string
	}{expected}

	Browse(x, func(input string) {
		result = append(result, input)
	})

	if result[0] != expected {
		t.Errorf("result '%s, expected '%s", result[0], expected)
	}

	if len(result) != 1 {
		t.Errorf("incorrect number of function calls: result %d, expected %d", len(result), 1)
	}
}