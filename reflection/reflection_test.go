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
		{
			"Slices",
			[]Profile{
				{25, "Bebedouro"},
				{26, "Santos"},
			},
			[]string{"Bebedouro", "Santos"},
		},
		{
			"Arrays",
			[2]Profile{
				{25, "Bebedouro"},
				{30, "Melbourne"},
			},
			[]string{"Bebedouro", "Melbourne"},
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


	t.Run("maps", func(t *testing.T) {
		mapA := map[string]string{
			"Bebedouro": "Santos",
			"Santos": "Melbourne",
		}

		var result []string
		Browse(mapA, func(input string) {
			result = append(result, input)
		})

		check(t, result, "Santos")
		check(t, result, "Melbourne")
	})
	
}

func check(t *testing.T, haystack []string, needle string) {
	contain := false
	for _, x := range haystack {
		if x == needle {
			contain = true
		}
	}

	if !contain {
		t.Errorf("'%v' was expected to contain '%s', but it didn't", haystack, needle)
	}
}