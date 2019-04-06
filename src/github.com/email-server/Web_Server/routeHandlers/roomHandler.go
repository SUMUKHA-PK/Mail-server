package routeHandlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/email-server/Web_Server/DB"
	"github.com/email-server/Web_Server/errorHandler"
	"github.com/email-server/Web_Server/util"
)

func RoomCreationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		util.RenderPage(w, "../webpages/static/createRoom.html")
	} else if r.Method == "POST" {
		r.ParseForm()

		Room.RoomName = util.GetString(r.Form["roomName"])
		Room.Admins = util.GetStringAndAppend(Room.Admins, r.Form["admins"])
		Room.Members = util.GetStringAndAppend(Room.Members, r.Form["users"])

		err := DB.CreateRoom(Room)
		if err != nil {
			util.RenderPage(w, "../webpages/static/createRoomErr.html")
		} else {
			fmt.Println(Room)
			util.RenderPage(w, "../webpages/static/room.html")
		}
	}
}

func RoomHandler(w http.ResponseWriter, r *http.Request, user []util.UserData) {
	r.ParseForm()

	email_body := util.GetString(r.Form["body"])

	DB.AddDataToDB(Room.RoomName, user[0].UserName, email_body)

	renderRoomData(w, r, Room.RoomName, user)
}

type Emails struct {
	From_addr string
	Body      string
}

type RoomsUser struct {
	Rooms string
}

type Rooms struct {
	RoomName string
	Members  string
	Admins   string
}

func renderRoomData(w http.ResponseWriter, r *http.Request, roomName string, user []util.UserData) {
	var template *template.Template
	template, err := template.ParseGlob("../webpages/static/*.html")

	errorHandler.ErrorHandler(err)

	var from_addr string
	var body string

	var emails []Emails

	rows := DB.GetRoomData(roomName)

	for rows.Next() {
		err = rows.Scan(&body, &from_addr)
		if err != nil {
			log.Println(err)
			http.Error(w, "there was an error", http.StatusInternalServerError)
			return
		}
		emails = append(emails, Emails{Body: body, From_addr: from_addr})
	}
	fmt.Println(emails)
	if DB.CheckAdmin(user, roomName) {
		template.ExecuteTemplate(w, "room.html", emails)
	} else {
		template.ExecuteTemplate(w, "room_na.html", emails)
	}

}

func RenderRoomChoicePage(w http.ResponseWriter, r *http.Request, user []util.UserData) {
	var template *template.Template
	template, err := template.ParseGlob("../webpages/static/*.html")

	errorHandler.ErrorHandler(err)

	var rooms []RoomsUser

	var rName string
	var uName string
	var admin string

	rows := DB.GetRoomsUser(user[0].UserName)

	for rows.Next() {
		err = rows.Scan(&rName, &uName, &admin)
		if err != nil {
			log.Println(err)
			http.Error(w, "there was an error", http.StatusInternalServerError)
			return
		}
		rooms = append(rooms, RoomsUser{Rooms: rName})
	}
	fmt.Println(rooms)
	template.ExecuteTemplate(w, "rooms.html", rooms)
}

func RenderUserRoom(w http.ResponseWriter, r *http.Request, user []util.UserData) {
	r.ParseForm()

	roomName := util.GetString(r.Form["roomName"])

	var members string
	var room string
	var admins string
	var rooms []Rooms

	rows, err := DB.GetUserRoomData(roomName)

	if err != nil {
		// Unauthorised room access/something
	}

	for rows.Next() {
		err = rows.Scan(&room, &members, &admins)
		if err != nil {
			log.Println(err)
			http.Error(w, "there was an error", http.StatusInternalServerError)
			return
		}

		rooms = append(rooms, Rooms{RoomName: room, Members: members, Admins: admins})
	}
	fmt.Println(rooms)

	Room.RoomName = roomName
	renderRoomData(w, r, roomName, user)
}
