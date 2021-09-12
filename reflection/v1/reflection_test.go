package v1
import "testing"

func TestWalk(t *testing.T) {
	expected := "maxu"
	var got []string
	x := struct {
		name string
	}{expected}
	Walk(x, func(input string) {
		got = append(got, input)
	})
	if got[0] != expected {
		t.Errorf("got '%s', want '%s'", got[0], expected)
	}
	if len(got) != 1 {
        t.Errorf("wrong number of function calls, got %d want %d", len(got), 1)
    }
}
