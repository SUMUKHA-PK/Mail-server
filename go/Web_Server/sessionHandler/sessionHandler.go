/*
Here are functions to manage already created sessions.

This currently can get the email data and render pages on a template
*/

package sessionHandler

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

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

/*
SessionHandlerNew is currently a botched job.

It must be replaced with the above SessionManager and possibly a utility function that does the job of the below,
along with managing the current session(Time outs, give user data to the invoking functions like Authorisers)

*/

func SessionHandlerNew(w http.ResponseWriter, r *http.Request, username string, decider string) {

	var template *template.Template
	template, err := template.ParseGlob("../webpages/static/*.html")

	errorHandler.ErrorHandler(err)

	var from_addr string
	var to_addr string
	var body string
	var inbox int
	var sent int

	var emails []Email

	rows := DB.GetEmails(username, decider)

	for rows.Next() {
		err = rows.Scan(&body, &from_addr, &to_addr, &inbox, &sent)
		if err != nil {
			log.Println(err)
			http.Error(w, "there was an error", http.StatusInternalServerError)
			return
		}

		emails = append(emails, Email{Body: body, From_addr: from_addr, To_addr: to_addr, Inbox: inbox, Sent: sent})
	}
	if decider == "1" {
		template.ExecuteTemplate(w, "loggedin.html", emails)
	} else {
		template.ExecuteTemplate(w, "sentmail.html", emails)
	}
}
