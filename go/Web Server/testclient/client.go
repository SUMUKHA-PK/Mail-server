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

	var decider string

	fmt.Print("Signup? 0, Login? 1: ")
	fmt.Scan(&decider)
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

	scanner := bufio.NewScanner(strings.NewReader(decider))
	fmt.Print("Client message: ")
	for scanner.Scan() {
		text := scanner.Text()
		_, errWrite := fmt.Fprintf(conn, text+"\n")
		if errWrite != nil {
			fmt.Print(err)
		}
		log.Print("Decider: " + text)
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

	// Send the query to the root server
	line := ip
	scanner = bufio.NewScanner(strings.NewReader(line))
	for scanner.Scan() {
		text := scanner.Text()
		_, errWrite := fmt.Fprintf(conn, text+"\n")
		if errWrite != nil {
			fmt.Print(err)
		}
		log.Print("Username sent to server: " + text)
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
		log.Print("Password sent to server: " + text)
		break
	}
}
