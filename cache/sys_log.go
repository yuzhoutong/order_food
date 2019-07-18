package cache

import (
	"github.com/astaxie/beego/context"
	"net/url"
	"time"
	"order_food/models"
	"order_food/store"
)

func RecordLogs(user_id, business_id int, username, displayname, action, logger, message string, input *context.BeegoInput) bool {
	ip := input.IP()
	urlpath := input.URL()
	querystrings := input.URI()
	fromparams, _ := url.QueryUnescape(string(input.RequestBody))
	log := &models.SysLog{UserId: user_id, UserName: username, UserDisplayName: displayname, UserIp: ip, Action: action, Logger: logger, UrlPath: urlpath, Message: message, FromParams: fromparams, QueryStrings: querystrings, CreateTime: time.Now(), BusinessId: business_id}
	if store.Re == nil {
		store.Rc.LPush(store.CacheKeySystemLogs, log)
		return true
	}
	return false
}
