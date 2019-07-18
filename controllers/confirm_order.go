package controllers

import (
	"order_food/models"
	"strconv"
	"order_food/cache"
)

type ConfirmOrder struct {
	HomeController
}

//结算购物车页面渲染
func (c *ConfirmOrder) ToOrderConfirm() {
	//用户的uid
	uid := c.OrderUser.OrderUsersId
	//获取用户的收货地址信息
	address, err := models.GetUserAddress(uid)
	if err != nil {
		cache.RecordLogs(c.OrderUser.OrderUsersId, 0, c.OrderUser.Name, c.OrderUser.Displayname, "获取用户的收货地址信息失败", "结算购物车页面渲染/AddShops", err.Error(), c.Ctx.Input)
		return
	}
	//获取用户的订单信息(菜名字，数量，价格)
	orderInf, err := models.GetUserCloseOrder(uid)
	if err != nil {
		cache.RecordLogs(c.OrderUser.OrderUsersId, 0, c.OrderUser.Name, c.OrderUser.Displayname, "获取用户的订单信息失败", "结算购物车页面渲染/AddShops", err.Error(), c.Ctx.Input)
		return
	}
	//获取总价
	all, err  := models.SumCount(uid)
	if err != nil {
		cache.RecordLogs(c.OrderUser.OrderUsersId, 0, c.OrderUser.Name, c.OrderUser.Displayname, "获取总价失败", "结算购物车页面渲染/AddShops", err.Error(), c.Ctx.Input)
		return
	}
	//公告
	notice, _ := models.GetNotic()
	c.Data["notice"] = notice
	c.Data["all"] = all
	c.Data["orderInf"] = orderInf
	c.Data["address"] = address
	c.TplName = "confirm_order.html"
}

//获取结算购物车所选商品
func (c *ConfirmOrder) AddShops() {
	resultMap := make(map[string]interface{})
	resultMap["ret"] = 403
	defer func() {
		c.Data["json"] = resultMap
		c.ServeJSON()
	}()
	//获取数据
	//获取用户的id
	uid := c.OrderUser.OrderUsersId
	//用户购物车所选的商品
	var name = c.GetStrings("name")
	var price = c.GetStrings("price")
	var count = c.GetStrings("count")
	var ids = c.GetStrings("ids")
	for i := 0; i < len(name); i++ {
		err := models.Addshops(uid, count[i], name[i], price[i], ids[i])
		if err != nil {
			resultMap["msg"] = "结算购物车商品失败！"
			cache.RecordLogs(c.OrderUser.OrderUsersId, 0, c.OrderUser.Name, c.OrderUser.Displayname, "结算购物车商品失败", "获取结算购物车所选商品/AddShops", err.Error(), c.Ctx.Input)
			return
		}
	}
	resultMap["ret"] = 200
	resultMap["msg"] = "结算购物车商品成功"
	return
}

