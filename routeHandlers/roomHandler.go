package routeHandlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/SUMUKHA-PK/Mail-Server/DataBase"
	"github.com/SUMUKHA-PK/Mail-Server/errorHandler"
	"github.com/SUMUKHA-PK/Mail-Server/sessionHandler"
	"github.com/SUMUKHA-PK/Mail-Server/util"
)

func RoomCreationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		util.RenderPage(w, "webpages/static/createRoom.html")
	} else if r.Method == "POST" {
		r.ParseForm()

		Room.RoomName = util.GetString(r.Form["roomName"])
		Room.Admins = util.GetStringAndAppend(Room.Admins, r.Form["admins"])
		Room.Members = util.GetStringAndAppend(Room.Members, r.Form["users"])

		err := DataBase.CreateRoom(Room)
		if err != nil {
			util.RenderPage(w, "webpages/static/createRoomErr.html")
		} else {
			fmt.Println(Room)
			util.RenderPage(w, "webpages/static/room.html")
		}
	}
}

func RoomHandler(w http.ResponseWriter, r *http.Request) {
	user, val := sessionHandler.CheckActiveSession(r)
	if val {
		log.Print("Routed to room page\n")
		RenderRoomData(w, r, user)
	} else {

	}
}

func RoomsHandler(w http.ResponseWriter, r *http.Request) {
	user, val := sessionHandler.CheckActiveSession(r)
	if val {
		log.Print("Routed to room page\n")
		RenderRoomChoicePage(w, r, user, 0)
	} else {

	}
}

func UserRoomHandler(w http.ResponseWriter, r *http.Request) {
	user, val := sessionHandler.CheckActiveSession(r)
	if val {
		log.Print("Routed to user room page\n")
		RenderUserRoom(w, r, user)
	}
}

func RenderRoomData(w http.ResponseWriter, r *http.Request, user []util.UserData) {
	r.ParseForm()

	email_body := util.GetString(r.Form["body"])

	err := DataBase.AddDataToDB(Room.RoomName, user[0].UserName, email_body)

	if err == nil {
		renderRoomData(w, r, Room.RoomName, user)
	} else {
		fmt.Println("WWW")
	}
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
	template, err := template.ParseGlob("webpages/static/*.html")

	errorHandler.ErrorHandler(err)

	var from_addr string
	var body string

	var emails []Emails

	rows, err := DataBase.GetRoomData(roomName)

	if err == nil {
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
		if DataBase.CheckAdmin(user, roomName) {
			template.ExecuteTemplate(w, "room.html", emails)
		} else {
			template.ExecuteTemplate(w, "room_na.html", emails)
		}
	} else {
		// util.RenderPage(w, "webpages/static/rooms.html")
		RenderRoomChoicePage(w, r, user, 1)
	}
}

func RenderRoomChoicePage(w http.ResponseWriter, r *http.Request, user []util.UserData, x int) {
	var template *template.Template
	template, err := template.ParseGlob("webpages/static/*.html")

	errorHandler.ErrorHandler(err)

	var rooms []RoomsUser

	var rName string
	var uName string
	var admin string

	rows := DataBase.GetRoomsUser(user[0].UserName)

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
	if x == 0 {
		template.ExecuteTemplate(w, "rooms.html", rooms)
	} else {
		template.ExecuteTemplate(w, "rooms1.html", rooms)
	}
}

func RenderUserRoom(w http.ResponseWriter, r *http.Request, user []util.UserData) {
	r.ParseForm()

	roomName := util.GetString(r.Form["roomName"])

	var members string
	var room string
	var admins string
	var rooms []Rooms

	rows, err := DataBase.GetUserRoomData(roomName)

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
