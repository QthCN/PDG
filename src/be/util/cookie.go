package util

import (
	"be/option"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/securecookie"
	log "github.com/sirupsen/logrus"
)

type cookieMgr struct {
	s *securecookie.SecureCookie
}

var CM *cookieMgr

func InitCM() {
	s := securecookie.New([]byte(*option.Cookie01), []byte(*option.Cookie02))
	CM = &cookieMgr{s: s}
}

func (c *cookieMgr) Set(key string, value string, res http.ResponseWriter) {
	expiration := time.Now().Add(time.Duration(24) * time.Hour)
	if encoded, err := c.s.Encode(key, value); err == nil {
		cookie := &http.Cookie{
			Name:    key,
			Value:   encoded,
			Expires: expiration,
			Path:    "/",
		}
		http.SetCookie(res, cookie)
	} else {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("Cookie Set失败")
	}
}

func (c *cookieMgr) Get(key string, req *http.Request) (string, error) {
	if cookie, err := req.Cookie(key); err == nil {
		value := ""
		if err = c.s.Decode(key, cookie.Value, &value); err == nil {
			return value, nil
		} else {
			log.WithFields(log.Fields{
				"err": err.Error(),
			}).Error("Cookie Get失败")
			return "", fmt.Errorf("Cookie Get失败")
		}
	} else {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("Cookie Get失败")
		return "", fmt.Errorf("Cookie Get失败")
	}
}

func (c *cookieMgr) Remove(key string, res http.ResponseWriter) {
	// 设置为空字符串即认为删除
	c.Set(key, "", res)
}
