package models

import (
	_ "github.com/go-sql-driver/mysql"
	"order_food/utils"
	"zcm_tools/orm"
)

func init() {
	orm.RegisterDataBase("default", "mysql", utils.MYSQL_URL)
}
