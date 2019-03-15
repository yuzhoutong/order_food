package controllers

import (
	"order_food/models"
)

type OrderManagementController struct {
	HomeController
}

func (c *OrderManagementController) OrderManagement() {
	c.IsNeedTemplate()
	condition := ""
	params := []string{}

	if name := c.GetString("name"); name != "" {
		condition += " and name = ?"
		params = append(params, name)
	}
	list, _ := models.GetOrderListAdmin(condition, params)
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
		return
	}
	resultMap["ret"] = 200
	resultMap["msg"] = "删除成功"
}
