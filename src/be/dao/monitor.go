package dao

import (
	"be/mysql"
	"be/structs"
	"strings"

	log "github.com/sirupsen/logrus"
)

type MonitorDAO struct {
}

func (d *MonitorDAO) UpdateMonitorItemDCCfg(id int64, dcType string, dcFakeCfg_ItemName string, dcFakeCfg_HostIp string) error {
	var err error

	// 更新类型
	if err = mysql.DB.SimpleExec("UPDATE MONITOR_ITEM SET dcType=? WHERE id=?", dcType, id); err != nil {
		log.Errorln(err.Error())
		return err
	}

	dcType = strings.ToUpper(dcType)

	if dcType == "FAKE" {
		// 先统一删除
		if err = mysql.DB.SimpleExec("DELETE FROM MONITOR_ITEM_DC_FAKE_CFG WHERE itemId=?", id); err != nil {
			log.Errorln(err.Error())
			return err
		}

		// 插入
		if err = mysql.DB.SimpleExec("INSERT INTO MONITOR_ITEM_DC_FAKE_CFG(itemId, fakeItemName, hostip) VALUES(?, ?, ?)", id, dcFakeCfg_ItemName, dcFakeCfg_HostIp); err != nil {
			log.Errorln(err.Error())
			return err
		}
	}

	return nil
}

