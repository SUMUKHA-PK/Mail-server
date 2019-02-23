package main

import (
	"bufio"
	"fmt"
	"net"

	"../../Web_Server/errorHandler"
)

func main() {
	fmt.Printf("GO!")

	var receive string
	go func() {
		link, err := net.Listen("tcp", "127.0.0.1:3001")
		errorHandler.ErrorHandler(err)

		for {
			fmt.Print("SMTP server serving at port 3001\n")
			conn, err := link.Accept()
			errorHandler.ErrorHandler(err)

			scanner := bufio.NewScanner(conn)
			for scanner.Scan() {
				receive = scanner.Text()
				fmt.Printf("Root server: ")
				fmt.Printf("IP received to map from client this: " + receive + "\n")
				break
			}
			if errReadConn := scanner.Err(); errReadConn != nil {
				fmt.Print(errReadConn)
				fmt.Printf("Root server: ")
				return
			}

			//Get the result of the mapping from the servers
			// result := start_servers(receive, IP_List_Name, IP_List_Addr)

			//Communicate back the result to the client on the same connection
			// scanner = bufio.NewScanner(strings.NewReader(result))

			// for scanner.Scan() {
			// 	text := scanner.Text()
			// 	_, err := fmt.Fprintf(conn, text+"\n")
			// 	if err != nil {
			// 		//Error exists due to sending in same connection, figure it out
			// 	}
			// 	fmt.Printf("Root server: ")
			// 	log.Print("Query mapping sent: " + text)
			// 	break
			// }
		}
	}()
}
