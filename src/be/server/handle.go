package server

import (
	"be/option"
	"net/http"
)

func (s *Server) registAPI() {
	// 初始化静态文件路径
	initStaticFileMapping(s)
	// 初始化Portal页面
	initPortal(s)
	// 初始化ajax接口
	initAjaxMapping(s)
}

func initStaticFileMapping(r *Server) {
	fs := http.FileServer(http.Dir(*option.StaticFilePath))
	r.GetRouter().PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
}

func initPortal(r *Server) {
	r.RegistURLMapping("/", "GET", showIndexHtml)
	r.RegistURLMapping("/login.html", "GET", showLoginHtml)
}

func initAjaxMapping(r *Server) {
	// 用户认证密码并生成token
	r.RegistURLMapping("/v1/ajax/auth/token", "POST", ajaxGenTokenByUMAndPassword)
	// 获取用户信息
	r.RegistURLMapping("/v1/ajax/auth/info", "GET", ajaxGetUserInfo)
}
