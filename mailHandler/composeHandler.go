package mailHandler

import (
	"net/http"

	"github.com/SUMUKHA-PK/Mail-Server/util"
)

func ComposeHandler(w http.ResponseWriter, r *http.Request) [][]string {

	// Handling a POST request /compose

	// Parse the form
	r.ParseForm()

	to_email := r.Form["to"]
	email_body := r.Form["body"]

	// sendEmail.SendEmail(getString(from_mail), getString(to_email), getString(email_body))

	// Sending data the SMTP Server
	// Later this will be done by send mail function
	// addr := "10.53.70.237:3001"
	// err := smtp.SendMail(addr, nil, from, to, []byte(msg))
	// errorHandler.ErrorHandler(err)

	// conn, err := net.Dial("tcp", addr)
	// errorHandler.ErrorHandler(err)

	var d [][]string

	d = append(d, util.GetStringArr(to_email, 1))
	d = append(d, util.GetStringArr(email_body, 0))

	// go sendData(conn, d)

	// var data []string

	// data = append(data, getData(conn))
	// data = append(data, getData(conn))
	// data = append(data, getData(conn))

	return d
}
