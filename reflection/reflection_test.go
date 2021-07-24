package reflectioness

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

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
		{
			"Struct with non string field",
			struct {
				Name string
				Age  int
			}{
				"Jo",
				33,
			},
			[]string{"Jo"},
		},
		{
			"Nested fieds",
			Person{
				"Aek",
				Profile{33, "London"},
			},
			[]string{"Aek", "London"},
		},
		{
			"Pointers to things",
			&Person{
				"Po",
				Profile{25, "Bangkok"},
			},
			[]string{"Po", "Bangkok"},
		},
		{
			"Slices",
			[]Profile{
				{33, "London"},
				{34, "New York"},
			},
			[]string{"London", "New York"},
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
