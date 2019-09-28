package mysql

import (
	"be/option"
	"database/sql"

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
