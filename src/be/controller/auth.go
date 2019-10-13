package controller

import (
	"be/dao"
	"be/option"
	"be/structs"
	"be/util"
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
	token := util.GetUUID()
	if err := m.dao.CreateToken(username, token); err != nil {
		return "", err
	}

	return token, nil
}

func (m *AuthMgr) GetUserInfoByToken(token string) (*structs.UserInfo, error) {
	if *option.Mode == "DEV" {
		return &structs.UserInfo{
			Username: "ADMIN",
			Role: "管理员",
		}, nil
	}
	return m.dao.GetUserInfoByToken(token)
}

func (m *AuthMgr) ListUsers() ([]*structs.UserInfo, error) {
	return m.dao.ListUsers()
}

func (m *AuthMgr) CreateUser(username string, password string, role string, mobile string, mail string, wx string) error {
	return m.dao.CreateUser(username, password, role, mobile, mail, wx)
}

func (m *AuthMgr) RemoveUser(username string) error {
	return m.dao.RemoveUser(username)
}
