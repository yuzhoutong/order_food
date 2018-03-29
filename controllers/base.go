package controllers

import (
	"github.com/astaxie/beego"
	"order_food/models"
)

type HomeController struct {
	beego.Controller
	OrderUser models.OrderUser
}
func (c *HomeController) IndexAdmin(){
	c.TplName = "index_admin.html"
}
func (c *HomeController) IndexUser(){
	c.Data["username"] = c.OrderUser.Name
	c.TplName = "index.html"
}