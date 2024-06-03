package strings

import (
	"testing"
)

func TestReplaceChar(t *testing.T) {
	s := "a, b, c"
	want := "b, b, c"
	got := ReplaceChar(s, "a", "b")

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
