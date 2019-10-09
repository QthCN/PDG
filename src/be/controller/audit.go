package controller

import (
	"be/dao"
	"be/structs"
)

type AuditMgr struct {
	dao *dao.AuditDAO
}

var Audit *AuditMgr

func init() {
	Audit = &AuditMgr{
		dao: &dao.AuditDAO{},
	}
}

func (m *AuditMgr) CreateRecord(username string, action string, url string, args string) error {
	return m.dao.CreateRecord(username, action, url, args)
}

func (m *AuditMgr) ListRecords(filter *structs.ListAuditRecordsCondition) (*structs.AuditRecords, error) {
	return m.dao.ListRecords(filter)
}
