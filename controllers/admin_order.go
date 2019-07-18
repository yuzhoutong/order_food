package controllers

import (
	"order_food/models"
	"order_food/cache"
)

type OrderManagementController struct {
	HomeController
}
//订单管理
func (c *OrderManagementController) OrderManagement() {
	c.IsNeedTemplate()
	condition := ""
	params := []string{}

	if name := c.GetString("name"); name != "" {
		condition += " and name = ?"
		params = append(params, name)
	}
	list, err := models.GetOrderListAdmin(condition, params)
	if err != nil {
		cache.RecordLogs(c.OrderUser.OrderUsersId, 0, c.OrderUser.Name, c.OrderUser.Displayname, "获取订单管理失败", "订单管理/OrderManagement", err.Error(), c.Ctx.Input)
		return
	}
	c.Data["OrderList"] = list
	c.TplName = "back/order_management.html"
}

//删除订单列表订单
func (c *OrderManagementController) DeleteOrderList() {
	resultMap := make(map[string]interface{})
	resultMap["ret"] = 403
	defer func() {
		c.Data["json"] = resultMap
		c.ServeJSON()
	}()
	//获取OrderId
	OrderId := c.GetString("orderId")
	err := models.DeleteOrderByOrderId(OrderId)
	if err != nil {
		resultMap["msg"] = "删除订单错误！"
		cache.RecordLogs(c.OrderUser.OrderUsersId, 0, c.OrderUser.Name, c.OrderUser.Displayname, "删除订单错误", "删除订单列表订单/DeleteOrderList", err.Error(), c.Ctx.Input)
		return
	}
	resultMap["ret"] = 200
	resultMap["msg"] = "删除成功"
}
