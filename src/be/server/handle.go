package server

import (
	"be/option"
	"net/http"
)

func (s *Server) registAPI() {
	// 初始化静态文件路径
	initStaticFileMapping(s)

	// 初始化api接口
	initApiMapping(s)

	// 初始化ajax接口
	initAjaxMapping(s)
}

func initStaticFileMapping(r *Server) {
	fs := http.FileServer(http.Dir(*option.StaticFilePath))
	r.GetRouter().PathPrefix("/js/").Handler(fs)
	r.GetRouter().PathPrefix("/css/").Handler(fs)
	r.GetRouter().PathPrefix("/img/").Handler(fs)
	r.GetRouter().Path("/favicon.ico").Handler(fs)
	r.GetRouter().Path("/login.html").Handler(fs)
	r.GetRouter().Path("/jquery-3.4.1.min.js").Handler(fs)
	r.GetRouter().Path("/vue.js").Handler(fs)
	r.GetRouter().Path("/element-ui.js").Handler(fs)
	r.GetRouter().Path("/theme-chalk.css").Handler(fs)

	r.GetRouter().Path("/").Handler(fs)
	r.GetRouter().NotFoundHandler = fs
}

func initApiMapping(r *Server) {
	r.RegistURLMapping("/v1/api/alert/record", "POST", apiRecordAlert)
}

