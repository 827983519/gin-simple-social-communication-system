package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sbs-entrytask-template/libs/encrypt"
)

var NOT_CHECK_URL = []string{"as"}


func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, _ := c.Request.Cookie("token")
		if cookie == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"msg": "请先登陆"})
			//c.Redirect(http.StatusUnauthorized, "login")
			c.Abort()
		}

		_, err := encrypt.AesDeCrypt([]byte(cookie.Value), encrypt.PwdKey)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"msg": "请先登陆"})
			//c.Redirect(http.StatusUnauthorized, "login")
			c.Abort()
		}

		c.Next()
	}
}
