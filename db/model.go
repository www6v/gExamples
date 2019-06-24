package db

type TDescribeNode struct {
	Id   uint `gorm:"primary_key"`
	NodeId string
	TopOrganizationId string
	OrganizationId string
	Channel string
	CountryCode string
	ProvinceCode string
	CityCode string
	DistrictCode string
	NetworkOperator string
	CpuPlatform string
	MachineType string
	UhostType string
	ImageId string
	BasicImageId string
	BasicImageName string
	Tag  string
	Name  string
	State  string
	CreateTime  int
	ChargeType  string
	ExpireTime  int
	Cpu  int
	Memory  int
	AutoRenew  string
	OsName  string
	OsType  string
	BootDiskState  string
	TotalDiskSpace  int
}

func (tdn TDescribeNode) TableName() string {
	return "t_describe_node"
}

type TMachineryRoom struct {
	Id   int `gorm:"primary_key"`
	MachineryId int
	OcName string
}

func (tmr TMachineryRoom) TableName() string {
	return "t_machinery_room"
}