package server

import (
	"be/controller"
	"be/structs"
	"be/util"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func isAdmin(token string) bool {
	userInfo, err := controller.Auth.GetUserInfoByToken(token)
	if err != nil {
		log.Errorln(err.Error())
		return false
	}
	if userInfo.Role == "管理员" {
		return true
	}
	return false
}

func audit(token string, action string, url string, args string) {
	userInfo, err := controller.Auth.GetUserInfoByToken(token)
	if err != nil {
		log.Errorln(err.Error())
	} else {
		controller.Audit.CreateRecord(userInfo.Username, action, url, args)
	}
}

func ajaxLogout(res http.ResponseWriter, req *http.Request) {
	util.CM.Remove("token", res)
	http.Redirect(res, req, "/login.html", http.StatusTemporaryRedirect)
}

func ajaxGenTokenByUMAndPassword(res http.ResponseWriter, req *http.Request) {
	reqContent, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		log.WithFields(log.Fields{}).Error("请求报文解析失败")
		ResInvalidRequestBody(res)
		return
	}

	type Request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	request := &Request{}
	if err := ParseJsonStr(string(reqContent), request); err != nil {
		log.Errorln("解析模板JSON失败")
		ResMsg(res, 400, err.Error())
		return
	}

	token, err := controller.Auth.GenTokenByUsernameAndPassword(request.Username, request.Password)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("GenToken失败")
		ResMsg(res, 400, err.Error())
		return
	}

	// audit
	audit(token, "登陆认证", "", fmt.Sprintf("用户名: %s", request.Username))

	// 在session中记录token
	util.CM.Set("token", token, res)
	ResSuccessMsg(res, 200, "token生成成功")
}

func ajaxGetUserInfo(res http.ResponseWriter, req *http.Request) {
	token, err := util.CM.Get("token", req)
	if err != nil || token == "" {
		ResMsg(res, 400, "请求中未包含token")
		return
	}
	userInfo, err := controller.Auth.GetUserInfoByToken(token)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("失败")
		ResMsg(res, 400, err.Error())
		return
	}
	b, err := json.Marshal(userInfo)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("JSON生成失败")
		ResMsg(res, 500, err.Error())
		return
	}
	ResMsg(res, 200, string(b))
}

func ajaxListUsers(res http.ResponseWriter, req *http.Request) {
	token, err := util.CM.Get("token", req)
	if err != nil || token == "" {
		ResMsg(res, 400, "请求中未包含token")
		return
	}

	users, err := controller.Auth.ListUsers()
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("根据token获取用户信息失败")
		ResMsg(res, 400, err.Error())
		return
	}
	b, err := json.Marshal(users)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("根据token获取用户信息失败 JSON生成失败")
		ResMsg(res, 500, err.Error())
		return
	}
	ResMsg(res, 200, string(b))

	// audit
	audit(token, "列出用户列表", "", "")
}

func ajaxCreateUser(res http.ResponseWriter, req *http.Request) {
	token, err := util.CM.Get("token", req)
	if err != nil || token == "" {
		ResMsg(res, 400, "请求中未包含token")
		return
	}

	if isAdmin(token) == false {
		ResMsg(res, 400, "权限不足")
		return
	}

	reqContent, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		log.WithFields(log.Fields{}).Error("请求报文解析失败")
		ResInvalidRequestBody(res)
		return
	}

	type Request struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Role     string `json:"role"`
		Mobile   string `json:"mobile"`
		Mail     string `json:"mail"`
		WX       string `json:"wx"`
	}

	request := &Request{}
	if err := ParseJsonStr(string(reqContent), request); err != nil {
		log.Errorln("解析模板JSON失败")
		ResMsg(res, 400, err.Error())
		return
	}

	// audit
	audit(token, "新建用户", "", fmt.Sprintf("用户名: %s", request.Username))

	err = controller.Auth.CreateUser(request.Username, request.Password, request.Role, request.Mobile, request.Mail, request.WX)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("失败")
		ResMsg(res, 400, err.Error())
		return
	}
	ResSuccessMsg(res, 200, "操作成功")
}

func ajaxRemoveUser(res http.ResponseWriter, req *http.Request) {
	token, err := util.CM.Get("token", req)
	if err != nil || token == "" {
		ResMsg(res, 400, "请求中未包含token")
		return
	}

	if isAdmin(token) == false {
		ResMsg(res, 400, "权限不足")
		return
	}

	reqContent, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		log.WithFields(log.Fields{}).Error("请求报文解析失败")
		ResInvalidRequestBody(res)
		return
	}

	type Request struct {
		Username string `json:"username"`
	}

	request := &Request{}
	if err := ParseJsonStr(string(reqContent), request); err != nil {
		log.Errorln("解析模板JSON失败")
		ResMsg(res, 400, err.Error())
		return
	}

	err = controller.Auth.RemoveUser(request.Username)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("失败")
		ResMsg(res, 400, err.Error())
		return
	}
	ResSuccessMsg(res, 200, "操作成功")

	// audit
	audit(token, "删除用户", "", fmt.Sprintf("被删除用户名: %s", request.Username))
}

func ajaxCreateDataCenter(res http.ResponseWriter, req *http.Request) {
	token, err := util.CM.Get("token", req)
	if err != nil || token == "" {
		ResMsg(res, 400, "请求中未包含token")
		return
	}

	if isAdmin(token) == false {
		ResMsg(res, 400, "权限不足")
		return
	}

	reqContent, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		log.WithFields(log.Fields{}).Error("请求报文解析失败")
		ResInvalidRequestBody(res)
		return
	}

	type Request struct {
		Name string `json:"name"`
	}

	request := &Request{}
	if err := ParseJsonStr(string(reqContent), request); err != nil {
		log.Errorln("解析模板JSON失败")
		ResMsg(res, 400, err.Error())
		return
	}

	err = controller.Device.CreateDataCenter(request.Name)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("失败")
		ResMsg(res, 400, err.Error())
		return
	}
	ResSuccessMsg(res, 200, "操作成功")

	// audit
	audit(token, "新建机房", "", fmt.Sprintf("机房名: %s", request.Name))
}

