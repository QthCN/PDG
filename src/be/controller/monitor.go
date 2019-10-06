package controller

import (
	"math/rand"
)

type MonitorMgr struct {
}

var Monitor *MonitorMgr

func init() {
	Monitor = &MonitorMgr{}
}

func (m *MonitorMgr) GetDeviceStatus(uuid string) string {
	if rand.Intn(2) == 0 {
		return "BAD"
	}
	return "GOOD"
}
