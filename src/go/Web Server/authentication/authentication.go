package authentication

import (
	"bufio"
	"fmt"
	"net"
	"strings"

	"../errorHandler"
)

func Authentication() {

	username := ""
	password := ""
	ack := 1

	//Starting the backend server
	link, err := net.Listen("tcp", "127.0.0.1:2345")
	errorHandler.ErrorHandler(err)

	for {
		fmt.Print("Backend server active on port 2345\n")
		conn, err := link.Accept()
		errorHandler.ErrorHandler(err)

		// Get data from the client

		scanner := bufio.NewScanner(conn)

		for scanner.Scan() {
			receive := scanner.Text()
			username = receive
			break
		}
		if errReadConn := scanner.Err(); errReadConn != nil {
			panic(errReadConn)
		}

		scanner = bufio.NewScanner(strings.NewReader("success"))

		for scanner.Scan() {
			text := scanner.Text()
			_, err := fmt.Fprintf(conn, text+"\n")
			errorHandler.ErrorHandler(err)
			break
		}

		scanner = bufio.NewScanner(conn)

		for scanner.Scan() {
			receive := scanner.Text()
			password = receive
			break
		}
		if errReadConn := scanner.Err(); errReadConn != nil {
			panic(errReadConn)
		}

		if (username == "") || (password == "") {
			ack = 0
		}

		scanner = bufio.NewScanner(strings.NewReader(string(ack)))

		for scanner.Scan() {
			text := scanner.Text()
			_, err := fmt.Fprintf(conn, text+"\n")
			errorHandler.ErrorHandler(err)
			break
		}
		fmt.Printf("Username is %s and password is %s\n", username, password)
	}
}