func initAjaxMapping(r *Server) {
	// 用户认证密码并生成token
	r.RegistURLMapping("/v1/ajax/auth/token", "POST", ajaxGenTokenByUMAndPassword)
	// 登出
	r.RegistURLMapping("/v1/ajax/auth/logout", "GET", ajaxLogout)
	// 获取用户信息
	r.RegistURLMapping("/v1/ajax/auth/info", "GET", ajaxGetUserInfo)
	// 列出用户列表
	r.RegistURLMapping("/v1/ajax/auth/user/list", "GET", ajaxListUsers)
	// 创建用户
	r.RegistURLMapping("/v1/ajax/auth/user/create", "POST", ajaxCreateUser)
	// 删除用户
	r.RegistURLMapping("/v1/ajax/auth/user/remove", "POST", ajaxRemoveUser)

	// DataCenter
	r.RegistURLMapping("/v1/ajax/device/datacenter/create", "POST", ajaxCreateDataCenter)
	r.RegistURLMapping("/v1/ajax/device/datacenter/remove", "POST", ajaxDeleteDataCenter)
	r.RegistURLMapping("/v1/ajax/device/datacenter/list", "GET", ajaxListDataCenters)

	// 列出物理拓扑
	r.RegistURLMapping("/v1/ajax/device/topology/physical", "GET", ajaxGetPhysicalTopology)
	// 列出资源拓扑
	r.RegistURLMapping("/v1/ajax/device/topology/resource", "GET", ajaxGetResourceTopology)

	// Rack
	r.RegistURLMapping("/v1/ajax/device/rack/create", "POST", ajaxCreateRack)
	r.RegistURLMapping("/v1/ajax/device/rack/remove", "POST", ajaxDeleteRack)
	r.RegistURLMapping("/v1/ajax/device/rack/list", "GET", ajaxListRacks)
	r.RegistURLMapping("/v1/ajax/device/rack/map/datacenter", "POST", ajaxMapRackAndDatacenter)

	r.RegistURLMapping("/v1/ajax/device/map/rack", "POST", ajaxMapDeviceAndRack)

	// ServerDevice
	r.RegistURLMapping("/v1/ajax/device/server/create", "POST", ajaxCreateServerDevice)
	r.RegistURLMapping("/v1/ajax/device/server/remove", "POST", ajaxDeleteServerDevice)
	r.RegistURLMapping("/v1/ajax/device/server/list", "GET", ajaxListServerDevices)

	// NetworkDevice
	r.RegistURLMapping("/v1/ajax/device/network/create", "POST", ajaxCreateNetworkDevice)
	r.RegistURLMapping("/v1/ajax/device/network/remove", "POST", ajaxDeleteNetworkDevice)
	r.RegistURLMapping("/v1/ajax/device/network/list", "GET", ajaxListNetworkDevices)

	// StorageDevice
	r.RegistURLMapping("/v1/ajax/device/storage/create", "POST", ajaxCreateStorageDevice)
	r.RegistURLMapping("/v1/ajax/device/storage/remove", "POST", ajaxDeleteStorageDevice)
	r.RegistURLMapping("/v1/ajax/device/storage/list", "GET", ajaxListStorageDevices)

	// CommonDevice
	r.RegistURLMapping("/v1/ajax/device/common/create", "POST", ajaxCreateCommonDevice)
	r.RegistURLMapping("/v1/ajax/device/common/remove", "POST", ajaxDeleteCommonDevice)
	r.RegistURLMapping("/v1/ajax/device/common/list", "GET", ajaxListCommonDevices)

	// 设备通用接口
	r.RegistURLMapping("/v1/ajax/device/list", "GET", ajaxListDevices)
	r.RegistURLMapping("/v1/ajax/device/get", "GET", ajaxGetDeviceByUUID)

	// IP
	r.RegistURLMapping("/v1/ajax/ip/create", "POST", ajaxCreateIPRecord)
	r.RegistURLMapping("/v1/ajax/ip/remove", "POST", ajaxDeleteIPRecord)
	r.RegistURLMapping("/v1/ajax/ip/list", "GET", ajaxListIPs)

	// IPSet
	r.RegistURLMapping("/v1/ajax/ipset/create", "POST", ajaxCreateIPSetRecord)
	r.RegistURLMapping("/v1/ajax/ipset/remove", "POST", ajaxDeleteIPSetRecord)
	r.RegistURLMapping("/v1/ajax/ipset/list", "GET", ajaxListIPSets)

	// Connection
	r.RegistURLMapping("/v1/ajax/connection/create", "POST", ajaxCreateConnection)
	r.RegistURLMapping("/v1/ajax/connection/remove", "POST", ajaxDeleteConnection)
	r.RegistURLMapping("/v1/ajax/connection/list", "GET", ajaxListConnections)

	// 查看审计记录
	r.RegistURLMapping("/v1/ajax/audit/list", "GET", ajaxListAuditRecords)

	// 监控
	r.RegistURLMapping("/v1/ajax/monitor/item/list", "GET", ajaxListMonitorItems)
	r.RegistURLMapping("/v1/ajax/monitor/item/create", "POST", ajaxCreateMonitorItem)
	r.RegistURLMapping("/v1/ajax/monitor/item/delete", "POST", ajaxDeleteMonitorItem)
	r.RegistURLMapping("/v1/ajax/monitor/item/update", "POST", ajaxUpdateMonitorItem)
	r.RegistURLMapping("/v1/ajax/monitor/item/detail", "GET", ajaxGetMonitorItemById)
	r.RegistURLMapping("/v1/ajax/monitor/item/dc/update", "POST", ajaxUpdateMonitorItemDCCfg)
	r.RegistURLMapping("/v1/ajax/monitor/item/device/list", "GET", ajaxListMonitorItemReleatedDevices)
	r.RegistURLMapping("/v1/ajax/monitor/device/item/list", "GET", ajaxListDeviceReleatedMonitorItems)
	r.RegistURLMapping("/v1/ajax/monitor/item/device/bind", "POST", ajaxBindMonitorItemAndDevice)

	// 监控服务
	r.RegistURLMapping("/v1/ajax/monitor/backend/list", "GET", ajaxListMonitorBackendCfgs)
	r.RegistURLMapping("/v1/ajax/monitor/backend/update", "POST", ajaxUpdateMonitorBackendCfg)

	// 监控数据获取
	r.RegistURLMapping("/v1/ajax/monitor/history/query", "GET", ajaxGetDeviceMonitorItemHistoryData)

	// 告警
	r.RegistURLMapping("/v1/ajax/alert/list", "GET", ajaxListAlertItems)
	r.RegistURLMapping("/v1/ajax/alert/create", "POST", ajaxCreateAlertItem)
	r.RegistURLMapping("/v1/ajax/alert/delete", "POST", ajaxDeleteAlertItem)
	r.RegistURLMapping("/v1/ajax/alert/event/list", "GET", ajaxListAlertEvent)
}
