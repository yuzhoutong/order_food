package models

import (
	"zcm_tools/orm"
)

type DishCart struct {
	Id       int
	Uid      int     //用户id
	DishId   int     //菜品id
	DishName string  //菜品名称
	PicPath  string  //菜品图片的路径
	Price    float64 //菜的价格
	Type     string  //菜品的种类

}

//加入购物车
func AddCar(uid, DishId int) (err error) {
	sql := `INSERT INTO my_car(uid,dish_id,create_time) values(?,?,now())`
	_, err = orm.NewOrm().Raw(sql, uid, DishId).Exec()
	return
}

//统计购物车菜品数量
func DishCartCount(uid int) (count int, err error) {
	sql := ` SELECT COUNT(*) FROM my_car WHERE uid = ?`
	err = orm.NewOrm().Raw(sql, uid).QueryRow(&count)
	return
}

//获取用户菜品列表
func DishCartList(uid int) (list []DishCart, err error) {
	sql := `SELECT * FROM my_car mc
			LEFT JOIN dish_table dt
			ON mc.dish_id = dt.dish_id
			WHERE uid = ? `
	_, err = orm.NewOrm().Raw(sql, uid).QueryRows(&list)
	return
}

//删除购物车商品
func DeleteCartShop(uid, Id int) (err error) {
	sql := `DELETE FROM my_car WHERE uid = ? AND id = ?`
	_, err = orm.NewOrm().Raw(sql, uid, Id).Exec()
	return
}

/*
//点击订餐页面的结算将数据存入add_order_car表
func AddToAddOrderCar(uid int, ids, name, price, count string) (err error){
	sql := `INSERT`
}*/
