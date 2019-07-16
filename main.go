package main

import (
	"gDemo/json"
	_ "gDemo/test"
)

func main(){
	//core.ArrayTest()
	//core.CopyTest()
	//core.MakeTest()
	//core.SelectTest()
	//core.SelectTest1()
	//core.SelectTest2()
	//core.Test_Sleep()
	//core.Test_Tick()
	//fmt.Printf("hello world\n")

	//gorm.DbConnection()
	//gorm.DbConnection1()
	//gorm.DbConnection2()


	//m.DescribeNode("200000548","200000933", "1")
	//m.QueryMachineryRoom(1)
	//m.InsertMachineryRoom(299,"机房299")
	//m.UpdateMachineryRoom(299,"机房299 改变")
	//m.DeleteMachineryRoom(299,"机房299 改变")

	json.JsonTest()
	json.Check()
	json.MyJsonTest()
	json.ProjectTest1()

	json.HttpQueryAllUsers("", "")
	json.HttpQueryAllUsers("URtc-h4r1txxy","26")
	json.UserTest()

	json.ArrayTest()
}