func ajaxDeleteDataCenter(res http.ResponseWriter, req *http.Request) {
	token, err := util.CM.Get("token", req)
	if err != nil || token == "" {
		ResMsg(res, 400, "请求中未包含token")
		return
	}

	if isAdmin(token) == false {
		ResMsg(res, 400, "权限不足")
		return
	}

	reqContent, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		log.WithFields(log.Fields{}).Error("请求报文解析失败")
		ResInvalidRequestBody(res)
		return
	}

	type Request struct {
		UUID string `json:"uuid"`
	}

	request := &Request{}
	if err := ParseJsonStr(string(reqContent), request); err != nil {
		log.Errorln("解析模板JSON失败")
		ResMsg(res, 400, err.Error())
		return
	}

	err = controller.Device.DeleteDataCenter(request.UUID)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("失败")
		ResMsg(res, 400, err.Error())
		return
	}
	ResSuccessMsg(res, 200, "操作成功")

	// audit
	audit(token, "删除机房", "", fmt.Sprintf("机房ID: %s", request.UUID))
}

func ajaxGetResourceTopology(res http.ResponseWriter, req *http.Request) {
	token, err := util.CM.Get("token", req)
	if err != nil || token == "" {
		ResMsg(res, 400, "请求中未包含token")
		return
	}

	topology, err := controller.Device.GetResourceTopology()
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("失败")
		ResMsg(res, 400, err.Error())
		return
	}
	b, err := json.Marshal(topology)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("JSON生成失败")
		ResMsg(res, 500, err.Error())
		return
	}
	ResMsg(res, 200, string(b))

	// audit
	audit(token, "查看资源拓扑", "", "")
}

func ajaxGetPhysicalTopology(res http.ResponseWriter, req *http.Request) {
	token, err := util.CM.Get("token", req)
	if err != nil || token == "" {
		ResMsg(res, 400, "请求中未包含token")
		return
	}

	reqContent, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		log.WithFields(log.Fields{}).Error("请求报文解析失败")
		ResInvalidRequestBody(res)
		return
	}

	type Request struct {
		UUID string `json:"uuid"`
	}

	request := &Request{}
	if err := ParseJsonStr(string(reqContent), request); err != nil {
		log.Errorln("解析模板JSON失败")
		ResMsg(res, 400, err.Error())
		return
	}

	topology, err := controller.Device.GetPhysicalTopology(request.UUID)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("失败")
		ResMsg(res, 400, err.Error())
		return
	}
	b, err := json.Marshal(topology)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("JSON生成失败")
		ResMsg(res, 500, err.Error())
		return
	}
	ResMsg(res, 200, string(b))

	// audit
	audit(token, "查看物理拓扑", "", "")
}

func ajaxListDataCenters(res http.ResponseWriter, req *http.Request) {
	token, err := util.CM.Get("token", req)
	if err != nil || token == "" {
		ResMsg(res, 400, "请求中未包含token")
		return
	}

	records, err := controller.Device.ListDataCenters()
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("失败")
		ResMsg(res, 400, err.Error())
		return
	}
	b, err := json.Marshal(records)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("JSON生成失败")
		ResMsg(res, 500, err.Error())
		return
	}
	ResMsg(res, 200, string(b))

	// audit
	audit(token, "查看数据中心列表", "", "")
}

func ajaxMapDeviceAndRack(res http.ResponseWriter, req *http.Request) {
	token, err := util.CM.Get("token", req)
	if err != nil || token == "" {
		ResMsg(res, 400, "请求中未包含token")
		return
	}

	if isAdmin(token) == false {
		ResMsg(res, 400, "权限不足")
		return
	}

	reqContent, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		log.WithFields(log.Fields{}).Error("请求报文解析失败")
		ResInvalidRequestBody(res)
		return
	}

	type Request struct {
		RackId     string `json:"rack_id"`
		DeviceId   string `json:"device_id"`
		DeviceType string `json:"device_type"`
		BegPos     int64  `json:"beg_pos"`
		EndPos     int64  `json:"end_pos"`
	}

	request := &Request{}
	if err := ParseJsonStr(string(reqContent), request); err != nil {
		log.Errorln("解析模板JSON失败")
		ResMsg(res, 400, err.Error())
		return
	}

	err = controller.Device.MapDeviceAndRack(request.RackId, request.DeviceId, request.DeviceType, request.BegPos, request.EndPos)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("失败")
		ResMsg(res, 400, err.Error())
		return
	}
	ResSuccessMsg(res, 200, "操作成功")

	// audit
	audit(token, "关联设备和机柜", "", fmt.Sprintf("机柜ID %s, 设备ID %s", request.RackId, request.DeviceId))
}

func ajaxMapRackAndDatacenter(res http.ResponseWriter, req *http.Request) {
	token, err := util.CM.Get("token", req)
	if err != nil || token == "" {
		ResMsg(res, 400, "请求中未包含token")
		return
	}

	if isAdmin(token) == false {
		ResMsg(res, 400, "权限不足")
		return
	}

	reqContent, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		log.WithFields(log.Fields{}).Error("请求报文解析失败")
		ResInvalidRequestBody(res)
		return
	}

	type Request struct {
		RackId       string `json:"rack_id"`
		DatacenterId string `json:"datacenter_id"`
		PositionX    int64  `json:"position_x"`
		PositionZ    int64  `json:"position_z"`
	}

	request := &Request{}
	if err := ParseJsonStr(string(reqContent), request); err != nil {
		log.Errorln("解析模板JSON失败")
		ResMsg(res, 400, err.Error())
		return
	}

	err = controller.Device.MapRackAndDatacenter(request.RackId, request.DatacenterId, request.PositionX, request.PositionZ)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("失败")
		ResMsg(res, 400, err.Error())
		return
	}
	ResSuccessMsg(res, 200, "操作成功")

	// audit
	audit(token, "关联机房机柜", "", fmt.Sprintf("机柜ID %s, 机房ID %s", request.RackId, request.DatacenterId))
}

func ajaxCreateRack(res http.ResponseWriter, req *http.Request) {
	token, err := util.CM.Get("token", req)
	if err != nil || token == "" {
		ResMsg(res, 400, "请求中未包含token")
		return
	}

	if isAdmin(token) == false {
		ResMsg(res, 400, "权限不足")
		return
	}

	reqContent, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		log.WithFields(log.Fields{}).Error("请求报文解析失败")
		ResInvalidRequestBody(res)
		return
	}

	type Request struct {
		Name string `json:"name"`
		Size int64  `json:"size"`
	}

	request := &Request{}
	if err := ParseJsonStr(string(reqContent), request); err != nil {
		log.Errorln("解析模板JSON失败")
		ResMsg(res, 400, err.Error())
		return
	}

	err = controller.Device.CreateRack(request.Name, request.Size)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("失败")
		ResMsg(res, 400, err.Error())
		return
	}
	ResSuccessMsg(res, 200, "操作成功")

	// audit
	audit(token, "创建机柜", "", fmt.Sprintf("机柜名: %s", request.Name))
}

