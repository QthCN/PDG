package fake

import (
	"be/monitor"
	"be/structs"

	log "github.com/sirupsen/logrus"
)

type FakeMonitorProxy struct {
}

func (p *FakeMonitorProxy) GetDeviceHistoryDataRecords(deviceIp string, monitorItem *structs.MonitorItem, filter *structs.HistoryDataFilter) ([]*structs.HistoryDataRecord, error) {
	records := []*structs.HistoryDataRecord{}

	dcCfg := monitorItem.DCFakeCfg

	// 参数
	itemName := monitor.GetValueByVariable(dcCfg.ItemName, deviceIp)
	hostIp := monitor.GetValueByVariable(dcCfg.HostIp, deviceIp)

	log.Debugf("Fake 监控服务开始查询数据， %s - %s", itemName, hostIp)

	return records, nil
}
