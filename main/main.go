package main

import (
	"fmt"
	"net/http"

	"github.com/SUMUKHA-PK/Mail-Server/authorisation"
	"github.com/SUMUKHA-PK/Mail-Server/routeHandlers"
)

func main() {
	var dbPass string
	fmt.Print("Enter the DB password: ")
	fmt.Scan(&dbPass)
	authorisation.GetPass(dbPass)

	go func() {
		mux := &http.ServeMux{}
		mux.HandleFunc("/", routeHandlers.HandlerFunc)
		fmt.Print("Web server Serving on port 3000\n")
		http.ListenAndServe("127.0.0.1:3000", mux)
	}()

	fmt.Scanln()
	fmt.Scanln()
	fmt.Print("Server shut down")
}
