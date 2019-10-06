package dao

import (
	"be/mysql"
	"be/structs"
	"be/util"

	log "github.com/sirupsen/logrus"
)

type ConnectionDAO struct {
}

func (d *ConnectionDAO) ListConnections() ([]*structs.Connection, error) {
	var err error
	tx := mysql.DB.GetTx()
	defer func() {
		if err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()

	records := []*structs.Connection{}

	sql := `SELECT CONNECTION.uuid, CONNECTION.sourceId, CONNECTION.sourcePort, CONNECTION.sourceDeviceType, CONNECTION.sourceDeviceName, CONNECTION.destinationId, CONNECTION.destinationPort, CONNECTION.destinationDeviceType, CONNECTION.destinationDeviceName
			FROM CONNECTION
			ORDER BY CONNECTION.sourceId, CONNECTION.destinationId`

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
		record := &structs.Connection{}
		if err = rows.Scan(&record.UUID, &record.SourceId, &record.SourcePort, &record.SourceDeviceType, &record.SourceDeviceName, &record.DestinationId, &record.DestinationPort, &record.DestinationDeviceType, &record.DestinationDeviceName); err != nil {
			log.Errorln(err.Error())
			return nil, err
		}
		records = append(records, record)
	}

	return records, nil
}

func (d *ConnectionDAO) CreateConnection(sourceId string, sourcePort string, sourceDeviceType string, sourceDeviceName string, destinationId string, destinationPort string, destinationDeviceType string, destinationDeviceName string) error {
	err := mysql.DB.SimpleExec("INSERT INTO CONNECTION(uuid, sourceId, sourcePort, sourceDeviceType, sourceDeviceName, destinationId, destinationPort, destinationDeviceType, destinationDeviceName) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)", util.GetUUID(), sourceId, sourcePort, sourceDeviceType, sourceDeviceName, destinationId, destinationPort, destinationDeviceType, destinationDeviceName)
	if err != nil {
		log.Errorln(err.Error())
		return err
	}
	return nil
}

func (d *ConnectionDAO) DeleteConnection(uuid string) error {
	err := mysql.DB.SimpleExec("DELETE FROM CONNECTION WHERE uuid=?", uuid)
	if err != nil {
		log.Errorln(err.Error())
		return err
	}
	return nil
}
