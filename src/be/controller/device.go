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

func (m *DeviceMgr) CreateRack(name string, size int64) error {
	return m.dao.CreateRack(name, size)
}

func (m *DeviceMgr) MapRackAndDatacenter(rackUUID string, datacenterUUID string, positionX int64, positionZ int64) error {
	return m.dao.MapRackAndDatacenter(rackUUID, datacenterUUID, positionX, positionZ)
}

func (m *DeviceMgr) MapDeviceAndRack(rackUUID string, deviceUUID string, deviceType string, begPos int64, endPos int64) error {
	return m.dao.MapDeviceAndRack(rackUUID, deviceUUID, deviceType, begPos, endPos)
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

func (m *DeviceMgr) GetPhysicalTopology(datacenterUUID string) (*structs.PhysicalTopology, error) {

	// 机房
	datacenters, err := m.ListDataCenters()
	if err != nil {
		return nil, err
	}

	// 如果入参没有提供机房ID则默认取第一个机房
	if datacenterUUID == "" {
		if len(datacenters) > 0 {
			datacenterUUID = datacenters[0].UUID
		}
	}

	// 机柜
	racks, err := m.ListRacks()
	if err != nil {
		return nil, err
	}
	// 主机
	servers, err := m.ListServerDevices()
	if err != nil {
		return nil, err
	}
	// 存储
	storages, err := m.ListStorageDevices()
	if err != nil {
		return nil, err
	}
	// 网络设备
	networks, err := m.ListNetworkDevices()
	if err != nil {
		return nil, err
	}
	// 其它设备
	others, err := m.ListCommonDevices()
	if err != nil {
		return nil, err
	}

	topology := &structs.PhysicalTopology{
		Size: &structs.PhysicalTopologySize{
			Height: 150,
			Width:  150,
		},
		Racks: []*structs.PhysicalTopologyRack{},
	}

	for _, rack := range racks {
		if rack.Position.DataCenterUUID != datacenterUUID {
			continue
		}

		rackRecord := &structs.PhysicalTopologyRack{
			Name:    rack.Name,
			X:       rack.Position.PositionX,
			Z:       rack.Position.PositionZ,
			U:       rack.SizeU,
			Servers: []*structs.PhysicalTopologyRackServer{},
		}

		for _, device := range servers {
			if device.Position.RackUUID == rack.UUID {
				serverRecord := &structs.PhysicalTopologyRackServer{
					UUID:   device.UUID,
					Name:   device.Hostname,
					BegU:   device.Position.BegPos,
					SizeU:  device.Position.EndPos - device.Position.BegPos,
					Type:   "SERVER",
					Status: Monitor.GetDeviceStatus(device.UUID),
				}
				rackRecord.Servers = append(rackRecord.Servers, serverRecord)
			}
		}

		for _, device := range networks {
			if device.Position.RackUUID == rack.UUID {
				serverRecord := &structs.PhysicalTopologyRackServer{
					UUID:   device.UUID,
					Name:   device.Name,
					BegU:   device.Position.BegPos,
					SizeU:  device.Position.EndPos - device.Position.BegPos,
					Type:   "NETWORK",
					Status: Monitor.GetDeviceStatus(device.UUID),
				}
				rackRecord.Servers = append(rackRecord.Servers, serverRecord)
			}
		}

		for _, device := range storages {
			if device.Position.RackUUID == rack.UUID {
				serverRecord := &structs.PhysicalTopologyRackServer{
					UUID:   device.UUID,
					Name:   device.Name,
					BegU:   device.Position.BegPos,
					SizeU:  device.Position.EndPos - device.Position.BegPos,
					Type:   "STORAGE",
					Status: Monitor.GetDeviceStatus(device.UUID),
				}
				rackRecord.Servers = append(rackRecord.Servers, serverRecord)
			}
		}

		for _, device := range others {
			if device.Position.RackUUID == rack.UUID {
				serverRecord := &structs.PhysicalTopologyRackServer{
					UUID:   device.UUID,
					Name:   device.Name,
					BegU:   device.Position.BegPos,
					SizeU:  device.Position.EndPos - device.Position.BegPos,
					Type:   "OTHER",
					Status: Monitor.GetDeviceStatus(device.UUID),
				}
				rackRecord.Servers = append(rackRecord.Servers, serverRecord)
			}
		}

		topology.Racks = append(topology.Racks, rackRecord)
	}

	return topology, nil
}