func ajaxDeleteRack(res http.ResponseWriter, req *http.Request) {
	token, err := util.CM.Get("token", req)
	if err != nil || token == "" {
		ResMsg(res, 400, "请求中未包含token")
		return
	}

	if isAdmin(token) == false {
		ResMsg(res, 400, "权限不足")
		return
	}

	reqContent, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		log.WithFields(log.Fields{}).Error("请求报文解析失败")
		ResInvalidRequestBody(res)
		return
	}

	type Request struct {
		UUID string `json:"uuid"`
	}

	request := &Request{}
	if err := ParseJsonStr(string(reqContent), request); err != nil {
		log.Errorln("解析模板JSON失败")
		ResMsg(res, 400, err.Error())
		return
	}

	err = controller.Device.DeleteRack(request.UUID)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("失败")
		ResMsg(res, 400, err.Error())
		return
	}
	ResSuccessMsg(res, 200, "操作成功")

	// audit
	audit(token, "删除机柜", "", fmt.Sprintf("机柜ID: %s", request.UUID))
}

func ajaxListRacks(res http.ResponseWriter, req *http.Request) {
	token, err := util.CM.Get("token", req)
	if err != nil || token == "" {
		ResMsg(res, 400, "请求中未包含token")
		return
	}

	records, err := controller.Device.ListRacks()
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("失败")
		ResMsg(res, 400, err.Error())
		return
	}
	b, err := json.Marshal(records)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("JSON生成失败")
		ResMsg(res, 500, err.Error())
		return
	}
	ResMsg(res, 200, string(b))

	// audit
	audit(token, "查看机柜列表", "", "")
}

func ajaxCreateServerDevice(res http.ResponseWriter, req *http.Request) {
	token, err := util.CM.Get("token", req)
	if err != nil || token == "" {
		ResMsg(res, 400, "请求中未包含token")
		return
	}

	if isAdmin(token) == false {
		ResMsg(res, 400, "权限不足")
		return
	}

	reqContent, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		log.WithFields(log.Fields{}).Error("请求报文解析失败")
		ResInvalidRequestBody(res)
		return
	}

	type Request struct {
		Brand          string `json:"brand"`
		Model          string `json:"model"`
		DiskCapacity   int64  `json:"disk_capacity"`
		MemoryCapacity int64  `json:"memory_capacity"`
		Hostname       string `json:"hostname"`
		EnableTime     string `json:"enable_time"`
		ExpireTime     string `json:"expire_time"`
		OS             string `json:"os"`
		Comment        string `json:"comment"`
	}

	request := &Request{}
	if err := ParseJsonStr(string(reqContent), request); err != nil {
		log.Errorln("解析模板JSON失败")
		ResMsg(res, 400, err.Error())
		return
	}

	err = controller.Device.CreateServerDevice(request.Brand, request.Model, request.DiskCapacity, request.MemoryCapacity, request.Hostname, request.EnableTime, request.ExpireTime, request.OS, request.Comment)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("失败")
		ResMsg(res, 400, err.Error())
		return
	}
	ResSuccessMsg(res, 200, "操作成功")

	// audit
	audit(token, "新建物理服务器", "", fmt.Sprintf("服务器名: %s", request.Hostname))
}

func ajaxDeleteServerDevice(res http.ResponseWriter, req *http.Request) {
	token, err := util.CM.Get("token", req)
	if err != nil || token == "" {
		ResMsg(res, 400, "请求中未包含token")
		return
	}

	if isAdmin(token) == false {
		ResMsg(res, 400, "权限不足")
		return
	}

	reqContent, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		log.WithFields(log.Fields{}).Error("请求报文解析失败")
		ResInvalidRequestBody(res)
		return
	}

	type Request struct {
		UUID string `json:"uuid"`
	}

	request := &Request{}
	if err := ParseJsonStr(string(reqContent), request); err != nil {
		log.Errorln("解析模板JSON失败")
		ResMsg(res, 400, err.Error())
		return
	}

	err = controller.Device.DeleteServerDevice(request.UUID)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("失败")
		ResMsg(res, 400, err.Error())
		return
	}
	ResSuccessMsg(res, 200, "操作成功")

	// audit
	audit(token, "删除物理服务器", "", fmt.Sprintf("服务器ID: %s", request.UUID))
}

func ajaxListServerDevices(res http.ResponseWriter, req *http.Request) {
	token, err := util.CM.Get("token", req)
	if err != nil || token == "" {
		ResMsg(res, 400, "请求中未包含token")
		return
	}

	records, err := controller.Device.ListServerDevices()
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("失败")
		ResMsg(res, 400, err.Error())
		return
	}
	b, err := json.Marshal(records)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("JSON生成失败")
		ResMsg(res, 500, err.Error())
		return
	}
	ResMsg(res, 200, string(b))

	// audit
	audit(token, "查看物理服务器列表", "", "")
}

func ajaxCreateNetworkDevice(res http.ResponseWriter, req *http.Request) {
	token, err := util.CM.Get("token", req)
	if err != nil || token == "" {
		ResMsg(res, 400, "请求中未包含token")
		return
	}

	if isAdmin(token) == false {
		ResMsg(res, 400, "权限不足")
		return
	}

	reqContent, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		log.WithFields(log.Fields{}).Error("请求报文解析失败")
		ResInvalidRequestBody(res)
		return
	}

	type Request struct {
		Brand      string `json:"brand"`
		Model      string `json:"model"`
		Name       string `json:"name"`
		EnableTime string `json:"enable_time"`
		ExpireTime string `json:"expire_time"`
		Comment    string `json:"comment"`
	}

	request := &Request{}
	if err := ParseJsonStr(string(reqContent), request); err != nil {
		log.Errorln("解析模板JSON失败")
		ResMsg(res, 400, err.Error())
		return
	}

	err = controller.Device.CreateNetworkDevice(request.Brand, request.Model, request.Name, request.EnableTime, request.ExpireTime, request.Comment)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("失败")
		ResMsg(res, 400, err.Error())
		return
	}
	ResSuccessMsg(res, 200, "操作成功")

	// audit
	audit(token, "创建网络设备", "", fmt.Sprintf("网络设备名: %s", request.Name))
}

