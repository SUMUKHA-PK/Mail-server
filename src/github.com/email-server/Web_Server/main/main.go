package main

import (
	"fmt"
	"net/http"

	"../authorisation"
	"../routeHandlers"
)

func main() {
	var dbPass string
	fmt.Print("Enter the DB password: ")
	fmt.Scan(&dbPass)
	authorisation.GetPass(dbPass)

	go func() {
		mux := &http.ServeMux{}
		mux.HandleFunc("/", routeHandlers.HandlerFunc) // Path and function to go to // Path matching
		fmt.Print("Web server Serving on port 3000\n")
		http.ListenAndServe("192.168.43.10:3000", mux)
	}()

	fmt.Scanln()
	fmt.Scanln()
	fmt.Print("Server shut down")
}
