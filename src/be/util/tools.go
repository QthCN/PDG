package util

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

func GetUUID() string {
	uuid, _ := uuid.NewV4()
	return fmt.Sprintf("%s", uuid)
}
