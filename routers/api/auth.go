package api

import "github.com/gin-gonic/gin"

func ChenHua(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "chenhua",
	})
}
