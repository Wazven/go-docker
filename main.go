package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello Aldi")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}