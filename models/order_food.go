package models

import (
	"zcm_tools/orm"
)

//菜品
type DishFood struct {
	DishId		int        	//菜品的ID
	Detail		string     	//菜品描述
	DishName	string		//菜品名称
	OrderCount	int			//被订次数
	PicPath		string		//菜品图片的路径
	Price		float64		//菜的价格
	Type		string		//菜品的种类
}
type List struct {
	Type  string
}
//获取菜品列表
func DishList(condition string, params []string)(list []DishFood, err error){
	sql := `SELECT * FROM dish_table WHERE 1=1 `
	if condition != ""{
		sql += condition
	}
	_,err = orm.NewOrm().Raw(sql,params).QueryRows(&list)
	return
}
//查询所有菜品种类
func DishKide()(list []List, err error){
	sql := `SELECT DISTINCT type FROM dish_table`
	_,err = orm.NewOrm().Raw(sql).QueryRows(&list)
	return
}
//删除用户所选事物数据
func DelUserChooseFood(uid, id int )(err error){
	sql := `DELETE FROM add_order_car WHERE uid = ? AND id = ?`
	_,err = orm.NewOrm().Raw(sql ,uid, id).Exec()
	return
}