package controllers

import (
	"github.com/astaxie/beego"
	"order_food/models"
)

type AboutController struct {
	beego.Controller
}

func (c *AboutController) About() {
	//公告
	notice, _ := models.GetNotic()
	c.Data["notice"] = notice
	c.TplName = "article_read.html"
}
