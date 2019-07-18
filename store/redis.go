package store

import (
	"log"
	"order_food/utils"
	"zcm_tools/cache"
)

var (
	Rc *cache.Cache //redis缓存
	Re error        //redis错误
)

const (
	CacheKeySystemLogs              = "CacheKeySystemLogs"
	CacheKeyUserInfo              = "CacheKeyUserInfo"

)
func init() {

	Rc, Re = cache.NewCache(utils.BEEGO_CACHE)
	if Re != nil {
		log.Fatal(Re)
	}else {
		log.Println("redis connection success")
	}
}
