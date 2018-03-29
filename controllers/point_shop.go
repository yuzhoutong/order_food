package controllers

import "github.com/astaxie/beego"

type PointShopController struct {
	beego.Controller
}
func (c *PointShopController) Shop(){
	c.TplName = "category.html"
}
