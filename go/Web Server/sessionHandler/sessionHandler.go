package sessionHandler

import (
	"fmt"
	"net/http"
	"os"

	"../../sessions"
	"../errorHandler"
)

func SessionManager(session *sessions.Session, w http.ResponseWriter, r *http.Request) {

	sess_name := session.Values["ID"]

	var ID = string(sess_name.([]uint8))

	fmt.Print(ID)

}

func HandlerFunc(w http.ResponseWriter, r *http.Request) {
	// renderPage(w, "../webpages/static/loggedin.html")
	fmt.Print(r.URL.Path)
}

func renderPage(w http.ResponseWriter, pageName string) {

	f, err := os.Open(pageName)
	errorHandler.ErrorHandler(err)
	b1 := make([]byte, 100000)
	_, err = f.Read(b1)
	errorHandler.ErrorHandler(err)
	fmt.Fprintf(w, string(b1))
}
