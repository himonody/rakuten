package router

import (
	"rakuten_backend/internal/api/xhttp"
	"rakuten_backend/internal/handler"
)

func setApiRouter() {
	userRouter()
	user := api.Group("/user")
	{
		xhttp.POST(user, "/register", handler.GetUserProductConfigList) //用户注册
		xhttp.POST(user, "/login", handler.GetUserProductConfigList)    //用户登录
		xhttp.GET(user, "/info", handler.GetUserProductConfigList)      //用户信息
		xhttp.PATCH(user, "/:id", handler.GetUserProductConfigList)     //修改用户信息
	}

	activity := api.Group("/activity")
	{
		xhttp.GET(activity, "/list", handler.AdminLogin)
	}

	product := api.Group("/product")
	{
		xhttp.GET(product, "/:id", handler.GetUserProductConfigList)            //获取当前用户活动
		xhttp.GET(product, "/flash-sale/:id", handler.GetUserProductConfigList) //获取当前用户活动
		xhttp.GET(product, "/group-buy/:id", handler.GetUserProductConfigList)  //获取当前用户活动
	}

	orders := api.Group("/order")
	{
		xhttp.POST(orders, "/login", handler.AdminLogin)
	}
	settings := api.Group("/setting")
	{
		xhttp.POST(settings, "/login", handler.AdminLogin)
	}

}
func apiUserRouter() {
	xhttp.POST(api, "/register", handler.GetUserProductConfigList) //用户注册
	xhttp.POST(api, "/login", handler.GetUserProductConfigList)    //用户登录
}
