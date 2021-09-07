package iteration

import (
	"fmt"
	"testing"
)

func ExampleRepeat() {
	repeatRes := Repeat("b",2)
	fmt.Println(repeatRes)
	// Output: bb
}

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 4)
	expected := "aaaa"

	if repeated != expected {
		t.Errorf("expected '%s' got '%s'", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	// TODO: Initialize
	for i := 0; i < b.N; i++ {
		Repeat("a", 2)
	}
}


//TODO: strings pkg进行一些学习和测试

