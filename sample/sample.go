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

<<<<<<< HEAD
	n.Use(mux)

=======
	n.UseHandler(mux)
>>>>>>> 4220a4435352515e238756de90d7fae69617052a
	n.Run(":3000")
}
