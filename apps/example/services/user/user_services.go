package services

import (
	"apps/example/repository"
	"github.com/gin-gonic/gin"
)



func Get_user_info(c *gin.Context) {
	user_id := c.Query("user_id")
	User
}