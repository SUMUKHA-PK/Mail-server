package main

import (
	"fmt"
	"net/http"

	"../authentication"

	"../routeHandlers"
)


func main() {
	// router := httprouter.New()
	go func() {
		authentication.Authentication()
	}()
	go func() {
		mux := &http.ServeMux{}
		mux.HandleFunc("/", routeHandlers.HandlerFunc) // Path and function to go to // Path matching
		fmt.Print("Serving on port 3000")
		http.ListenAndServe(":3000", mux)
	}()

	fmt.Scanln()
	fmt.Print("Server shut down")
}
