package helloworld

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying yolo to the kikis", func(t *testing.T) {
		got := Hello("les kikis!", "")
		want := "Yolo les kikis!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("saying yolo to the kikis with an empty string", func(t *testing.T) {
		got := Hello("", "")
		want := "Yolo les kikis!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("saying yolo to the kikis in spanish", func(t *testing.T) {
		got := Hello("les kikis!", "Spanish")
		want := "Hola les kikis!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("saying yolo to the kikis in french", func(t *testing.T) {
		got := Hello("les kikis!", "French")
		want := "Coucou les kikis!"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
