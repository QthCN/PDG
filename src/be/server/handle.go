package server

import (
	"be/option"
	"net/http"
)

func (s *Server) registAPI() {
	// 初始化静态文件路径
	initStaticFileMapping(s)

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

	r.GetRouter().Path("/").Handler(fs)
	r.GetRouter().NotFoundHandler = fs
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
}
