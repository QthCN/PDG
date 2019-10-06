package controller

import (
	"be/dao"
	"be/structs"
)

type ConnectionMgr struct {
	dao *dao.ConnectionDAO
}

var Connection *ConnectionMgr

func init() {
	Connection = &ConnectionMgr{
		dao: &dao.ConnectionDAO{},
	}
}

func (m *ConnectionMgr) ListConnections() ([]*structs.Connection, error) {
	return m.dao.ListConnections()
}

func (m *ConnectionMgr) CreateConnection(sourceId string, sourcePort string, sourceDeviceType string, sourceDeviceName string, destinationId string, destinationPort string, destinationDeviceType string, destinationDeviceName string) error {
	return m.dao.CreateConnection(sourceId, sourcePort, sourceDeviceType, sourceDeviceName, destinationId, destinationPort, destinationDeviceType, destinationDeviceName)
}

func (m *ConnectionMgr) DeleteConnection(uuid string) error {
	return m.dao.DeleteConnection(uuid)
}
