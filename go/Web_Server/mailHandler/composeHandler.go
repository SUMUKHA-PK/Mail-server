package mailHandler

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"strings"

	"../errorHandler"
)

func ComposeHandler(w http.ResponseWriter, r *http.Request) [][]string {

	// Handling a POST request /compose

	// Parse the form
	r.ParseForm()

	from_mail := r.Form["from"]
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

	d = append(d, getString(from_mail, 0))
	d = append(d, getString(to_email, 1))
	d = append(d, getString(email_body, 0))

	// go sendData(conn, d)

	// var data []string

	// data = append(data, getData(conn))
	// data = append(data, getData(conn))
	// data = append(data, getData(conn))

	return d
}

func getString(input []string, decider int) []string {

	var data []string
	result := ""
	for i := 0; i < len(input); i++ {
		result = result + input[i]
	}
	if decider == 0 {
		data = append(data, result)
	} else {
		data = strings.Split(result, ";")
	}
	return data
}

func sendData(conn net.Conn, message []string) {

	for i := 0; i < len(message); i++ {
		scanner := bufio.NewScanner(strings.NewReader(message[i]))

		for scanner.Scan() {
			text := scanner.Text()
			fmt.Print(text)
			_, errWrite := fmt.Fprintf(conn, text+"\n")
			errorHandler.ErrorHandler(errWrite)
			break
		}
	}

}

func getData(conn net.Conn) string {
	var receive string
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		receive = scanner.Text()
		break
	}
	if errReadConn := scanner.Err(); errReadConn != nil {
		fmt.Print(errReadConn)
		return ""
	}
	return receive
}
