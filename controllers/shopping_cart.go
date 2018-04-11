package controllers

import (
	"github.com/astaxie/beego"
	"order_food/models"
	"strconv"
)

type ShopCartCtroller struct {
	beego.Controller
}
//跳转到购物车
func (c *ShopCartCtroller) INShopCart(){

	//用户的uid
	var id = c.Ctx.GetCookie("id")
	uid,_:= strconv.Atoi(id)
	//获取购物车列表
	cartList, _ := models.DishCartList(uid)
	c.Data["cartList"] = cartList
	c.TplName = "cart.html"
}
//将菜加入购物车
func (c *ShopCartCtroller) AddShopCart(){
	resultMap := make(map[string]interface{})
	resultMap["ret"] = 403
	defer func() {
		c.Data["json"] = resultMap
		c.ServeJSON()

	}()
	//用户的uid
	var id = c.Ctx.GetCookie("id")
	uid,_ := strconv.Atoi(id)
	//菜品的dish_id
	var dishId,_  = c.GetInt("dishId")
	err := models.AddCar(uid,dishId)
	if err != nil{
		resultMap["msg"] = "加入购物车失败！"
		return
	}
	resultMap["msg"] = "加入购物车成功!"
	resultMap["ret"] = 200
	return
}
//删除购物车的商品
func (c ShopCartCtroller) DeleteShop(){
	resultMap := make(map[string]interface{})
	resultMap["ret"] = 403
	defer func() {
		c.Data["json"] = resultMap
		c.ServeJSON()

	}()
	//获取用户id
	var id = c.Ctx.GetCookie("id")
	uid,_ := strconv.Atoi(id)
	//购物车中当前选中商品的id
	Id,_ := c.GetInt("Id")
	err := models.DeleteCartShop(uid, Id)
	if err != nil{
		resultMap["msg"] = "删除商品失败！"
		return
	}
	resultMap["ret"] = 200
	resultMap["msg"] = "删除成功"

}


//点击订餐页面的结算将数据存入add_order_car表
func (c ShopCartCtroller) AddOrderDataToAddOrderCar(){
	resultMap := make(map[string]interface{})
	resultMap["ret"] = 403
	defer func() {
		c.Data["json"] = resultMap
		c.ServeJSON()
	}()
	//用户的uid
	var id = c.Ctx.GetCookie("id")
	uid,_ := strconv.Atoi(id)
	//获取菜品信息
	var name = c.GetString("name")
	var price = c.GetString("price")
	var count = c.GetString("count")
	var ids = c.GetString("dishId")
	err := models.Addshops(uid, count, name, price, ids)
	if err != nil {
		resultMap["msg"] = "添加下单信息失败"
	}
	//查询下单信息
	orderInf,_ := models.GetUserCloseOrder(uid)
	//c.Data["json"] = orderInf
	resultMap["ret"] = 200
	resultMap["msg"] = "下单信息添加成功"
	resultMap["data"] = orderInf
}
















