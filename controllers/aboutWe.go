package controllers

import "github.com/astaxie/beego"

type AboutController struct {
	beego.Controller
}
func (c *AboutController) About(){
	c.TplName = "article_read.html"
}
