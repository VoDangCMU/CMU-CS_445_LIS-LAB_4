package check

import (
	"github.com/gin-gonic/gin"
)

func CheckWithToken(c *gin.Context) {
	user_id := c.MustGet("user_id").(string)
	c.JSON(200, gin.H{
		"status":  200,
		"message": "Check with token!",
		"user_id": user_id,
	})
	return
}
