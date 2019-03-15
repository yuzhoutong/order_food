package controllers

import (
	"github.com/astaxie/beego"
	"order_food/models"
	"order_food/utils"
)

type HomeController struct {
	beego.Controller
	OrderUser models.OrderUser
}

func (c *HomeController) Prepare() {
	var id = c.Ctx.GetCookie("id")
	if id == "" {
		c.Ctx.Redirect(302, "/login")
	}
}

//用户页面
func (c *HomeController) IndexUser() {
	list, _ := models.GetOrderTableINf()
	allEvaluate, _ := models.GetAllUserEvaluateList()
	//公告
	notice, _ := models.GetNotic()
	//lunbo
	OrderList, _ := models.GetNamePic()
	c.Data["OrderList"] = OrderList
	c.Data["notice"] = notice
	c.Data["list"] = list
	c.Data["allEvaluate"] = allEvaluate
	c.Data["username"] = c.OrderUser.Name
	c.TplName = "index.html"
}

//管理员主页
func (c *HomeController) IndexAdmin() {
	c.IsNeedTemplate()
	condition := ""
	params := []string{}
	index := []int{}
	if name := c.GetString("name"); name != "" {
		condition = " and name = ?"
		params = append(params, name)
	}

	//用户列表
	userList, _ := models.GetUserList(condition, params)
	//获取注册总人数
	RegisterCount, _ := models.GetAllRegisterCount()
	pages := utils.PageCount(RegisterCount, utils.PAGE_SIZE10)
	for i := 1; i <= pages; i++ {
		index = append(index, i)
	}
	//获取今日注册人数
	TodayRegisterCount, _ := models.GetTodayRegisterCount()
	//获取下单总人数
	OrderCount, _ := models.GetAllOrderCount()
	//获取今日下单人数
	TodayOrderCount, _ := models.GetTodayOrderCount()
	c.Data["userList"] = userList
	c.Data["RegisterCount"] = RegisterCount
	c.Data["TodayRegisterCount"] = TodayRegisterCount
	c.Data["OrderCount"] = OrderCount
	c.Data["TodayOrderCount"] = TodayOrderCount
	c.Data["index"] = index
	c.TplName = "back/index.html"
}

func (this *HomeController) IsNeedTemplate() {
	name := this.Ctx.GetCookie("name")
	this.Data["name"] = name
	this.Layout = "back/starter.html"
}

type msgResp struct {
	Status bool
	Object interface{}
	Page   int
	Pages  int
	Msg    string
	Next   int
	Prev   int
}

func (this *HomeController) UserList() {
	var resp msgResp
	defer func() {
		this.Data["json"] = resp
		this.ServeJSON()
	}()
	condition := ""
	params := []string{}
	if name := this.GetString("name"); name != "" {
		condition = " and name = ?"
		params = append(params, name)
	}
	//用户列表
	userList, err := models.GetUserList(condition, params)
	if err != nil {
		resp.Msg = err.Error()
		return
	}
	if err != nil {
		resp.Msg = err.Error()
		return
	}
	resp.Object = userList
	return

}
