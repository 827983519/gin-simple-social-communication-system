package routers


import (
	"github.com/gin-gonic/gin"
	follow_services "sbs-entrytask-template/apps/example/services/follow"
	login_services "sbs-entrytask-template/apps/example/services/login"
	"sbs-entrytask-template/apps/example/services/user"
)


func SetupRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/user")
	{
		v1.GET("info/search", user_services.Get_user_info)
		v1.GET("/follow_list/search", user_services.Get_user_follow_list)
		v1.GET("/view_other/search", user_services.View_other_user_info)
	}

	v2 := r.Group("login")
	{
		v2.GET("", login_services.Login)
	}

	v3 := r.Group("follow")
	{
		v3.POST("/create", follow_services.Follow_user)
		v3.POST("/delete", follow_services.Get_off_user)
	}

	return r
}
