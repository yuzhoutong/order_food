package models

import (
	"zcm_tools/orm"
)

/*type User struct {
	AdminId 		int			//管理员id
	Name			string		//管理员登录名
	Password		string		//管理员登录密码
	Displayname		string		//管理员角色
	Accountstatus	int			//管理员状态
	CreateTime		time.Time	//创建时间
}*/
type OrderUser struct {
	OrderUsersId  int    //uid
	Name          string //登录名
	Password      string //登录密码
	Phone         string //电话
	UserExp       int    //积分
	Displayname   string //用户角色
	Accountstatus int    //用户状态
	CreateTime    string //创建时间
}
type LoginRecord struct {
	Id        int
	Uid       int    //用户id
	LoginTime string //登录时间
}

//用户登录 账号，密码验证
func Login(username, password string) (u *OrderUser, err error) {
	o := orm.NewOrm()
	sql := `SELECT * FROM order_users WHERE name = ? and password = ? `
	err = o.Raw(sql, username, password).QueryRow(&u)
	return
}

//根据用户名获取用户信息
func GetUserByName(name string) (u *OrderUser, err error) {
	sql := `SELECT * FROM order_users ou
	LEFT JOIN detail_address da
	ON ou.order_users_id = da.uid
	WHERE NAME = ?`
	err = orm.NewOrm().Raw(sql, name).QueryRow(&u)
	return
}

//注册账号
func RegisterUser(name, password, phone string) (err error) {
	sql := `INSERT INTO order_users(NAME, PASSWORD, phone, create_time) VALUES(?,?,?,now())`
	o := orm.NewOrm()
	_, err = o.Raw(sql, name, password, phone).Exec()
	return
}

//检查该用户是否注册过
func CheckUserName(name string) (u *OrderUser, err error) {
	sql := `SELECT * from order_users where name = ?`
	err = orm.NewOrm().Raw(sql, name).QueryRow(&u)
	return
}

//修改用户信息
func ModifyUserPwd(name, password string) (err error) {
	sql := `UPDATE order_users SET password = ? where name = ?`
	_, err = orm.NewOrm().Raw(sql, password, name).Exec()
	return
}

//添加登录记录
func InsertLoginRecord(uid int) (err error) {
	sql := `INSERT INTO login_record(uid, login_time) VALUES( ?, NOW())`
	_, err = orm.NewOrm().Raw(sql, uid).Exec()
	return
}

//获取上一次登录时间
func GetUploginTime(uid int) (time string, err error) {
	sql := `SELECT  login_time FROM login_record where uid = ? ORDER BY login_time DESC LIMIT 1, 1 `
	err = orm.NewOrm().Raw(sql, uid).QueryRow(&time)
	return
}

//根据用户uid 获取用户名
func GetUsernameByUid(uid int) (name string, err error) {
	sql := `SELECT name FROM order_users
			where order_users_id = ?`
	err = orm.NewOrm().Raw(sql, uid).QueryRow(&name)
	return
}

//获取用户订单数
func GetUserOrderCount(uid int) (count int, err error) {
	sql := `SELECT count(1) FROM order_table
			where uid = ?`
	err = orm.NewOrm().Raw(sql, uid).QueryRow(&count)
	return
}

////获取用户待付款数
func GetNotBuyCount(uid int) (count int, err error) {
	sql := `SELECT count(1) FROM order_table
			where uid = ? AND is_buy = 2`
	err = orm.NewOrm().Raw(sql, uid).QueryRow(&count)
	return
}
