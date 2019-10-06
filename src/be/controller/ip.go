package controller

import (
	"be/dao"
	"be/structs"
)

type IpMgr struct {
	dao *dao.IpDAO
}

var Ip *IpMgr

func init() {
	Ip = &IpMgr{
		dao: &dao.IpDAO{},
	}
}

func (m *IpMgr) ListIPRecords() ([]*structs.IP, error) {
	return m.dao.ListIPRecords()
}

func (m *IpMgr) CreateIPRecord(ipAddress string, ipType string, role string, targetId string, ipSetId string) error {
	// todo 检查IP是否属于该网段
	return m.dao.CreateIPRecord(ipAddress, ipType, role, targetId, ipSetId)
}

func (m *IpMgr) DeleteIPRecord(uuid string) error {
	return m.dao.DeleteIPRecord(uuid)
}

func (m *IpMgr) CreateIPSet(cidr string, comment string) error {
	return m.dao.CreateIPSet(cidr, comment)
}

func (m *IpMgr) DeleteIPSet(uuid string) error {
	return m.dao.DeleteIPSet(uuid)
}

func (m *IpMgr) ListIPSets() ([]*structs.IPSet, error) {
	return m.dao.ListIPSets()
}
