package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)



func Response_data(c *gin.Context, data gin.H, retcode int32, message string){
	c.JSON(http.StatusOK, gin.H{
		"data":    data,
		"retcode": retcode,
		"message": message,
	})
}