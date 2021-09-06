package main

import (
	"testing"

	"github.com/mxage18/learn-go-with-tests/hello/entity"
)


func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	}
	t.Run("say hello to people", func(t *testing.T) {
		got := Hello("maxu", "")
		want := entity.EnglishHelloPrefix + "maxu"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello world when empty string", func(t *testing.T) {
		got := Hello("", "")
		want := entity.EnglishHelloPrefix + entity.DefaultHelloSuffix
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := entity.SpanishHelloPrefix + "Elodie"
		assertCorrectMessage(t, got, want)
	})
	
}
