package main

import (
	"be/mysql"
	"be/option"
	"be/server"
	"be/util"
	"math/rand"
	"net/http"
	"os"
	"runtime"
	"runtime/debug"
	"strings"
	"time"

	go_log "log"

	log "github.com/sirupsen/logrus"
)

func initLog() {
	log.SetReportCaller(true)

	customFormatter := new(log.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	customFormatter.FullTimestamp = true
	log.SetFormatter(customFormatter)

	if *option.LogFile != "" {
		f, err := os.OpenFile(*option.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			panic("无法打开日志文件")
		}
		log.SetOutput(f)
	} else {
		log.SetOutput(os.Stdout)
	}

	switch strings.ToUpper(*option.LogLevel) {
	case "DEBUG":
		log.SetLevel(log.DebugLevel)
	case "INFO":
		log.SetLevel(log.InfoLevel)
	case "WARN":
		log.SetLevel(log.WarnLevel)
	case "ERROR":
		log.SetLevel(log.ErrorLevel)
	default:
		log.SetLevel(log.DebugLevel)
	}

	log.Infoln("日志文件初始化成功")
}

func doServe() {
	defer func() {
		if err := recover(); err != nil {
			log.Errorf("%s", err)
			log.Errorln(string(debug.Stack()))
			doServe()
		}
	}()

	// 初始化DB
	mysql.DB.InitConn()

	// 初始化服务,并启动服务
	httpServer := server.New()

	// http服务
	srv := &http.Server{
		Handler:      httpServer.GetCORSHandler(),
		Addr:         *option.HTTPAddress,
		WriteTimeout: 60 * time.Second,
		ReadTimeout:  60 * time.Second,
		ErrorLog:     go_log.New(log.StandardLogger().Writer(), "", 0),
	}

	// 启动
	log.Fatal(srv.ListenAndServe())
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	rand.Seed(time.Now().UTC().UnixNano())

	// 初始化配置
	option.InitOption()

	// 初始化日志
	initLog()

	// 初始化cookie
	util.InitCM()

	// 启动服务
	doServe()
}
