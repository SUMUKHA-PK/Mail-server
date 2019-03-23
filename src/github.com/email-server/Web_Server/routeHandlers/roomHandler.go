package routeHandlers

import (
	"fmt"
	"net/http"

	"github.com/email-server/Web_Server/util"
)


func RoomCreationHandler(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET" {
		util.RenderPage(w,"../webpages/static/createRoom.html")
	}else if r.Method == "POST" {
		r.ParseForm()

		var roomData util.RoomData 

		roomData.RoomName = util.GetString(r.Form["roomName"])
		roomData.Admins = util.GetStringAndAppend(roomData.Admins,r.Form["admins"])
		roomData.Members = util.GetStringAndAppend(roomData.Members,r.Form["users"])

		fmt.Println(roomData)

	}
	
}