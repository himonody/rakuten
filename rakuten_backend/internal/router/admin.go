package router

import (
	"rakuten_backend/internal/api/xhttp"
	"rakuten_backend/internal/handler"
)

func setAdminRouter() {

	//仪表盘
	dashboardRouter()
	//数据统计
	statisticsRouter()
	//管理员管理
	userRouter()
	//会员管理
	memberRouter()
	//代理管理
	agentRouter()
	//交易管理
	transactionRouter()
	//订单管理
	// orderRouter()
	//消息管理
	messageRouter()
	//配置管理
	settingRouter()

}
func dashboardRouter() {
	dashboard := admin.Group("/dashboard")
	{
		xhttp.GET(dashboard, "/statistics", handler.GetDashboard)
	}
}
func statisticsRouter() {
	statistics := admin.Group("/statistics")
	{
		xhttp.GET(statistics, "/daily", handler.GetDailyStatistics)
		xhttp.GET(statistics, "/weekly", handler.GetWeeklyStatistics)
		xhttp.GET(statistics, "/monthly", handler.GetMonthlyStatistics)

	}
}

func userRouter() {

	xhttp.POST(admin, "/login", handler.AdminLogin)

	user := admin.Group("/user")
	{
		xhttp.GET(user, "/list", handler.AdminList)
		xhttp.POST(user, "/add", handler.CreateAdmin)
		xhttp.PATCH(user, "/:id", handler.EditAdminUser)
		xhttp.DELETE(user, "/:id", handler.DeleteAdminUser)
	}

	userOperationLog := admin.Group("/user-operation-log")
	{
		xhttp.GET(userOperationLog, "/list", handler.AdminUserOperationLogList)
	}
}
func agentRouter() {
	agent := admin.Group("/agent")
	{
		xhttp.GET(agent, "/list", handler.AgentList)
		xhttp.POST(agent, "/add", handler.CreateAgent)
		xhttp.PATCH(agent, "/:id", handler.EditAdminUser)
		xhttp.DELETE(agent, "/:id", handler.DeleteAdminUser)
	}

	agentOperationLog := admin.Group("/agent-operation-log")
	{
		xhttp.GET(agentOperationLog, "/list", handler.AdminAgentOperationLogList)
	}
}
func memberRouter() {
	member := admin.Group("/member")
	{
		xhttp.GET(member, "/list", handler.AdminLogin)
		xhttp.POST(member, "/add", handler.AdminLogin)
	}
	memberLoginLog := admin.Group("/member-login-log")
	{
		xhttp.GET(memberLoginLog, "/list", handler.CreateAdmin)
		xhttp.POST(memberLoginLog, "/edit", handler.CreateAdmin)
		xhttp.POST(memberLoginLog, "/delete", handler.CreateAdmin)
	}
	memberOperationLog := admin.Group("/member-operation-log")
	{
		xhttp.GET(memberOperationLog, "/list", handler.CreateAdmin)
		xhttp.POST(memberOperationLog, "/edit", handler.CreateAdmin)
		xhttp.POST(memberOperationLog, "/delete", handler.CreateAdmin)
	}
}

func transactionRouter() {
	transaction := admin.Group("/transaction")
	topUp := transaction.Group("/top-up")
	{
		xhttp.GET(topUp, "/list", handler.CreateAdmin)
	}
	payOut := transaction.Group("/payout")
	{
		xhttp.GET(payOut, "/list", handler.CreateAdmin)
	}
	accountChange := transaction.Group("/account-change")
	{
		xhttp.GET(accountChange, "/list", handler.CreateAdmin)
	}

}
func orderRouter() {

	order := admin.Group("/order")
	snapUpEvent := order.Group("/snap-up-event")
	{
		xhttp.GET(snapUpEvent, "/list", handler.AdminLogin)
		xhttp.POST(snapUpEvent, "/add", handler.AdminLogin)
		xhttp.PATCH(snapUpEvent, "/:id", handler.EditAdminUser)
		xhttp.GET(snapUpEvent, "/details", handler.AdminLogin)
		xhttp.PATCH(snapUpEvent, "/:id", handler.EditAdminUser)
	}
	groupBuyEvent := order.Group("/group-buy-event")
	{
		xhttp.GET(groupBuyEvent, "/list", handler.AdminLogin)
		xhttp.POST(groupBuyEvent, "/add", handler.AdminLogin)
		xhttp.PATCH(groupBuyEvent, "/:id", handler.EditAdminUser)
		xhttp.GET(groupBuyEvent, "/details", handler.AdminLogin)
		xhttp.PATCH(groupBuyEvent, "/:id", handler.EditAdminUser)

	}
}
func messageRouter() {
	message := admin.Group("/message")
	{
		xhttp.GET(message, "/list", handler.AdminLogin)
		xhttp.POST(message, "/add", handler.AdminLogin)
		xhttp.PATCH(message, "/edit", handler.AdminLogin)
		xhttp.DELETE(message, "/delete", handler.AdminLogin)
	}
}
func settingRouter() {
	setting := admin.Group("/setting")
	activity := setting.Group("/activity")
	{
		xhttp.GET(activity, "/list", handler.GetAdminProductConfigList)
		xhttp.PATCH(activity, "/:id", handler.AdminLogin)
		xhttp.DELETE(activity, "/:id", handler.AdminLogin)
	}
	topUp := setting.Group("/top-up")
	{
		xhttp.GET(topUp, "/list", handler.GetAdminProductConfigList)
		xhttp.PATCH(topUp, "/:id", handler.AdminLogin)
		xhttp.DELETE(topUp, "/:id", handler.AdminLogin)
	}
	product := setting.Group("/product")
	{
		xhttp.GET(product, "/list", handler.GetAdminProductConfigList)
		xhttp.PATCH(product, "/:id", handler.AdminLogin)
		xhttp.DELETE(product, "/:id", handler.AdminLogin)
	}
}
