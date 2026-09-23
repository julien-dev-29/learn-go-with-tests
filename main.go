package main

import (
	"log"
	"net/http"

	"example.com/hello/injection"
)

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(injection.MyGreeterHandler)))
}
