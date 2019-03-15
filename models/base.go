package models

import "zcm_tools/orm"

type Usertable struct {
	Id          int
	Uid         int    //用户id
	OrderId     string //订单id
	Name        string //收件人姓名
	IsProcessed int    //订单是否被处理 1:已处理，2:未处理
	OrderTime   string //订单时间
	OrderPrice  string //订单金额
	IsBuy       int    //是否付款:1:已支付:2未支付
	AddressId   int    //地址id
	IsEvaluate  int    //是否评价1:已评价 2:未评价 0:表示为付款没有评价功能
}

func GetOrderTableINf() (list []Usertable, err error) {
	sql := `SELECT * FROM order_table where is_buy = 1`
	_, err = orm.NewOrm().Raw(sql).QueryRows(&list)
	return
}

//用户评价
func GetAllUserEvaluateList() (EvaList []UserEvaluate, err error) {
	sql := `SELECT * FROM order_evaluate order by create_time desc`
	_, err = orm.NewOrm().Raw(sql).QueryRows(&EvaList)
	return
}

//查询公告
func GetNotic() (notice string, err error) {
	sql := `SELECT context FROM notice where is_user = 1`
	err = orm.NewOrm().Raw(sql).QueryRow(&notice)
	return
}

//获取销售量前五的(轮播)
func GetNamePic() (list []DishCart, err error) {
	sql := `SELECT dish_name,pic_path FROM dish_table ORDER BY order_count DESC LIMIT 0,5`
	_, err = orm.NewOrm().Raw(sql).QueryRows(&list)
	return
}
