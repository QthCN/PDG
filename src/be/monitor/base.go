package monitor

import (
	"be/structs"
	"strings"
)

type MonitorProxyBase interface {
	GetDeviceHistoryDataRecords(deviceIp string, monitorItem *structs.MonitorItem, filter *structs.HistoryDataFilter) ([]*structs.HistoryDataRecord, error)
}

func GetValueByVariable(variable string, deviceIp string) string {
	variable = strings.TrimSpace(variable)
	switch variable {
	case "${ip}", "${hostip}":
		return deviceIp
	default:
		return variable
	}
}
