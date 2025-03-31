package reflection

import (
	"reflect"
	"testing"
)


func TestBrowse(t *testing.T) {
	cases := []struct {
		Name string
		Input interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with a string",
			struct {
				Name string
			}{"Nathan"},
			[]string{"Nathan"},
		},
		{
			"Struct with two string type",
			struct {
				Name string
				City string
			}{"Nathan", "Santos"},
			[]string{"Nathan", "Santos"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var result []string
			Browse(test.Input, func(input string) {
				result = append(result, input)
			})
			if !reflect.DeepEqual(result, test.ExpectedCalls) {
				t.Errorf("result %v, expected %v", result, test.ExpectedCalls)
			}
		})
	}
	
}