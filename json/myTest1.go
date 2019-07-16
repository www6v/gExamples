package json

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type Room struct {
	AppID string `json:"app_id"`
	Rooms []struct {
		RoomID  string `json:"roomId"`
		Start   string `json:"start"`
		UserNum int    `json:"userNum"`
	} `json:"rooms"`
}

func ProjectTest() {

body :=	httpQueryAllProject()
result := parseQueryAllProjectBody(body)


	for _, t := range result.Datas {
		appId := t.AppID
		room := httpQueryProject(appId)

		if room.Rooms != nil {
			fmt.Println(  len(room.Rooms) )
			rooms := room.Rooms
			for _, room :=range rooms {
				roomID := room.RoomID
				start := room.Start
				userNum := room.UserNum

				fmt.Println( roomID + ":" + start + ":" + strconv.Itoa(userNum))
				//ch <- prometheus.NewMetricWithTimestamp(
				//	tm,
				//	prometheus.MustNewConstHistogram(
				//		c.UserCountDesc,
				//		uint64(userNum),
				//		float64(i),
				//		testmap,
				//
				//		c.Hostname,
				//		appId,
				//		roomID,
				//		start,
				//	),
				//)
			}

		}
	}

}

func httpQueryAllProject() []byte {
	postRequestBody := " {\"action\":\"queryAllProject\",\"user_id\":\"\",\"rpc_id\":\"222\",\"data\":{ \"statue\":0 } }"

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	client := &http.Client{Transport: tr}
	url:= "https://urtc.com.cn/urtcdbproxy"
	contentType:="application/json"
	resp, err := client.Post(url,
		contentType,
		strings.NewReader(postRequestBody))
	if err != nil {
		fmt.Println("Get error:", err)
		//return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return body
}


func parseQueryAllProjectBody(body []byte) (MyJsonName) {
	//data, _ := json.Marshal(&p)

	var p1 MyJsonName
	err := json.Unmarshal([]byte(body), &p1)
	if err != nil {
		fmt.Println("err: ", err)
	} else {
		//fmt.Printf("name=%s, c=%i, email=%s\n", p1.Action, len(p1.Datas), p1.RPCID)
	}
	fmt.Printf("%+v\n", p1)

	return p1
}


func httpQueryProject(appid string) Room {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	client := &http.Client{Transport: tr}
	url := "https://106.75.223.113:6005/getrooms?appid=" + appid + "&key=%242a%2410%24hxndLUX%2FnrHX34QGkF19DOFlAuy5dYyu7x%2Fi.ig1vkWbxGbRlD7iy"
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println("Get error:", err)
		//return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)



	var room Room
	roomErr := json.Unmarshal([]byte(body), &room)
	if roomErr != nil {
		fmt.Println("roomErr: ", roomErr)
	} else {
		//fmt.Printf("name=%s, c=%i, email=%s\n", room.Action, len(room.Datas), room.RPCID)
	}
	fmt.Printf("%+v\n", room)

	//fmt.Println( url )
	fmt.Println( string(body) )

	return room
}