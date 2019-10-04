package controller

import (
	"be/dao"
	"be/option"
	"be/structs"
	"fmt"

	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
)

type AuthMgr struct {
	dao *dao.AuthDAO
}

var Auth *AuthMgr

func init() {
	Auth = &AuthMgr{
		dao: &dao.AuthDAO{},
	}
}

func (m *AuthMgr) GenTokenByUsernameAndPassword(username string, password string) (string, error) {
	// 判断用户和密码是否有效
	if err := m.dao.ValidateUsernameAndPassword(username, password); err != nil {
		return "", err
	}

	// 生成token
	uuid, err := uuid.NewV4()
	if err != nil {
		log.Errorln(err.Error())
		return "", err
	}
	token := fmt.Sprintf("%s", uuid)
	if err := m.dao.CreateToken(username, token); err != nil {
		return "", err
	}

	return token, nil
}

func (m *AuthMgr) GetUserInfoByToken(token string) (*structs.UserInfo, error) {
	if *option.Mode == "DEV" {
		return &structs.UserInfo{
			Username: "ADMIN",
		}, nil
	}
	return m.dao.GetUserInfoByToken(token)
}

func (m *AuthMgr) ListUsers() ([]*structs.UserInfo, error) {
	return m.dao.ListUsers()
}

func (m *AuthMgr) CreateUser(username string, password string) error {
	return m.dao.CreateUser(username, password)
}

func (m *AuthMgr) RemoveUser(username string) error {
	return m.dao.RemoveUser(username)
}
