package reflectioness

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with a string field",
			struct{ Name string }{
				"Chris Array",
			},
			[]string{"Chris Array"},
		},
		{
			"Struct with 2 string fields",
			struct {
				Name string
				City string
			}{
				"Chris Array",
				"London",
			},
			[]string{"Chris Array", "London"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("want %q, but got %q", test.ExpectedCalls, got)
			}

		})
	}
}
