package authentication

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	_ "../../mysql"
	"../errorHandler"
)

var DBPass string

func GetPass(dbPass string) {
	DBPass = dbPass
}

func Authentication(username string, password string, decider int) int {

	// username := ""
	// password := ""
	// decider := ""

	// //Starting the backend server
	// link, err := net.Listen("tcp", "127.0.0.1:2345")
	// errorHandler.ErrorHandler(err)

	// for {
	// 	fmt.Print("Backend server active on port 2345\n")
	// 	conn, err := link.Accept()
	// 	errorHandler.ErrorHandler(err)

	// 	// Get data from client
	// 	decider = getData(conn)
	// 	sendData(conn, "success")
	// 	username = getData(conn)
	// 	sendData(conn, "success")
	// 	password = getData(conn)

	// 	if (username == "") || (password == "") {
	// 		sendData(conn, "invalid")
	// 		break
	// 	}

	// 	fmt.Printf("Decider is %s, Username is %s and password is %s\n", decider, username, password)

	// 	if decider == "0" {
	// 		if SignUpHelper(dbPass, username, password) == 1 {
	// 			fmt.Printf("Successfully Signed Up user %s!\n", username)
	// 		} else {
	// 			fmt.Print("Error in signing up\n")
	// 		}
	// 	} else if decider == "1" {
	// 		if LoginHelper(dbPass, username, password) == 1 {
	// 			fmt.Printf("Successfully Logged in user %s!\n", username)
	// 		} else {
	// 			fmt.Print("Error in logging in\n")
	// 		}
	// 	} else {
	// 		sendData(conn, "invalid")
	// 		break
	// 	}
	// }
	if decider == 0 {
		if SignUpHelper(DBPass, username, password) == 1 {
			return 1
		} else {
			return -1
		}
	} else if decider == 1 {
		if LoginHelper(DBPass, username, password) == 1 {
			return 2
		} else {
			return -2
		}
	} else {
		return 0
	}
}

func getData(conn io.ReadWriter) string {

	var receive string

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		receive = scanner.Text()
		break
	}
	if errReadConn := scanner.Err(); errReadConn != nil {
		panic(errReadConn)
	}

	return receive
}

func sendData(conn io.ReadWriter, data string) {

	scanner := bufio.NewScanner(strings.NewReader("success"))

	for scanner.Scan() {
		text := scanner.Text()
		_, err := fmt.Fprintf(conn, text+"\n")
		errorHandler.ErrorHandler(err)
		break
	}
}
