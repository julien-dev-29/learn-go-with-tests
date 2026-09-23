package injection

import (
	"fmt"
	"io"
	"net/http"
)

const (
	greetPrefix = "yolo"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "%s %s", greetPrefix, name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "les kikis!")
}
