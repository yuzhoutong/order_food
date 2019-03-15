package utils

import "time"

// 缓存时间
const (
	RedisCacheTime_User    = 15 * time.Minute
	RedisCacheTime_TwoHour = 2 * time.Hour
	RedisCacheTime_Role    = 15 * time.Second
)

// 常量配置
const (
	PAGE_SIZE10    = 10 //列表页每页数据量
	PAGE_SIZE15    = 15
	PAGE_SIZE20    = 20
	PAGE_SIZE40    = 40
	PAGE_SIZE500   = 500
	FormatTime     = "15:04:05"            //时间格式
	FormatDate     = "2006-01-02"          //日期格式
	FormatDateTime = "2006-01-02 15:04:05" //完整时间格式
	WITHDRAWDAY    = "Thursday"
)

//验证配置
const (
	Regular = "^((13[0-9])|(14[5|7])|(15([0-3]|[5-9]))|(18[0-9])|(17[0-9]))\\d{8}$"
)
