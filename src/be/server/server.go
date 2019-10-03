package server

import (
	"net/http"
	"reflect"
	"runtime"
	"strings"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	r *mux.Router
}

func New() *Server {
	server := &Server{r: mux.NewRouter()}

	// api
	server.registAPI()

	// middleware
	server.r.Use(loggingMiddleware)
	return server
}

func (m *Server) GetRouter() *mux.Router {
	return m.r
}

func (m *Server) GetCORSHandler() http.Handler {
	headersOk := handlers.AllowedHeaders([]string{"*"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"*"})

	return handlers.CORS(originsOk, headersOk, methodsOk)(m.r)
}

func (m *Server) RegistURLMapping(path string, method string, handle func(http.ResponseWriter, *http.Request)) {
	log.WithFields(log.Fields{
		"path":   path,
		"method": method,
		"handle": runtime.FuncForPC(reflect.ValueOf(handle).Pointer()).Name(),
	}).Info("注册URL映射")
	m.r.HandleFunc(path, handle).Methods(method)
	if strings.ToUpper(method) == "GET" {
		m.r.HandleFunc(path, handle).Methods("POST")
	}
}
