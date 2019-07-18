package controllers

import (
	"order_food/models"
	"order_food/cache"
)

type AdminUserController struct {
	HomeController
}

//用户管理-用户登录记录
func (c *AdminUserController) UserManagement() {
	c.IsNeedTemplate()
	condition := ""
	parars := []string{}
	if name := c.GetString("name"); name != "" {
		condition += " and os.name = ?"
		parars = append(parars, name)
	}
	list, err := models.GetUserLoginInfo(condition, parars)
	if err != nil {
		cache.RecordLogs(c.OrderUser.OrderUsersId, 0, c.OrderUser.Name, c.OrderUser.Displayname, "获取用户登录记录失败", "用户登录记录/UserManagement", err.Error(), c.Ctx.Input)
		return
	}
	c.Data["list"] = list
	c.TplName = "back/user_management.html"
}

//管理员-删除用户
func (c *AdminUserController) DeleteUserInfo() {
	resultMap := make(map[string]interface{})
	resultMap["ret"] = 403
	defer func() {
		c.Data["json"] = resultMap
		c.ServeJSON()
	}()
	//获取用户的uid
	uid, _ := c.GetInt("uid")
	err := models.DeleteUserByUid(uid)
	if err != nil {
		resultMap["msg"] = "删除用户错误！"
		cache.RecordLogs(c.OrderUser.OrderUsersId, 0, c.OrderUser.Name, c.OrderUser.Displayname, "删除用户错误", "删除用户/DeleteUserInfo", err.Error(), c.Ctx.Input)
		return
	}
	resultMap["ret"] = 200
	resultMap["msg"] = "删除成功"
}

//管理员修改用户的状态
func (c *AdminUserController) UpdateUserStatus() {
	resultMap := make(map[string]interface{})
	resultMap["ret"] = 403
	defer func() {
		c.Data["json"] = resultMap
		c.ServeJSON()
	}()
	//获取用户的uid
	uid, _ := c.GetInt("uid")
	i, _ := c.GetInt("i")
	if i == 1 {
		i = 2
	} else if i == 2 {
		i = 1
	}
	err := models.UpdateUserStatus(uid, i)
	if err != nil {
		resultMap["msg"] = "修改用户状态错误！"
		cache.RecordLogs(c.OrderUser.OrderUsersId, 0, c.OrderUser.Name, c.OrderUser.Displayname, "修改用户状态错误", "管理员修改用户的状态/UpdateUserStatus", err.Error(), c.Ctx.Input)
		return
	}
	resultMap["ret"] = 200
	resultMap["msg"] = "修改用户状态成功"
}

//管理员设置
func (c *AdminUserController) SetAdmin() {
	c.IsNeedTemplate()
	admin, _ := models.GetadminSet()
	c.Data["admin"] = admin
	c.TplName = "back/admin_set.html"
}

//修改管理员信息
func (c *AdminUserController) UpdateAdmin() {
	resultMap := make(map[string]interface{})
	resultMap["ret"] = 403
	defer func() {
		c.Data["json"] = resultMap
		c.ServeJSON()
	}()
	//查询旧密码是否正确
	oldpwd := c.GetString("oldpwd")
	newpwd := c.GetString("newpwd")
	phone := c.GetString("phone")
	name, _ := models.Getadmin(oldpwd)
	/*if err !=nil {
		resultMap["msg"] = "查询错误"
		return
	}*/
	if name == "" {
		resultMap["msg"] = "原密码错误"
		return
	}
	err := models.UpdateAdmin(newpwd, phone)
	if err != nil {
		resultMap["msg"] = "修改密码失败"
		cache.RecordLogs(c.OrderUser.OrderUsersId, 0, c.OrderUser.Name, c.OrderUser.Displayname, "修改密码失败", "修改管理员信息/UpdateAdmin", err.Error(), c.Ctx.Input)
		return
	}
	resultMap["ret"] = 200
	resultMap["msg"] = "修改成功"

}
