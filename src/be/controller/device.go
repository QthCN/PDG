package controller

import (
	"be/dao"
	"be/structs"
)

type DeviceMgr struct {
	dao *dao.DeviceDAO
}

var Device *DeviceMgr

func init() {
	Device = &DeviceMgr{
		dao: &dao.DeviceDAO{},
	}
}

func (m *DeviceMgr) CreateDataCenter(name string) error {
	return m.dao.CreateDataCenter(name)
}

func (m *DeviceMgr) DeleteDataCenter(uuid string) error {
	return m.dao.DeleteDataCenter(uuid)
}

func (m *DeviceMgr) ListDataCenters() ([]*structs.DataCenter, error) {
	return m.dao.ListDataCenters()
}

func (m *DeviceMgr) CreateRack(name string) error {
	return m.dao.CreateRack(name)
}

func (m *DeviceMgr) DeleteRack(uuid string) error {
	return m.dao.DeleteRack(uuid)
}

func (m *DeviceMgr) ListRacks() ([]*structs.Rack, error) {
	return m.dao.ListRacks()
}

func (m *DeviceMgr) CreateServerDevice(brand string, model string, diskCapacity int64, memoryCapacity int64, hostname string, enableTime string, expireTime string, os string, comment string) error {
	return m.dao.CreateServerDevice(brand, model, diskCapacity, memoryCapacity, hostname, enableTime, expireTime, os, comment)
}

func (m *DeviceMgr) DeleteServerDevice(uuid string) error {
	return m.dao.DeleteServerDevice(uuid)
}

func (m *DeviceMgr) ListServerDevices() ([]*structs.ServerDevice, error) {
	return m.dao.ListServerDevices()
}

func (m *DeviceMgr) CreateNetworkDevice(brand string, model string, name string, enableTime string, expireTime string, comment string) error {
	return m.dao.CreateNetworkDevice(brand, model, name, enableTime, expireTime, comment)
}

func (m *DeviceMgr) DeleteNetworkDevice(uuid string) error {
	return m.dao.DeleteNetworkDevice(uuid)
}

func (m *DeviceMgr) ListNetworkDevices() ([]*structs.NetworkDevice, error) {
	return m.dao.ListNetworkDevices()
}

func (m *DeviceMgr) CreateStorageDevice(brand string, model string, name string, enableTime string, expireTime string, comment string) error {
	return m.dao.CreateStorageDevice(brand, model, name, enableTime, expireTime, comment)
}

func (m *DeviceMgr) DeleteStorageDevice(uuid string) error {
	return m.dao.DeleteStorageDevice(uuid)
}

func (m *DeviceMgr) ListStorageDevices() ([]*structs.StorageDevice, error) {
	return m.dao.ListStorageDevices()
}

func (m *DeviceMgr) CreateCommonDevice(brand string, model string, name string, enableTime string, expireTime string, comment string) error {
	return m.dao.CreateCommonDevice(brand, model, name, enableTime, expireTime, comment)
}

func (m *DeviceMgr) DeleteCommonDevice(uuid string) error {
	return m.dao.DeleteCommonDevice(uuid)
}

func (m *DeviceMgr) ListCommonDevices() ([]*structs.CommonDevice, error) {
	return m.dao.ListCommonDevices()
}
