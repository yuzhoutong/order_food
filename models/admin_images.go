package models

import "zcm_tools/orm"

//查询留言管理列表
func GetLeaveList() (list []UserEvaluate, err error) {
	sql := `SELECT * FROM order_evaluate  order by create_time desc`
	_, err = orm.NewOrm().Raw(sql).QueryRows(&list)
	return
}

//删除留言
func DeleteInfo(id int) (err error) {
	sql := `delete from order_evaluate where id = ?`
	_, err = orm.NewOrm().Raw(sql, id).Exec()
	return
}

//公告展示
type Notice struct {
	Id         int
	Name       string //操作人
	Context    string //公告内容
	IsUser     int    //是否使用
	CreateTime string //创建时间
}

func GetNoticeList() (list []Notice, err error) {
	sql := `SELECT * FROM notice  order by create_time desc`
	_, err = orm.NewOrm().Raw(sql).QueryRows(&list)
	return
}

//删除公告
func DeleteNotice(id int) (err error) {
	sql := `delete from notice where id = ?`
	_, err = orm.NewOrm().Raw(sql, id).Exec()
	return
}

//添加公告
func AddNotice(name, context string, using int) (err error) {
	sql := `INSERT INTO notice (NAME, context, is_user,create_time) VALUE (?, ?, ?, NOW())`
	_, err = orm.NewOrm().Raw(sql, name, context, using).Exec()
	return
}

//添加公告 如果添加启用 则修改其他公告状态 为禁用
func UpdateNoticeUsing() (err error) {
	sql := `Update notice set is_user = 2`
	_, err = orm.NewOrm().Raw(sql).Exec()
	return
}

//禁用-》启用
func UpdateNotice(id, user int) (err error) {
	sql := `Update notice set is_user = ? where id = ?`
	_, err = orm.NewOrm().Raw(sql, user, id).Exec()
	return
}
