package controllers

import (
	"order_food/models"
)

type AdminImagesController struct {
	HomeController
}

//信息管理
func (c *AdminImagesController) ImagesManagement() { //留言管理
	c.IsNeedTemplate()
	//留言
	list, _ := models.GetLeaveList()

	c.Data["list"] = list
	c.TplName = "back/images_management.html"
}
func (c *AdminImagesController) DeleteInfo() { //删除留言
	resultMap := make(map[string]interface{})
	resultMap["ret"] = 403
	defer func() {
		c.Data["json"] = resultMap
		c.ServeJSON()
	}()
	id, _ := c.GetInt("id")
	err := models.DeleteInfo(id)
	if err != nil {
		resultMap["msg"] = "删除留言失败！！"
	}
	resultMap["ret"] = 200
	resultMap["msg"] = "删除留言成功！！"
}
func (c *AdminImagesController) InformManagement() { //公告管理
	c.IsNeedTemplate()
	//公告
	notice, _ := models.GetNoticeList()
	c.Data["notice"] = notice
	c.TplName = "back/infor_management.html"
}
func (c *AdminImagesController) DeleteNotice() { //删除公告
	resultMap := make(map[string]interface{})
	resultMap["ret"] = 403
	defer func() {
		c.Data["json"] = resultMap
		c.ServeJSON()
	}()
	id, _ := c.GetInt("id")
	err := models.DeleteNotice(id)
	if err != nil {
		resultMap["msg"] = "删除公告失败！！"
	}
	resultMap["ret"] = 200
	resultMap["msg"] = "删除公告成功！！"
}

//添加公告
func (c *AdminImagesController) AddNotice() { //添加公告
	resultMap := make(map[string]interface{})
	resultMap["ret"] = 403
	defer func() {
		c.Data["json"] = resultMap
		c.ServeJSON()
	}()
	var name = c.GetString("name")
	var context = c.GetString("context")
	var using, _ = c.GetInt("using")
	if using == 1 {
		//修改数据库中所有启用的状态为不启用
		err := models.UpdateNoticeUsing()
		if err != nil {
			resultMap["msg"] = "修改状态失败"
		}
	}
	err := models.AddNotice(name, context, using)
	if err != nil {
		resultMap["msg"] = "添加公告失败！"
	}

	resultMap["ret"] = 200
	resultMap["msg"] = "添加公告成功!"
}

//修改公告为启用或者禁用
func (c *AdminImagesController) UpdateNotice() {
	resultMap := make(map[string]interface{})
	resultMap["ret"] = 403
	defer func() {
		c.Data["json"] = resultMap
		c.ServeJSON()
	}()
	id, _ := c.GetInt("id")
	user, _ := c.GetInt("user")
	if user == 1 { //当前是启用 ->到禁用
		user = 2
		err := models.UpdateNotice(id, user)
		if err != nil {
			resultMap["msg"] = "修改状态失败！"
		}
	} else if user == 2 { //当前是禁用 -> 启用 (其他全为禁用)
		user = 1
		models.UpdateNoticeUsing() //将公告全部设为禁用
		err := models.UpdateNotice(id, user)
		if err != nil {
			resultMap["msg"] = "修改状态失败！"
		}
	}
	resultMap["ret"] = 200
	resultMap["msg"] = "修改状态成功！"
}
