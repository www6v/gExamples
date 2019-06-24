package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	model "gDemo/db"
	"strconv"

	//"uframework/common"
	//"uframework/log"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func DescribeNode(topOrganizationId string,organizationId string, channel string) ([]model.TDescribeNode, error) {
	node  := queryNode(topOrganizationId, organizationId, channel)

	return node, nil
}

func queryNode(topOrganizationId string, organizationId string, channel string) ( []model.TDescribeNode ) {
	db := dbConnect()

	var dns []model.TDescribeNode
	db.Where("top_organization_id = ? and organization_id = ? and channel = ?", topOrganizationId, organizationId, channel).Find(&dns)

	fmt.Println("len(dns) :" + strconv.Itoa(len(dns)))
	return dns
}

func QueryMachineryRoom(MachineryId int)  model.TMachineryRoom {
	db := dbConnect()

	var tmr model.TMachineryRoom
	db.Where("machinery_id = ?", MachineryId).First(&tmr)

	fmt.Println( tmr.OcName)

	return tmr
}

func InsertMachineryRoom(MachineryId int, Name string)  bool {
	db := dbConnect()

	tmr := model.TMachineryRoom{
		MachineryId:MachineryId,
		OcName:Name,
	}
	inserted := db.NewRecord(tmr)
	db.Create(&tmr)
	inserted1 := db.NewRecord(tmr)

	fmt.Println( inserted )
	fmt.Println( inserted1 )

	return  inserted
}

func UpdateMachineryRoom(MachineryId int, Name string)   {
	db := dbConnect()

	tmr := model.TMachineryRoom{
		MachineryId:MachineryId,
	}
	db.Where(&tmr).First(&tmr)

    fmt.Println(strconv.Itoa(tmr.Id) +":"+ strconv.Itoa(tmr.MachineryId)+":"+  tmr.OcName)

	//tmr.MachineryId = MachineryId
	tmr.OcName = Name

	db.Save(&tmr)
}


func DeleteMachineryRoom(MachineryId int, Name string)   {
	db := dbConnect()

	tmr := model.TMachineryRoom{
		MachineryId:MachineryId,
	}
	db.Where(&tmr).First(&tmr)

	//tmr := model.TMachineryRoom{}
	//tmr.MachineryId = MachineryId

	db.Delete(&tmr)
}

func dbConnect() *gorm.DB {
	url := "root:1qaZxsw23edcvfr4@tcp(localhost:3306)/uedn?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", url)

	//url, _ := ufcommon.GetConfigByKey("sqlite.conn_str")  /// 测试用的
	//db, err := gorm.Open("sqlite3", url)
	if err != nil {
		fmt.Println(err)
		panic("Connet db  fails.")
	}
	//defer db.Close()

	err = db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(TABLES...).Error
	if err != nil {
		//uflog.ERROR("AutoMigrate db error :", err)
		//return err
	}

	db.LogMode(true)

	return db
}


//注册数据库表
var TABLES = []interface{}{
	&model.TDescribeNode{},
	&model.TMachineryRoom{},
}


