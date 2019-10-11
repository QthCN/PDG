package controller

import (
	"be/dao"
	"be/structs"
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

func (m *MonitorMgr) CreateMonitorItem(name string, dcType string, alertType string) error {
	return m.dao.CreateMonitorItem(name, dcType, alertType)
}

func (m *MonitorMgr) UpdateMonitorItem(id int64, name string, dcType string, alertType string) error {
	return m.dao.UpdateMonitorItem(id, name, dcType, alertType)
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
	return m.dao.BindMonitorItemAndDevice(itemId, itemName, deviceUUID, deviceType, deviceName)
}