func ajaxDeleteNetworkDevice(res http.ResponseWriter, req *http.Request) {
	token, err := util.CM.Get("token", req)
	if err != nil || token == "" {
		ResMsg(res, 400, "请求中未包含token")
		return
	}

	if isAdmin(token) == false {
		ResMsg(res, 400, "权限不足")
		return
	}

	reqContent, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		log.WithFields(log.Fields{}).Error("请求报文解析失败")
		ResInvalidRequestBody(res)
		return
	}

	type Request struct {
		UUID string `json:"uuid"`
	}

	request := &Request{}
	if err := ParseJsonStr(string(reqContent), request); err != nil {
		log.Errorln("解析模板JSON失败")
		ResMsg(res, 400, err.Error())
		return
	}

	err = controller.Device.DeleteNetworkDevice(request.UUID)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("失败")
		ResMsg(res, 400, err.Error())
		return
	}
	ResSuccessMsg(res, 200, "操作成功")

	// audit
	audit(token, "删除网络设备", "", fmt.Sprintf("网络设备ID: %s", request.UUID))
}

func ajaxListNetworkDevices(res http.ResponseWriter, req *http.Request) {
	token, err := util.CM.Get("token", req)
	if err != nil || token == "" {
		ResMsg(res, 400, "请求中未包含token")
		return
	}

	records, err := controller.Device.ListNetworkDevices()
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("失败")
		ResMsg(res, 400, err.Error())
		return
	}
	b, err := json.Marshal(records)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("JSON生成失败")
		ResMsg(res, 500, err.Error())
		return
	}
	ResMsg(res, 200, string(b))

	// audit
	audit(token, "查看网络设备列表", "", "")
}

func ajaxCreateStorageDevice(res http.ResponseWriter, req *http.Request) {
	token, err := util.CM.Get("token", req)
	if err != nil || token == "" {
		ResMsg(res, 400, "请求中未包含token")
		return
	}

	if isAdmin(token) == false {
		ResMsg(res, 400, "权限不足")
		return
	}

	reqContent, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		log.WithFields(log.Fields{}).Error("请求报文解析失败")
		ResInvalidRequestBody(res)
		return
	}

	type Request struct {
		Brand      string `json:"brand"`
		Model      string `json:"model"`
		Name       string `json:"name"`
		EnableTime string `json:"enable_time"`
		ExpireTime string `json:"expire_time"`
		Comment    string `json:"comment"`
	}

	request := &Request{}
	if err := ParseJsonStr(string(reqContent), request); err != nil {
		log.Errorln("解析模板JSON失败")
		ResMsg(res, 400, err.Error())
		return
	}

	err = controller.Device.CreateStorageDevice(request.Brand, request.Model, request.Name, request.EnableTime, request.ExpireTime, request.Comment)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("失败")
		ResMsg(res, 400, err.Error())
		return
	}
	ResSuccessMsg(res, 200, "操作成功")

	// audit
	audit(token, "创建存储设备", "", fmt.Sprintf("存储设备名: %s", request.Name))
}

func ajaxDeleteStorageDevice(res http.ResponseWriter, req *http.Request) {
	token, err := util.CM.Get("token", req)
	if err != nil || token == "" {
		ResMsg(res, 400, "请求中未包含token")
		return
	}

	if isAdmin(token) == false {
		ResMsg(res, 400, "权限不足")
		return
	}

	reqContent, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		log.WithFields(log.Fields{}).Error("请求报文解析失败")
		ResInvalidRequestBody(res)
		return
	}

	type Request struct {
		UUID string `json:"uuid"`
	}

	request := &Request{}
	if err := ParseJsonStr(string(reqContent), request); err != nil {
		log.Errorln("解析模板JSON失败")
		ResMsg(res, 400, err.Error())
		return
	}

	err = controller.Device.DeleteStorageDevice(request.UUID)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("失败")
		ResMsg(res, 400, err.Error())
		return
	}
	ResSuccessMsg(res, 200, "操作成功")

	// audit
	audit(token, "删除存储设备", "", fmt.Sprintf("存储设备ID: %s", request.UUID))
}

func ajaxListStorageDevices(res http.ResponseWriter, req *http.Request) {
	token, err := util.CM.Get("token", req)
	if err != nil || token == "" {
		ResMsg(res, 400, "请求中未包含token")
		return
	}

	records, err := controller.Device.ListStorageDevices()
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("失败")
		ResMsg(res, 400, err.Error())
		return
	}
	b, err := json.Marshal(records)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("JSON生成失败")
		ResMsg(res, 500, err.Error())
		return
	}
	ResMsg(res, 200, string(b))

	// audit
	audit(token, "查看存储设备列表", "", "")
}

func ajaxCreateCommonDevice(res http.ResponseWriter, req *http.Request) {
	token, err := util.CM.Get("token", req)
	if err != nil || token == "" {
		ResMsg(res, 400, "请求中未包含token")
		return
	}

	if isAdmin(token) == false {
		ResMsg(res, 400, "权限不足")
		return
	}

	reqContent, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		log.WithFields(log.Fields{}).Error("请求报文解析失败")
		ResInvalidRequestBody(res)
		return
	}

	type Request struct {
		Brand      string `json:"brand"`
		Model      string `json:"model"`
		Name       string `json:"name"`
		EnableTime string `json:"enable_time"`
		ExpireTime string `json:"expire_time"`
		Comment    string `json:"comment"`
	}

	request := &Request{}
	if err := ParseJsonStr(string(reqContent), request); err != nil {
		log.Errorln("解析模板JSON失败")
		ResMsg(res, 400, err.Error())
		return
	}

	err = controller.Device.CreateCommonDevice(request.Brand, request.Model, request.Name, request.EnableTime, request.ExpireTime, request.Comment)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("失败")
		ResMsg(res, 400, err.Error())
		return
	}
	ResSuccessMsg(res, 200, "操作成功")

	// audit
	audit(token, "创建其它设备", "", fmt.Sprintf("设备名: %s", request.Name))
}

func ajaxDeleteCommonDevice(res http.ResponseWriter, req *http.Request) {
	token, err := util.CM.Get("token", req)
	if err != nil || token == "" {
		ResMsg(res, 400, "请求中未包含token")
		return
	}

	if isAdmin(token) == false {
		ResMsg(res, 400, "权限不足")
		return
	}

	reqContent, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		log.WithFields(log.Fields{}).Error("请求报文解析失败")
		ResInvalidRequestBody(res)
		return
	}

	type Request struct {
		UUID string `json:"uuid"`
	}

	request := &Request{}
	if err := ParseJsonStr(string(reqContent), request); err != nil {
		log.Errorln("解析模板JSON失败")
		ResMsg(res, 400, err.Error())
		return
	}

	err = controller.Device.DeleteCommonDevice(request.UUID)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("失败")
		ResMsg(res, 400, err.Error())
		return
	}
	ResSuccessMsg(res, 200, "操作成功")

	// audit
	audit(token, "删除其它设备", "", fmt.Sprintf("设备ID: %s", request.UUID))
}

