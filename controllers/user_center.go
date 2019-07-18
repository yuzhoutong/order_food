package controllers

import (
	"order_food/models"
	"strconv"
	"order_food/cache"
)

type UserCenterController struct {
	HomeController
	User *models.UserCenter
}

//用户中心-我的订单
func (c *UserCenterController) UserOrderList() {
	id := c.Ctx.GetCookie("id")
	uid, _ := strconv.Atoi(id)
	orderList, _ := models.GetUserOrder(uid)
	//公告
	notice, _ := models.GetNotic()
	c.Data["notice"] = notice
	c.Data["orderList"] = orderList
	c.TplName = "user_orderlist.html"
}

//取消订单
func (c *UserCenterController) DeleteOrder() {
	resultMap := make(map[string]interface{})
	resultMap["ret"] = 403
	defer func() {
		c.Data["json"] = resultMap
		c.ServeJSON()
	}()
	//获取参数订单id
	var orderId = c.GetString("orderId")
	//获取用户uid
	uid := c.OrderUser.OrderUsersId
	err := models.DeleteOrder(uid, orderId)
	if err != nil {
		resultMap["msg"] = "删除订单失败！"
		return
	}
	resultMap["ret"] = 200
	resultMap["msg"] = "删除订单成功！！"
}

//点击订单号显示订单详情

type OrderList struct {
	Address     *models.UserAddress
	OrderDetail []models.OrderDetail
	Count       int
}

//查询用户的订单详情
func (c *UserCenterController) GetOrderDetail() {
	var resp OrderList
	defer func() {
		c.Data["json"] = resp
		c.ServeJSON()
	}()
	//订单号
	orderId := c.GetString("orderId")
	uid := c.OrderUser.OrderUsersId
	//获取地址id
	addressId, _ := models.GetAddressIdByOrderId(orderId)
	//根据地址id获地址信息
	address, _ := models.GetAddressByAddressId(addressId)
	OrderDetail, _ := models.GetOrderImanages(uid, orderId)
	//获取付款情况
	i, _ := models.IsBuyByUIdAndOrderId(uid, orderId)
	resp.Count = i
	resp.Address = address
	resp.OrderDetail = OrderDetail
}

//用户中心-收货地址
func (c *UserCenterController) ShippingAddress() {
	id := c.Ctx.GetCookie("id")
	id1, _ := strconv.Atoi(id)
	address, _ := models.GetUserAddress(id1)
	c.Data["address"] = address
	//公告
	notice, _ := models.GetNotic()
	c.Data["notice"] = notice
	c.TplName = "user_address.html"

}

//修改用户收货地址信息
func (c *UserCenterController) AddShippingAddress() {
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
	user, _ := models.GetUserByName(name)
	var uid = user.OrderUsersId
	id, _ := c.GetInt("id")
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

//修改用户的地址
func (c *UserCenterController) UpdateAddress() {
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
	//详细地址
	var address = c.GetString("address")
	//收件人名字
	var rename = c.GetString("name")
	//电话号码
	var phone = c.GetString("phone")
	//用户的id
	uid := c.OrderUser.OrderUsersId
	id, _ := c.GetInt("id")
	//修改地址表
	err := models.UpdateAddress(uid, id, address, rename, phone, provice, city, district)
	if err != nil {
		resultMap["msg"] = "修改用户信息失败"
		return
	}
	//修改订单表姓名
	err1 := models.UpdateOrderNameByRname(rename, uid, id)
	if err1 != nil {
		resultMap["msg"] = "修改用户订单信息失败"
		return
	}
	resultMap["ret"] = 200
	resultMap["msg"] = "修改用户信息成功"
	return
}

//删除用户地址信息
func (c *UserCenterController) DeleteAddressInf() {
	resultMap := make(map[string]interface{})
	resultMap["ret"] = 403
	defer func() {
		c.Data["json"] = resultMap
		c.ServeJSON()
	}()
	//获取用户uid与id
	uid := c.OrderUser.OrderUsersId
	var id, _ = c.GetInt("id")
	err := models.DelAddress(uid, id)
	if err != nil {
		resultMap["msg"] = "删除地址失败！"
		return
	}
	resultMap["ret"] = 200
	resultMap["msg"] = "删除成功！"
	return
}

//新增用户收货地址
func (c *UserCenterController) AddAddress() {
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
	uid := c.OrderUser.OrderUsersId
	err := models.AddAddress(uid, address, rename, phone, provice, city, district)
	if err != nil {
		resultMap["msg"] = "新增收货地址失败！"
		return
	}
	resultMap["ret"] = 200
	resultMap["msg"] = "新增地址成功！"
}

//用户中心-我的留言
func (c *UserCenterController) UserMessage() {
	//获取订单号
	uid := c.OrderUser.OrderUsersId
	OrderId := c.GetString("orderId")
	c.Data["orderId"] = OrderId
	evaList, _ := models.GetUserEvaluate(uid)
	c.Data["evaList"] = evaList
	//公告
	notice, _ := models.GetNotic()
	c.Data["notice"] = notice
	c.TplName = "user_message.html"
}

//用户提交评价
func (c *UserCenterController) UserSubmitEvaluate() {
	resultMap := make(map[string]interface{})
	resultMap["ret"] = 403
	defer func() {
		c.Data["json"] = resultMap
		c.ServeJSON()
	}()
	//获取参数
	uid := c.OrderUser.OrderUsersId
	name, _ := models.GetUsernameByUid(uid)
	var orderId = c.GetString("orderId")
	var remark = c.GetString("inputText")
	var taste, _ = c.GetInt("taste")
	var manner, _ = c.GetInt("manner")
	var speed, _ = c.GetInt("speed")
	err := models.AddUserEvaluate(orderId, remark, uid, speed, manner, taste, name)
	if err != nil {
		resultMap["msg"] = "插入数据错误"
		return
	}
	//用户提交评价修改用户订单中订单评价的状态
	err1 := models.UpdateUserEvaluateStatus(orderId)
	if err1 != nil {
		resultMap["msg"] = "修改评价状态失败"
		return
	}
	resultMap["ret"] = 200
	resultMap["msg"] = "添加用户评论成功！"
}

//删除用户评价
func (c *UserCenterController) DelUserevaluate() {
	resultMap := make(map[string]interface{})
	resultMap["ret"] = 403
	defer func() {
		c.Data["json"] = resultMap
		c.ServeJSON()
	}()
	//获取参数
	var orderId = c.GetString("orderId")
	var time = c.GetString("time")
	err := models.DeleteUserEvaluate(orderId, time)
	if err != nil {
		resultMap["msg"] = "删除用户评价失败"
		cache.RecordLogs(c.OrderUser.OrderUsersId, 0, c.OrderUser.Name, c.OrderUser.Displayname, "删除错误", "删除点餐页面用户所选的food/DeleteUserChooseFood", err.Error(), c.Ctx.Input)
		return
	}
	resultMap["ret"] = 200
	resultMap["msg"] = "删除评价成功！！"
}
