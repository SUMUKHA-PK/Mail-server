package routeHandlers

import (
	"os"
	"fmt"
	"net/http"
	"../errorHandler"
)


func renderPage(w http.ResponseWriter, pageName string) {

	f, err := os.Open(pageName)
	errorHandler.ErrorHandler(err)
	b1:= make([]byte, 100000)
	 n1, err := f.Read(b1)
	 errorHandler.ErrorHandler(err)
	fmt.Fprintf(w, string(b1))
	fmt.Printf("n1 = %d\n", n1)
	// check(err)
}


func HandlerFunc(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		renderPage(w, "../webpages/static/index.html")
		
	} else if r.URL.Path == "/login.html" {
		renderPage(w, "../webpages/authentication/login.html")

	} else {
		w.WriteHeader(http.StatusNotFound) // Status code 404
		fmt.Fprint(w, "<h1>Error 404 : Page not found</h1>")
	}
}


