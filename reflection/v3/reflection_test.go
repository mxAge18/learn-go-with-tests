package v3

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
		{
			"Maps",
			map[string]string{
				"Foo": "Bar",
				"Baz": "Boz",
			},
			[]string{"Bar", "Boz"},
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
}
