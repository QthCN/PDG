package option

import (
	"flag"

	"github.com/vharitonsky/iniflags"
)

var (
	Mode = flag.String("Mode", "PRD", "运行模式(DEV/PRD)")

	LogLevel = flag.String("LogLevel", "DEBUG", "日志级别")
	LogFile  = flag.String("LogFile", "/tmp/service.log", "日志文件")

	Cookie01 = flag.String("Cookie01", "qaz2sxedcrfv0987z54321poastgdfew", "HashKey")
	Cookie02 = flag.String("Cookie02", "t6s5gghd8uzh4red", "BlockKey")

	HTTPAddress = flag.String("HTTPAddress", "0.0.0.0:18080", "HTTP服务地址")

	StaticFilePath   = flag.String("StaticFilePath", "", "静态文件路径")
	TemplateFilePath = flag.String("TemplateFilePath", "", "web页面文件路径")

	DataSourceName = flag.String("DataSourceName", "root:rootroot@tcp(127.0.0.1:3306)/PDG?autocommit=0&collation=utf8_general_ci", "数据库信息")
	DBMaxOpenConn  = flag.Int("DBMaxOpenConn", 32, "最大DB连接数")
	DBMaxIdleConn  = flag.Int("DBMaxIdleConn", 16, "日常DB连接数")

	TokenExpireInterval = flag.Int("TokenExpireInterval", 720, "token过期时间(小时)")
)

func InitOption() {
	iniflags.Parse()
}
