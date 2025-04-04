package auth

import (
	"github.com/VoDangCMU/CMU-CS_445_LIS-LAB_4/api/user"
	"github.com/VoDangCMU/CMU-CS_445_LIS-LAB_4/providers"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var req providers.ClientReq
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("UserRequest binding error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "UserRequest binding error: " + err.Error()})
		return
	}
	*req.Password = providers.HashPassword(*req.Password)
	if (req.Fullname == nil && req.Email == nil) ||
		req.Password == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email/Username and Password must be provided"})
		return
	}

	log.Println("Parsed Request:", req)

	user.CreateUser(c, req)
}
