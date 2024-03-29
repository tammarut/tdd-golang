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
		{
			"Array",
			[2]Profile{
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

	t.Run("Maps", func(t *testing.T) {
		inputMap := map[string]string{
			"key1": "value1",
			"key2": "value2",
		}

		var got []string
		walk(inputMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "value1")
		assertContains(t, got, "value2")
	})

	t.Run("Channel", func(t *testing.T) {
		channel := make(chan Profile)

		go func() {
			channel <- Profile{33, "John"}
			channel <- Profile{34, "Ben"}
			close(channel)
		}()

		var got []string
		want := []string{"John", "Ben"}

		walk(channel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("Function", func(t *testing.T) {
		function := func() (Profile, Profile) {
			return Profile{33, "Ki"}, Profile{34, "Bom"}
		}

		var got []string
		want := []string{"Ki", "Bom"}

		walk(function, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("want %q, but got %q", want, got)
		}
	})

}

func assertContains(t testing.TB, got []string, needle string) {
	t.Helper()
	contains := false
	for _, value := range got {
		if value == needle {
			contains = true
		}
	}

	if !contains {
		t.Errorf("expected %+v to contain %q but it didn't", got, needle)
	}
}
