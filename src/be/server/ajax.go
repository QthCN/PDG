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