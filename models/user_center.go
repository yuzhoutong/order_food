package models

import (
	"time"
	"zcm_tools/orm"
)

type UserCenter struct {
	OrderUsersId		int
	Name			string		//登录名
	Password		string		//登录密码
	UserExp			int			//积分
	Displayname		string		//用户角色
	Accountstatus	int			//用户状态
	Id				int			//用户id
	Address			string		//地址
	Phone 			string		//电话
	Rname			string		//收件人名字
	CreateTime		time.Time	//创建时间
}
type UserAddress struct {
	Id				int
	Uid				int			//用户id
	Address			string		//地址
	RePhone 		string		//电话
	Rname			string		//收件人名字
	Provice			string		//省份
	City			string		//城市
	District		string		//区
	CreateTime		time.Time	//创建时间
}
//修改收货地址
func UpdateAddress(uid, id int,address, rename, phone, provice, city, distract string) (err error){
	sql := `UPDATE detail_address da SET da.address = ? ,da.rname = ? ,da.re_phone = ?, da.provice = ?,
			da.city = ?, da.district = ? WHERE da.uid = ? and da.id = ?`
	_,err = orm.NewOrm().Raw(sql, address, rename, phone, provice, city, distract, uid, id).Exec()
	return
}
//查询用户的信息(地址+名字+手机号）
func GetUserAddress(id int)(u []UserAddress,err error){
	sql := `SELECT * FROM detail_address da
			LEFT JOIN order_users ou
			ON da.uid = ou.order_users_id
 			WHERE uid = ?`
	_,err = orm.NewOrm().Raw(sql,id).QueryRows(&u)
	return
}
//添加收货地址
func AddAddress(uid int, address, rename, phone , provice, city, distract string)(err error){
	sql := `INSERT INTO detail_address (uid,re_phone,address,rname,createtime,provice,city,district)
		VALUES(?, ?, ?, ?, NOW(), ?, ?, ?)`
	_,err = orm.NewOrm().Raw(sql, uid, phone, address, rename , provice, city, distract).Exec()
	return
}
//删除地址
func DelAddress(uid, id int)(err error){
	sql := `delete from detail_address where id = ? and uid = ?`
	_,err = orm.NewOrm().Raw(sql, id, uid).Exec()
	return
}
















