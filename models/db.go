package models

import (
	"zcm_tools/orm"
	"order_food/utils"
	_ "github.com/go-sql-driver/mysql"
)

func init(){
	orm.RegisterDataBase("default","mysql",utils.MYSQL_URL)
}