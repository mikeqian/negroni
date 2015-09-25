package main

import (
	"fmt"
	"github.com/mikeqian/negroni"
	"net/http"
)

func main() {
	n := negroni.Classic()

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Welcome to the home page!")
	})

	n.UseHandler(mux)
	n.Run(":3000")
}
