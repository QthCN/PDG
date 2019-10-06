package dao

import (
	"be/mysql"
	"be/structs"
	"be/util"

	log "github.com/sirupsen/logrus"
)

type DeviceDAO struct {
}

func (d *DeviceDAO) CreateDataCenter(name string) error {
	if err := mysql.DB.SimpleExec("INSERT INTO DATACENTER(uuid, name, isDeleted) VALUES(?, ?, 0)", util.GetUUID(), name); err != nil {
		log.Errorln(err.Error())
		return err
	}
	return nil
}

func (d *DeviceDAO) DeleteDataCenter(uuid string) error {
	if err := mysql.DB.SimpleExec("UPDATE DATACENTER SET isDeleted=1 WHERE uuid=?", uuid); err != nil {
		log.Errorln(err.Error())
		return err
	}
	return nil
}

func (d *DeviceDAO) ListDataCenters() ([]*structs.DataCenter, error) {
	var err error
	tx := mysql.DB.GetTx()
	defer func() {
		if err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()

	records := []*structs.DataCenter{}

	sql := `SELECT uuid, name FROM DATACENTER WHERE isDeleted=0 ORDER BY name`
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
		record := &structs.DataCenter{}
		if err = rows.Scan(&record.UUID, &record.Name); err != nil {
			log.Errorln(err.Error())
			return nil, err
		}
		records = append(records, record)
	}

	return records, nil
}

func (d *DeviceDAO) CreateRack(name string, size int64) error {
	if err := mysql.DB.SimpleExec("INSERT INTO RACK(uuid, name, size, isDeleted) VALUES(?, ?, ?, 0)", util.GetUUID(), name, size); err != nil {
		log.Errorln(err.Error())
		return err
	}
	return nil
}

func (d *DeviceDAO) MapDeviceAndRack(rackUUID string, deviceUUID string, deviceType string, begPos int64, endPos int64) error {
	if err := mysql.DB.SimpleExec("DELETE FROM MAPPING_RACK_DEVICE WHERE deviceId=?", deviceUUID); err != nil {
		log.Errorln(err.Error())
		return err
	}

	if err := mysql.DB.SimpleExec("INSERT INTO MAPPING_RACK_DEVICE(uuid, rackId, deviceId, deviceType, begPos, endPos) VALUES(?, ?, ?, ?, ?, ?)", util.GetUUID(), rackUUID, deviceUUID, deviceType, begPos, endPos); err != nil {
		log.Errorln(err.Error())
		return err
	}

	return nil
}

func (d *DeviceDAO) MapRackAndDatacenter(rackUUID string, datacenterUUID string, positionX int64, positionZ int64) error {
	if err := mysql.DB.SimpleExec("DELETE FROM MAPPING_DATACENTER_RACK WHERE rackId=?", rackUUID); err != nil {
		log.Errorln(err.Error())
		return err
	}

	if err := mysql.DB.SimpleExec("INSERT INTO MAPPING_DATACENTER_RACK(uuid, datacenterId, rackId, positionX, positionZ) VALUES(?, ?, ?, ?, ?)", util.GetUUID(), datacenterUUID, rackUUID, positionX, positionZ); err != nil {
		log.Errorln(err.Error())
		return err
	}

	return nil
}

func (d *DeviceDAO) DeleteRack(uuid string) error {
	if err := mysql.DB.SimpleExec("UPDATE RACK SET isDeleted=1 WHERE uuid=?", uuid); err != nil {
		log.Errorln(err.Error())
		return err
	}
	return nil
}

func (d *DeviceDAO) ListRacks() ([]*structs.Rack, error) {
	var err error
	tx := mysql.DB.GetTx()
	defer func() {
		if err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()

	records := []*structs.Rack{}

	sql := `SELECT RACK.uuid, 
				   RACK.name, 
				   RACK.size,
				   IFNULL(DATACENTER.name, "") AS dc_name,
				   IFNULL(MAPPING_DATACENTER_RACK.datacenterId, "") AS dc_uuid,
				   IFNULL(MAPPING_DATACENTER_RACK.positionX, 0) AS position_x,
				   IFNULL(MAPPING_DATACENTER_RACK.positionZ, 0) AS position_z
	FROM RACK LEFT JOIN MAPPING_DATACENTER_RACK ON RACK.uuid=MAPPING_DATACENTER_RACK.rackId LEFT JOIN DATACENTER ON MAPPING_DATACENTER_RACK.datacenterId=DATACENTER.uuid 
	WHERE RACK.isDeleted=0 ORDER BY RACK.name`
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
		record := &structs.Rack{
			Position: &structs.RackPosition{},
		}
		if err = rows.Scan(&record.UUID, &record.Name, &record.SizeU, &record.Position.DataCenterName, &record.Position.DataCenterUUID, &record.Position.PositionX, &record.Position.PositionZ); err != nil {
			log.Errorln(err.Error())
			return nil, err
		}
		records = append(records, record)
	}

	return records, nil
}

