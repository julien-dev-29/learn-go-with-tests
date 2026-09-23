package injection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "les kikis!")

	got := buffer.String()
	want := "yolo les kikis!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
