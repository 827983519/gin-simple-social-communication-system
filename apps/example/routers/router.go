package routers


import (
	"github.com/gin-gonic/gin"
	"sbs-entrytask-template/apps/example/services/user"
)


func SetupRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/user")
	{
		v1.GET("get", user_services.Get_user_info())
	}

	return r
}
