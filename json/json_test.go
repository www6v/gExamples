package json

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/bitly/go-simplejson"
)

type personInfo struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email" xml:"email"`
}

type personInfo1 struct {
	Name  string `json:"name"`
	Email string `json:"email" xml:"email"`
	C     string
}

func TestJson(t *testing.T) {
	// 创建数据
	p := personInfo{Name: "Piao", Age: 10, Email: "piaoyunsoft@163.com"}

	// 序列化
	data, _ := json.Marshal(&p)
	fmt.Println(string(data))

	// 反序列化
	var p1 personInfo1
	err := json.Unmarshal([]byte(data), &p1) // 貌似这种解析方法需要提前知道 json 结构
	if err != nil {
		fmt.Println("err: ", err)
	} else {
		fmt.Printf("name=%s, c=%s, email=%s\n", p1.Name, p1.C, p1.Email)
	}
	fmt.Printf("%+v\n", p1)

	// 反序列化
	res, err := simplejson.NewJson([]byte(data))
	if err != nil {
		fmt.Println("err: ", err)
	} else {
		fmt.Printf("%+v\n", res)
	}
}