func ajaxListCommonDevices(res http.ResponseWriter, req *http.Request) {
	token, err := util.CM.Get("token", req)
	if err != nil || token == "" {
		ResMsg(res, 400, "请求中未包含token")
		return
	}

	records, err := controller.Device.ListCommonDevices()
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("失败")
		ResMsg(res, 400, err.Error())
		return
	}
	b, err := json.Marshal(records)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("JSON生成失败")
		ResMsg(res, 500, err.Error())
		return
	}
	ResMsg(res, 200, string(b))

	// audit
	audit(token, "查看其它设备列表", "", "")
}

func ajaxListIPs(res http.ResponseWriter, req *http.Request) {
	token, err := util.CM.Get("token", req)
	if err != nil || token == "" {
		ResMsg(res, 400, "请求中未包含token")
		return
	}

	records, err := controller.Ip.ListIPRecords()
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("失败")
		ResMsg(res, 400, err.Error())
		return
	}
	b, err := json.Marshal(records)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("JSON生成失败")
		ResMsg(res, 500, err.Error())
		return
	}
	ResMsg(res, 200, string(b))

	// audit
	audit(token, "创建IP列表", "", "")
}

func ajaxCreateIPRecord(res http.ResponseWriter, req *http.Request) {
	token, err := util.CM.Get("token", req)
	if err != nil || token == "" {
		ResMsg(res, 400, "请求中未包含token")
		return
	}

	if isAdmin(token) == false {
		ResMsg(res, 400, "权限不足")
		return
	}

	reqContent, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		log.WithFields(log.Fields{}).Error("请求报文解析失败")
		ResInvalidRequestBody(res)
		return
	}

	type Request struct {
		IPAddress string `json:"ip_address"`
		Type      string `json:"ip_type"`
		Role      string `json:"ip_role"`
		TargetId  string `json:"target_id"`
		IpSetId   string `json:"ip_set_id"`
	}

	request := &Request{}
	if err := ParseJsonStr(string(reqContent), request); err != nil {
		log.Errorln("解析模板JSON失败")
		ResMsg(res, 400, err.Error())
		return
	}

	err = controller.Ip.CreateIPRecord(request.IPAddress, request.Type, request.Role, request.TargetId, request.IpSetId)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("失败")
		ResMsg(res, 400, err.Error())
		return
	}
	ResSuccessMsg(res, 200, "操作成功")

	// audit
	audit(token, "新增IP记录", "", fmt.Sprintf("IP: %s", request.IPAddress))
}

func ajaxDeleteIPRecord(res http.ResponseWriter, req *http.Request) {
	token, err := util.CM.Get("token", req)
	if err != nil || token == "" {
		ResMsg(res, 400, "请求中未包含token")
		return
	}

	if isAdmin(token) == false {
		ResMsg(res, 400, "权限不足")
		return
	}

	reqContent, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		log.WithFields(log.Fields{}).Error("请求报文解析失败")
		ResInvalidRequestBody(res)
		return
	}

	type Request struct {
		UUID string `json:"uuid"`
	}

	request := &Request{}
	if err := ParseJsonStr(string(reqContent), request); err != nil {
		log.Errorln("解析模板JSON失败")
		ResMsg(res, 400, err.Error())
		return
	}

	err = controller.Ip.DeleteIPRecord(request.UUID)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("失败")
		ResMsg(res, 400, err.Error())
		return
	}
	ResSuccessMsg(res, 200, "操作成功")

	// audit
	audit(token, "删除IP记录", "", fmt.Sprintf("IP ID: %s", request.UUID))
}

func ajaxListIPSets(res http.ResponseWriter, req *http.Request) {
	token, err := util.CM.Get("token", req)
	if err != nil || token == "" {
		ResMsg(res, 400, "请求中未包含token")
		return
	}

	records, err := controller.Ip.ListIPSets()
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("失败")
		ResMsg(res, 400, err.Error())
		return
	}
	b, err := json.Marshal(records)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("JSON生成失败")
		ResMsg(res, 500, err.Error())
		return
	}
	ResMsg(res, 200, string(b))

	// audit
	audit(token, "查看网段列表", "", "")
}

func ajaxCreateIPSetRecord(res http.ResponseWriter, req *http.Request) {
	token, err := util.CM.Get("token", req)
	if err != nil || token == "" {
		ResMsg(res, 400, "请求中未包含token")
		return
	}

	if isAdmin(token) == false {
		ResMsg(res, 400, "权限不足")
		return
	}

	reqContent, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		log.WithFields(log.Fields{}).Error("请求报文解析失败")
		ResInvalidRequestBody(res)
		return
	}

	type Request struct {
		CIDR    string `json:"cidr"`
		Comment string `json:"comment"`
	}

	request := &Request{}
	if err := ParseJsonStr(string(reqContent), request); err != nil {
		log.Errorln("解析模板JSON失败")
		ResMsg(res, 400, err.Error())
		return
	}

	err = controller.Ip.CreateIPSet(request.CIDR, request.Comment)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("失败")
		ResMsg(res, 400, err.Error())
		return
	}
	ResSuccessMsg(res, 200, "操作成功")

	// audit
	audit(token, "创建IP网段", "", fmt.Sprintf("CIRD: %s", request.CIDR))
}

func ajaxDeleteIPSetRecord(res http.ResponseWriter, req *http.Request) {
	token, err := util.CM.Get("token", req)
	if err != nil || token == "" {
		ResMsg(res, 400, "请求中未包含token")
		return
	}

	if isAdmin(token) == false {
		ResMsg(res, 400, "权限不足")
		return
	}

	reqContent, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		log.WithFields(log.Fields{}).Error("请求报文解析失败")
		ResInvalidRequestBody(res)
		return
	}

	type Request struct {
		UUID string `json:"uuid"`
	}

	request := &Request{}
	if err := ParseJsonStr(string(reqContent), request); err != nil {
		log.Errorln("解析模板JSON失败")
		ResMsg(res, 400, err.Error())
		return
	}

	err = controller.Ip.DeleteIPSet(request.UUID)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("失败")
		ResMsg(res, 400, err.Error())
		return
	}
	ResSuccessMsg(res, 200, "操作成功")

	// audit
	audit(token, "删除IP网段", "", fmt.Sprintf("网段ID: %s", request.UUID))
}

