package basic

//《25 直播：如何写出优雅的 Go 代码-更多课程【ubkz.com】_一手IT课程资源+微信2268731[2]》
//1. simple
import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestT(t *testing.T) {
	got := 2
	expected := 2
	assert.Equal(t, got, expected, "they should be equal")
}

//2.表驱动模式

//3. test suite

//4. mock

//5. 接口测试
//内置的httptest， 在不占用端口的前提下进行http接口测试

// https://blog.csdn.net/c_circle/article/details/89846724
const (
	ADDRESS = "shanghai"
)

type Person struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Age     int    `json:"age"`
}

var personResponse = []Person{
	{
		Name:    "wahaha",
		Address: "shanghai",
		Age:     20,
	},
	{
		Name:    "lebaishi",
		Address: "shanghai",
		Age:     10,
	},
}

var personResponseBytes, _ = json.Marshal(personResponse)

func TestHttptest(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(personResponseBytes)
		if r.Method != "GET" {
			t.Errorf("Expected 'GET' request, got '%s'", r.Method)
		}
		if r.URL.EscapedPath() != "/person" {
			t.Errorf("Expected request to '/person', got '%s'", r.URL.EscapedPath())
		}
		r.ParseForm()
		topic := r.Form.Get("addr")
		if topic != "shanghai" {
			t.Errorf("Expected request to have 'addr=shanghai', got: '%s'", topic)
		}
	}))

	defer ts.Close()
	api := ts.URL
	fmt.Println("url:", api)
	resp, _ := GetInfo(api)

	fmt.Println("reps:", resp)
}

func GetInfo(api string) ([]Person, error) {
	url := fmt.Sprintf("%s/person?addr=%s", api, ADDRESS)
	resp, err := http.Get(url)

	if err != nil {
		return []Person{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return []Person{}, fmt.Errorf("get info didn’t respond 200 OK: %s", resp.Status)
	}

	return nil, nil
}
