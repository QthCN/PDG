package controller

import (
	"be/structs"
)

type AuthMgr struct {
}

var Auth *AuthMgr

func init() {
	Auth = &AuthMgr{}
}

func (m *AuthMgr) GenTokenByUsernameAndPassword(username string, password string) (string, error) {
	return "", nil
}
func (m *AuthMgr) GetUserInfoByToken(token string) (*structs.UserInfo, error) {
	return nil, nil
}
