package controllers

import (
	"github.com/astaxie/beego"
	"order_food/models"
	"strconv"
)

type AccountController struct {
	beego.Controller
}

//登录界面
func (c *AccountController) Login() {
	c.TplName = "login.html"
}

//登录检查

func (c *AccountController) CheckPassword() {
	resultMap := make(map[string]interface{})
	resultMap["ret"] = 403
	defer func() {
		c.Data["json"] = resultMap
		c.ServeJSON()
	}()
	//获取用户输入的用户名与密码
	username := c.GetString("username")
	password := c.GetString("password")
	//验证登录
	users, _ := models.Login(username, password)
	if users == nil {
		resultMap["msg"] = "用户名或者密码错误"
		return
	}
	//判断用户是否被冻结
	if users.Accountstatus == 2 {
		resultMap["msg"] = "账号已被冻结，请联系管理员"
		return
	}
	//管理员与用户的区分
	if users.Displayname == "管理员" {
		resultMap["displayname"] = "管理员"
	} else {
		resultMap["displayname"] = "用户"
	}
	//存入Cookie
	c.Ctx.SetCookie("name", users.Name)
	c.Ctx.SetCookie("id", strconv.Itoa(users.OrderUsersId))
	c.Ctx.SetCookie("uid", strconv.Itoa(users.OrderUsersId))
	uid := users.OrderUsersId
	//添加登录记录
	err := models.InsertLoginRecord(uid)
	if err != nil {
		resultMap["msg"] = "添加登录历史错误"
		return
	}
	resultMap["ret"] = 200
	resultMap["msg"] = "登录成功"
	resultMap["users"] = users
	return
}

//退出登录
func (c *AccountController) LogOut() {
	name := c.Ctx.GetCookie("name")
	users, _ := models.GetUserByName(name)

	if name != "" && users != nil {
		//清楚cookie
		c.Ctx.SetCookie("name", "", -1)
	}
	c.Ctx.Redirect(302, "/login")
}

//用户注册

func (c *AccountController) Register() {
	resultMap := make(map[string]interface{})
	resultMap["ret"] = 403
	defer func() {
		c.Data["json"] = resultMap
		c.ServeJSON()
	}()
	//获取用户输入的姓名，密码，手机号
	username := c.GetString("username")
	password := c.GetString("password")
	phone := c.GetString("phone_r")
	//判断该用户名字是否注册
	user, _ := models.CheckUserName(username)
	if user != nil {
		resultMap["msg"] = "该名字已被注册过了"
		return
	}
	//注册用户
	err := models.RegisterUser(username, password, phone)
	if err != nil {
		resultMap["msg"] = "注册用户失败!"
		return
	}
	resultMap["ret"] = 200
	resultMap["msg"] = "注册成功!"
	return
}

//修改用户信息
func (c *AccountController) ToModifyUser() {
	//公告
	notice, _ := models.GetNotic()
	c.Data["notice"] = notice
	c.TplName = "user_account.html"
}
func (c AccountController) ModifyUser() {
	resultMap := make(map[string]interface{})
	resultMap["ret"] = 403
	defer func() {
		c.Data["json"] = resultMap
		c.ServeJSON()
	}()
	name := c.GetString("name")
	oldpwd := c.GetString("Opassword")
	newpwd := c.GetString("Npassword")
	if oldpwd == newpwd {
		resultMap["msg"] = "新密码与原密码不能一样"
		return
	}
	user, _ := models.Login(name, oldpwd)
	if user == nil {
		resultMap["msg"] = "原密码错误"
		return
	}
	err := models.ModifyUserPwd(name, newpwd)
	if err != nil {
		resultMap["msg"] = "修改密码失败"
		return
	}
	resultMap["ret"] = 200
	resultMap["msg"] = "修改密码成功"
	return
}

//用户中心
func (c *AccountController) UserInformation() {
	//获取参数
	var id = c.Ctx.GetCookie("id")
	uid, _ := strconv.Atoi(id)
	time, _ := models.GetUploginTime(uid)
	//获取用户订单数
	OrderCount, _ := models.GetUserOrderCount(uid)
	//获取待付款数
	NotBuyCount, _ := models.GetNotBuyCount(uid)
	//公告
	notice, _ := models.GetNotic()
	c.Data["notice"] = notice
	c.Data["OrderCount"] = OrderCount
	c.Data["NotBuyCount"] = NotBuyCount
	c.Data["time"] = time
	c.TplName = "user_center.html"
}
