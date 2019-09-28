package server

import (
	"be/option"
	"net/http"
	"path/filepath"
	"text/template"
)

func templateRealPath(path string) string {
	return filepath.Join(*option.TemplateFilePath, path)
}

func showIndexHtml(res http.ResponseWriter, req *http.Request) {
	tmpl, _ := template.ParseFiles(templateRealPath("index.html"))
	tmpl.ExecuteTemplate(res, "index", nil)
}

func showLoginHtml(res http.ResponseWriter, req *http.Request) {
	tmpl, _ := template.ParseFiles(templateRealPath("login.html"))
	tmpl.ExecuteTemplate(res, "login", nil)
}
