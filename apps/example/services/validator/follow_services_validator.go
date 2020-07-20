package Validator

import "github.com/gin-gonic/gin"

type FollowForm struct {
	UserId   int32 `form:"user_id" json:"user_id" binding:"required"`
	FollowBy int32 `form:"follow_by" json:"follow_by" binding:"required"`
}


func CheckFollowForm(c *gin.Context) (FollowForm, error) {
	var validateForm FollowForm
	err := c.Bind(&validateForm)
	return validateForm, err
}
