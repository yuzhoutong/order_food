package controllers

import (
	"github.com/astaxie/beego"
	"order_food/models"
	"order_food/cache"
)

type AboutController struct {
	beego.Controller
}

func (c *AboutController) About() {
	//公告
	notice, err := models.GetNotic()
	if err != nil {
		cache.RecordLogs(0, 0, "ss", "ss", "获取公告失败", "公告/About", err.Error(), c.Ctx.Input)
	}
	c.Data["notice"] = notice
	c.TplName = "article_read.html"
}
