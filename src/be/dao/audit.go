package dao

import (
	"be/mysql"
	"be/structs"
	"fmt"

	log "github.com/sirupsen/logrus"
)

type AuditDAO struct {
}

func (d *AuditDAO) CreateRecord(username string, action string, url string, args string) error {
	err := mysql.DB.SimpleExec("INSERT INTO AUDIT(username, action, url, args, actionTime) VALUES(?, ?, ?, ?, NOW())", username, action, url, args)
	if err != nil {
		log.Errorln(err.Error())
		return err
	}
	return nil
}

func (d *AuditDAO) ListRecords(filter *structs.ListAuditRecordsCondition) (*structs.AuditRecords, error) {
	records := &structs.AuditRecords{
		Records: []*structs.AuditRecord{},
	}

	sqlTemplate := "SELECT %s FROM AUDIT ORDER BY %s DESC LIMIT %s"

	dataSql := fmt.Sprintf(sqlTemplate, "id, username, action, url, args, actionTime", "actionTime", fmt.Sprintf(" %d, %d ", filter.RecordsPerPage*(filter.CurrentPage-1), filter.RecordsPerPage))
	cntSql := fmt.Sprintf(sqlTemplate, "COUNT(id) AS CNT", "CNT", "1")

	queryData := func(sql string) ([]*structs.AuditRecord, error) {
		var err error
		tx := mysql.DB.GetTx()
		defer func() {
			if err == nil {
				tx.Commit()
			} else {
				tx.Rollback()
			}
		}()

		records := []*structs.AuditRecord{}
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
			record := &structs.AuditRecord{}
			if err = rows.Scan(&record.Id, &record.Username, &record.Action, &record.URL, &record.Args, &record.ActionTime); err != nil {
				log.Errorln(err.Error())
				return nil, err
			}
			records = append(records, record)
		}

		return records, nil
	}

	rs, err := queryData(dataSql)
	if err != nil {
		log.Errorln(err.Error())
		return nil, err
	}
	records.Records = rs

	queryCnt := func(sql string) (int64, error) {
		var cnt int64
		var err error
		tx := mysql.DB.GetTx()
		defer func() {
			if err == nil {
				tx.Commit()
			} else {
				tx.Rollback()
			}
		}()

		stmt, err := tx.Prepare(sql)
		if err != nil {
			log.Errorln(err.Error())
			return -1, err
		}
		defer stmt.Close()

		rows, err := stmt.Query()
		if err != nil {
			log.Errorln(err.Error())
			return -1, err
		}
		defer rows.Close()

		for rows.Next() {
			if err = rows.Scan(&cnt); err != nil {
				log.Errorln(err.Error())
				return -1, err
			}
		}

		return cnt, nil
	}

	cs, err := queryCnt(cntSql)
	if err != nil {
		log.Errorln(err.Error())
		return nil, err
	}
	records.TotalCnt = cs

	return records, nil
}
