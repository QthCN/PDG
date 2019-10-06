package dao

import (
	"be/mysql"
	"be/structs"
	"be/util"

	log "github.com/sirupsen/logrus"
)

type IpDAO struct {
}

func (d *IpDAO) ListIPSets() ([]*structs.IPSet, error) {
	var err error
	tx := mysql.DB.GetTx()
	defer func() {
		if err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()

	records := []*structs.IPSet{}

	sql := `SELECT IPSet.uuid, IPSet.cidr, IPSet.comment FROM IPSet WHERE IPSet.isDeleted=0 ORDER BY IPSet.cidr`
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
		record := &structs.IPSet{}

		if err = rows.Scan(&record.UUID, &record.CIDR, &record.Comment); err != nil {
			log.Errorln(err.Error())
			return nil, err
		}

		records = append(records, record)
	}

	return records, nil
}

func (d *IpDAO) ListIPRecords() ([]*structs.IP, error) {
	var err error
	tx := mysql.DB.GetTx()
	defer func() {
		if err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()

	records := []*structs.IP{}

	sql := `SELECT IP.uuid, IP.ipAddress, IP.type, IP.role, IP.targetId, IP.ipSetId FROM IP ORDER BY IP.ipAddress`
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
		record := &structs.IP{}

		if err = rows.Scan(&record.UUID, &record.IPAddress, &record.Type, &record.Role, &record.TargetId, &record.IPSetId); err != nil {
			log.Errorln(err.Error())
			return nil, err
		}

		records = append(records, record)
	}

	return records, nil
}

func (d *IpDAO) CreateIPSet(cidr string, comment string) error {
	err := mysql.DB.SimpleExec("INSERT INTO IPSet(uuid, cidr, comment, isDeleted) VALUES(?, ?, ?, 0)", util.GetUUID(), cidr, comment)
	if err != nil {
		log.Errorln(err.Error())
		return err
	}
	return nil
}

func (d *IpDAO) DeleteIPSet(uuid string) error {
	err := mysql.DB.SimpleExec("UPDATE IPSet SET isDeleted=1 WHERE uuid=?", uuid)
	if err != nil {
		log.Errorln(err.Error())
		return err
	}
	return nil
}

func (d *IpDAO) CreateIPRecord(ipAddress string, ipType string, role string, targetId string, ipSetId string) error {
	err := mysql.DB.SimpleExec("INSERT INTO IP(uuid, ipAddress, type, role, targetId, ipSetId) VALUES(?, ?, ?, ?, ?, ?)", util.GetUUID(), ipAddress, ipType, role, targetId, ipSetId)
	if err != nil {
		log.Errorln(err.Error())
		return err
	}
	return nil
}

func (d *IpDAO) DeleteIPRecord(uuid string) error {
	err := mysql.DB.SimpleExec("DELETE FROM IP WHERE uuid=?", uuid)
	if err != nil {
		log.Errorln(err.Error())
		return err
	}
	return nil
}
