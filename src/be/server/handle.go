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
}
