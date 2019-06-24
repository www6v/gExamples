package gorm


type DescribeNode struct {
	Id   uint `gorm:"primary_key"`

	NodeId string
	Access
	User
	NodeDetail
}

type User struct {
	RequestUuid       string `field:"optional=true"`
	TopOrganizationId int    /// 公司
	OrganizationId    int    /// 组织
	Channel           int    /// 渠道
}

type Access struct {
	CountryCode     string // 国家代码
	ProvinceCode    string // 省代码
	CityCode        string // 市代码
	DistrictCode    string // 区县代码
	NetworkOperator string // 运营商
}
type NodeDetail struct {
	CpuPlatform    string
	MachineType    string
	NodeId         string
	UHostType      string
	IPSet          []NodeIPDetail
	ImageId        string
	BasicImageId   string
	BasicImageName string
	Tag            string
	Name           string
	State          string
	CreateTime     int
	ChargeType     string
	ExpireTime     int
	CPU            int
	Memory         int
	AutoRenew      string
	OsName         string
	OsType         string
	DiskSet        []NodeDiskDetail
	BootDiskState  string
	TotalDiskSpace int
}


type NodeIPDetail struct {
	Type      string
	IPId      string
	IP        string
	Bandwidth int
	Default   string
}

type NodeDiskDetail struct {
	Type   string
	DiskId string
	Name   string
	Drive  string
	Size   int
}

type TDescribeNode struct {
	//gorm.Model

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

func (u TDescribeNode) TableName() string {
		return "t_describe_node"
}

func (u DescribeNode) TableName() string {
	return "t_describe_node"
}
