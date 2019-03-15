package controllers

import (
	"github.com/astaxie/beego"
	"order_food/models"
	"strconv"
)

type OrderFoodController struct {
	beego.Controller
}

func (c *OrderFoodController) OrderFoodList() {
	//获取菜单列表
	condition := ""
	params := []string{}
	//获取菜品列表
	dishList, _ := models.DishList(condition, params)
	//获取菜品种类
	dishkides, _ := models.DishKide()
	//获取用户id
	var id = c.Ctx.GetCookie("id")
	uid, _ := strconv.Atoi(id)
	orderlist, _ := models.GetUserCloseOrder(uid)
	c.Data["diskide"] = dishkides
	c.Data["dishList"] = dishList
	c.Data["orderList"] = orderlist
	//公告
	notice, _ := models.GetNotic()
	c.Data["notice"] = notice
	c.TplName = "list.html"
}

///条件筛选
var count = 2

func (c *OrderFoodController) OrderFoodListJson() {
	defer c.ServeJSON()
	condition := ""
	params := []string{}
	if type1 := c.GetString("type", "全部"); type1 != "" {
		if type1 == "全部" {
			condition += " AND 1=1"
		} else {
			condition += " AND type = ?"
			params = append(params, type1)
			c.Data["type"] = type1
		}

	}
	//获取价格
	if price := c.GetString("price"); price != "" {
		if price == "全部" {
			condition += " AND 1=1"
		} else if price == "20元以下" {
			condition += " AND price <= 20"
		} else if price == "20-40元" {
			condition += " AND price >20 AND price <=40"
		} else if price == "40-60元" {
			condition += " AND price >40 AND price <=60"
		} else if price == "60-80元" {
			condition += " AND price >60 AND price <=80"
		} else if price == "80-100元" {
			condition += " AND price >80 AND price <=100"
		} else if price == "100元以上" {
			condition += " AND price >100"
		}
	}
	//获取筛选条件的属性值
	//销量排序
	var s = c.GetString("s")
	if s == "2" {
		condition += " order by order_count desc"
	} else if s == "4" {
		condition += " order by create_time desc"
	} else if s == "3" {
		if count%2 == 0 {
			condition += " order by price desc"
		} else {
			condition += " order by price "
		}
		count++
	}
	//时间
	if time := c.GetString(""); time != "" {
		condition += ""
		params = append(params, time)
	}
	//获取菜品列表
	dishList, _ := models.DishList(condition, params)
	c.Data["json"] = dishList
}

//删除点餐页面用户所选的food
func (c *OrderFoodController) DeleteUserChooseFood() {
	resultMap := make(map[string]interface{})
	resultMap["ret"] = 403
	defer func() {
		c.Data["json"] = resultMap
		c.ServeJSON()
	}()
	//获取用户id
	var id = c.Ctx.GetCookie("id")
	uid, _ := strconv.Atoi(id)
	//获取id 表add_order_car 的 id
	id1, _ := c.GetInt("id")
	err := models.DelUserChooseFood(uid, id1)
	if err != nil {
		resultMap["msg"] = "删除错误！！"
		return
	}
	resultMap["ret"] = 200
	resultMap["msg"] = "删除成功！！"
	return
}
