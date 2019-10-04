package server

import (
	"be/controller"
	"be/util"
	"encoding/json"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

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
}

func ajaxCreateUser(res http.ResponseWriter, req *http.Request) {
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
		Username string `json:"username"`
		Password string `json:"password"`
	}

	request := &Request{}
	if err := ParseJsonStr(string(reqContent), request); err != nil {
		log.Errorln("解析模板JSON失败")
		ResMsg(res, 400, err.Error())
		return
	}

	err = controller.Auth.CreateUser(request.Username, request.Password)
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
}

func ajaxCreateDataCenter(res http.ResponseWriter, req *http.Request) {
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
}

func ajaxDeleteDataCenter(res http.ResponseWriter, req *http.Request) {
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

	err = controller.Device.DeleteDataCenter(request.UUID)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("失败")
		ResMsg(res, 400, err.Error())
		return
	}
	ResSuccessMsg(res, 200, "操作成功")
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
}

func ajaxCreateRack(res http.ResponseWriter, req *http.Request) {
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
		Name string `json:"name"`
	}

	request := &Request{}
	if err := ParseJsonStr(string(reqContent), request); err != nil {
		log.Errorln("解析模板JSON失败")
		ResMsg(res, 400, err.Error())
		return
	}

	err = controller.Device.CreateRack(request.Name)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("失败")
		ResMsg(res, 400, err.Error())
		return
	}
	ResSuccessMsg(res, 200, "操作成功")
}

func ajaxDeleteRack(res http.ResponseWriter, req *http.Request) {
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

	err = controller.Device.DeleteRack(request.UUID)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("失败")
		ResMsg(res, 400, err.Error())
		return
	}
	ResSuccessMsg(res, 200, "操作成功")
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
}

func ajaxCreateServerDevice(res http.ResponseWriter, req *http.Request) {
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
}

func ajaxDeleteServerDevice(res http.ResponseWriter, req *http.Request) {
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

	err = controller.Device.DeleteServerDevice(request.UUID)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("失败")
		ResMsg(res, 400, err.Error())
		return
	}
	ResSuccessMsg(res, 200, "操作成功")
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
}

func ajaxCreateNetworkDevice(res http.ResponseWriter, req *http.Request) {
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
}

func ajaxDeleteNetworkDevice(res http.ResponseWriter, req *http.Request) {
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

	err = controller.Device.DeleteNetworkDevice(request.UUID)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("失败")
		ResMsg(res, 400, err.Error())
		return
	}
	ResSuccessMsg(res, 200, "操作成功")
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
}

func ajaxCreateStorageDevice(res http.ResponseWriter, req *http.Request) {
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
}

func ajaxDeleteStorageDevice(res http.ResponseWriter, req *http.Request) {
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

	err = controller.Device.DeleteStorageDevice(request.UUID)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("失败")
		ResMsg(res, 400, err.Error())
		return
	}
	ResSuccessMsg(res, 200, "操作成功")
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
}

func ajaxCreateCommonDevice(res http.ResponseWriter, req *http.Request) {
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
}

func ajaxDeleteCommonDevice(res http.ResponseWriter, req *http.Request) {
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

	err = controller.Device.DeleteCommonDevice(request.UUID)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("失败")
		ResMsg(res, 400, err.Error())
		return
	}
	ResSuccessMsg(res, 200, "操作成功")
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
}