func (d *MonitorDAO) GetMonitorItemById(id int64) (*structs.MonitorItem, error) {
	record := &structs.MonitorItem{
		DCFakeCfg: &structs.DCFakeCfg{},
	}

	var err error
	tx := mysql.DB.GetTx()
	defer func() {
		if err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()

	sql := `SELECT MONITOR_ITEM.id AS mi_id, MONITOR_ITEM.name AS mi_name, MONITOR_ITEM.isInternal AS mi_isinternal, MONITOR_ITEM.dcType AS mi_dctype, MONITOR_ITEM.alertType AS mi_alerttype, IFNULL(MONITOR_ITEM_DC_FAKE_CFG.fakeItemName, ""), IFNULL(MONITOR_ITEM_DC_FAKE_CFG.hostip, "") AS DC_fake_hostip 
			FROM MONITOR_ITEM LEFT JOIN MONITOR_ITEM_DC_FAKE_CFG ON MONITOR_ITEM.id=MONITOR_ITEM_DC_FAKE_CFG.itemId
			WHERE MONITOR_ITEM.id=?`

	stmt, err := tx.Prepare(sql)
	if err != nil {
		log.Errorln(err.Error())
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(id)
	if err != nil {
		log.Errorln(err.Error())
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&record.Id, &record.Name, &record.IsInternal, &record.DCType, &record.AlertType, &record.DCFakeCfg.ItemName, &record.DCFakeCfg.HostIp); err != nil {
			log.Errorln(err.Error())
			return nil, err
		}
	}

	return record, nil
}

func (d *MonitorDAO) ListMonitorItems() ([]*structs.MonitorItem, error) {
	records := []*structs.MonitorItem{}
	var err error
	tx := mysql.DB.GetTx()
	defer func() {
		if err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()

	sql := `SELECT MONITOR_ITEM.id, MONITOR_ITEM.name, MONITOR_ITEM.isInternal, MONITOR_ITEM.dcType, MONITOR_ITEM.alertType FROM MONITOR_ITEM`
	stmt, err := tx.Prepare(sql)
	if err != nil {
		log.Errorln(err.Error())
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Errorln(err.Error())
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		record := &structs.MonitorItem{}
		if err = rows.Scan(&record.Id, &record.Name, &record.IsInternal, &record.DCType, &record.AlertType); err != nil {
			log.Errorln(err.Error())
			return nil, err
		}
		records = append(records, record)
	}

	return records, nil
}

func (d *MonitorDAO) CreateMonitorItem(name string, dcType string, alertType string) error {
	err := mysql.DB.SimpleExec("INSERT INTO MONITOR_ITEM(name, isInternal, dcType, alertType) VALUES(?, 0, ?, ?)", name, dcType, alertType)
	if err != nil {
		log.Errorln(err.Error())
		return err
	}
	return nil
}

func (d *MonitorDAO) UpdateMonitorItem(id int64, name string, dcType string, alertType string) error {
	err := mysql.DB.SimpleExec("UPDATE MONITOR_ITEM SET name=?, dcType=?, alertType=? WHERE id=?", name, dcType, alertType, id)
	if err != nil {
		log.Errorln(err.Error())
		return err
	}
	return nil
}

func (d *MonitorDAO) DeleteMonitorItem(id int64) error {
	err := mysql.DB.SimpleExec("DELETE FROM MONITOR_ITEM WHERE id=? AND isInternal=0", id)
	if err != nil {
		log.Errorln(err.Error())
		return err
	}
	return nil
}

func (d *MonitorDAO) ListMonitorItemReleatedDevices(itemId int64) ([]*structs.MonitorItemReleatedDevice, error) {
	records := []*structs.MonitorItemReleatedDevice{}
	var err error
	tx := mysql.DB.GetTx()
	defer func() {
		if err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()

	sql := `SELECT MONITOR_ITEM_DEVICE_MAPPING.id, MONITOR_ITEM_DEVICE_MAPPING.itemId, MONITOR_ITEM_DEVICE_MAPPING.itemName, MONITOR_ITEM_DEVICE_MAPPING.deviceUUID, MONITOR_ITEM_DEVICE_MAPPING.deviceType, MONITOR_ITEM_DEVICE_MAPPING.deviceName
			FROM MONITOR_ITEM_DEVICE_MAPPING
			WHERE itemId=?`
	stmt, err := tx.Prepare(sql)
	if err != nil {
		log.Errorln(err.Error())
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(itemId)
	if err != nil {
		log.Errorln(err.Error())
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		record := &structs.MonitorItemReleatedDevice{}
		if err = rows.Scan(&record.MappingId, &record.MonitorItemId, &record.MonitorItemName, &record.DeviceUUID, &record.DeviceType, &record.DeviceName); err != nil {
			log.Errorln(err.Error())
			return nil, err
		}
		records = append(records, record)
	}

	return records, nil
}

func (d *MonitorDAO) ListDeviceReleatedMonitorItems(uuid string) ([]*structs.MonitorItemReleatedDevice, error) {
	records := []*structs.MonitorItemReleatedDevice{}
	var err error
	tx := mysql.DB.GetTx()
	defer func() {
		if err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()

	sql := `SELECT MONITOR_ITEM_DEVICE_MAPPING.id, MONITOR_ITEM_DEVICE_MAPPING.itemId, MONITOR_ITEM_DEVICE_MAPPING.itemName, MONITOR_ITEM_DEVICE_MAPPING.deviceUUID, MONITOR_ITEM_DEVICE_MAPPING.deviceType, MONITOR_ITEM_DEVICE_MAPPING.deviceName
			FROM MONITOR_ITEM_DEVICE_MAPPING
			WHERE deviceUUID=?`
	stmt, err := tx.Prepare(sql)
	if err != nil {
		log.Errorln(err.Error())
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(uuid)
	if err != nil {
		log.Errorln(err.Error())
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		record := &structs.MonitorItemReleatedDevice{}
		if err = rows.Scan(&record.MappingId, &record.MonitorItemId, &record.MonitorItemName, &record.DeviceUUID, &record.DeviceType, &record.DeviceName); err != nil {
			log.Errorln(err.Error())
			return nil, err
		}
		records = append(records, record)
	}

	return records, nil
}

func (d *MonitorDAO) UnBindMonitorItemReleatedDevices(itemId int64) error {
	err := mysql.DB.SimpleExec("DELETE FROM MONITOR_ITEM_DEVICE_MAPPING WHERE itemId=?", itemId)
	if err != nil {
		log.Errorln(err.Error())
		return err
	}

	return nil
}

func (d *MonitorDAO) BindMonitorItemAndDevice(itemId int64, itemName string, deviceUUID string, deviceType string, deviceName string) error {
	err := mysql.DB.SimpleExec("DELETE FROM MONITOR_ITEM_DEVICE_MAPPING WHERE itemId=? AND deviceUUID=?", itemId, deviceUUID)
	if err != nil {
		log.Errorln(err.Error())
		return err
	}

	err = mysql.DB.SimpleExec("INSERT INTO MONITOR_ITEM_DEVICE_MAPPING(itemId, itemName, deviceUUID, deviceType, deviceName) VALUES(?, ?, ?, ?, ?)", itemId, itemName, deviceUUID, deviceType, deviceName)
	if err != nil {
		log.Errorln(err.Error())
		return err
	}

	return nil
}

func (d *MonitorDAO) ListMonitorBackendCfgs() ([]*structs.MonitorBackendCfg, error) {
	records := []*structs.MonitorBackendCfg{}
	var err error
	tx := mysql.DB.GetTx()
	defer func() {
		if err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()

	sql := `SELECT MONITOR_BACKEND_CFG.id, MONITOR_BACKEND_CFG.backendName, MONITOR_BACKEND_CFG.cfgStr  FROM MONITOR_BACKEND_CFG`
	stmt, err := tx.Prepare(sql)
	if err != nil {
		log.Errorln(err.Error())
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Errorln(err.Error())
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		record := &structs.MonitorBackendCfg{
			FakeCfg:   &structs.MonitorBackendFakeCfg{},
			ZabbixCfg: &structs.MonitorBackendZabbixCfg{},
		}
		if err = rows.Scan(&record.Id, &record.Name, &record.Cfg); err != nil {
			log.Errorln(err.Error())
			return nil, err
		}
		records = append(records, record)
	}

	return records, nil
}

func (d *MonitorDAO) UpdateMonitorBackendCfg(backendName string, cfg string) error {
	err := mysql.DB.SimpleExec("DELETE FROM MONITOR_BACKEND_CFG WHERE backendName=?", backendName)
	if err != nil {
		log.Errorln(err.Error())
		return err
	}

	err = mysql.DB.SimpleExec("INSERT INTO MONITOR_BACKEND_CFG(backendName, cfgStr) VALUES(?, ?)", backendName, cfg)
	if err != nil {
		log.Errorln(err.Error())
		return err
	}

	return nil
}
