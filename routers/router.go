package routers

import (
	"github.com/astaxie/beego"
	"order_food/controllers"
)

func init() {
	//---------------------------------------------用户登录注册-----------------------------------------------------------
	beego.Router("/login", &controllers.AccountController{}, "get:Login;post:CheckPassword") //登录
	beego.Router("/logout", &controllers.AccountController{}, "get:LogOut")                  //退出登录
	beego.Router("/index_user", &controllers.HomeController{}, "get:IndexUser")              //用户页面
	beego.Router("/register", &controllers.AccountController{}, "post:Register")             //注册界面

	//---------------------------------------------对用户的管理-----------------------------------------------------------
	beego.Router("/modifyUser", &controllers.AccountController{}, "get:ToModifyUser;post:ModifyUser") //修改密码
	beego.Router("/userInformation", &controllers.AccountController{}, "get:UserInformation")         //to用户中心

	//----------------------------------------------用户中心操作----------------------------------------------------------
	beego.Router("/userOrderList", &controllers.UserCenterController{}, "get:UserOrderList")    //to我的订单
	beego.Router("/deleteOrder", &controllers.UserCenterController{}, "post:DeleteOrder")       //取消订单
	beego.Router("/getDetailOrder", &controllers.UserCenterController{}, "post:GetOrderDetail") //点击订单号获取订单详细信息
	//-----------------------------------------------用户中心 修改地址--------------------------------------------------------------
	beego.Router("/userAddress", &controllers.UserCenterController{}, "get:ShippingAddress;post:AddShippingAddress") //to我的地址:修改用户地址
	beego.Router("/updateAddress", &controllers.UserCenterController{}, "post:UpdateAddress")                        //修改用户地址(1)
	beego.Router("/deleteAddress", &controllers.UserCenterController{}, "post:DeleteAddressInf")                     //删除用户地址
	beego.Router("/addAddress", &controllers.UserCenterController{}, "post:AddAddress")                              //新增用户地址
	//-----------------------------------------------用户中心 用户评价--------------------------------------------------------------
	beego.Router("/userLeaveWords", &controllers.UserCenterController{}, "get:UserMessage")          //to我的留言
	beego.Router("/addUserEvaluate", &controllers.UserCenterController{}, "post:UserSubmitEvaluate") //添加留言
	beego.Router("/delUserEvaluate", &controllers.UserCenterController{}, "post:DelUserevaluate")    //删除用户评价

	//-----------------------------------------------购物车 下单--------------------------------------------------------------
	beego.Router("/shoppingCart", &controllers.ShopCartCtroller{}, "get:INShopCart;post:AddShopCart") //to购物车
	beego.Router("/deleteCart", &controllers.ShopCartCtroller{}, "post:DeleteShop")                   //删除购物车商品
	beego.Router("/placeAnOrder", &controllers.ShopCartCtroller{}, "post:AddOrderDataToAddOrderCar")  //订餐页面点击下单加入数据库中

	//------------------------------------------------点餐---------------------------------------------------------------
	beego.Router("/orderFood", &controllers.OrderFoodController{}, "get:OrderFoodList;post:OrderFoodListJson") //根据条件展现不同的列
	beego.Router("/deleteUserOrderFood", &controllers.OrderFoodController{}, "post:DeleteUserChooseFood")      //删除点餐页面用户所选的food(点击X)
	//------------------------------------------------订单结算-----------------------------------------------------------
	beego.Router("/toOrderConfirm", &controllers.ConfirmOrder{}, "get:ToOrderConfirm")            //跳转到订单结算页面
	beego.Router("/shopsClose", &controllers.ConfirmOrder{}, "post:AddShops")                     //购物车结算
	beego.Router("/addOrderTable", &controllers.ConfirmOrder{}, "post:SubmitOrderAddDatabase")    //将订单添加到数据库中
	beego.Router("/deleteShopTable", &controllers.ConfirmOrder{}, "post:DeleteshopsFromDatabase") //当点击支付宝支付时将购物车表和暂存表删除
	beego.Router("/updateIsBuy", &controllers.ConfirmOrder{}, "post:UpdateIsBuy")                 //点击未付款修改购买状态

	//------------------------------------------------关于我们-----------------------------------------------------------
	beego.Router("/aboutWe", &controllers.AboutController{}, "get:About")

	//--------------------------------------------------------------管理员后台管理----------------------------------------------------------
	beego.Router("/admin_index", &controllers.HomeController{}, "get:IndexAdmin")                    //管理员首页页面 -用户管理
	beego.Router("/delete_user", &controllers.AdminUserController{}, "post:DeleteUserInfo")          //删除用户
	beego.Router("/update_user_status", &controllers.AdminUserController{}, "post:UpdateUserStatus") //修改用户状态

	//-----------------------------------------------菜品-----------------------------------------------------------------
	beego.Router("/admin_footmanagement", &controllers.AdminDishController{}, "get:FoodManagement") //管理员菜品管理
	beego.Router("/updateimg", &controllers.AdminDishController{}, "post:UpdateImg")                //添加菜品上传图片
	beego.Router("/deletedish", &controllers.AdminDishController{}, "post:DeleteDish")              //删除菜品

	//-----------------------------------------------------留言-----------------------------------------------------------
	beego.Router("/admin_imagesmanagement", &controllers.AdminImagesController{}, "get:ImagesManagement") //管理员信息管理 -留言管理
	beego.Router("/delete_info", &controllers.AdminImagesController{}, "post:DeleteInfo")                 //管理员信息管理 -删除留言管理

	//------------------------------------------------------------------公告------------------------------------------------------
	beego.Router("/admin_infotmanagement", &controllers.AdminImagesController{}, "get:InformManagement") //管理员信息管理 -公告管理
	beego.Router("/delete_notice", &controllers.AdminImagesController{}, "post:DeleteNotice")            //管理员信息管理 -删除公告
	beego.Router("/add_notice", &controllers.AdminImagesController{}, "post:AddNotice")                  //管理员信息管理 -添加公告
	beego.Router("/update_notice", &controllers.AdminImagesController{}, "post:UpdateNotice")            //管理员信息管理 -修改公告状态

	//-----------------------------------------------------------管理___订单管理----------------------------------------------------
	beego.Router("/admin_ordermanagement", &controllers.OrderManagementController{}, "get:OrderManagement") //管理员订单管理
	beego.Router("/delete_orderList", &controllers.OrderManagementController{}, "post:DeleteOrderList")     //管理员删除订单

	//-----------------------------------------------------------管理___用户管理----------------------------------------------------
	beego.Router("/admin_usermanagement", &controllers.AdminUserController{}, "get:UserManagement") //用户登录记录
	beego.Router("/adminSet", &controllers.AdminUserController{}, "get:SetAdmin")                   //用户登录记录
	beego.Router("/update_admin", &controllers.AdminUserController{}, "post:UpdateAdmin")           //用户登录记录

}
