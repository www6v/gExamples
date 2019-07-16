package json

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Data struct {
	AppID     string `json:"AppId"`
	AppStatue bool   `json:"AppStatue"`
	UserID    string `json:"UserId"`
}


type MyJsonName struct {
	Action string `json:"action"`

	Datas []Data `json:"data"`

	Err   int    `json:"err"`
	Msg   string `json:"msg"`
	RPCID string `json:"rpc_id"`
}

//type MyJsonName struct {
//	Action string `json:"action"`
//	Data   []struct {
//		AppID     string `json:"AppId"`
//		AppStatue bool   `json:"AppStatue"`
//		UserID    string `json:"UserId"`
//	} `json:"data"`
//	Err   int    `json:"err"`
//	Msg   string `json:"msg"`
//	RPCID string `json:"rpc_id"`
//}

//{
//"action": ,
//"rpc_id": "222",
//"err": 0,
//"msg": "succ",
//"data": [{
//"AppId": "urtc-i24qcwoo",
//"AppStatue": true,
//"UserId": "37475"
//}, {
//"AppId": "urtc-x0hxv4na",
//"AppStatue": false,
//"UserId": "37475"
//}]
//}

func MyJsonTest() {
	// 创建数据
	//p := personInfo{Name: "Piao", Age: 10, Email: "piaoyunsoft@163.com"}

	datas := make([]Data, 0)
	datas = append(datas, Data{
		AppID:     "urtc-i24qcwoo",
		AppStatue: true,
		UserID:    "37475",
	})

	p := MyJsonName {
		Action : "queryAllProject",
		Datas : datas,
		Err:0,
		Msg:"succ",
		RPCID: "222",
	}

	// 序列化
	data, _ := json.Marshal(&p)
	fmt.Println(string(data))

	// 反序列化
	var p1 MyJsonName
	err := json.Unmarshal([]byte(data), &p1)
	if err != nil {
		fmt.Println("err: ", err)
	} else {
		//fmt.Printf("name=%s, c=%i, email=%s\n", p1.Action, len(p1.Datas), p1.RPCID)
	}
	fmt.Printf("%+v\n", p1)

	//// 反序列化
	//res, err := simplejson.NewJson([]byte(data))
	//if err != nil {
	//	fmt.Println("err: ", err)
	//} else {
	//	fmt.Printf("%+v\n", res)
	//}
}

func HttpQueryAllProject() []byte {
	postRequestBody := " {\"action\":\"queryAllProject\",\"user_id\":\"\",\"rpc_id\":\"222\",\"data\":{ \"statue\":0 } }"

	fmt.Println( "postRequestBody:" + postRequestBody)
	//设置tls配置信息
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
			//InsecureSkipVerify: false,
		},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Post("https://urtc.com.cn/urtcdbproxy",
		//"application/json;charset=UTF-8",
		"application/json",
		strings.NewReader(postRequestBody))
	if err != nil {
		fmt.Println("Get error:", err)
		//return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println( string(body) )
	return body
}


func ParseQueryAllProjectBody(body []byte)   MyJsonName {
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


func HttpQueryProject(appid string) []byte {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
			//InsecureSkipVerify: false,
		},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://120.132.106.50:6005/getrooms?appid=" + appid + "&key=%242a%2410%24hxndLUX%2FnrHX34QGkF19DOFlAuy5dYyu7x%2Fi.ig1vkWbxGbRlD7iy")
	if err != nil {
		fmt.Println("Get error:", err)
		//return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println( string(body) )
	return body
}