func ajaxListConnections(res http.ResponseWriter, req *http.Request) {
	token, err := util.CM.Get("token", req)
	if err != nil || token == "" {
		ResMsg(res, 400, "请求中未包含token")
		return
	}

	records, err := controller.Connection.ListConnections()
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("失败")
		ResMsg(res, 400, err.Error())
		return
	}
	b, err := json.Marshal(records)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("JSON生成失败")
		ResMsg(res, 500, err.Error())
		return
	}
	ResMsg(res, 200, string(b))

	// audit
	audit(token, "查看布线信息列表", "", "")
}

func ajaxCreateConnection(res http.ResponseWriter, req *http.Request) {
	token, err := util.CM.Get("token", req)
	if err != nil || token == "" {
		ResMsg(res, 400, "请求中未包含token")
		return
	}

	if isAdmin(token) == false {
		ResMsg(res, 400, "权限不足")
		return
	}

	reqContent, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		log.WithFields(log.Fields{}).Error("请求报文解析失败")
		ResInvalidRequestBody(res)
		return
	}

	type Request struct {
		SourceId              string `json:"source_id"`
		SourcePort            string `json:"source_port"`
		SourceDeviceType      string `json:"source_device_type"`
		SourceDeviceName      string `json:"source_device_name"`
		DestinationId         string `json:"destination_id"`
		DestinationPort       string `json:"destination_port"`
		DestinationDeviceType string `json:"destination_device_type"`
		DestinationDeviceName string `json:"destination_device_name"`
	}

	request := &Request{}
	if err := ParseJsonStr(string(reqContent), request); err != nil {
		log.Errorln("解析模板JSON失败")
		ResMsg(res, 400, err.Error())
		return
	}

	err = controller.Connection.CreateConnection(request.SourceId, request.SourcePort, request.SourceDeviceType, request.SourceDeviceName, request.DestinationId, request.DestinationPort, request.DestinationDeviceType, request.DestinationDeviceName)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("失败")
		ResMsg(res, 400, err.Error())
		return
	}
	ResSuccessMsg(res, 200, "操作成功")

	// audit
	audit(token, "新增布线", "", fmt.Sprintf("A端: %s B端: %s", request.SourceDeviceName, request.DestinationDeviceName))
}

func ajaxDeleteConnection(res http.ResponseWriter, req *http.Request) {
	token, err := util.CM.Get("token", req)
	if err != nil || token == "" {
		ResMsg(res, 400, "请求中未包含token")
		return
	}

	if isAdmin(token) == false {
		ResMsg(res, 400, "权限不足")
		return
	}

	reqContent, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		log.WithFields(log.Fields{}).Error("请求报文解析失败")
		ResInvalidRequestBody(res)
		return
	}

	type Request struct {
		UUID string `json:"uuid"`
	}

	request := &Request{}
	if err := ParseJsonStr(string(reqContent), request); err != nil {
		log.Errorln("解析模板JSON失败")
		ResMsg(res, 400, err.Error())
		return
	}

	err = controller.Connection.DeleteConnection(request.UUID)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("失败")
		ResMsg(res, 400, err.Error())
		return
	}
	ResSuccessMsg(res, 200, "操作成功")

	// audit
	audit(token, "删除布线", "", fmt.Sprintf("布线ID: %s", request.UUID))
}

func ajaxListAuditRecords(res http.ResponseWriter, req *http.Request) {
	token, err := util.CM.Get("token", req)
	if err != nil || token == "" {
		ResMsg(res, 400, "请求中未包含token")
		return
	}

	reqContent, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		log.WithFields(log.Fields{}).Error("请求报文解析失败")
		ResInvalidRequestBody(res)
		return
	}

	request := &structs.ListAuditRecordsCondition{}
	if err := ParseJsonStr(string(reqContent), request); err != nil {
		log.Errorln("解析模板JSON失败")
		ResMsg(res, 400, err.Error())
		return
	}

	records, err := controller.Audit.ListRecords(request)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("失败")
		ResMsg(res, 400, err.Error())
		return
	}
	b, err := json.Marshal(records)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("JSON生成失败")
		ResMsg(res, 500, err.Error())
		return
	}
	ResMsg(res, 200, string(b))

	// audit
	audit(token, "查看审计记录", "", "")
}

func ajaxListMonitorItems(res http.ResponseWriter, req *http.Request) {
	token, err := util.CM.Get("token", req)
	if err != nil || token == "" {
		ResMsg(res, 400, "请求中未包含token")
		return
	}

	records, err := controller.Monitor.ListMonitorItems()
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("失败")
		ResMsg(res, 400, err.Error())
		return
	}
	b, err := json.Marshal(records)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("JSON生成失败")
		ResMsg(res, 500, err.Error())
		return
	}
	ResMsg(res, 200, string(b))

	// audit
	audit(token, "查看监控项", "", "")
}

func ajaxCreateMonitorItem(res http.ResponseWriter, req *http.Request) {
	token, err := util.CM.Get("token", req)
	if err != nil || token == "" {
		ResMsg(res, 400, "请求中未包含token")
		return
	}

	if isAdmin(token) == false {
		ResMsg(res, 400, "权限不足")
		return
	}

	reqContent, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		log.WithFields(log.Fields{}).Error("请求报文解析失败")
		ResInvalidRequestBody(res)
		return
	}

	type Request struct {
		Name      string `json:"name"`
		DCType    string `json:"dc_type"`
		AlertType string `json:"alert_type"`
	}

	request := &Request{}
	if err := ParseJsonStr(string(reqContent), request); err != nil {
		log.Errorln("解析模板JSON失败")
		ResMsg(res, 400, err.Error())
		return
	}

	err = controller.Monitor.CreateMonitorItem(request.Name, request.DCType, request.AlertType)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("失败")
		ResMsg(res, 400, err.Error())
		return
	}
	ResSuccessMsg(res, 200, "操作成功")

	// audit
	audit(token, "新增监控项", "", fmt.Sprintf("监控项名称: %s", request.Name))
}

