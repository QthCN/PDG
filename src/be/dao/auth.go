package dao

import (
	"be/mysql"
	"be/option"
	"be/structs"
	"crypto/md5"
	"fmt"
	"io"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

type AuthDAO struct {
}

func (d *AuthDAO) ValidateUsernameAndPassword(username string, password string) error {
	w := md5.New()
	io.WriteString(w, password)

	var c int64 = 0
	_, err := mysql.DB.SingleRowQuery("SELECT COUNT(username) as CNT FROM USER WHERE UPPER(username)=? AND epassword=?", []interface{}{strings.ToUpper(strings.TrimSpace(username)), fmt.Sprintf("%x", w.Sum(nil))}, &c)
	if err != nil {
		log.Errorln(err.Error())
		return err
	}

	if c == 0 {
		return fmt.Errorf("认证失败")
	} else {
		return nil
	}
}

func (d *AuthDAO) CreateToken(username string, token string) error {
	err := mysql.DB.SimpleExec("INSERT INTO TOKEN(token, username, expireTime) VALUES(?, ?, ?)", token, strings.ToUpper(strings.TrimSpace(username)), time.Now().Add(time.Duration(*option.TokenExpireInterval)*time.Hour).Format("2006-01-02 15:04:05"))
	if err != nil {
		log.Errorln(err.Error())
		return err
	}
	return nil
}

func (d *AuthDAO) GetUserInfoByToken(token string) (*structs.UserInfo, error) {
	userInfo := &structs.UserInfo{}

	_, err := mysql.DB.SingleRowQuery("SELECT USER.username FROM USER, TOKEN WHERE TOKEN.expireTime>NOW() AND TOKEN.token=? AND TOKEN.username=USER.username", []interface{}{token}, &userInfo.Username)
	if err != nil {
		log.Errorln(err.Error())
		return nil, err
	}

	return userInfo, nil
}

func (d *AuthDAO) ListUsers() ([]*structs.UserInfo, error) {
	var err error
	tx := mysql.DB.GetTx()
	defer func(){
		if err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()

	users := []*structs.UserInfo{}

	sql := `SELECT username FROM USER`
	
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
		user := &structs.UserInfo{}
		if err = rows.Scan(&user.Username); err != nil {
			log.Errorln(err.Error())
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (d *AuthDAO) CreateUser(username string, password string) error {
	w := md5.New()
	io.WriteString(w, password)

	err := mysql.DB.SimpleExec("INSERT INTO USER(username, epassword) VALUES(?, ?)", username, fmt.Sprintf("%x", w.Sum(nil)))
	if err != nil {
		log.Errorln(err.Error())
		return err
	}
	return nil
}

func (d *AuthDAO) RemoveUser(username string) error {
	err := mysql.DB.SimpleExec("DELETE FROM USER WHERE username=?", username)
	if err != nil {
		log.Errorln(err.Error())
		return err
	}
	return nil
}