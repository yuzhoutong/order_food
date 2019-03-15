package models

import (
	"zcm_tools/orm"
)

//查询用户列表
func GetUserList(condition string, params []string) (list []OrderUser, err error) {
	sql := `SELECT * FROM order_users where 1=1`
	if condition != "" {
		sql += condition
	}
	_, err = orm.NewOrm().Raw(sql, params).QueryRows(&list)
	return
}

//查询总注册用户
func GetAllRegisterCount() (count int, err error) {
	sql := `SELECT COUNT(1) FROM order_users`
	err = orm.NewOrm().Raw(sql).QueryRow(&count)
	return
}

//查询今日注册用户
func GetTodayRegisterCount() (count int, err error) {
	sql := `SELECT COUNT(1) FROM order_users WHERE DATE_FORMAT(create_time,'%Y-%m-%d') = DATE_FORMAT(NOW(),'%Y-%m-%d')`
	err = orm.NewOrm().Raw(sql).QueryRow(&count)
	return
}

//查询总下单用户
func GetAllOrderCount() (count int, err error) {
	sql := `SELECT COUNT(1) FROM order_table`
	err = orm.NewOrm().Raw(sql).QueryRow(&count)
	return
}

//查询今日下单人数
func GetTodayOrderCount() (count int, err error) {
	sql := `SELECT COUNT(1) FROM order_table WHERE DATE_FORMAT(order_time,'%Y-%m-%d') = DATE_FORMAT(NOW(),'%Y-%m-%d')`
	err = orm.NewOrm().Raw(sql).QueryRow(&count)
	return
}

//根据uid删除用户
func DeleteUserByUid(uid int) (err error) {
	sql := `delete from order_users where order_users_id = ?`
	_, err = orm.NewOrm().Raw(sql, uid).Exec()
	return
}

//根据uid修改用户的状态
func UpdateUserStatus(uid, i int) (err error) {
	sql := `UPDATE  order_users SET accountstatus = ? WHERE order_users_id = ?`
	_, err = orm.NewOrm().Raw(sql, i, uid).Exec()
	return
}

type LoginRecord1 struct {
	Id          int
	Uid         int    //用户id
	LoginTime   string //登录时间
	Name        string //用户名
	Displayname string //角色
}

//获取用户的登录信息
func GetUserLoginInfo(condition string, params []string) (list []LoginRecord1, err error) {
	sql := `SELECT lr.uid, lr.login_time, os.name, os.displayname FROM login_record lr
		LEFT JOIN order_users os
		ON lr.uid = os.order_users_id where 1=1	`
	if condition != " " {
		sql += condition
	}
	sql += " order By login_time desc"
	_, err = orm.NewOrm().Raw(sql, params).QueryRows(&list)
	return
}

//管理员列表(设置)
func GetadminSet() (admin []OrderUser, err error) {
	sql := `SELECT * FROM order_users where displayname = "管理员"`
	_, err = orm.NewOrm().Raw(sql).QueryRows(&admin)
	return
}

//查询管理员信息
func Getadmin(password string) (name string, err error) {
	sql := `select name from order_users where name = "admin" and password = ?`
	err = orm.NewOrm().Raw(sql, password).QueryRow(&name)
	return
}

//修改管理员信息
func UpdateAdmin(password, phone string) (err error) {
	sql := `update order_users set password=?,phone=? where name = "admin" `
	_, err = orm.NewOrm().Raw(sql, password, phone).Exec()
	return
}
