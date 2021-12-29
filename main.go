package main

import (
	"fmt"
	"gDemo/core"
	"gDemo/json"
	_ "gDemo/test"
)

func main() {
	coreMethod()
	// gormMethod()
	// mMethod()
	// jsonMethod()
}

func coreMethod() {
	core.ContextTest()
	core.StructTest()
	core.ArrayTest()
	core.CopyTest()
	core.MakeTest()
	core.SelectTest()
	core.SelectTest1()
	core.SelectTest2()
	core.Test_Sleep()
	core.FuctionTest()
	// core.Test_Tick()

	fmt.Printf("hello world\n")
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
	json.JsonTest()
	json.Check()
	json.MyJsonTest()
	json.ProjectTest1()

	json.HttpQueryAllUsers("", "")
	json.HttpQueryAllUsers("URtc-h4r1txxy", "26")
	json.UserTest()

	json.ArrayTest()
}
