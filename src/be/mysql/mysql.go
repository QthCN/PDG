package mysql

import (
	"be/option"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

type MySQLUtil struct {
	db          *sql.DB
	initialized bool
}

var DB = MySQLUtil{db: nil, initialized: false}

func (m *MySQLUtil) InitConn() {
	m.CloseConn()

	db, err := sql.Open("mysql", *option.DataSourceName)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("MySQL conn error")
		panic(err)
	}

	db.SetMaxOpenConns(*option.DBMaxOpenConn)
	db.SetMaxIdleConns(*option.DBMaxIdleConn)

	err = db.Ping()
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("MySQL conn error")
		panic(err)
	}

	m.db = db
	m.initialized = true
	log.WithFields(log.Fields{}).Info("MySQL conn successful")
}

func (m *MySQLUtil) CloseConn() {
	// 这个方法正常应该不会被调用，并且如果被调用的话会影响正常的使用
	if m.initialized {
		m.db.Close()
		m.db = nil
		m.initialized = false
	}
}

func (m *MySQLUtil) GetConn() *sql.DB {
	if m.initialized == false {
		log.WithFields(log.Fields{
			"err": "DB还没有初始化",
		}).Error("MySQL GET conn error")
		return nil
	}
	return m.db
}

func (m *MySQLUtil) GetTx() *sql.Tx {
	if m.initialized == false {
		log.WithFields(log.Fields{
			"err": "DB还没有初始化",
		}).Error("Get txn error")
		return nil
	}

	tx, err := m.db.Begin()
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("Get txn error")
		return nil
	}
	return tx
}

// 查询返回0行或者1行的数据
func (m *MySQLUtil) SingleRowQuery(sql string, args []interface{}, result ...interface{}) (int64, error) {
	if m.initialized == false {
		log.WithFields(log.Fields{
			"err": "DB还没有初始化",
		}).Error("Get txn error")
		return -1, fmt.Errorf("DB ERROR")
	}

	tx := m.GetTx()
	stmt, err := tx.Prepare(sql)
	if err != nil {
		log.WithFields(log.Fields{
			"sql": sql,
			"err": err.Error(),
		}).Error("SingleRowQuery prepare错误")
		tx.Rollback()
		return -1, fmt.Errorf("DB ERROR")
	}
	rows, err := stmt.Query(args...)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("SingleRowQuery query错误")
		stmt.Close()
		tx.Rollback()
		return -1, fmt.Errorf("DB ERROR")
	}
	var cnt int64 = 0
	for rows.Next() {
		err := rows.Scan(result...)
		if err != nil {
			log.WithFields(log.Fields{
				"err": err.Error(),
			}).Error("SingleRowQuery rows.Next错误")
			rows.Close()
			stmt.Close()
			tx.Rollback()
			return -1, fmt.Errorf("DB ERROR")
		} else {
			cnt += 1
			break
		}
	}
	err = rows.Err()
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("SingleRowQuery rows.Err错误")
		rows.Close()
		stmt.Close()
		tx.Rollback()
		return -1, fmt.Errorf("DB ERROR")
	}
	rows.Close()
	stmt.Close()
	tx.Commit()
	return cnt, nil

}

func (m *MySQLUtil) SimpleExec(sql string, args ...interface{}) error {
	if m.initialized == false {
		log.WithFields(log.Fields{
			"err": "DB还没有初始化",
		}).Error("Get txn error")
		return fmt.Errorf("DB ERROR")
	}

	tx := m.GetTx()
	stmt, err := tx.Prepare(sql)
	if err != nil {
		log.WithFields(log.Fields{
			"sql": sql,
			"err": err.Error(),
		}).Error("SimpleExec prepare错误")
		tx.Rollback()
		return err
	}
	_, err = stmt.Exec(args...)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("SimpleExec exec错误")
		stmt.Close()
		tx.Rollback()
		return err
	}
	stmt.Close()
	err = tx.Commit()
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("SimpleExec commit错误")
		return err
	}
	return nil
}
