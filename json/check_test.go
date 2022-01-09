package json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Person struct {
	ID   string
	Name string
}

type Class struct {
	ID    string
	Name  string
	stuff []Person
}

func TestCheck(t *testing.T) {

	pa := Person{ID: "001", Name: "赵大"}
	pb := Person{"002", "刘二"}
	ca := Class{"001", "初一（3）班", []Person{pa, pb}}

	fmt.Println(ca)

	data, _ := json.Marshal(ca)

	fmt.Printf("%s\n", data)

}
