package json

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Users struct {
	AppID string `json:"app_id"`
	Rooms []struct {
		RoomID string   `json:"room_id"`
		Users  []string `json:"users"`
	} `json:"rooms"`
}

const (
	USER_IP_PRE_1 = "120.132.106.136"
	USER_IP_PRE_2 = "120.132.102.168"
	USER_IP_PRE_3 = "117.50.25.102"
	USER_IP_PRE_4 = "106.75.99.16"

	USER_KEY_PRE = "%242a%2410%24hxndLUX%2FnrHX34QGkF19DOFlAuy5dYyu7x%2Fi.ig1vkWbxGbRlD7iy"
)

func HttpQueryAllUsers(appid string, roomid string) []Users {
	users1 := httpQueryUsers(USER_IP_PRE_1, appid, roomid)
	users2 := httpQueryUsers(USER_IP_PRE_2, appid, roomid)
	users3 := httpQueryUsers(USER_IP_PRE_3, appid, roomid)
	users4 := httpQueryUsers(USER_IP_PRE_4, appid, roomid)

	users := append( append( append(users1, users2...), users3...), users4...)
	fmt.Println("---")
	fmt.Println(users)
	return users
}

func httpQueryUsers(userIp string, appid string, roomid string) []Users {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	client := &http.Client{Transport: tr}
	url := "https://" + userIp +":5005/getusers?appid=" + appid + "&roomid=" + roomid + "&key=" + USER_KEY_PRE
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println("Get error:", err)
		//return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)


	var users []Users
	userErr := json.Unmarshal([]byte(body), &users)
	if userErr != nil {
		fmt.Println("userErr: ", userErr)
	} else {
		//fmt.Printf("name=%s, c=%i, email=%s\n", users.Action, len(users.Datas), users.RPCID)
	}
	fmt.Printf("%+v\n", users)

	//fmt.Println( url )
	//fmt.Println( string(body) )

	return users
}