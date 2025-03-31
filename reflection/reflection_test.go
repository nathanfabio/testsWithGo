package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name string
	Profile Profile
}

type Profile struct {
	Age int
	City string
}

func TestBrowse(t *testing.T) {
	cases := []struct {
		Name string
		Input interface{}
		ExpectedCalls []string
	}{
		{
			"Nested fields",
			Person {
				"Nathan",
				Profile{25, "Santos"},
			},
			[]string{"Nathan", "Santos"},
		},
		{
			"Poiters to things",
			&Person{
				"Nathan",
				Profile{25, "Santos"},
			},
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