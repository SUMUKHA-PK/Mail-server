package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Bitch</h1>")
	} else if r.URL.Path == "/contact" {
		fmt.Fprint(w, "<h1>Hello world</h1>")
	} else {
		w.WriteHeader(http.StatusNotFound) // Status code 404
		fmt.Fprint(w, "<h1>Error 404 : Page not found</h1>")
	}
}

func main() {
	// router := httprouter.New()
	mux := &http.ServeMux{}
	mux.HandleFunc("/", handlerFunc) // Path and function to go to // Path matching
	fmt.Print("Serving on port 3000")
	http.ListenAndServe(":3000", mux)
}
