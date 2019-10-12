package fake

import (
	"be/monitor"
	"be/structs"
	"fmt"
	"math/rand"

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

	for _, minute := range []int64{11, 12, 13, 14, 15, 16, 17, 18, 19, 20} {
		d := fmt.Sprintf("2019-10-10 09:03:%d", minute)
		records = append(records, &structs.HistoryDataRecord{
			Key:   d,
			Value: 50.0 + rand.Float64()*50.0,
		})
	}

	return records, nil
}
