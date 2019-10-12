package util

import (
	"encoding/json"
	"fmt"

	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
)

func GetUUID() string {
	uuid, _ := uuid.NewV4()
	return fmt.Sprintf("%s", uuid)
}

func ParseJsonStr(str string, body interface{}) error {
	data := []byte(str)
	err := json.Unmarshal(data, body)
	if err != nil {
		log.WithFields(log.Fields{
			"body": str,
			"err":  err.Error(),
		}).Error("报文解析失败")
		return err
	} else {
		return nil
	}
}
