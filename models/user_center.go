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
type UserOders struct {
	Id  			int
	Uid  			int			//用户id
	OrderId			string		//订单id
	Name 			string		//收件人姓名
	IsProcessed		int			//订单是否被处理 1:已处理，2:未处理
	OrderTime		string		//订单时间
	OrderPrice		string		//订单金额
	IsBuy			int         //是否付款:1:已支付:2未支付
}
type OrderDetail struct {
	Id  			int
	Uid  			int			//用户id
	OrderId			string		//订单id
	Name 			string		//菜名
	Count 			int 	    //数量
	Price			string		//价格
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

//获取该用户的订单
func GetUserOrder(uid int) (Order []UserOders, err error){
	sql := `SELECT * FROM order_table
			WHERE uid = ?`
	_,err = orm.NewOrm().Raw(sql, uid).QueryRows(&Order)
	return
}

//(取消)删除用户订单
func DeleteOrder(uid int, orderId string) (err error){
	sql := `DELETE FROM order_table WHERE uid = ? AND order_id = ?`
	_, err = orm.NewOrm().Raw(sql, uid, orderId).Exec()
	return
}
//点击订单号显示订单信息
func GetOrderImanages(uid int, orderId string) (list []OrderDetail, err error){
	sql := `SELECT * FROM order_detail  WHERE uid = ? And order_id = ?`
	_,err = orm.NewOrm().Raw(sql, uid, orderId).QueryRows(&list)
	return
}