func (d *DeviceDAO) CreateServerDevice(brand string, model string, diskCapacity int64, memoryCapacity int64, hostname string, enableTime string, expireTime string, os string, comment string) error {
	if err := mysql.DB.SimpleExec("INSERT INTO SERVER_DEVICE(uuid, brand, model, diskCapacity, memoryCapacity, hostname, createTime, enableTime, expireTime, os, comment, isDeleted) VALUES(?, ?, ?, ?, ?, ?, NOW(), ?, ?, ?, ?, 0)", util.GetUUID(), brand, model, diskCapacity, memoryCapacity, hostname, enableTime, expireTime, os, comment); err != nil {
		log.Errorln(err.Error())
		return err
	}
	return nil
}

func (d *DeviceDAO) DeleteServerDevice(uuid string) error {
	if err := mysql.DB.SimpleExec("UPDATE SERVER_DEVICE SET isDeleted=1 WHERE uuid=?", uuid); err != nil {
		log.Errorln(err.Error())
		return err
	}
	return nil
}

func (d *DeviceDAO) ListServerDevices() ([]*structs.ServerDevice, error) {
	var err error
	tx := mysql.DB.GetTx()
	defer func() {
		if err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()

	records := []*structs.ServerDevice{}

	sql := `SELECT SERVER_DEVICE.uuid, 
				   SERVER_DEVICE.brand, 
				   SERVER_DEVICE.model, 
				   SERVER_DEVICE.diskCapacity, 
				   SERVER_DEVICE.memoryCapacity, 
				   SERVER_DEVICE.hostname, 
				   SERVER_DEVICE.createTime, 
				   SERVER_DEVICE.enableTime, 
				   SERVER_DEVICE.expireTime, 
				   SERVER_DEVICE.os, 
				   SERVER_DEVICE.comment, 
				   IFNULL(IP.uuid, "") AS ip_uuid, 
				   IFNULL(IP.ipAddress, "") AS ip_address, 
				   IFNULL(IP.type, "") AS ip_type,
				   IFNULL(IP.role, "") AS ip_role,
				   IFNULL(IP.ipSetId, "") AS ip_set_id,
				   IFNULL(RACK.name, "") AS rack_name,
				   IFNULL(MAPPING_RACK_DEVICE.rackId, "") AS rack_id,
				   IFNULL(MAPPING_RACK_DEVICE.begPos, 0) AS beg_pos,
				   IFNULL(MAPPING_RACK_DEVICE.endPos, 0) AS end_pos
	FROM SERVER_DEVICE LEFT JOIN IP ON SERVER_DEVICE.uuid=IP.targetId 
		 LEFT JOIN MAPPING_RACK_DEVICE ON MAPPING_RACK_DEVICE.deviceId=SERVER_DEVICE.uuid
		 LEFT JOIN RACK ON RACK.uuid=MAPPING_RACK_DEVICE.rackId
	WHERE SERVER_DEVICE.isDeleted=0 ORDER BY SERVER_DEVICE.uuid`
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

	record := &structs.ServerDevice{
		IPAddresses: []*structs.IP{},
	}
	for rows.Next() {
		tmpRecord := &structs.ServerDevice{
			Position:    &structs.DevicePosition{},
			IPAddresses: []*structs.IP{},
		}
		ipAddress := &structs.IP{}
		if err = rows.Scan(&tmpRecord.UUID, &tmpRecord.Brand, &tmpRecord.Model, &tmpRecord.DiskCapacity, &tmpRecord.MemoryCapacity, &tmpRecord.Hostname, &tmpRecord.CreateTime, &tmpRecord.EnableTime, &tmpRecord.ExpireTime, &tmpRecord.OS, &tmpRecord.Comment, &ipAddress.UUID, &ipAddress.IPAddress, &ipAddress.Type, &ipAddress.Role, &ipAddress.IPSetId, &tmpRecord.Position.RackName, &tmpRecord.Position.RackUUID, &tmpRecord.Position.BegPos, &tmpRecord.Position.EndPos); err != nil {
			log.Errorln(err.Error())
			return nil, err
		}
		ipAddress.TargetId = tmpRecord.UUID

		if tmpRecord.UUID == record.UUID {
			if ipAddress.UUID != "" {
				record.IPAddresses = append(record.IPAddresses, ipAddress)
			}
		} else {
			record = tmpRecord
			if ipAddress.UUID != "" {
				record.IPAddresses = append(record.IPAddresses, ipAddress)
			}
			records = append(records, record)
		}
	}

	return records, nil
}

func (d *DeviceDAO) CreateNetworkDevice(brand string, model string, name string, enableTime string, expireTime string, comment string) error {
	if err := mysql.DB.SimpleExec("INSERT INTO NETWORK_DEVICE(uuid, brand, model, name, createTime, enableTime, expireTime, comment, isDeleted) VALUES(?, ?, ?, ?, NOW(), ?, ?, ?, 0)", util.GetUUID(), brand, model, name, enableTime, expireTime, comment); err != nil {
		log.Errorln(err.Error())
		return err
	}
	return nil
}

func (d *DeviceDAO) DeleteNetworkDevice(uuid string) error {
	if err := mysql.DB.SimpleExec("UPDATE NETWORK_DEVICE SET isDeleted=1 WHERE uuid=?", uuid); err != nil {
		log.Errorln(err.Error())
		return err
	}
	return nil
}

func (d *DeviceDAO) ListNetworkDevices() ([]*structs.NetworkDevice, error) {
	var err error
	tx := mysql.DB.GetTx()
	defer func() {
		if err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()

	records := []*structs.NetworkDevice{}

	sql := `SELECT NETWORK_DEVICE.uuid, 
				   NETWORK_DEVICE.brand, 
				   NETWORK_DEVICE.model, 
				   NETWORK_DEVICE.name, 
				   NETWORK_DEVICE.createTime, 
				   NETWORK_DEVICE.enableTime, 
				   NETWORK_DEVICE.expireTime, 
				   NETWORK_DEVICE.comment, 
				   IFNULL(IP.uuid, "") AS ip_uuid, 
				   IFNULL(IP.ipAddress, "") AS ip_address, 
				   IFNULL(IP.type, "") AS ip_type,
				   IFNULL(IP.role, "") AS ip_role,
				   IFNULL(IP.ipSetId, "") AS ip_set_id,
				   IFNULL(RACK.name, "") AS rack_name,
				   IFNULL(MAPPING_RACK_DEVICE.rackId, "") AS rack_id,
				   IFNULL(MAPPING_RACK_DEVICE.begPos, 0) AS beg_pos,
				   IFNULL(MAPPING_RACK_DEVICE.endPos, 0) AS end_pos
			FROM NETWORK_DEVICE LEFT JOIN IP ON NETWORK_DEVICE.uuid=IP.targetId
			     LEFT JOIN MAPPING_RACK_DEVICE ON MAPPING_RACK_DEVICE.deviceId=NETWORK_DEVICE.uuid
		         LEFT JOIN RACK ON RACK.uuid=MAPPING_RACK_DEVICE.rackId
			WHERE NETWORK_DEVICE.isDeleted=0 ORDER BY NETWORK_DEVICE.uuid`
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

	record := &structs.NetworkDevice{
		IPAddresses: []*structs.IP{},
	}
	for rows.Next() {
		tmpRecord := &structs.NetworkDevice{
			Position:    &structs.DevicePosition{},
			IPAddresses: []*structs.IP{},
		}
		ipAddress := &structs.IP{}
		if err = rows.Scan(&tmpRecord.UUID, &tmpRecord.Brand, &tmpRecord.Model, &tmpRecord.Name, &tmpRecord.CreateTime, &tmpRecord.EnableTime, &tmpRecord.ExpireTime, &tmpRecord.Comment, &ipAddress.UUID, &ipAddress.IPAddress, &ipAddress.Type, &ipAddress.Role, &ipAddress.IPSetId, &tmpRecord.Position.RackName, &tmpRecord.Position.RackUUID, &tmpRecord.Position.BegPos, &tmpRecord.Position.EndPos); err != nil {
			log.Errorln(err.Error())
			return nil, err
		}
		ipAddress.TargetId = tmpRecord.UUID

		if tmpRecord.UUID == record.UUID {
			if ipAddress.UUID != "" {
				record.IPAddresses = append(record.IPAddresses, ipAddress)
			}
		} else {
			record = tmpRecord
			if ipAddress.UUID != "" {
				record.IPAddresses = append(record.IPAddresses, ipAddress)
			}
			records = append(records, record)
		}
	}

	return records, nil
}

func (d *DeviceDAO) CreateStorageDevice(brand string, model string, name string, enableTime string, expireTime string, comment string) error {
	if err := mysql.DB.SimpleExec("INSERT INTO STORAGE_DEVICE(uuid, brand, model, name, createTime, enableTime, expireTime, comment, isDeleted) VALUES(?, ?, ?, ?, NOW(), ?, ?, ?, 0)", util.GetUUID(), brand, model, name, enableTime, expireTime, comment); err != nil {
		log.Errorln(err.Error())
		return err
	}
	return nil
}

func (d *DeviceDAO) DeleteStorageDevice(uuid string) error {
	if err := mysql.DB.SimpleExec("UPDATE STORAGE_DEVICE SET isDeleted=1 WHERE uuid=?", uuid); err != nil {
		log.Errorln(err.Error())
		return err
	}
	return nil
}

func (d *DeviceDAO) ListStorageDevices() ([]*structs.StorageDevice, error) {
	var err error
	tx := mysql.DB.GetTx()
	defer func() {
		if err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()

	records := []*structs.StorageDevice{}

	sql := `SELECT STORAGE_DEVICE.uuid, 
				   STORAGE_DEVICE.brand, 
				   STORAGE_DEVICE.model, 
				   STORAGE_DEVICE.name, 
				   STORAGE_DEVICE.createTime, 
				   STORAGE_DEVICE.enableTime, 
				   STORAGE_DEVICE.expireTime, 
				   STORAGE_DEVICE.comment, 
				   IFNULL(IP.uuid, "") AS ip_uuid, 
				   IFNULL(IP.ipAddress, "") AS ip_address, 
				   IFNULL(IP.type, "") AS ip_type,
				   IFNULL(IP.role, "") AS ip_role,
				   IFNULL(IP.ipSetId, "") AS ip_set_id,
				   IFNULL(RACK.name, "") AS rack_name,
				   IFNULL(MAPPING_RACK_DEVICE.rackId, "") AS rack_id,
				   IFNULL(MAPPING_RACK_DEVICE.begPos, 0) AS beg_pos,
				   IFNULL(MAPPING_RACK_DEVICE.endPos, 0) AS end_pos
			FROM STORAGE_DEVICE LEFT JOIN IP ON STORAGE_DEVICE.uuid=IP.targetId
			     LEFT JOIN MAPPING_RACK_DEVICE ON MAPPING_RACK_DEVICE.deviceId=STORAGE_DEVICE.uuid
		         LEFT JOIN RACK ON RACK.uuid=MAPPING_RACK_DEVICE.rackId
			WHERE STORAGE_DEVICE.isDeleted=0 ORDER BY STORAGE_DEVICE.uuid`
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

	record := &structs.StorageDevice{
		IPAddresses: []*structs.IP{},
	}
	for rows.Next() {
		tmpRecord := &structs.StorageDevice{
			Position:    &structs.DevicePosition{},
			IPAddresses: []*structs.IP{},
		}
		ipAddress := &structs.IP{}
		if err = rows.Scan(&tmpRecord.UUID, &tmpRecord.Brand, &tmpRecord.Model, &tmpRecord.Name, &tmpRecord.CreateTime, &tmpRecord.EnableTime, &tmpRecord.ExpireTime, &tmpRecord.Comment, &ipAddress.UUID, &ipAddress.IPAddress, &ipAddress.Type, &ipAddress.Role, &ipAddress.IPSetId, &tmpRecord.Position.RackName, &tmpRecord.Position.RackUUID, &tmpRecord.Position.BegPos, &tmpRecord.Position.EndPos); err != nil {
			log.Errorln(err.Error())
			return nil, err
		}
		ipAddress.TargetId = tmpRecord.UUID

		if tmpRecord.UUID == record.UUID {
			if ipAddress.UUID != "" {
				record.IPAddresses = append(record.IPAddresses, ipAddress)
			}
		} else {
			record = tmpRecord
			if ipAddress.UUID != "" {
				record.IPAddresses = append(record.IPAddresses, ipAddress)
			}
			records = append(records, record)
		}
	}

	return records, nil
}

func (d *DeviceDAO) CreateCommonDevice(brand string, model string, name string, enableTime string, expireTime string, comment string) error {
	if err := mysql.DB.SimpleExec("INSERT INTO COMMON_DEVICE(uuid, brand, model, name, createTime, enableTime, expireTime, comment, isDeleted) VALUES(?, ?, ?, ?, NOW(), ?, ?, ?, 0)", util.GetUUID(), brand, model, name, enableTime, expireTime, comment); err != nil {
		log.Errorln(err.Error())
		return err
	}
	return nil
}

func (d *DeviceDAO) DeleteCommonDevice(uuid string) error {
	if err := mysql.DB.SimpleExec("UPDATE COMMON_DEVICE SET isDeleted=1 WHERE uuid=?", uuid); err != nil {
		log.Errorln(err.Error())
		return err
	}
	return nil
}

func (d *DeviceDAO) ListCommonDevices() ([]*structs.CommonDevice, error) {
	var err error
	tx := mysql.DB.GetTx()
	defer func() {
		if err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()

	records := []*structs.CommonDevice{}

	sql := `SELECT COMMON_DEVICE.uuid, 
	               COMMON_DEVICE.brand, 
	               COMMON_DEVICE.model, 
	               COMMON_DEVICE.name, 
	               COMMON_DEVICE.createTime, 
	               COMMON_DEVICE.enableTime, 
	               COMMON_DEVICE.expireTime, 
	               COMMON_DEVICE.comment, 
				   IFNULL(IP.uuid, "") AS ip_uuid, 
				   IFNULL(IP.ipAddress, "") AS ip_address, 
				   IFNULL(IP.type, "") AS ip_type,
				   IFNULL(IP.role, "") AS ip_role,
				   IFNULL(IP.ipSetId, "") AS ip_set_id,
				   IFNULL(RACK.name, "") AS rack_name,
				   IFNULL(MAPPING_RACK_DEVICE.rackId, "") AS rack_id,
				   IFNULL(MAPPING_RACK_DEVICE.begPos, 0) AS beg_pos,
				   IFNULL(MAPPING_RACK_DEVICE.endPos, 0) AS end_pos
			FROM COMMON_DEVICE LEFT JOIN IP ON COMMON_DEVICE.uuid=IP.targetId
			     LEFT JOIN MAPPING_RACK_DEVICE ON MAPPING_RACK_DEVICE.deviceId=COMMON_DEVICE.uuid
		         LEFT JOIN RACK ON RACK.uuid=MAPPING_RACK_DEVICE.rackId
			WHERE COMMON_DEVICE.isDeleted=0 ORDER BY COMMON_DEVICE.uuid`
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

	record := &structs.CommonDevice{
		IPAddresses: []*structs.IP{},
	}
	for rows.Next() {
		tmpRecord := &structs.CommonDevice{
			Position:    &structs.DevicePosition{},
			IPAddresses: []*structs.IP{},
		}
		ipAddress := &structs.IP{}
		if err = rows.Scan(&tmpRecord.UUID, &tmpRecord.Brand, &tmpRecord.Model, &tmpRecord.Name, &tmpRecord.CreateTime, &tmpRecord.EnableTime, &tmpRecord.ExpireTime, &tmpRecord.Comment, &ipAddress.UUID, &ipAddress.IPAddress, &ipAddress.Type, &ipAddress.Role, &ipAddress.IPSetId, &tmpRecord.Position.RackName, &tmpRecord.Position.RackUUID, &tmpRecord.Position.BegPos, &tmpRecord.Position.EndPos); err != nil {
			log.Errorln(err.Error())
			return nil, err
		}
		ipAddress.TargetId = tmpRecord.UUID

		if tmpRecord.UUID == record.UUID {
			if ipAddress.UUID != "" {
				record.IPAddresses = append(record.IPAddresses, ipAddress)
			}
		} else {
			record = tmpRecord
			if ipAddress.UUID != "" {
				record.IPAddresses = append(record.IPAddresses, ipAddress)
			}
			records = append(records, record)
		}
	}

	return records, nil
}
