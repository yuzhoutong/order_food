package controllers

import (
	"order_food/models"
	"order_food/utils"
	"time"
)

type AdminDishController struct {
	HomeController
}

//菜品管理
func (c *AdminDishController) FoodManagement() {
	c.IsNeedTemplate()
	//获取菜品列表
	condition := ""
	params := []string{}
	if name := c.GetString("name"); name != "" { //菜名
		condition = " and dish_name like ?"
		name = "%" + name + "%"
		params = append(params, name)
	}
	FoodList, _ := models.GetDishList(condition, params)
	c.Data["FoodList"] = FoodList
	c.TplName = "back/Food_management.html"
}

//添加菜品上传图片
func (c *AdminDishController) UpdateImg() {
	resultMap := make(map[string]interface{})
	resultMap["ret"] = 403
	defer func() {
		c.Data["json"] = resultMap
		c.ServeJSON()
	}()
	//获取菜品信息
	dish := c.GetString("dishname")
	detail := c.GetString("detail")
	price, _ := c.GetFloat("price")
	types := c.GetString("type")
	s, f, err := c.GetFile("images")
	//上传图片不能为空
	if err != nil {
		resultMap["msg"] = err.Error()
		return
	}
	//以时间作为图片的名字
	id := int(time.Now().Unix())
	name, err := utils.InputImages(s, f, id)
	if name == "" && err != nil {
		resultMap["msg"] = "图片上传失败" + err.Error()
		return
	}
	//图片的路径
	pic_path := "../static/tupian/" + name
	//根据用户名查重
	//重复的菜品不可上传
	lis, _ := models.GetInformationByDishName(dish)
	if lis != nil {
		resultMap["msg"] = "该菜名已经存在！"
		return
	}
	//上传菜品信息
	err1 := models.AddDishDate(dish, detail, pic_path, types, price)
	if err1 != nil {
		resultMap["msg"] = "插入数据库数据失败！"
		return
	}
	resultMap["ret"] = 200
	resultMap["msg"] = "添加数据成功"

}

//删除菜品
func (c *AdminDishController) DeleteDish() {
	resultMap := make(map[string]interface{})
	resultMap["ret"] = 403
	defer func() {
		c.Data["json"] = resultMap
		c.ServeJSON()
	}()
	DishId, _ := c.GetInt("id")
	err := models.DeleteDishByDishId(DishId)
	if err != nil {
		resultMap["msg"] = "删除菜品错误"
		return
	}
	resultMap["ret"] = 200
	resultMap["msg"] = "删除成功"
}
