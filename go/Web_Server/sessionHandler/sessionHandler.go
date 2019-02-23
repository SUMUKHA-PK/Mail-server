package sessionHandler

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"../../sessions"
	"../DB"
	"../errorHandler"
)

type Email struct {
	Body      string
	From_addr string
	To_addr   string
	Inbox     int
	Sent      int
}

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

func SessionHandlerNew(w http.ResponseWriter, r *http.Request, username string) {

	var template *template.Template
	template, err := template.ParseGlob("../webpages/static/*.html")

	errorHandler.ErrorHandler(err)

	var from_addr string
	var to_addr string
	var body string
	var inbox int
	var sent int

	var emails []Email

	rows := DB.GetEmails(username, "1")

	for rows.Next() {
		err = rows.Scan(&body, &from_addr, &to_addr, &inbox, &sent)
		if err != nil {
			log.Println(err)
			http.Error(w, "there was an error", http.StatusInternalServerError)
			return
		}

		emails = append(emails, Email{Body: body, From_addr: from_addr, To_addr: to_addr, Inbox: inbox, Sent: sent})
	}
	template.ExecuteTemplate(w, "loggedin.html", emails)
}
