package models

import (
	"zcm_tools/orm"
)

//查询订单列表
func GetOrderListAdmin(condition string, params []string) (list []UserOders, err error) {
	sql := `SELECT * FROM order_table where 1=1`
	sql += condition
	_, err = orm.NewOrm().Raw(sql, params).QueryRows(&list)
	return
}

//根据orderId删除用户
func DeleteOrderByOrderId(OrderId string) (err error) {
	sql := `delete from order_table where order_id = ?`
	_, err = orm.NewOrm().Raw(sql, OrderId).Exec()
	return
}
