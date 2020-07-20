package Validator

import "github.com/gin-gonic/gin"


func Check(c *gin.Context, validatorForm interface{}) (interface{}, error){
	err := c.ShouldBindJSON(&validatorForm)
	return validatorForm, err
}


type GetUserForm struct {
	UserId int32  `form:"user_id" binding:"required"`
}

type GetOtherUserForm struct {
	UserId int32  `form:"user_id" binding:"required"`
	ViewBy int32  `form:"view_by" binding:"required"`
}

func CheckGetUserForm(c *gin.Context) (GetUserForm, error) {
	var validateForm GetUserForm
	err := c.ShouldBindJSON(&validateForm)
	return validateForm, err
}

func CheckGetOtherUserForm(c *gin.Context) (GetOtherUserForm, error) {
	var validateForm GetOtherUserForm
	err := c.ShouldBindJSON(&validateForm)
	return validateForm, err
}












