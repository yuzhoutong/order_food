package models

import (
	"time"
	"zcm_tools/orm"
)

type UserCenter struct {
	OrderUsersId  int
	Name          string    //登录名
	Password      string    //登录密码
	UserExp       int       //积分
	Displayname   string    //用户角色
	Accountstatus int       //用户状态
	Id            int       //用户id
	Address       string    //地址
	Phone         string    //电话
	Rname         string    //收件人名字
	CreateTime    time.Time //创建时间
}
type UserAddress struct {
	Id         int
	Uid        int       //用户id
	Address    string    //地址
	RePhone    string    //电话
	Rname      string    //收件人名字
	Provice    string    //省份
	City       string    //城市
	District   string    //区
	CreateTime time.Time //创建时间
}
type UserOders struct {
	Id          int
	Uid         int    //用户id
	OrderId     string //订单id
	Name        string //收件人姓名
	IsProcessed int    //订单是否被处理 1:已处理，2:未处理
	OrderTime   string //订单时间
	OrderPrice  string //订单金额
	IsBuy       int    //是否付款:1:已支付:2未支付
	AddressId   int    //地址id
	IsEvaluate  int    //是否评价1:已评价 2:未评价 0:表示为付款没有评价功能
}
type OrderDetail struct {
	Id         int
	Uid        int       //用户id
	OrderId    string    //订单id
	Name       string    //菜名
	Count      int       //数量
	Price      string    //价格
	CreateTime time.Time //创建时间
}

//修改收货地址
func UpdateAddress(uid, id int, address, rename, phone, provice, city, distract string) (err error) {
	sql := `UPDATE detail_address da SET da.address = ? ,da.rname = ? ,da.re_phone = ?, da.provice = ?,
			da.city = ?, da.district = ? WHERE da.uid = ? and da.id = ?`
	_, err = orm.NewOrm().Raw(sql, address, rename, phone, provice, city, distract, uid, id).Exec()
	return
}

//查询用户的信息(地址+名字+手机号）
func GetUserAddress(id int) (u []UserAddress, err error) {
	sql := `SELECT * FROM detail_address da
			LEFT JOIN order_users ou
			ON da.uid = ou.order_users_id
 			WHERE uid = ?`
	_, err = orm.NewOrm().Raw(sql, id).QueryRows(&u)
	return
}

//添加收货地址
func AddAddress(uid int, address, rename, phone, provice, city, distract string) (err error) {
	sql := `INSERT INTO detail_address (uid,re_phone,address,rname,createtime,provice,city,district)
		VALUES(?, ?, ?, ?, NOW(), ?, ?, ?)`
	_, err = orm.NewOrm().Raw(sql, uid, phone, address, rename, provice, city, distract).Exec()
	return
}

//删除地址
func DelAddress(uid, id int) (err error) {
	sql := `delete from detail_address where id = ? and uid = ?`
	_, err = orm.NewOrm().Raw(sql, id, uid).Exec()
	return
}

//获取该用户的订单
func GetUserOrder(uid int) (Order []UserOders, err error) {
	sql := `SELECT * FROM order_table
			WHERE uid = ? order by order_time desc`
	_, err = orm.NewOrm().Raw(sql, uid).QueryRows(&Order)
	return
}

//(取消)删除用户订单
func DeleteOrder(uid int, orderId string) (err error) {
	sql := `DELETE FROM order_table WHERE uid = ? AND order_id = ?`
	_, err = orm.NewOrm().Raw(sql, uid, orderId).Exec()
	return
}

//点击订单号显示订单信息
func GetOrderImanages(uid int, orderId string) (list []OrderDetail, err error) {
	sql := `SELECT * FROM order_detail  WHERE uid = ? And order_id = ?`
	_, err = orm.NewOrm().Raw(sql, uid, orderId).QueryRows(&list)
	return
}

//根据订单号获取地址id
func GetAddressIdByOrderId(orderId string) (list int, err error) {
	sql := `SELECT address_id FROM order_table  WHERE order_id = ?`
	err = orm.NewOrm().Raw(sql, orderId).QueryRow(&list)
	return
}

//根据地址id获取地址信息
func GetAddressByAddressId(address int) (addr *UserAddress, err error) {
	sql := `SELECT *  FROM detail_address  WHERE id = ?`
	err = orm.NewOrm().Raw(sql, address).QueryRow(&addr)
	return
}

//根据uid与order_id获取该用户是否付款
func IsBuyByUIdAndOrderId(uid int, orderId string) (i int, err error) {
	sql := `SELECT is_buy FROM order_table
			WHERE uid = ? AND order_id = ?`
	err = orm.NewOrm().Raw(sql, uid, orderId).QueryRow(&i)
	return
}

//根据地址id 查询出收货地址的姓名
func GetNameByAddressId(addressId int) (name string, err error) {
	sql := `select rname from detail_address where id = ?`
	err = orm.NewOrm().Raw(sql, addressId).QueryRow(&name)
	return
}

//根据查出的收货地址姓名来修改orderby的订单人的姓名
func UpdateOrderNameByRname(rname string, uid, id int) (err error) {
	sql := `Update order_table set name = ? where uid = ? and address_id = ?`
	_, err = orm.NewOrm().Raw(sql, rname, uid, id).Exec()
	return
}

type UserEvaluate struct {
	Id          int
	OrderId     string //订单id
	Uid         int    //点单用户的名字
	Speed       int    //配送速度  1心:差,2-3心:一般,4心:好,5心:非常好
	Server      int    //服务态度 1心:差,2-3心:一般,4心:好,5心:非常好
	DishesTaste int    //菜品口味 1心:差,2-3心:一般,4心:好,5心:非常好
	Remark      string //备注
	CreateTime  string //创建时间
	Name        string //用户名
}

//用户评价
func GetUserEvaluate(uid int) (EvaList []UserEvaluate, err error) {
	sql := `SELECT * FROM order_evaluate where uid = ? order by create_time desc`
	_, err = orm.NewOrm().Raw(sql, uid).QueryRows(&EvaList)
	return
}

//添加用户评价
func AddUserEvaluate(orderId, remark string, uid, speed, server, taste int, name string) (err error) {
	sql := `INSERT INTO order_evaluate (order_id, uid, speed, SERVER, dishes_taste, remark, create_time, name)
			VALUES(?, ?, ?, ?, ?, ?, NOW(), ?);`
	_, err = orm.NewOrm().Raw(sql, orderId, uid, speed, server, taste, remark, name).Exec()
	return
}

//修改用户订单表评价的状态为已评价
func UpdateUserEvaluateStatus(orderId string) (err error) {
	sql := `UPDATE order_table SET is_evaluate = 1 where order_id = ?`
	_, err = orm.NewOrm().Raw(sql, orderId).Exec()
	return
}

//删除用户评价
func DeleteUserEvaluate(orderId, time string) (err error) {
	sql := `delete from order_evaluate where order_id = ? and create_time = ?`
	_, err = orm.NewOrm().Raw(sql, orderId, time).Exec()
	return
}
