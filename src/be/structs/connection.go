package structs

type Connection struct {
	UUID                  string `json:"uuid"`
	SourceId              string `json:"source_id"`
	SourcePort            string `json:"source_port"`
	SourceDeviceType      string `json:"source_device_type"`
	SourceDeviceName      string `json:"source_device_name"`
	DestinationId         string `json:"destination_id"`
	DestinationPort       string `json:"destination_port"`
	DestinationDeviceType string `json:"destination_device_type"`
	DestinationDeviceName string `json:"destination_device_name"`
}
