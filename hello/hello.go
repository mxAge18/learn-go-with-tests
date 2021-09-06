package main

import (
	"fmt"

	"github.com/mxage18/learn-go-with-tests/hello/entity"
)

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "Spanish":
		prefix = entity.SpanishHelloPrefix
	case "English":
		prefix = entity.EnglishHelloPrefix
	default:
		prefix = entity.EnglishHelloPrefix
	}
	return
}

func Hello(name string, language string) string {
	if name == "" {
        name = entity.DefaultHelloSuffix
    }
	return greetingPrefix(language) + name
}
func main() {
	fmt.Println(Hello("world", "English"))
}