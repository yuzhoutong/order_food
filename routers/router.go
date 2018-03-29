package routers

import (
	"order_food/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//---------------------------------------------用户登录注册-----------------------------------------------------------
    beego.Router("/login",&controllers.AccountController{},"get:Login;post:CheckPassword")//登录
    beego.Router("/logout",&controllers.AccountController{},"get:LogOut")//退出登录
    beego.Router("/index_admin",&controllers.HomeController{},"get:IndexAdmin")//管理员页面
	beego.Router("/index_user",&controllers.HomeController{},"get:IndexUser")//用户页面
    beego.Router("/register",&controllers.AccountController{},"post:Register")//注册界面

    //---------------------------------------------对用户的管理-----------------------------------------------------------
    beego.Router("/modifyUser",&controllers.AccountController{},"get:ToModifyUser;post:ModifyUser")//修改密码
    beego.Router("/userInformation",&controllers.AccountController{},"get:UserInformation")//to用户中心

    //----------------------------------------------用户中心操作----------------------------------------------------------
    beego.Router("/userOrderList",&controllers.UserCenterController{},"get:UserOrderList")//to我的订单
    beego.Router("/userAddress",&controllers.UserCenterController{},"get:ShippingAddress;post:AddShippingAddress")//to我的地址:修改用户地址
    beego.Router("/deleteAddress",&controllers.UserCenterController{},"post:DeleteAddressInf")//删除用户地址
    beego.Router("/addAddress",&controllers.UserCenterController{},"post:AddAddress")//新增用户地址
	beego.Router("/userLeaveWords",&controllers.UserCenterController{},"get:UserMessage")//to我的留言
	beego.Router("/userCoupon",&controllers.UserCenterController{},"get:UserCoupon")//to我的优惠券
	beego.Router("/userCollect",&controllers.UserCenterController{},"get:UserCollect")//to我的收藏

	//-----------------------------------------------购物车--------------------------------------------------------------
	beego.Router("/shoppingCart",&controllers.ShopCartCtroller{},"get:INShopCart;post:AddShopCart")//to购物车
	beego.Router("/deleteCart",&controllers.ShopCartCtroller{},"post:DeleteShop")//删除购物车商品

	//------------------------------------------------点餐---------------------------------------------------------------
	beego.Router("/orderFood",&controllers.OrderFoodController{},"get:OrderFoodList;post:OrderFoodListJson")//根据条件展现不同的列

	//------------------------------------------------订单结算-----------------------------------------------------------
	beego.Router("/toOrderConfirm",&controllers.ConfirmOrder{},"get:ToOrderConfirm")//跳转到订单结算页面
	beego.Router("/shopsClose",&controllers.ConfirmOrder{},"post:AddShops")//购物车结算
	beego.Router("/addOrderTable",&controllers.ConfirmOrder{},"post:SubmitOrderAddDatabase")//将订单添加到数据库中

	//------------------------------------------------积分商城-----------------------------------------------------------
	beego.Router("/pointShop",&controllers.PointShopController{},"get:Shop")

	//------------------------------------------------关于我们-----------------------------------------------------------
	beego.Router("/aboutWe",&controllers.AboutController{},"get:About")

}