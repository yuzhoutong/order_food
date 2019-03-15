package models

import (
	"zcm_tools/orm"
)

type DishTable struct {
	DishId     int     //菜品id
	Detail     string  //菜品描述
	DishName   string  //菜品名称
	PicPath    string  //菜品图片的路径
	Price      float64 //菜的价格
	Type       string  //菜品的种类
	OrderCount int     //菜品下单次数
	CreateTime string  //创建时间
}

//菜品列表
func GetDishList(condition string, params []string) (list []DishTable, err error) {
	sql := `select * from dish_table where 1=1 order by create_time desc`
	if condition != "" {
		sql += condition
	}
	_, err = orm.NewOrm().Raw(sql, params).QueryRows(&list)
	return
}

//添加菜品
func AddDishDate(dishname, detail, picpath, type1 string, price float64) (err error) {
	sql := `insert into dish_table(dish_name, detail, pic_path, price, TYPE, order_count, create_time)value(?, ?, ?, ?, ?, 0, now())`
	_, err = orm.NewOrm().Raw(sql, dishname, detail, picpath, price, type1).Exec()
	return
}

//删除菜品
func DeleteDishByDishId(dishId int) (err error) {
	sql := `delete from dish_table where dish_id = ?`
	_, err = orm.NewOrm().Raw(sql, dishId).Exec()
	return
}

//根据添加的菜名查 (不可上传重复菜名）
func GetInformationByDishName(dishname string) (list *DishTable, err error) {
	sql := `select * from dish_table where dish_name = ?`
	err = orm.NewOrm().Raw(sql, dishname).QueryRow(&list)
	return
}
