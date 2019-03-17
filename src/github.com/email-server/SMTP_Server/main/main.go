package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"

	"github.com/email-server/Web_Server/errorHandler"
)

func main() {
	fmt.Printf("GO!")

	go func() {
		link, err := net.Listen("tcp", "10.53.70.237:3001")
		errorHandler.ErrorHandler(err)

		var d []string

		for {
			fmt.Print("SMTP server serving at port 3001\n")
			conn, err := link.Accept()
			errorHandler.ErrorHandler(err)

			c := make(chan string)

			go handleConnection(conn, c)
			d = append(d, <-c)

			fmt.Print(d)
		}
	}()

	fmt.Scanln()
	fmt.Println("done")
}

func handleConnection(conn net.Conn, c chan string) {

	fmt.Printf("Serving %s\n", conn.RemoteAddr().String())
	for {
		fmt.Print("WW")
		netData, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		temp := strings.TrimSpace(string(netData))
		if temp == "STOP" {
			break
		}
		fmt.Print(temp)
		c <- temp
		// sendData(conn, temp)
		// close(c)
	}
	close(c)
	conn.Close()
	// for i := 0; i < len(data); i++ {
	// 	sendData(conn, data[i])
	// }
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

func sendData(conn net.Conn, message string) {

	fmt.Print("ETF")
	scanner := bufio.NewScanner(strings.NewReader(message))

	for scanner.Scan() {
		text := scanner.Text()
		_, errWrite := fmt.Fprintf(conn, text+"\n")
		errorHandler.ErrorHandler(errWrite)
		break
	}
}
