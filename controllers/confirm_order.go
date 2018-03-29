package controllers

import (
	"github.com/astaxie/beego"
	"order_food/models"
	"strconv"
)

type ConfirmOrder struct {
	beego.Controller
}
//结算购物车页面渲染
func (c *ConfirmOrder) ToOrderConfirm(){
	//用户的uid
	var id= c.Ctx.GetCookie("id")
	uid,_ := strconv.Atoi(id)
	//获取用户的收货地址信息
	address, _ := models.GetUserAddress(uid)
	//获取用户的订单信息(菜名字，数量，价格)
	orderInf,_ := models.GetUserCloseOrder(uid)
	//获取总价
	all,_ := models.SumCount(uid)
	c.Data["all"] = all
	c.Data["orderInf"] = orderInf
	c.Data["address"] = address
	c.TplName = "confirm_order.html"
}
//获取结算购物车所选商品
func (c *ConfirmOrder) AddShops(){
	resultMap := make(map[string]interface{})
	resultMap["ret"] = 403
	defer func() {
		c.Data["json"] = resultMap
		c.ServeJSON()
	}()
	//获取数据
	//获取用户的id
	var id = c.Ctx.GetCookie("id")
	uid, _ := strconv.Atoi(id)
	//用户购物车所选的商品
	var name = c.GetStrings("name")
	var price = c.GetStrings("price")
	var count = c.GetStrings("count")
	for i := 0;i<len(name);i++{
		err := models.Addshops(uid, count[i], name[i], price[i])
		if err != nil{
			resultMap["msg"] = "结算购物车商品失败！"
			return
		}
	}
	resultMap["ret"] = 200
	resultMap["msg"] = "结算购物车商品成功"
	return
}
//点击提交订单将该用户的订单号存在数据库中
func (c *ConfirmOrder) SubmitOrderAddDatabase(){
	resultMap := make(map[string]interface{})
	resultMap["ret"] = 403
	defer func() {
		c.Data["json"] = resultMap
		c.ServeJSON()
	}()
	//用户的id
	var id = c.Ctx.GetCookie("id")
	var uid, _ = strconv.Atoi(id)
	//订单号
	var orderId = c.GetString("code")
	//订单的总价
	var price =  c.GetString("price")
	err := models.AddOrderTable(uid,orderId,price)
	if err != nil{
		resultMap["msg"] = "添加到订单失败"
	}
	resultMap["ret"] = 200
	resultMap["msg"] = "添加订单成功"
}