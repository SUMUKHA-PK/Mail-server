/*
This contains all the utility functions necesssary.

This is done to remove recursive imports mainly.
*/
package util

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"net/http"
	"os"
	"strings"

	"github.com/SUMUKHA-PK/Mail-Server/errorHandler"
)

/*
Has all the data a session of an User must have.
*/

type UserData struct {
	LoggedIn string
	UserName string
	Password string
	PhoneNo  string
	OTP      string
	ID       string
	Auth     bool
	Rooms    []RoomData
	Admin    bool
}

type RoomData struct {
	RoomName string
	Members  []string
	RoomID   string
	Admins   []string
}

/*
RenderPage is used to Render any webpage when called
*/

func RenderPage(w http.ResponseWriter, pageName string) {

	f, err := os.Open(pageName)
	errorHandler.ErrorHandler(err)
	b1 := make([]byte, 100000)
	_, err = f.Read(b1)
	errorHandler.ErrorHandler(err)
	fmt.Fprintf(w, string(b1))
}

/*
GetString returns a string type object from a []string type.

Used in routeHandlers for getting FORM data
*/

func GetString(input []string) string {

	result := ""
	for i := 0; i < len(input); i++ {
		result = result + input[i]
	}
	return result
}

/*
GetStringArr gets the data from []string type , converts them based on []string is needed or string

Used in routeHandler to obtain COMPOSE data
*/

func GetStringArr(input []string, decider int) []string {

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

/*
SendData and GetData are utility functions used to move data through sockets from a pre-existing connection

Still work under progress
*/

func SendData(conn net.Conn, message []string) {

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

func GetData(conn net.Conn) string {
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

// func GenerateRandomBytes(n int) ([]byte, error) {
// 	b := make([]byte, n)
// 	_, err := rand.Read(b)
// 	// Note that err == nil only if we read len(b) bytes.
// 	if err != nil {
// 		return nil, err
// 	}

// 	return b, nil
// }

func RandomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func GenerateRandomString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(RandomInt(65, 90))
	}
	return string(bytes)
}

func GetStringAndAppend(arr []string, data []string) []string {
	var addr []string
	newData := GetStringArr(data, 1)
	for i := 0; i < len(newData); i++ {
		addr = append(addr, newData[i])
	}
	return addr
}
