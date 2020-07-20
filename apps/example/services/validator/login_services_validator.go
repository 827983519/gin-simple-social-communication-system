package Validator

import "github.com/gin-gonic/gin"

type LoginForm struct {
	Username string  `form:"username" binding:"required"`
	Password string  `form:"password" binding:"required"`
}


func CheckLoginForm(c *gin.Context) (LoginForm, error) {
	var validateForm LoginForm
	err := c.Bind(&validateForm)
	return validateForm, err
}