func ajaxUpdateMonitorItem(res http.ResponseWriter, req *http.Request) {
	token, err := util.CM.Get("token", req)
	if err != nil || token == "" {
		ResMsg(res, 400, "请求中未包含token")
		return
	}

	if isAdmin(token) == false {
		ResMsg(res, 400, "权限不足")
		return
	}

	reqContent, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		log.WithFields(log.Fields{}).Error("请求报文解析失败")
		ResInvalidRequestBody(res)
		return
	}

	type Request struct {
		Id        int64  `json:"id"`
		Name      string `json:"name"`
		DCType    string `json:"dc_type"`
		AlertType string `json:"alert_type"`
	}

	request := &Request{}
	if err := ParseJsonStr(string(reqContent), request); err != nil {
		log.Errorln("解析模板JSON失败")
		ResMsg(res, 400, err.Error())
		return
	}

	err = controller.Monitor.UpdateMonitorItem(request.Id, request.Name, request.DCType, request.AlertType)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("失败")
		ResMsg(res, 400, err.Error())
		return
	}
	ResSuccessMsg(res, 200, "操作成功")

	// audit
	audit(token, "更新监控项", "", fmt.Sprintf("被更新监控项名称: %s", request.Name))
}

func ajaxDeleteMonitorItem(res http.ResponseWriter, req *http.Request) {
	token, err := util.CM.Get("token", req)
	if err != nil || token == "" {
		ResMsg(res, 400, "请求中未包含token")
		return
	}

	if isAdmin(token) == false {
		ResMsg(res, 400, "权限不足")
		return
	}

	reqContent, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		log.WithFields(log.Fields{}).Error("请求报文解析失败")
		ResInvalidRequestBody(res)
		return
	}

	type Request struct {
		Id int64 `json:"id"`
	}

	request := &Request{}
	if err := ParseJsonStr(string(reqContent), request); err != nil {
		log.Errorln("解析模板JSON失败")
		ResMsg(res, 400, err.Error())
		return
	}

	err = controller.Monitor.DeleteMonitorItem(request.Id)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("失败")
		ResMsg(res, 400, err.Error())
		return
	}
	ResSuccessMsg(res, 200, "操作成功")

	// audit
	audit(token, "删除监控项", "", fmt.Sprintf("被删除监控项名称: %s", request.Id))
}

func ajaxUpdateMonitorItemDCCfg(res http.ResponseWriter, req *http.Request) {
	token, err := util.CM.Get("token", req)
	if err != nil || token == "" {
		ResMsg(res, 400, "请求中未包含token")
		return
	}

	if isAdmin(token) == false {
		ResMsg(res, 400, "权限不足")
		return
	}

	reqContent, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		log.WithFields(log.Fields{}).Error("请求报文解析失败")
		ResInvalidRequestBody(res)
		return
	}

	type Request struct {
		Id                int64  `json:"id"`
		DCType            string `json:"dc_type"`
		DCFakeCfgItemName string `json:"dc_fake_cfg_item_name"`
		DCFakeCfgHostIp   string `json:"dc_fake_cfg_host_ip"`
	}

	request := &Request{}
	if err := ParseJsonStr(string(reqContent), request); err != nil {
		log.Errorln("解析模板JSON失败")
		ResMsg(res, 400, err.Error())
		return
	}

	err = controller.Monitor.UpdateMonitorItemDCCfg(request.Id, request.DCType, request.DCFakeCfgItemName, request.DCFakeCfgHostIp)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("失败")
		ResMsg(res, 400, err.Error())
		return
	}
	ResSuccessMsg(res, 200, "操作成功")

	// audit
	audit(token, "更新监控项数据收集模块信息", "", fmt.Sprintf("被更新监控项ID: %d", request.Id))
}

func ajaxGetMonitorItemById(res http.ResponseWriter, req *http.Request) {
	token, err := util.CM.Get("token", req)
	if err != nil || token == "" {
		ResMsg(res, 400, "请求中未包含token")
		return
	}

	if isAdmin(token) == false {
		ResMsg(res, 400, "权限不足")
		return
	}

	reqContent, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		log.WithFields(log.Fields{}).Error("请求报文解析失败")
		ResInvalidRequestBody(res)
		return
	}

	type Request struct {
		Id int64 `json:"id"`
	}

	request := &Request{}
	if err := ParseJsonStr(string(reqContent), request); err != nil {
		log.Errorln("解析模板JSON失败")
		ResMsg(res, 400, err.Error())
		return
	}

	records, err := controller.Monitor.GetMonitorItemById(request.Id)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("失败")
		ResMsg(res, 400, err.Error())
		return
	}
	b, err := json.Marshal(records)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("JSON生成失败")
		ResMsg(res, 500, err.Error())
		return
	}
	ResMsg(res, 200, string(b))

	// audit
	audit(token, "查看监控项数据收集模块详情", "", fmt.Sprintf("被查看监控项ID: %d", request.Id))
}

func ajaxListDeviceReleatedMonitorItems(res http.ResponseWriter, req *http.Request) {
	token, err := util.CM.Get("token", req)
	if err != nil || token == "" {
		ResMsg(res, 400, "请求中未包含token")
		return
	}

	if isAdmin(token) == false {
		ResMsg(res, 400, "权限不足")
		return
	}

	reqContent, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		log.WithFields(log.Fields{}).Error("请求报文解析失败")
		ResInvalidRequestBody(res)
		return
	}

	type Request struct {
		UUID string `json:"uuid"`
	}

	request := &Request{}
	if err := ParseJsonStr(string(reqContent), request); err != nil {
		log.Errorln("解析模板JSON失败")
		ResMsg(res, 400, err.Error())
		return
	}

	records, err := controller.Monitor.ListDeviceReleatedMonitorItems(request.UUID)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("失败")
		ResMsg(res, 400, err.Error())
		return
	}
	b, err := json.Marshal(records)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("JSON生成失败")
		ResMsg(res, 500, err.Error())
		return
	}
	ResMsg(res, 200, string(b))

	// audit
	audit(token, "查看设备关联的监控项", "", fmt.Sprintf("被查的设备ID: %s", request.UUID))
}

func ajaxListMonitorItemReleatedDevices(res http.ResponseWriter, req *http.Request) {
	token, err := util.CM.Get("token", req)
	if err != nil || token == "" {
		ResMsg(res, 400, "请求中未包含token")
		return
	}

	if isAdmin(token) == false {
		ResMsg(res, 400, "权限不足")
		return
	}

	reqContent, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		log.WithFields(log.Fields{}).Error("请求报文解析失败")
		ResInvalidRequestBody(res)
		return
	}

	type Request struct {
		ItemId int64 `json:"id"`
	}

	request := &Request{}
	if err := ParseJsonStr(string(reqContent), request); err != nil {
		log.Errorln("解析模板JSON失败")
		ResMsg(res, 400, err.Error())
		return
	}

	records, err := controller.Monitor.ListMonitorItemReleatedDevices(request.ItemId)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("失败")
		ResMsg(res, 400, err.Error())
		return
	}
	b, err := json.Marshal(records)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("JSON生成失败")
		ResMsg(res, 500, err.Error())
		return
	}
	ResMsg(res, 200, string(b))

	// audit
	audit(token, "查看监控项关联的设备", "", fmt.Sprintf("被查的监控项ID: %s", request.ItemId))
}

