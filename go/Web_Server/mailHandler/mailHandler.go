package mailHandler

import (
	"fmt"
	"net/http"
)

func ComposeHandler(w http.ResponseWriter, r *http.Request) {

	// Handling a POST request /compose

	// Parse the form
	r.ParseForm()

	from_mail := r.Form["from"]
	to_email := r.Form["to"]
	email_body := r.Form["body"]

	fmt.Print(from_mail)
	fmt.Print(to_email)
	fmt.Print(email_body)

}
