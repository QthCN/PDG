package structs

type MonitorItemReleatedDevice struct {
	MappingId       int64  `json:"mappint_id"`
	MonitorItemId   int64  `json:"monitor_item_id"`
	MonitorItemName string `json:"monitor_item_name"`
	DeviceUUID      string `json:"device_uuid"`
	DeviceType      string `json:"device_type"`
	DeviceName      string `json:"device_name"`
}

type HistoryDataFilter struct {
	DeviceUUID     string `json:"device_uuid"`
	ItemId         int64  `json:"item_id"`
	QueryBeginDate string `json:"query_begin_date"`
	QueryEndDate   string `json:"query_end_date"`
}

type HistoryDataRecord struct {
	Key   string  `json:"key"`
	Value float64 `json:"value"`
}

type MonitorBackendCfg struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Cfg  string `json:"cfg"`

	FakeCfg   *MonitorBackendFakeCfg   `json:"cfg_fake"`
	ZabbixCfg *MonitorBackendZabbixCfg `json:"cfg_zabbix"`
}

type MonitorBackendFakeCfg struct {
	Address string `json:"address"`
}

type MonitorBackendZabbixCfg struct {
	Address  string `json:"address"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// 监控项
type MonitorItem struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	// 是否为自带监控项
	IsInternal int64 `json:"is_internal"`

	// 数据采集
	DCType    string     `json:"dc_type"`
	DCFakeCfg *DCFakeCfg `json:"dc_fake_cfg"`

	// 告警
	AlertType string `json:"alert_type"`
}

type DCFakeCfg struct {
	ItemName string `json:"item_name"`
	HostIp   string `json:"host_ip"`
}
