package structs

type PhysicalTopologySize struct {
	Height int64 `json:"height"`
	Width  int64 `json:"width"`
}

type PhysicalTopologyRackServer struct {
	UUID   string `json:"uuid"`
	Name   string `json:"name"`
	BegU   int64  `json:"begU"`
	SizeU  int64  `json:"sizeU"`
	Type   string `json:"type"`
	Status string `json:"status"`
}

type PhysicalTopologyRack struct {
	Name    string                        `json:"name"`
	X       int64                         `json:"x"`
	Z       int64                         `json:"z"`
	U       int64                         `json:"u"`
	Servers []*PhysicalTopologyRackServer `json:"servers"`
}

type PhysicalTopology struct {
	Size  *PhysicalTopologySize   `json:"size"`
	Racks []*PhysicalTopologyRack `json:"racks"`
}

type ResourceTopology struct {
	UUID       string              `json:"uuid"`
	DeviceName string              `json:"label"`
	DeviceType string              `json:"device_type"`
	Childrens  []*ResourceTopology `json:"children"`
}
