package main

import (
	"gDemo/gorm"
	_ "gDemo/test"
)

func main(){
	//core.ArrayTest()
	//core.CopyTest()
	//core.MakeTest()
	//fmt.Printf("hello world\n")

	gorm.DbConnection()
	//gorm.DbConnection1()
	//gorm.DbConnection2()


	//m.DescribeNode("200000548","200000933", "1")
	//m.QueryMachineryRoom(1)
	//m.InsertMachineryRoom(299,"机房299")
	//m.UpdateMachineryRoom(299,"机房299 改变")
	//m.DeleteMachineryRoom(299,"机房299 改变")
}

