package utils

import (
	"zcm_tools/cache"

	"github.com/astaxie/beego"
)

// MYSQL
var (
	RunMode   string // 运行模式
	MYSQL_URL string // 微融主库
)

// Redis
var (
	Rc          *cache.Cache // redis缓存
	Re          error        // redis错误
	BEEGO_CACHE string       // redis地址
)

// base
var (
	Enablexsrf string // XSRF校验开关
	H5Encoded  string // H5接口base64编码开关
)

func init() {
	RunMode = beego.AppConfig.String("run_mode")
	config, err := beego.AppConfig.GetSection(RunMode)
	if err != nil {
		panic("配置文件读取错误 " + err.Error())
	}
	beego.Info("access")
	// mysql
	MYSQL_URL = config["mysql_url"]
	// redis
	// show
	beego.Info("┌───────────────────")
	beego.Info("│模式:" + RunMode)
	beego.Info("│XSRF校验:" + Enablexsrf)
	beego.Info("└───────────────────")
}
