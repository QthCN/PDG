package structs

type DataCenter struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}

type Rack struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}

type ServerDevice struct {
	UUID           string `json:"uuid"`
	Brand          string `json:"brand"`
	Model          string `json:"model"`
	DiskCapacity   int64  `json:"disk_capacity"`
	MemoryCapacity int64  `json:"memory_capacity"`
	Hostname       string `json:"hostname"`
	CreateTime     string `json:"create_time"`
	EnableTime     string `json:"enable_time"`
	ExpireTime     string `json:"expire_time"`
	OS             string `json:"os"`
	Comment        string `json:"comment"`
	IPAddresses    []*IP  `json:"ips"`
}

type NetworkDevice struct {
	UUID        string `json:"uuid"`
	Brand       string `json:"brand"`
	Model       string `json:"model"`
	Name        string `json:"name"`
	CreateTime  string `json:"create_time"`
	EnableTime  string `json:"enable_time"`
	ExpireTime  string `json:"expire_time"`
	Comment     string `json:"comment"`
	IPAddresses []*IP  `json:"ips"`
}

type StorageDevice struct {
	UUID        string `json:"uuid"`
	Brand       string `json:"brand"`
	Model       string `json:"model"`
	Name        string `json:"name"`
	CreateTime  string `json:"create_time"`
	EnableTime  string `json:"enable_time"`
	ExpireTime  string `json:"expire_time"`
	Comment     string `json:"comment"`
	IPAddresses []*IP  `json:"ips"`
}

type CommonDevice struct {
	UUID        string `json:"uuid"`
	Brand       string `json:"brand"`
	Model       string `json:"model"`
	Name        string `json:"name"`
	CreateTime  string `json:"create_time"`
	EnableTime  string `json:"enable_time"`
	ExpireTime  string `json:"expire_time"`
	Comment     string `json:"comment"`
	IPAddresses []*IP  `json:"ips"`
}

type IP struct {
	UUID      string `json:"uuid"`
	IPAddress string `json:"ip_address"`
	Type   string `json:"type"`
	TargetId  string `json:"target_id"`
}
