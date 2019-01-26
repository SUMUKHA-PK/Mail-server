package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

var receive string

func main() {

	// Username
	fmt.Printf("Enter the userame: ")
	var ip string
	fmt.Scan(&ip)

	//Backend server is always hosted here
	conn, err := net.Dial("tcp", "127.0.0.1:2345")
	if err != nil {
		fmt.Print(err)
	} else {
		log.Print("Connected")
	}

	// Send the query to the root server
	line := ip
	scanner := bufio.NewScanner(strings.NewReader(line))
	for scanner.Scan() {
		text := scanner.Text()
		_, errWrite := fmt.Fprintf(conn, text+"\n")
		if errWrite != nil {
			fmt.Print(err)
		}
		log.Print("IP sent to server: " + text)
		break
	}

	// Receive mapping from the same connection
	scanner = bufio.NewScanner(conn)
	for scanner.Scan() {
		receive = scanner.Text()
		fmt.Printf("Mapping received: " + receive + "\n")

		break
	}
	if errReadConn := scanner.Err(); errReadConn != nil {
		fmt.Print(errReadConn)
		return
	}

	fmt.Printf("Enter the password: ")
	fmt.Scan(&ip)

	// Send the query to the root server
	line = ip
	scanner = bufio.NewScanner(strings.NewReader(line))
	fmt.Print("Client message: ")
	for scanner.Scan() {
		text := scanner.Text()
		_, errWrite := fmt.Fprintf(conn, text+"\n")
		if errWrite != nil {
			fmt.Print(err)
		}
		log.Print("IP sent to server: " + text)
		break
	}

	// Receive mapping from the same connection
	scanner = bufio.NewScanner(conn)
	for scanner.Scan() {
		receive = scanner.Text()
		fmt.Printf("Mapping received: " + receive + "\n")

		break
	}
	if errReadConn := scanner.Err(); errReadConn != nil {
		fmt.Print(errReadConn)
		return
	}

}
