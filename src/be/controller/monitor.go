package controller

import (
	"be/dao"
	"be/monitor"
	monitor_fake "be/monitor/fake"
	"be/structs"
	"be/util"
	"fmt"
	"math/rand"
)

type MonitorMgr struct {
	dao *dao.MonitorDAO
}

var Monitor *MonitorMgr

func init() {
	Monitor = &MonitorMgr{
		dao: &dao.MonitorDAO{},
	}
}

func (m *MonitorMgr) GetDeviceStatus(uuid string) string {
	if rand.Intn(2) == 0 {
		return "BAD"
	}
	return "GOOD"
}

func (m *MonitorMgr) ListMonitorItems() ([]*structs.MonitorItem, error) {
	return m.dao.ListMonitorItems()
}

func (m *MonitorMgr) CreateMonitorItem(name string, dcType string) error {
	return m.dao.CreateMonitorItem(name, dcType)
}

func (m *MonitorMgr) UpdateMonitorItem(id int64, name string, dcType string) error {
	return m.dao.UpdateMonitorItem(id, name, dcType)
}

func (m *MonitorMgr) DeleteMonitorItem(id int64) error {
	return m.dao.DeleteMonitorItem(id)
}

func (m *MonitorMgr) UpdateMonitorItemDCCfg(id int64, dcType string, dcFakeCfg_ItemName string, dcFakeCfg_HostIp string) error {
	return m.dao.UpdateMonitorItemDCCfg(id, dcType, dcFakeCfg_ItemName, dcFakeCfg_HostIp)
}

func (m *MonitorMgr) GetMonitorItemById(id int64) (*structs.MonitorItem, error) {
	return m.dao.GetMonitorItemById(id)
}

func (m *MonitorMgr) ListDeviceReleatedMonitorItems(uuid string) ([]*structs.MonitorItemReleatedDevice, error) {
	return m.dao.ListDeviceReleatedMonitorItems(uuid)
}

func (m *MonitorMgr) ListMonitorItemReleatedDevices(itemId int64) ([]*structs.MonitorItemReleatedDevice, error) {
	return m.dao.ListMonitorItemReleatedDevices(itemId)
}

func (m *MonitorMgr) BindMonitorItemAndDevice(itemId int64, itemName string, deviceUUID string, deviceType string, deviceName string) error {
	if err := m.dao.UnBindMonitorItemReleatedDevices(itemId); err != nil {
		return err
	}

	return m.dao.BindMonitorItemAndDevice(itemId, itemName, deviceUUID, deviceType, deviceName)
}

func (m *MonitorMgr) GetMonitorBackendCfgByName(name string) (*structs.MonitorBackendCfg, error) {
	cfgs, err := m.ListMonitorBackendCfgs()
	if err != nil {
		return nil, err
	}

	for _, cfg := range cfgs {
		if cfg.Name == name {
			return cfg, nil
		}
	}

	return nil, fmt.Errorf("获取监控服务失败，无相关名称的监控服务")
}

func (m *MonitorMgr) ListMonitorBackendCfgs() ([]*structs.MonitorBackendCfg, error) {
	cfgs, err := m.dao.ListMonitorBackendCfgs()
	if err != nil {
		return nil, err
	}

	for _, cfg := range cfgs {
		switch cfg.Name {
		case "FAKE":
			if err := util.ParseJsonStr(cfg.Cfg, &cfg.FakeCfg); err != nil {
				return nil, err
			}
			break
		case "ZABBIX":
			if err := util.ParseJsonStr(cfg.Cfg, &cfg.ZabbixCfg); err != nil {
				return nil, err
			}
			break
		default:
			break
		}
	}

	return cfgs, nil
}

func (m *MonitorMgr) UpdateMonitorBackendCfg(backendName string, cfg string) error {
	return m.dao.UpdateMonitorBackendCfg(backendName, cfg)
}

func (m *MonitorMgr) GetDeviceMonitorItemHistoryData(filter *structs.HistoryDataFilter) ([]*structs.HistoryDataRecord, error) {
	records := []*structs.HistoryDataRecord{}

	monitorItem, err := m.dao.GetMonitorItemById(filter.ItemId)
	if err != nil {
		return nil, err
	}

	var monitorProxy monitor.MonitorProxyBase
	switch monitorItem.DCType {
	case "FAKE":
		monitorProxy = &monitor_fake.FakeMonitorProxy{}
		break
	}

	if monitorProxy == nil {
		return nil, fmt.Errorf("设备未关联监控服务")
	}

	var deviceIp string
	deviceIpRecords, err := Ip.GetIpByTargetID(filter.DeviceUUID)
	if err != nil {
		return nil, err
	}
	for _, deviceIpRecord := range deviceIpRecords {
		if deviceIpRecord.Role == "业务" {
			deviceIp = deviceIpRecord.IPAddress
			break
		}
	}
	if deviceIp == "" {
		return nil, fmt.Errorf("设备IP信息获取无效")
	}

	records, err = monitorProxy.GetDeviceHistoryDataRecords(deviceIp, monitorItem, filter)
	if err != nil {
		return nil, err
	}

	return records, nil
}

func (m *MonitorMgr) ListAlertItems() ([]*structs.AlertItem, error) {
	return m.dao.ListAlertItems()
}

func (m *MonitorMgr) DeleteAlertItem(id int64) error {
	return m.dao.DeleteAlertItem(id)
}

func (m *MonitorMgr) CreateAlertItem(itemName string, alertType string, eventId string) error {
	return m.dao.CreateAlertItem(itemName, alertType, eventId)
}

func (m *MonitorMgr) RecordAlert(alertType string, eventId string, alertId string, alertMsg string, alertHost string, isRecover bool) error {
	return m.dao.RecordAlert(alertType, eventId, alertId, alertMsg, alertHost, isRecover)
}

func (m *MonitorMgr) ListAlertEvent() ([]*structs.AlertEvent, error) {
	return m.dao.ListAlertEvent()
}
