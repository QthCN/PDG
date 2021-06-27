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

	sql := `SELECT MONITOR_ITEM.id AS mi_id, MONITOR_ITEM.name AS mi_name, MONITOR_ITEM.isInternal AS mi_isinternal, MONITOR_ITEM.dcType AS mi_dctype, IFNULL(MONITOR_ITEM_DC_FAKE_CFG.fakeItemName, ""), IFNULL(MONITOR_ITEM_DC_FAKE_CFG.hostip, "") AS DC_fake_hostip 
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
		if err = rows.Scan(&record.Id, &record.Name, &record.IsInternal, &record.DCType, &record.DCFakeCfg.ItemName, &record.DCFakeCfg.HostIp); err != nil {
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

	sql := `SELECT MONITOR_ITEM.id, MONITOR_ITEM.name, MONITOR_ITEM.isInternal, MONITOR_ITEM.dcType FROM MONITOR_ITEM`
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
		if err = rows.Scan(&record.Id, &record.Name, &record.IsInternal, &record.DCType); err != nil {
			log.Errorln(err.Error())
			return nil, err
		}
		records = append(records, record)
	}

	return records, nil
}

func (d *MonitorDAO) CreateMonitorItem(name string, dcType string) error {
	err := mysql.DB.SimpleExec("INSERT INTO MONITOR_ITEM(name, isInternal, dcType) VALUES(?, 0, ?)", name, dcType)
	if err != nil {
		log.Errorln(err.Error())
		return err
	}
	return nil
}

func (d *MonitorDAO) UpdateMonitorItem(id int64, name string, dcType string) error {
	err := mysql.DB.SimpleExec("UPDATE MONITOR_ITEM SET name=?, dcType=? WHERE id=?", name, dcType, id)
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

func (d *MonitorDAO) ListAlertItems() ([]*structs.AlertItem, error) {
	records := []*structs.AlertItem{}
	var err error
	tx := mysql.DB.GetTx()
	defer func() {
		if err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()

	sql := `SELECT ALERT_ITEM.id, ALERT_ITEM.itemName, ALERT_ITEM.alertType, ALERT_ITEM.eventId FROM ALERT_ITEM ORDER BY id DESC`
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
		record := &structs.AlertItem{}
		if err = rows.Scan(&record.Id, &record.ItemName, &record.AlertType, &record.EventId); err != nil {
			log.Errorln(err.Error())
			return nil, err
		}
		records = append(records, record)
	}

	return records, nil
}

func (d *MonitorDAO) DeleteAlertItem(id int64) error {
	err := mysql.DB.SimpleExec("DELETE FROM ALERT_ITEM WHERE id=?", id)
	if err != nil {
		log.Errorln(err.Error())
		return err
	}

	return nil
}

func (d *MonitorDAO) CreateAlertItem(itemName string, alertType string, eventId string) error {
	err := mysql.DB.SimpleExec("INSERT INTO ALERT_ITEM(itemName, alertType, eventId) VALUES(?, ?, ?)", itemName, alertType, eventId)
	if err != nil {
		log.Errorln(err.Error())
		return err
	}

	return nil
}

func (d *MonitorDAO) RecordAlert(alertType string, eventId string, alertId string, alertMsg string, alertHost string, isRecover bool) error {
	var err error
	tx := mysql.DB.GetTx()
	defer func() {
		if err == nil {
			unlockTableSql := `UNLOCK TABLES`
			tx.Exec(unlockTableSql)
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()

	// 锁表
	lockTableSql := `LOCK TABLES ALERT_ITEM WRITE, ALERT_EVENT WRITE`
	_, err = tx.Exec(lockTableSql)
	if err != nil {
		log.Errorln(err.Error())
		return err
	}

	// 判断事件是否关注
	checkIfEventExistSql := `SELECT ALERT_ITEM.id FROM ALERT_ITEM WHERE eventId=? AND alertType=? `
	checkIfEventExistStmt, err := tx.Prepare(checkIfEventExistSql)
	if err != nil {
		log.Errorln(err.Error())
		return err
	}
	defer checkIfEventExistStmt.Close()

	checkIfEventExistRows, err := checkIfEventExistStmt.Query(eventId, alertType)
	if err != nil {
		log.Errorln(err.Error())
		return err
	}
	defer checkIfEventExistRows.Close()

	var eventCnt = 0
	for checkIfEventExistRows.Next() {
		if err = checkIfEventExistRows.Scan(&eventCnt); err != nil {
			log.Errorln(err.Error())
			return err
		}
	}
	if eventCnt == 0 {
		log.Infof("告警事件 %s %s 未配置映射", alertType, eventId)
		return nil
	}

	// 判断是否已经存在该事件
	var recordId = 0
	checkIfRecordExistSql := `SELECT ALERT_EVENT.id FROM ALERT_EVENT WHERE ALERT_EVENT.alertType=? AND alertId=?`
	checkIfRecordExistStmt, err := tx.Prepare(checkIfRecordExistSql)
	if err != nil {
		log.Errorln(err.Error())
		return err
	}
	defer checkIfRecordExistStmt.Close()

	checkIfRecordExistRows, err := checkIfRecordExistStmt.Query(alertType, alertId)
	if err != nil {
		log.Errorln(err.Error())
		return err
	}
	defer checkIfRecordExistRows.Close()

	for checkIfRecordExistRows.Next() {
		if err = checkIfRecordExistRows.Scan(&recordId); err != nil {
			log.Errorln(err.Error())
			return err
		}
	}

	if recordId != 0 {
		// 如果记录不为0，则判断是否为恢复事件，如果是的话则修改状态
		if isRecover {
			setOKSql := `UPDATE ALERT_EVENT SET status="已恢复", endTime=NOW() WHERE id=?`
			setOKStmt, err := tx.Prepare(setOKSql)
			if err != nil {
				log.Errorln(err.Error())
				return err
			}
			defer setOKStmt.Close()
			_, err = setOKStmt.Exec(recordId)
			if err != nil {
				log.Errorln(err.Error())
				return err
			}
		}
	} else {
		if isRecover == false {
			insertSql := `INSERT INTO ALERT_EVENT(alertType, eventId, alertId, alertMsg, alertHost, createTime, endTime, status) VALUES(?, ?, ?, ?, ?, NOW(), "1970-01-01 00:00:00", "告警中")`
			insertStmt, err := tx.Prepare(insertSql)
			if err != nil {
				log.Errorln(err.Error())
				return err
			}
			defer insertStmt.Close()
			_, err = insertStmt.Exec(alertType, eventId, alertId, alertMsg, alertHost)
			if err != nil {
				log.Errorln(err.Error())
				return err
			}
		}
	}

	return nil
}

func (d *MonitorDAO) ListAlertEvent() ([]*structs.AlertEvent, error) {
	records := []*structs.AlertEvent{}
	var err error
	tx := mysql.DB.GetTx()
	defer func() {
		if err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()

	sql := `SELECT ALERT_EVENT.id, ALERT_EVENT.alertType, ALERT_EVENT.eventId, ALERT_EVENT.alertId, ALERT_EVENT.alertMsg, ALERT_EVENT.alertHost, ALERT_EVENT.createTime, ALERT_EVENT.endTime, ALERT_EVENT.status FROM ALERT_EVENT WHERE status = "告警中" ORDER BY ALERT_EVENT.createTime DESC`
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
		record := &structs.AlertEvent{}
		if err = rows.Scan(&record.Id, &record.AlertType, &record.EventId, &record.AlertId, &record.AlertMsg, &record.AlertHost, &record.CreateTime, &record.EndTime, &record.Status); err != nil {
			log.Errorln(err.Error())
			return nil, err
		}
		records = append(records, record)
	}

	return records, nil
}
