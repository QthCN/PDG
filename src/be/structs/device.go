package structs

type Device struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
	DeviceType string `json:"device_type"`
}

type DataCenter struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}

type RackPosition struct {
	DataCenterUUID string `json:"datacenter_uuid"`
	DataCenterName string `json:"datacenter_name"`
	PositionX      int64  `json:"position_x"`
	PositionZ      int64  `json:"position_z"`
}

type Rack struct {
	UUID     string        `json:"uuid"`
	Name     string        `json:"name"`
	SizeU    int64         `json:"size_u"`
	Position *RackPosition `json:"position"`
}

type DevicePosition struct {
	RackUUID string `json:"rack_uuid"`
	RackName string `json:"rack_name"`
	BegPos   int64  `json:"beg_pos"`
	EndPos   int64  `json:"end_pos"`
}

type ServerDevice struct {
	UUID           string          `json:"uuid"`
	Brand          string          `json:"brand"`
	Model          string          `json:"model"`
	DiskCapacity   int64           `json:"disk_capacity"`
	MemoryCapacity int64           `json:"memory_capacity"`
	Hostname       string          `json:"hostname"`
	CreateTime     string          `json:"create_time"`
	EnableTime     string          `json:"enable_time"`
	ExpireTime     string          `json:"expire_time"`
	OS             string          `json:"os"`
	Comment        string          `json:"comment"`
	IPAddresses    []*IP           `json:"ips"`
	Position       *DevicePosition `json:"position"`
}

type NetworkDevice struct {
	UUID        string          `json:"uuid"`
	Brand       string          `json:"brand"`
	Model       string          `json:"model"`
	Name        string          `json:"name"`
	CreateTime  string          `json:"create_time"`
	EnableTime  string          `json:"enable_time"`
	ExpireTime  string          `json:"expire_time"`
	Comment     string          `json:"comment"`
	IPAddresses []*IP           `json:"ips"`
	Position    *DevicePosition `json:"position"`
}

type StorageDevice struct {
	UUID        string          `json:"uuid"`
	Brand       string          `json:"brand"`
	Model       string          `json:"model"`
	Name        string          `json:"name"`
	CreateTime  string          `json:"create_time"`
	EnableTime  string          `json:"enable_time"`
	ExpireTime  string          `json:"expire_time"`
	Comment     string          `json:"comment"`
	IPAddresses []*IP           `json:"ips"`
	Position    *DevicePosition `json:"position"`
}

type CommonDevice struct {
	UUID        string          `json:"uuid"`
	Brand       string          `json:"brand"`
	Model       string          `json:"model"`
	Name        string          `json:"name"`
	CreateTime  string          `json:"create_time"`
	EnableTime  string          `json:"enable_time"`
	ExpireTime  string          `json:"expire_time"`
	Comment     string          `json:"comment"`
	IPAddresses []*IP           `json:"ips"`
	Position    *DevicePosition `json:"position"`
}