//点击提交订单将该用户的订单号存在数据库中 删除add_order_car数据
func (c *ConfirmOrder) SubmitOrderAddDatabase() {
	resultMap := make(map[string]interface{})
	resultMap["ret"] = 403
	defer func() {
		c.Data["json"] = resultMap
		c.ServeJSON()
	}()
	//用户的id
	uid := c.OrderUser.OrderUsersId
	//订单号
	var orderId = c.GetString("code")
	//订单的总价
	var price = c.GetString("price")
	//获取收件人姓名
	var name = c.GetString("name")
	//获取菜名
	var dishName = c.GetStrings("dishName")
	//获取菜品数量
	var dishCount = c.GetStrings("dishCount")
	//获取菜品单价
	var dishPrice = c.GetStrings("dishPrice")
	//获取地址的id
	var addressId, _ = c.GetInt("addressId")
	for i := 0; i < len(dishName); i++ {
		//添加到订单详情表
		err := models.AddOrderDetail(uid, orderId, dishName[i], dishCount[i], dishPrice[i])
		if err != nil {
			resultMap["msg"] = "添加数据失败！"
			cache.RecordLogs(c.OrderUser.OrderUsersId, 0, c.OrderUser.Name, c.OrderUser.Displayname, "添加数据失败", "点击提交订单将该用户的订单号存在数据库中/SubmitOrderAddDatabase", err.Error(), c.Ctx.Input)
			return
		}
		p, _ := strconv.Atoi(dishCount[i])
		err1 := models.UpdateCountClick(dishName[i], p)
		if err1 != nil {
			resultMap["msg"] = "添加数量失败！！"
			cache.RecordLogs(c.OrderUser.OrderUsersId, 0, c.OrderUser.Name, c.OrderUser.Displayname, "添加数量失败", "点击提交订单将该用户的订单号存在数据库中/SubmitOrderAddDatabase", err.Error(), c.Ctx.Input)
			return
		}
	}
	//添加到订单表
	err := models.AddOrderTable(uid, addressId, orderId, price, name)
	if err != nil {
		resultMap["msg"] = "添加到订单失败！"
		cache.RecordLogs(c.OrderUser.OrderUsersId, 0, c.OrderUser.Name, c.OrderUser.Displayname, "添加到订单失败", "点击提交订单将该用户的订单号存在数据库中/SubmitOrderAddDatabase", err.Error(), c.Ctx.Input)
		return
	}
	//add_order_car表中的ids
	var ids = c.GetStrings("ids")
	//删除数据
	for i := 0; i < len(ids); i++ {
		idsInt, _ := strconv.Atoi(ids[i])
		//err := models.DelShops(uid, idsInt)
		err1 := models.DELAddOrderCar(uid, idsInt)
		if err != nil || err1 != nil {
			resultMap["msg"] = "删除数据库数据异常！"
			cache.RecordLogs(c.OrderUser.OrderUsersId, 0, c.OrderUser.Name, c.OrderUser.Displayname, "删除数据库数据异常", "点击提交订单将该用户的订单号存在数据库中/SubmitOrderAddDatabase", err.Error(), c.Ctx.Input)
			return
		}
	}
	resultMap["ret"] = 200
	resultMap["msg"] = "添加订单成功!"
}

//当单击付款时(支付成功)删除数据库中所选商品在购物车car数据
//修改数据库中订单表中is_buy的状态2 - 1
func (c *ConfirmOrder) DeleteshopsFromDatabase() {
	resultMap := make(map[string]interface{})
	resultMap["ret"] = 403
	defer func() {
		c.Data["json"] = resultMap
		c.ServeJSON()
	}()
	//用户的uid
	uid := c.OrderUser.OrderUsersId
	//获取订单id
	var orderId = c.GetString("orderId")
	//add_order_car表中的ids
	var ids = c.GetStrings("ids")
	//删除数据
	for i := 0; i < len(ids); i++ {
		idsInt, _ := strconv.Atoi(ids[i])
		err := models.DelShops(uid, idsInt)
		//err1 := models.DELAddOrderCar(uid, idsInt)
		if err != nil {
			resultMap["msg"] = "删除数据库数据异常！"
			cache.RecordLogs(c.OrderUser.OrderUsersId, 0, c.OrderUser.Name, c.OrderUser.Displayname, "删除数据库数据异常", "当单击付款时(支付成功)/DeleteshopsFromDatabase", err.Error(), c.Ctx.Input)
			return
		}
	}
	//修改order_table 订单状态
	err := models.UpdateIsBuy(uid, orderId)
	if err != nil {
		resultMap["msg"] = "修改用户的购买状态失败！"
		cache.RecordLogs(c.OrderUser.OrderUsersId, 0, c.OrderUser.Name, c.OrderUser.Displayname, "修改用户的购买状态失败", "当单击付款时(支付成功)/DeleteshopsFromDatabase", err.Error(), c.Ctx.Input)
		return
	}
	resultMap["ret"] = 200
	resultMap["msg"] = "删除成功！！"
}

//点击未付款修改购买状态
func (c *ConfirmOrder) UpdateIsBuy() {
	resultMap := make(map[string]interface{})
	resultMap["ret"] = 403
	defer func() {
		c.Data["json"] = resultMap
		c.ServeJSON()
	}()
	//用户的uid
	uid := c.OrderUser.OrderUsersId
	//获取订单id
	var orderId = c.GetString("orderId")
	err := models.UpdateIsBuy(uid, orderId)
	if err != nil {
		resultMap["msg"] = "修改用户的购买状态失败！"
		cache.RecordLogs(c.OrderUser.OrderUsersId, 0, c.OrderUser.Name, c.OrderUser.Displayname, "修改用户的购买状态失败", "点击未付款修改购买状态(支付成功)/UpdateIsBuy", err.Error(), c.Ctx.Input)
		return
	}
	resultMap["ret"] = 200
	resultMap["msg"] = "付款成功！！"
}