func ajaxListDevices(res http.ResponseWriter, req *http.Request) {
	token, err := util.CM.Get("token", req)
	if err != nil || token == "" {
		ResMsg(res, 400, "请求中未包含token")
		return
	}

	records, err := controller.Device.ListDevices()
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("失败")
		ResMsg(res, 400, err.Error())
		return
	}
	b, err := json.Marshal(records)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("JSON生成失败")
		ResMsg(res, 500, err.Error())
		return
	}
	ResMsg(res, 200, string(b))

	// audit
	audit(token, "查看所有设备列表", "", "")
}

func ajaxBindMonitorItemAndDevice(res http.ResponseWriter, req *http.Request) {
	token, err := util.CM.Get("token", req)
	if err != nil || token == "" {
		ResMsg(res, 400, "请求中未包含token")
		return
	}

	if isAdmin(token) == false {
		ResMsg(res, 400, "权限不足")
		return
	}

	reqContent, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		log.WithFields(log.Fields{}).Error("请求报文解析失败")
		ResInvalidRequestBody(res)
		return
	}

	type Request struct {
		ItemId     int64  `json:"item_id"`
		ItemName   string `json:"item_name"`
		DeviceUUID string `json:"device_uuid"`
		DeviceType string `json:"device_type"`
		DeviceName string `json:"device_name"`
	}

	request := &Request{}
	if err := ParseJsonStr(string(reqContent), request); err != nil {
		log.Errorln("解析模板JSON失败")
		ResMsg(res, 400, err.Error())
		return
	}

	err = controller.Monitor.BindMonitorItemAndDevice(request.ItemId, request.ItemName, request.DeviceUUID, request.DeviceType, request.DeviceName)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("失败")
		ResMsg(res, 400, err.Error())
		return
	}
	ResSuccessMsg(res, 200, "操作成功")

	// audit
	audit(token, "绑定监控项和设备", "", fmt.Sprintf("绑定监控项 %s 和设备 %s", request.ItemName, request.DeviceName))
}

func ajaxGetDeviceByUUID(res http.ResponseWriter, req *http.Request) {
	token, err := util.CM.Get("token", req)
	if err != nil || token == "" {
		ResMsg(res, 400, "请求中未包含token")
		return
	}

	if isAdmin(token) == false {
		ResMsg(res, 400, "权限不足")
		return
	}

	reqContent, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		log.WithFields(log.Fields{}).Error("请求报文解析失败")
		ResInvalidRequestBody(res)
		return
	}

	type Request struct {
		UUID string `json:"uuid"`
	}

	request := &Request{}
	if err := ParseJsonStr(string(reqContent), request); err != nil {
		log.Errorln("解析模板JSON失败")
		ResMsg(res, 400, err.Error())
		return
	}

	records, err := controller.Device.GetDeviceByUUID(request.UUID)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("失败")
		ResMsg(res, 400, err.Error())
		return
	}
	b, err := json.Marshal(records)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("JSON生成失败")
		ResMsg(res, 500, err.Error())
		return
	}
	ResMsg(res, 200, string(b))

	// audit
	audit(token, "查看设备信息信息", "", fmt.Sprintf("设备ID %s", request.UUID))
}

func ajaxGetDeviceMonitorItemHistoryData(res http.ResponseWriter, req *http.Request) {
	token, err := util.CM.Get("token", req)
	if err != nil || token == "" {
		ResMsg(res, 400, "请求中未包含token")
		return
	}

	if isAdmin(token) == false {
		ResMsg(res, 400, "权限不足")
		return
	}

	reqContent, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		log.WithFields(log.Fields{}).Error("请求报文解析失败")
		ResInvalidRequestBody(res)
		return
	}

	request := &structs.HistoryDataFilter{}
	if err := ParseJsonStr(string(reqContent), request); err != nil {
		log.Errorln("解析模板JSON失败")
		ResMsg(res, 400, err.Error())
		return
	}

	records, err := controller.Monitor.GetDeviceMonitorItemHistoryData(request)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("失败")
		ResMsg(res, 400, err.Error())
		return
	}
	b, err := json.Marshal(records)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("JSON生成失败")
		ResMsg(res, 500, err.Error())
		return
	}
	ResMsg(res, 200, string(b))

	// audit
	audit(token, "查看监控信息", "", fmt.Sprintf("被查的监控项ID: %s, 设备项 %s", request.ItemId, request.DeviceUUID))
}

func ajaxListMonitorBackendCfgs(res http.ResponseWriter, req *http.Request) {
	token, err := util.CM.Get("token", req)
	if err != nil || token == "" {
		ResMsg(res, 400, "请求中未包含token")
		return
	}

	records, err := controller.Monitor.ListMonitorBackendCfgs()
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("失败")
		ResMsg(res, 400, err.Error())
		return
	}
	b, err := json.Marshal(records)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("JSON生成失败")
		ResMsg(res, 500, err.Error())
		return
	}
	ResMsg(res, 200, string(b))

	// audit
	audit(token, "查看监控服务信息", "", "")
}

func ajaxUpdateMonitorBackendCfg(res http.ResponseWriter, req *http.Request) {
	token, err := util.CM.Get("token", req)
	if err != nil || token == "" {
		ResMsg(res, 400, "请求中未包含token")
		return
	}

	if isAdmin(token) == false {
		ResMsg(res, 400, "权限不足")
		return
	}

	reqContent, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		log.WithFields(log.Fields{}).Error("请求报文解析失败")
		ResInvalidRequestBody(res)
		return
	}

	type Request struct {
		BackendName string `json:"backend_name"`
		Cfg         string `json:"cfg"`
	}

	request := &Request{}
	if err := ParseJsonStr(string(reqContent), request); err != nil {
		log.Errorln("解析模板JSON失败")
		ResMsg(res, 400, err.Error())
		return
	}

	err = controller.Monitor.UpdateMonitorBackendCfg(request.BackendName, request.Cfg)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("失败")
		ResMsg(res, 400, err.Error())
		return
	}
	ResSuccessMsg(res, 200, "操作成功")

	// audit
	audit(token, "更新监控服务信息", "", fmt.Sprintf("监控服务 %s", request.BackendName))
}
