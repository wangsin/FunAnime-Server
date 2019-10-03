package controller

import "github.com/gin-gonic/gin"

func TestController(c *gin.Context) {
	userAgent := c.GetHeader("user-agent")

	c.JSON(200, gin.H{
		"message": userAgent,
	})
}
