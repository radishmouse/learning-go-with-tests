package reflection

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	// For my table based test, I care about:
	// - a name for my test
	// - the input
	// - the expected output
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			// just a name for my test
			"struct with one string field",
			// what am I passing into the function I'm testing?
			// It's just a struct with a .Name field
			// and the value of that field is "Chris"
			struct {
				Name string
			}{"Chris"},
			// the result should be a 1 element slice of strings
			[]string{"Chris"},
		},
		{
			// second test case
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		},
		{
			"struct with non string fields",
			struct {
				Name string
				Age  int
			}{"Chris", 44},
			[]string{"Chris"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})
			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}

}
