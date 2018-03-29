package utils

import "time"

// 缓存key
const (
	CACHE_KEY_USER_PREFIX         = "wr_hezuo_CACHE_KEY_USER_PREFIX_"         //根据code获取用户ticket
	CACHE_KEY_USER_INFO           = "wr_hezuo_CACHE_KEY_USER_INFO_"           //根据code获取用户信息
	CACHE_KEY_USER_ACCOUNT_INFO   = "wr_hezuo_CACHE_KEY_USER_ACCOUNT_INFO_"   //根据用户id获取用户账号信息
	CACHE_KEY_SYSTEM_MENU         = "wr_hezuo_CACHE_KEY_SYSTEM_MENU"          //所有菜单
	CACHE_KEY_ROLE_ADD_RIGHT      = "wr_hezuo_CACHE_KEY_ROLE_ADD_RIGHT"       //所有角色添加账号权限
	CACHE_KEY_USER_ROLE_ADD_RIGHT = "wr_hezuo_CACHE_KEY_USER_ROLE_ADD_RIGHT_" //角色添加账号权限
	CACHE_KEY_ROLE_MENU_LIST      = "wr_hezuo_CACHE_KEY_ROLE_MENU_LIST_"      //根据角色id获取用户菜单
	CACHE_KEY_SYSTEM_LOGS         = "wr_hezuo_CACHE_KEY_SYSTEM_LOGS"          //微融合作平台日志
	CACHE_KEY_ACCOUNT_OPERATIONAL = "wr_hezuo_CACHE_KEY_ACCOUNT_OPERATIONAL"  //账号添加/编辑频繁提交控制
)

// 缓存时间
const (
	RedisCacheTime_User         = 15 * time.Minute
	RedisCacheTime_TwoHour      = 2 * time.Hour
	RedisCacheTime_Role         = 15 * time.Second
	RedisCacheTime_Organization = 24 * time.Hour //24 * time.Hour //组织架构信息缓存时间
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

// 邮件
const (
	ToUsers = "lijr@zcmlc.com"
)
