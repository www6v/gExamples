package main

import (
	"gDemo/core"
	_ "gDemo/test"
)

func main() {
	coreMethod()
	// gormMethod()
	// mMethod()
	// jsonMethod()
}

func coreMethod() {
	core.ArrayTest()
	core.CopyTest()
	core.MakeTest()
}

func gormMethod() {
	// gorm.DbConnection()
	// gorm.DbConnection1()
	// gorm.DbConnection2()
}

func mMethod() {
	// m.DescribeNode("200000548", "200000933", "1")
	// m.QueryMachineryRoom(1)
	// m.InsertMachineryRoom(299, "机房299")
	// m.UpdateMachineryRoom(299, "机房299 改变")
	// m.DeleteMachineryRoom(299, "机房299 改变")
}

func jsonMethod() {
	// json.JsonTest()
	// json.Check()
	// json.MyJsonTest()
	// json.ProjectTest1()

	// json.HttpQueryAllUsers("", "")
	// json.HttpQueryAllUsers("URtc-h4r1txxy", "26")
	// json.UserTest()

	// json.ArrayTest()
}
