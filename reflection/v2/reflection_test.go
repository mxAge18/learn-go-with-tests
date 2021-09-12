package v2

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
	cases := []struct{
		Name string
		Input interface{}
		ExpectedCalls []string
	} {
		{
			Name: "Struct with one string field",
			Input: struct{
				Name string
			}{ "Chris"},
			ExpectedCalls: []string{"Chris"},
		},
		{
			Name: "Struct with multi string field",
			Input: struct{
				Name string
				Addr string
			}{ "maxu", "beijing 1001"},
			ExpectedCalls: []string{"maxu", "beijing 1001"},
		},
		{
			Name: "Struct with multi string field and int",
			Input: struct{
				Name string
				Addr string
				Age int
			}{ "maxu", "beijing 1001", 23},
			ExpectedCalls: []string{"maxu", "beijing 1001"},
		},
		{
			Name: "Struct with string field and struct",
			Input: Person{ "maxu", Profile{Age:23, City:"London"}},
			ExpectedCalls: []string{"maxu", "London"},
		},
		{
			"Pointers to things",
			&Person{
				"Chris",
				Profile{33, "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			"Slices",
			[]Profile {
				{33, "London"},
				{34, "Reykjavík"},
			},
			[]string{"London", "Reykjavík"},
		},
		{
			"Arrays",
			[2]Profile {
				{33, "London"},
				{34, "Reykjavík"},
			},
			[]string{"London", "Reykjavík"},
		},
	}
	for _, tt := range cases {
		t.Run(tt.Name, func(t *testing.T) {
			var got []string
			Walk(tt.Input, func(s string) {
				got = append(got, s)
			})
			if !reflect.DeepEqual(got, tt.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, tt.ExpectedCalls)
			}
		})
	}
	t.Run("with map", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}
	
		var got []string
		Walk(aMap, func(input string) {
			got = append(got, input)
		})
		assertContains(t, got, "Bar")
		assertContains(t, got, "Boz")
	
	})

}
func assertContains(t *testing.T, haystack []string, containStr string) {
	contains := false
	for _, val := range haystack {
		if val == containStr {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %+v to contain '%s' but it didnt", haystack, containStr)
	}
}