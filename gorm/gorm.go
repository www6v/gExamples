package gorm

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"strconv"
)

type Product struct {
	gorm.Model
	Code string
	Price uint
}

func DbConnection() {

	//url := "root:1qaZxsw23edcvfr4@tcp(localhost:3306)/uedn?charset=utf8&parseTime=True&loc=Local"
	////url := "root:1qaZxsw23edcvfr4@tcp(192.168.189.238:3306)/test?charset=utf8&parseTime=True&loc=Local"
	//
	////db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	//db, err := gorm.Open("mysql", url)
	//
	//if err != nil {
	//	fmt.Println(err)
	//	//panic("连接数据库失败")
	//}
	//defer db.Close()

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("连接数据库失败")
	}
	defer db.Close()

	db.LogMode(true)


	db.AutoMigrate(&Product{})


	db.Create(&Product{Code: "L1212", Price: 1000})


	var product Product
	db.First(&product, 1) // 查询id为1的product
	db.First(&product, "code = ?", "L1212") // 查询code为l1212的product

	db.Model(&product).Update("Price", 2000)

	//db.Delete(&product)
}


func DbConnection1() {

	url := "root:1qaZxsw23edcvfr4@tcp(localhost:3306)/uedn?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", url)

	if err != nil {
		fmt.Println(err)
		//panic("连接数据库失败")
	}
	defer db.Close()


	//db.AutoMigrate(&Product{})

	db.AutoMigrate(&TDescribeNode{})


	//db.Create(&DescribeNode{
	//	NodeId : "n1",
	//	TopOrganizationId : "t1",
	//	OrganizationId : "oi",
	//	Channel : "c1",
	//	CountryCode : "cc1",
	//	ProvinceCode : "pc1",
	//	CityCode : "cc1",
	//	DistrictCode : "dc",
	//	NetworkOperator : "no",
	//	CpuPlatform : "cp",
	//	MachineType : "mt",
	//	UhostType : "uht",
	//	ImageId : "ii",
	//	BasicImageId : "bii",
	//	BasicImageName : "bin",
	//	Tag  : "t",
	//	Name  : "n",
	//	State  : "s",
	//	CreateTime : 111,
	//	ChargeType  : "ct",
	//	ExpireTime : 111,
	//	Cpu  : 11,
	//	Memory  : 22,
	//	AutoRenew  : "ar",
	//	OsName  : "on",
	//	OsType  : "ot",
	//	BootDiskState  : "bds",
	//	TotalDiskSpace  : 123,
	//})

	//var dn TDescribeNode
	//db.First(&dn, "top_organization_id = ? and organization_id = ? and channel = ?", "200000548","200000933","1")
	//fmt.Println("dn.AutoRenew :" + dn.AutoRenew)

	var dns []TDescribeNode
	//db.First(&dns, "top_organization_id = ? and organization_id = ? and channel = ?", "200000548","200000933","1")
	db.Where("top_organization_id = ? and organization_id = ? and channel = ?", "200000548","200000933","1").Find(&dns)

	fmt.Println( len(dns) )

	//db.First(&product, 1) // 查询id为1的product
	//db.First(&dn, "code = ?", "L1212")

	//db.Model(&product).Update("Price", 2000)
	//db.Delete(&product)
}


func DbConnection2() {
	db, err := gorm.Open("sqlite3", "t_uedn")
	if err != nil {
		fmt.Println(err)
		panic("连接数据库失败")
	}
	//defer db.Close()

	db.LogMode(true)

	db.AutoMigrate(&HostInfo{})

	//var product Product
	//db.First(&product, 1)
	//db.First(&product, "code = ?", "L1212")

	var hi []HostInfo
	//db.First(&hi, "resource_type = ? and state = ?", "100", "1")

	db.Where("resource_type = ? and state = ?", "100", "1").Find(&hi)

	fmt.Println( "hi.xxx :" + strconv.Itoa( len(hi) )  )

	//db.Model(&product).Update("Price", 2000)

	//db.Delete(&product)
}


type HostInfo struct {
	Id   int `gorm:"primary_key"`

	State            int
	HostId           int
	ResourceType     int
	IP               string
	ActionIP         string
	ActionPort       int
	CPUInfo          string
	NicCap           int
	DiskSpaceCap     int
	DiskIOCap        int
	BlockDevices     string
	CoreCountUsed    int
	MemoryCapUsed    int
	CoreCountCap     int
	MemoryCap        int
	NicCapUsed       int
	DiskSpaceCapUsed int
	DiskIOCapUsed    int
	NodeCount        int
	CreateTime       int
	ModifyTime       int
}


func (tmr HostInfo) TableName() string {
	return "t_host_resource"
}
