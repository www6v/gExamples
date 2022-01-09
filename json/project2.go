package json

import "fmt"

// func ProjectTest1() {
// 	project := HttpQueryAllProject()
// 	result := ParseQueryAllProjectBody(project)
// 	for _, t := range result.Datas {
// 		appId := t.AppID
// 		HttpQueryProject(appId)
// 	}
// 	ProjectTest()
// }

func TestUser() {
	roomID := "1111"
	users := TestHttpQueryAllUsers("URtc-h4r1txxy", "1111")
	var usersInRoom []string
	for _, user := range users {
		rooms := user.Rooms
		for _, room := range rooms {
			if room.RoomID == roomID {
				usersInRoom = room.Users
				break
			}
		}
	}
	fmt.Print(usersInRoom)
}
