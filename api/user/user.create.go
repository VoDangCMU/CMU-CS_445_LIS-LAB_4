package user

import (
	"github.com/VoDangCMU/CMU-CS_445_LIS-LAB_4/providers"
	"log"
	"net/http"

	"github.com/VoDangCMU/CMU-CS_445_LIS-LAB_4/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateUser(c *gin.Context, r providers.ClientReq) {
	db := c.MustGet("db").(*gorm.DB)
	var userName *string
	if r.Username != nil && *r.Username != "" {
		userName = r.Username
	} else if r.Email != nil && *r.Email != "" {
		userName = r.Email
	}

	err := db.Transaction(func(tx *gorm.DB) error {
		userAuth := models.UserAuthentication{
			Username:   userName,
			Password:   r.Password,
			Permission: "member",
		}
		if err := tx.Create(&userAuth).Error; err != nil {
			return err
		}

		userInfor := models.UserInformation{
			FullName:    r.Fullname,
			Email:       providers.CheckNullString(r.Email),
			Phone:       providers.CheckNullString(r.Phone),
			DateOfBirth: r.DateOfBirth,
			UserId:      userAuth.UserId,
		}

		if err := tx.Create(&userInfor).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Println("Transaction error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot create user1: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User successfully created"})
}
