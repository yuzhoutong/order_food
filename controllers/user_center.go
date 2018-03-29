package controllers

import (
	"github.com/astaxie/beego"
	"order_food/models"
	"strconv"
	"fmt"
)

type UserCenterController struct {
	beego.Controller
	User *models.UserCenter
}
//用户中心-我的订单
func (c *UserCenterController) UserOrderList(){
	c.TplName = "user_orderlist.html"
}

//用户中心-收货地址
func (c *UserCenterController) ShippingAddress(){
	id := c.Ctx.GetCookie("id")
	id1,_:= strconv.Atoi(id)
	address,_ := models.GetUserAddress(id1)
	c.Data["address"] = address
	fmt.Println(address)
	c.TplName = "user_address.html"

}
//修改用户收货地址信息
func (c *UserCenterController) AddShippingAddress(){
	resultMap := make(map[string]interface{})
	resultMap["ret"] = 403
	defer func() {
		c.Data["json"] = resultMap
		c.ServeJSON()
	}()
	//获取值
	//省
	var provice = c.GetString("provice")
	//市
	var city = c.GetString("city")
	//区
	var district = c.GetString("district")
	//登录用户的名字
	var name = c.GetString("name")
	//详细地址
	var address = c.GetString("detail_address")
	//收件人名字
	var rename = c.GetString("rename")
	//电话号码
	var phone = c.GetString("phone")
	//查询用户详细信息
	user,_ := models.GetUserByName(name)
	var uid = user.OrderUsersId
	 id,_:= c.GetInt("id")
	err := models.UpdateAddress(uid, id, address, rename, phone, provice, city, district)
	if err != nil {
		resultMap["msg"] = "修改用户信息失败"
		return
	}
	resultMap["ret"] = 200
	resultMap["user"] = user
	resultMap["msg"] = "修改用户信息成功"
	return
}
//删除用户地址信息
func (c *UserCenterController) DeleteAddressInf(){
	resultMap := make(map[string]interface{})
	resultMap["ret"] = 403
	defer func() {
		c.Data["json"] = resultMap
		c.ServeJSON()
	}()
	//获取用户uid与id
	var id1 = c.Ctx.GetCookie("id")
	var uid, _ = strconv.Atoi(id1)
	var id,_ = c.GetInt("id")
	err := models.DelAddress(uid,id)
	if err != nil{
		resultMap["msg"] = "删除地址失败！"
		return
	}
	resultMap["ret"] = 200
	resultMap["msg"] = "删除成功！"
	return
}
//新增用户收货地址
func (c *UserCenterController) AddAddress(){
	resultMap := make(map[string]interface{})
	resultMap["ret"] = 403
	defer func() {
		c.Data["json"] = resultMap
		c.ServeJSON()
	}()
	//获取值
	//省
	var provice = c.GetString("province")
	//市
	var city = c.GetString("city")
	//区
	var district = c.GetString("district")
	//详细地址
	var address = c.GetString("address")
	//收件人名字
	var rename = c.GetString("name")
	//电话号码
	var phone = c.GetString("phone")
	//获取用户uid
	var id = c.Ctx.GetCookie("id")
	uid, _ := strconv.Atoi(id)
	err := models.AddAddress(uid,address,rename,phone,provice,city,district)
	if err != nil{
		resultMap["msg"] = "新增收货地址失败！"
		return
	}
	resultMap["ret"] = 200
	resultMap["msg"] = "新增地址成功！"
}

//用户中心-我的留言
func (c *UserCenterController) UserMessage(){
	c.TplName = "user_message.html"
}

//用户中心-我的优惠卷
func (c *UserCenterController) UserCoupon(){
	c.TplName = "user_coupon.html"
}
//用户中心-我的收藏
func (c *UserCenterController) UserCollect(){
	c.TplName = "user_favorites.html"
}
